package path

import (
	"strings"
)

/*
	"" -> "", ""
	"xxx" -> "", "xxx"
	"/" -> "/", ""
	"/text/" -> "/text", "/"
	"/prefix/text/" ->"/prefix", "/text/"
*/
func PopPrefix(url string) (first string, rest string) {
	if len(url) == 0 || url[0] != '/' {
		return "", url
	}
	i := strings.Index(url[1:], "/")
	if i < 0 {
		return url, ""
	}
	return url[:i+1], url[i+1:]
}
