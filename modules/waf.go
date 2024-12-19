package modules

import (
	"GoInjection/helper"
	"GoInjection/waf"
	"strings"
)

var wafPayload = "' OR 1=1; UNION SELECT NULL, NULL, CONCAT('<img src=\"x\" onerror=\"alert(document.cookie)\">', table_name) FROM information_schema.tables WHERE 'a'='a' -- -; EXEC xp_cmdshell('cat /etc/passwd')--\n"

func DetectWAF(url string) (bool, string) {
	responseWithPayload, statusCodeWithPayload := helper.SendRequest(strings.Replace(url, helper.PayloadReplaceString, helper.URLEncode(wafPayload), -1))
	responseWithoutPayload, statusCodeWithoutPayload := helper.SendRequest(url)

	if responseWithPayload == responseWithoutPayload {
		return false, "Response changed"
	}

	if statusCodeWithPayload != statusCodeWithoutPayload {
		return true, "Status code changed"
	}

	var isWAF, wafDetection = waf.DetectWAF(url)
	if isWAF {
		return true, wafDetection
	}

	return false, ""
}
