package fileHandler

import (
	"fmt"
	"os"
	"path/filepath"
)

func CreateFile(filePath string, fileName string) (*os.File, error) {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		os.MkdirAll(filePath, 0755)
	}
	absoluteFilePath := filepath.Join(filePath, fileName)

	_, e := os.Stat(absoluteFilePath)

	if e == nil {
		fmt.Printf("File Path already exisits.... Overwriting %s", absoluteFilePath)
		return os.Open(absoluteFilePath)
	}

	return os.Create(absoluteFilePath)
}
