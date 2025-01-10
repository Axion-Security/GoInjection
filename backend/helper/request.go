package helper

import (
	"github.com/corpix/uarand" // Random User-Agent
	"io"
	"log"
	"net/http"
	"net/url"
)

func DefaultHeaders() map[string]string {
	return map[string]string{
		"User-Agent":      uarand.GetRandom(),
		"Accept":          "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8",
		"Connection":      "keep-alive",
		"Accept-Language": "en-US,en;q=0.5",
	}
}

func SendRequest(urlStr string) (string, int) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", urlStr, nil)
	LogError(err)

	for key, value := range DefaultHeaders() {
		req.Header.Set(key, value)
	}

	resp, err := client.Do(req)
	if err != nil {
		return "", resp.StatusCode
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		LogError(err)
	}(resp.Body)
	if resp.StatusCode != http.StatusOK {
		return "", resp.StatusCode
	}

	body, err := io.ReadAll(resp.Body)
	LogError(err)

	return string(body), resp.StatusCode
}

func URLEncode(payload string) string {
	if payload == "" {
		log.Println("URLEncode received an empty payload")
		return ""
	}
	return url.QueryEscape(payload)
}

func URLEncodeQuery(input string) string {
	urlStr, _ := url.Parse(input)
	query := urlStr.Query()
	urlStr.RawQuery = query.Encode()
	return urlStr.String()
}

func ExtractDomain(urlStr string) string {
	u, err := url.Parse(urlStr)
	LogError(err)
	return u.Hostname()
}
