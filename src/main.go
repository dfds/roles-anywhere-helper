package main

import (
	"fmt"
	"os"
)

func main() {
	linuxDir, err := os.UserHomeDir()

	file, err := create(linuxDir+"/.aws", "credentials2")

	if err != nil {
		fmt.Println("Error writing credential file:", err)
	} else {
		fmt.Println("Credential file written successfully!")
	}

	defer file.Close()
}

func create(filePath string, fileName string) (*os.File, error) {

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		os.MkdirAll(filePath, 0700)
	}

	_, e := os.Stat(filePath + "/" + fileName)

	if e == nil {
		return nil, os.ErrExist
	}
	return os.Create(filePath + "/" + fileName)
}
