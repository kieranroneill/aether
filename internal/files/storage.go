package files

import (
	"aether/internal/constants"
	"aether/internal/errors"
	"fmt"
	"os"
)

// CreateRootFilesDirectory Creates the root files directory if it doesn't exist
func CreateRootFilesDirectory() *errors.WriteError {
	_, err := os.Stat(constants.RootFileDirectory)
	if os.IsNotExist(err) {
		fmt.Printf("root file directory %s does not exist, creating a new one", constants.RootFileDirectory)

		err = os.Mkdir(constants.RootFileDirectory, 0755)
		if err != nil {
			return errors.NewWriteError(fmt.Sprintf("failed to create directory %s", constants.RootFileDirectory), err)
		}

		fmt.Printf("created root file directory %s", constants.RootFileDirectory)
	}

	return nil
}

// GetRootFilesDirectory Checks if the root file directory exists, if it doesn't, it creates it.
func GetRootFilesDirectory() (string, *errors.WriteError) {
	_, err := os.Stat(constants.RootFileDirectory)
	if os.IsNotExist(err) {
		writeError := CreateRootFilesDirectory()
		if writeError != nil {
			return "", writeError
		}
	}

	return constants.RootFileDirectory, nil
}
