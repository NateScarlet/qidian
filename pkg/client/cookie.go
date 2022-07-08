package client

import (
	"net/http"
)

func parseCookies(s string) []*http.Cookie {
	var h = http.Header{
		"Cookie": {s},
	}
	var req = http.Request{
		Header: h,
	}
	return req.Cookies()
}
