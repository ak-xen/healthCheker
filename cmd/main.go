package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/ak-xen/healthCheker/Concurency"
	"github.com/ak-xen/healthCheker/FileManager"
)

func main() {
	fmt.Println("ClI start")
	var pathTo = flag.String("path", "urls.txt", "pathToFile to file")
	wd, _ := os.Getwd()
	pathToFile := filepath.Join(wd, "cmd", *pathTo)
	flag.Parse()
	if _, err := os.Stat(pathToFile); os.IsNotExist(err) {
		fmt.Printf("Файл не найден: %s\n", pathToFile)
		os.Exit(1) // Завершаем программу с ошибкой
	}
	var urls []string = FileManager.GetData(pathToFile)
	Concurency.GetAndDrawData(urls)

}
