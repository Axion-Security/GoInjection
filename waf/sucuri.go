package waf

import (
	"net/http"
)

func Sucuri(headers http.Header, response string) bool {
	XSucuriCacheHeader := headers.Get("x-sucuri-cache")
	XSucuriIdHeader := headers.Get("x-sucuri-id")

	if XSucuriCacheHeader != "" || XSucuriIdHeader != "" {
		return true
	}
	return false
}
