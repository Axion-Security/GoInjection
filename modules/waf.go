package modules

import (
	"GoInjection/helper"
	"GoInjection/waf"
	"strings"
)

var wafPayload = "' OR 1=1; UNION SELECT NULL, NULL, CONCAT('<img src=\"x\" onerror=\"alert(document.cookie)\">', table_name) FROM information_schema.tables WHERE 'a'='a' -- -; EXEC xp_cmdshell('cat /etc/passwd')--\n"

func DetectWAF(url string) (bool, string) {
	resp_1, statusCode_1 := helper.SendRequest(strings.Replace(url, helper.PayloadReplaceString, helper.URLEncode(wafPayload), -1))
	resp_2, statusCode_2 := helper.SendRequest(url)

	if resp_1 == resp_2 {
		return false, "Response changed"
	}

	/*
		// unstable detection
		resp := strings.ToLower(resp_1)
		if strings.Contains(resp, "alert") || strings.Contains(resp, "forbidden") || strings.Contains(resp, "error") || strings.Contains(resp, "waf") {
			return true, "Keyword detected"
		}
	*/

	if statusCode_1 != statusCode_2 {
		return true, "Status code changed"
	}

	var isWAF, wafDetection = waf.DetectWAF(url)
	if isWAF {
		return true, wafDetection
	}

	return false, ""
}
