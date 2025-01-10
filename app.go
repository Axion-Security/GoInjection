package main

import (
	"GoInjection/backend/helper"
	"GoInjection/backend/injections"
	"GoInjection/backend/injections/blind"
	helper2 "GoInjection/backend/injections/helper"
	"GoInjection/backend/modules"
	"GoInjection/backend/structs"
	"context"
	"strconv"
	"strings"
)

type App struct {
	ctx context.Context
}

func NewApp() *App {
	return &App{}
}

func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// SetTarget Input: (string: Option) (string: Value)
func (a *App) SetTarget(option string, value string) {
	switch option {
	case "url":
		structs.TargetURL = value
		break
	case "dbms":
		structs.TargetDBMS = value
		break
	case "syntax":
		structs.TargetSyntax = value
		break
	case "columns":
		target, _ := strconv.Atoi(value)
		structs.TargetColumns = target
		break
	case "database":
		structs.TargetDatabaseName = value
		break
	}
}

// GetTarget Input: (string: Option)
//
// GetTarget Returns: (string: Value)
func (a *App) GetTarget(option string) string {
	switch option {
	case "url":
		return structs.TargetURL
	case "dbms":
		return structs.TargetDBMS
	case "syntax":
		return structs.TargetSyntax
	case "columns":
		return strconv.Itoa(structs.TargetColumns)
	case "database":
		return structs.TargetDatabaseName
	case "union":
		return strconv.FormatBool(structs.TargetUnion)
	case "error":
		return strconv.FormatBool(structs.TargetError)
	case "boolean":
		return strconv.FormatBool(structs.TargetBoolean)
	case "time":
		return strconv.FormatBool(structs.TargetTime)
	case "unionLinks":
		return strconv.Itoa(structs.UnionLinks)
	case "errorLinks":
		return strconv.Itoa(structs.ErrorLinks)
	case "booleanLinks":
		return strconv.Itoa(structs.BooleanLinks)
	case "timeLinks":
		return strconv.Itoa(structs.TimeLinks)
	case "fingerprintLinks":
		return strconv.Itoa(structs.FingerprintLinks)
	case "wafLinks":
		return strconv.Itoa(structs.WAFLinks)
	}
	return ""
}

func CheckVariables(vars interface{}) bool {
	switch v := vars.(type) {
	case string:
		return v == ""
	case int:
		return v == 0
	default:
		return false
	}
}

// CheckWAF Needs: -/-
//
// CheckWAF Returns: (string: Url, string: Detection)
func (a *App) CheckWAF() map[string]string {
	helper.LogLine("Function called.")

	cleanedURLs, err := helper.GetUrls(structs.TargetURL)
	helper.LogError(err)

	for _, cleanedURL := range cleanedURLs {
		isWaf, wafType := modules.DetectWAF(cleanedURL, structs.TargetURL)
		cleanedURL = helper.ExtractDomain(cleanedURL)

		isWafParsed, _ := strconv.ParseBool(isWaf)

		helper.LogLine("URL: " + cleanedURL)
		helper.LogLine("WAF: " + strconv.FormatBool(isWafParsed))
		helper.LogLine("WAF Type: " + wafType)

		return map[string]string{
			"isWaf":   isWaf,
			"wafType": wafType,
			"domain":  cleanedURL,
		}

	}

	return map[string]string{
		"isWaf":   "",
		"wafType": "",
		"domain":  "",
	}
}

// Fingerprint Needs: structs.TargetColumns
//
// Fingerprint Returns: (string: Url, string: DBMS)
//
// Fingerprint Input: (string: Method)
func (a *App) Fingerprint(method string) map[string]string {
	helper.LogLine("Function called.")

	cleanedURLs, err := helper.GetUrls(structs.TargetURL)
	helper.LogError(err)

	for _, cleanedURL := range cleanedURLs {
		detectedDB := modules.FingerprintDB(cleanedURL, method)
		cleanedURL = helper.ExtractDomain(cleanedURL)

		helper.LogLine("URL: " + cleanedURL)
		helper.LogLine("DBMS: " + detectedDB)

		structs.TargetDBMS = detectedDB

		return map[string]string{
			"dbms":   detectedDB,
			"domain": cleanedURL,
		}
	}

	return map[string]string{
		"dbms":   "",
		"domain": "",
	}
}

// Interpreter Needs: structs.TargetDBMS
//
// Interpreter Returns: (string: DBMS)
func (a *App) Interpreter() string {
	if CheckVariables(structs.TargetDBMS) {
		return ""
	}

	helper.LogLine("Function called.")

	structs.TargetSyntax = strings.ReplaceAll(helper2.Interpreter(structs.TargetDBMS), " (default)", "")

	helper.LogLine("DBMS: " + structs.TargetDBMS)
	helper.LogLine("Syntax: " + structs.TargetSyntax)

	return structs.TargetSyntax
}

// Resolver Needs: -/-
//
// Resolver Returns: (string: Columns, string: Database Name)
func (a *App) Resolver() map[string]string {
	helper.LogLine("Function called.")
	cleanedURLs, err := helper.GetUrls(structs.TargetURL)
	helper.LogError(err)

	for _, cleanedURL := range cleanedURLs {
		_, structs.TargetColumns = helper2.GetColumnCount(cleanedURL)
		_, structs.TargetDatabaseName = helper2.GetDatabase(cleanedURL, structs.TargetColumns)

		cleanedURL = helper.ExtractDomain(cleanedURL)
		helper.LogLine("URL:" + cleanedURL)
		helper.LogLine("Columns: " + strconv.Itoa(structs.TargetColumns))
		helper.LogLine("Database: " + structs.TargetDatabaseName)

		return map[string]string{
			"domain":   cleanedURL,
			"columns":  strconv.Itoa(structs.TargetColumns),
			"database": structs.TargetDatabaseName,
		}
	}

	return map[string]string{
		"domain":   "",
		"columns":  "",
		"database": "",
	}
}

// Injection Needs: structs.TargetURL, structs.TargetColumns, structs.TargetSyntax
//
// Injection Returns: (string: Domain, string: Bool, string: Result)
//
// Injection Input: (string: Method)
func (a *App) Injection(method string) map[string]string {
	if CheckVariables(structs.TargetURL) || CheckVariables(structs.TargetColumns) || CheckVariables(structs.TargetSyntax) {
		return map[string]string{
			"domain": "",
			"bool":   "",
			"result": "",
		}
	}

	var injectionFunc func(string) (bool, string)
	switch method {
	case "union":
		injectionFunc = injections.UnionInjection
	case "error":
		injectionFunc = injections.ErrorInjection
	case "time":
		injectionFunc = blind.TimeBasedInjection
	case "boolean":
		injectionFunc = blind.BooleanBasedInjection
	default:
		return map[string]string{
			"domain": "",
			"bool":   "",
			"result": "",
		}
	}

	injectionBool, injectionResult := injectionFunc(structs.TargetURL)
	cleanedURL := helper.ExtractDomain(structs.TargetURL)
	helper.LogLine("URL: " + cleanedURL)
	helper.LogLine("Injection:" + injectionResult)
	helper.LogLine("Injection Bool: " + strconv.FormatBool(injectionBool))
	helper.LogLine("Injection Method: " + method)

	return map[string]string{
		"domain": cleanedURL,
		"bool":   strconv.FormatBool(injectionBool),
		"result": injectionResult,
	}
}
