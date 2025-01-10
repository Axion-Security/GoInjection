package waf

import (
	"GoInjection/backend/helper"
	"GoInjection/backend/structs"
	"io"
	"net/http"
)

/*
 * https://github.com/xscorp/wappalyzer/tree/master/src/technologies
 */

var Wafs = []struct {
	Name string
	Func func(http.Header, string) bool
}{
	{Name: "Cloudflare", Func: Cloudflare},
	{Name: "Cloudfront", Func: Cloudfront},
	{Name: "Akamai", Func: Akamai},
	{Name: "Sucuri", Func: Sucuri},
}

func sendRequest(urlStr string) (string, http.Header, int) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", urlStr, nil)
	helper.LogError(err)

	for key, value := range helper.DefaultHeaders() {
		req.Header.Set(key, value)
	}

	resp, err := client.Do(req)
	if err != nil {
		return "", resp.Header, resp.StatusCode
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		helper.LogError(err)
	}(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return "", resp.Header, resp.StatusCode
	}

	body, err := io.ReadAll(resp.Body)
	helper.LogError(err)

	return string(body), resp.Header, resp.StatusCode
}

func DetectWAF(url string) (bool, string) {
	body, headers, status := sendRequest(url)
	structs.WAFLinks++

	if status != http.StatusOK {
		return true, "Site not reachable"
	}

	for _, waf := range Wafs {
		if waf.Func(headers, body) {
			return true, waf.Name
		}
	}

	return false, ""
}
