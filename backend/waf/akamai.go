package waf

import (
	"net/http"
)

func Akamai(headers http.Header, response string) bool {
	XAkamaiTransformedHeader := headers.Get("X-Akamai-Transformed")
	XEdgeConnectMidMileRTTHeader := headers.Get("X-EdgeConnect-MidMile-RTT")
	XEdgeConnectOriginMEXLatencyHeader := headers.Get("X-EdgeConnect-Origin-MEX-Latency")

	if XAkamaiTransformedHeader != "" || XEdgeConnectMidMileRTTHeader != "" || XEdgeConnectOriginMEXLatencyHeader != "" {
		return true
	}

	return false
}
