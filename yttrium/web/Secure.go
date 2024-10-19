package yttrium_web

import "net/http"

var (
	Key string
)

func CheckKey(request *http.Request) bool {
	targetKey := request.Header.Get("key")

	return targetKey == Key
}

func KeyExist(request *http.Request) bool {
	return request.Header.Get("key") != ""
}
