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

func GetAllFileDirectories() (map[string][]*types.FileDirectoryItem, error) {
	var fileDirectory []*types.FileDirectoryItem
	var fileDirectories map[string][]*types.FileDirectoryItem

	directories, err := os.ReadDir(constants.RootFileDirectory)
	if err != nil {
		return nil, err
	}

	fileDirectories = map[string][]*types.FileDirectoryItem{}

	for _, dir := range directories {
		// read the directory file
		data, err := os.ReadFile(fmt.Sprintf("%s/%s/%s", constants.RootFileDirectory, dir.Name(), constants.DirectoryFileName))
		if err != nil {
			return nil, err
		}

		err = json.Unmarshal(data, &fileDirectory)
		if err != nil {
			return nil, err
		}

		fileDirectories[dir.Name()] = fileDirectory
	}

	return fileDirectories, nil
}

func WriteFileDirectoryJSONToStorage(dir string, fileDirectory []*types.FileDirectoryItem) error {
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
