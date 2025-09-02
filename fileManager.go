package main

import (
	"bufio"
	"os"
	"strings"
)

type File struct {
	Data *os.File
}

func (f *File) ReadFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	f.Data = file

}

func GetData(path string) []string {

	file := new(File)

	file.ReadFile(path)

	return parseFile(file)

}

func parseFile(file *File) []string {
	var urls []string
	fileData := file.Data
	scanner := bufio.NewScanner(fileData)
	for scanner.Scan() {
		line := scanner.Text()
		urls = append(urls, strings.TrimSpace(line))
	}
	return urls
}
