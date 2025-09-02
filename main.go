package main

import (
	"fmt"
	"sync"
)

func main() {
	//fmt.Println("Enter the path to the file...")
	var path string = "urls.txt"
	/*_, err := fmt.Scan(&path)
	if err != nil {
		return
	}*/

	urls := GetData(path)

	wg := sync.WaitGroup{}
	mu := sync.Mutex{}
	wg.Add(len(urls))
	resps := make(map[string]string)
	for _, url := range urls {
		go getStatusCode(url, &wg, resps, &mu)
	}

	wg.Wait()

	fmt.Println(resps)
}
