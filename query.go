package main

import (
	"net/http"
	"strconv"
	"sync"
)

func getStatusCode(url string, wg *sync.WaitGroup, resps map[string]string, mutex *sync.Mutex) {
	defer wg.Done()
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	mutex.Lock()
	resps[url] = strconv.Itoa(resp.StatusCode)
	mutex.Unlock()
}
