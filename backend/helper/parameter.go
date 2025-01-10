package helper

import (
	"fmt"
	"net/url"
	"strings"
)

var PayloadReplaceString = "§PAYLOAD§"

// ToDo: Save also Parameters with their values
func GetUrls(inputURL string) ([]string, error) {
	parsedURL, err := url.Parse(inputURL)
	LogError(err)

	queryParams := parsedURL.Query()

	var cleanedURLs []string

	for key := range queryParams {
		var newQueryParams []string
		for qKey := range queryParams {
			if qKey == key {
				newQueryParams = append(newQueryParams, fmt.Sprintf("%s=%s", qKey, PayloadReplaceString))
			} else {
				newQueryParams = append(newQueryParams, fmt.Sprintf("%s=%s", qKey, queryParams.Get(qKey)))
			}
		}

		parsedURL.RawQuery = strings.Join(newQueryParams, "&")
		cleanedURLs = append(cleanedURLs, parsedURL.String())
	}

	return cleanedURLs, nil
}
