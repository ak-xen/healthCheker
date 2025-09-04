package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gookit/color"
)

func main() {
	//fmt.Println("Enter the path to the file...")
	var path string = "urls.txt"
	/*_, err := fmt.Scan(&path)
	if err != nil {
		return
	}*/

	var urls []string = GetData(path)

	var data map[string]int

	ticker := time.NewTicker(3 * time.Second)
	done := make(chan os.Signal)
	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)

	defer ticker.Stop()
	for {
		select {
		case <-done:
			fmt.Println("Exiting...")
			return
		case <-ticker.C:
			data = queryManager(urls)
			for k, v := range data {

				color.Blue.Print(k)
				fmt.Print(":")

				if isRightCode(v) {
					color.Green.Println(v)
				} else {
					color.Red.Println(v)
				}

			}
		}
	}

}
