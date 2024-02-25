package utils

import (
	"fmt"
	"os"
)

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

func DoesDirExist(dirName string) bool {
	_, err := os.Stat(dirName)
	if os.IsNotExist(err) {
		return false
	}

	return true
}

func RemoveDir(dirName string) {
	err := os.RemoveAll(dirName)
	if err != nil {
		fmt.Println(fmt.Sprintf("failed to remove directory %s", dirName))
	}
}
