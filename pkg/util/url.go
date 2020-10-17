package util

import "strings"

// AbsoluteURL add missing protocol to url.
func AbsoluteURL(v string) (ret string) {
	ret = strings.Trim(v, " \n")
	if ret == "" {
		return
	}
	if ret[0] == '/' {
		ret = "https:" + ret
	}
	return
}
