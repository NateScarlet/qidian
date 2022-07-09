package client

import (
	"fmt"
	"net/http"
)

func parseSetCookie(s string) (_ *http.Cookie, err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("parseSetCookie('%s'): %w", s, err)
		}
	}()

	var h = http.Header{
		"Set-Cookie": {s},
	}
	var req = http.Response{
		Header: h,
	}
	var cookies = req.Cookies()
	if len(cookies) != 1 {
		err = fmt.Errorf("invalid value")
		return
	}
	return cookies[0], nil
}
