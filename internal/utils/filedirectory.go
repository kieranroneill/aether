package utils

import (
	"aether/internal/constants"
	"aether/internal/types"
	"encoding/json"
	"fmt"
	"os"
)

func ExtractHashesFromFileDirectory(fileDirectory []*types.FileDirectoryItem) []string {
	hashes := make([]string, len(fileDirectory))

	for i := range fileDirectory {
		hashes[i] = fileDirectory[i].Hash
	}

	return hashes
}

func WriteFileDirectoryToStorage(dir string, fileDirectory []*types.FileDirectoryItem) error {
	// create a file at the directory
	file, err := os.Create(fmt.Sprintf("%s/%s", dir, constants.DirectoryFileName))
	if err != nil {
		return err
	}
	defer file.Close()

	// convert the file directory to a json
	data, err := json.Marshal(fileDirectory)
	if err != nil {
		return err
	}

	// write the json data to the file
	_, err = file.Write(data)
	if err != nil {
		return err
	}

	return nil
}
