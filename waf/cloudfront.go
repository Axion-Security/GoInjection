package waf

import (
	"net/http"
	"strings"
)

func Cloudfront(headers http.Header, response string) bool {
	ViaHeader := headers.Get("Via")
	XAmzCfIdHeader := headers.Get("X-Amz-Cf-Id")

	if strings.Contains(ViaHeader, "cloudfront") || XAmzCfIdHeader != "" {
		return true
	}

	return false
}
