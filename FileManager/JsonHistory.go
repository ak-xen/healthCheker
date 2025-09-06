package FileManager

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type LogQuery struct {
	TimeQuery         string `json:"timeQuery"`
	SiteWithoutAccess string `json:"siteWithoutAccess"`
}

func getToPathFile() string {
	fileName := "jsonHistory.json"
	wd, err := os.Getwd()
	pathToFile := filepath.Join(wd, fileName)
	if err != nil {
		panic(err)
	}
	return pathToFile
}

func CreateJsonHistory() {

	file, err := os.Create(getToPathFile())
	if err != nil {
		fmt.Println("Error creating history file")
		return
	}

	defer file.Close()
}

func LogAccessDenied(url string, time string) {
	pathToFile := getToPathFile()
	if _, err := os.Stat(pathToFile); os.IsNotExist(err) {
		go CreateJsonHistory()
		fmt.Println("Created history file")
	}
	file, err := os.OpenFile(pathToFile, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return
	}
	lg := LogQuery{TimeQuery: time, SiteWithoutAccess: url}

	encoder := json.NewEncoder(file)
	err = encoder.Encode(lg)
	if err != nil {
		fmt.Println("Error marshalling history")
		return
	}

}
