package main

import (
	"flag"

	"github.com/ak-xen/healthCheker/Concurency"
	"github.com/ak-xen/healthCheker/FileManager"
)

func main() {

	var pathToFile = flag.String("path", "urls.txt", "pathToFile to file")
	flag.Parse()
	var urls []string = FileManager.GetData(*pathToFile)
	Concurency.GetAndDrawData(urls)

}
