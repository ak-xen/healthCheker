package main

import (
	"net/http"
	"sync"
)

func queryManager(urls []string) map[string]int {
	wg := sync.WaitGroup{}
	mu := sync.Mutex{}
	wg.Add(len(urls))

	responses := make(map[string]int)
	for _, url := range urls {
		go getStatusCode(url, &wg, responses, &mu)
	}

	wg.Wait()
	return responses

}

func getStatusCode(url string, wg *sync.WaitGroup, resps map[string]int, mutex *sync.Mutex) {
	defer wg.Done()
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	mutex.Lock()
	resps[url] = resp.StatusCode
	mutex.Unlock()
}
