package waf

import (
	"net/http"
	"strings"
)

func Cloudflare(headers http.Header, response string) bool {
	ServerHeader := strings.Contains(strings.ToLower(headers.Get("Server")), "cloudflare")
	if ServerHeader {
		return true
	}

	CfRayHeader := headers.Get("cf-ray")
	if CfRayHeader != "" {
		return true
	}

	return false
}
