package main

import (
	"github.com/ak-xen/healthCheker/Concurency"
	"github.com/ak-xen/healthCheker/FileManager"
)

func main() {
	//fmt.Println("Enter the path to the file...")
	var path string = "urls.txt"
	/*_, err := fmt.Scan(&path)
	if err != nil {
		return
	}*/

	var urls []string = FileManager.GetData(path)
	Concurency.GetAndDrawData(urls)

}
