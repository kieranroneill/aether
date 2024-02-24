package files

import (
	"aether/internal/constants"
	"aether/internal/errors"
	"fmt"
	"log"
	"os"
)

// GetRootFilesDirectory Checks if the root file directory exists, if it doesn't, it creates it.
func GetRootFilesDirectory() (string, *errors.WriteError) {
	_, err := os.Stat(constants.RootFileDirectory)
	if os.IsNotExist(err) {
		log.Printf("root file directory %s does not exist, creating a new one", constants.RootFileDirectory)

		err = os.Mkdir(constants.RootFileDirectory, 0755)
		if err != nil {
			return "", errors.NewWriteError(fmt.Sprintf("failed to create directory %s", constants.RootFileDirectory), err)
		}

		log.Printf("created root file directory %s", constants.RootFileDirectory)
	}

	return constants.RootFileDirectory, nil
}
