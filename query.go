package main

import (
	"net/http"
)

func getStatusCode(url string) int {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	return resp.StatusCode
}
