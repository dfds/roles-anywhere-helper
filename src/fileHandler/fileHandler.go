package fileHandler

import (
	"os"
)

func CreateFile(filePath string) (*os.File, error) {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		os.MkdirAll(filePath, 0755)
	}

	_, e := os.Stat(filePath)

	if e == nil {
		return nil, os.ErrExist
	}
	return os.Create(filePath)
}
