package helper

import (
	"GoInjection/helper"
	"GoInjection/structs"
	"strings"
)

func RunInjection(url string, payload string) (bool, string) {
	response, status_code := helper.SendRequest(strings.Replace(url, helper.PayloadReplaceString, helper.URLEncodeQuery(payload), -1))
	// (Error-based Injection)
	if status_code == 500 || strings.Contains(response, "SQL syntax") || strings.Contains(response, "error") {
		structs.TargetError = true
	}

	// (Union-based Injection)
	if strings.Contains(response, "Union select") || strings.Contains(response, "id") {
		structs.TargetUnion = true
	}

	// (Blind-based Injection)
	if strings.Contains(response, "1=1") || strings.Contains(response, "0=0") {
		structs.TargetBlind = true
	}

	return true, "Done."
}
