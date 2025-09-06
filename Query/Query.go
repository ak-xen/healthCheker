package Query

import (
	"net/http"
	"sync"
	"time"

	"github.com/ak-xen/healthCheker/FileManager"
	"github.com/ak-xen/healthCheker/Utils"
)

func QueryManager(urls []string) map[string]int {
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
	now := time.Now()
	timeQuery := now.Format("2006-01-02 15:04:05")
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	mutex.Lock()
	if Utils.IsRightCode(resp.StatusCode) {
		resps[url] = resp.StatusCode
	} else {
		FileManager.LogAccessDenied(url, timeQuery)
		resps[url] = 0
	}

	mutex.Unlock()
}
