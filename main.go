package main

import "fmt"

func main() {
	//fmt.Println("Enter the path to the file...")
	var path string = "urls.txt"
	/*_, err := fmt.Scan(&path)
	if err != nil {
		return
	}*/

	urls := GetData(path)
	fmt.Println(urls)
}
