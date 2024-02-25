package utils

import (
	"aether/internal/types"
	"fmt"
	"io"
	"os"
)

func SaveFilesToDir(dirName string, files []*types.FileReadData) error {
	for _, file := range files {
		destFile, err := os.Create(fmt.Sprintf("%s/%s", dirName, file.Name))
		if err != nil {
			return err
		}
		defer destFile.Close()

		_, err = io.Copy(destFile, file.File)
		if err != nil {
			return err
		}
	}

	return nil
}

// CreateDir Creates a directory folder to storage if it doesn't exist
func CreateDir(dirName string) error {
	_, err := os.Stat(dirName)
	if os.IsNotExist(err) {
		fmt.Printf("directory %s does not exist, creating a new one", dirName)

		err = os.Mkdir(dirName, 0755)
		if err != nil {
			return err
		}

		fmt.Printf("created new directory %s", dirName)
	}

	return nil
}
