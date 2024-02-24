package files

import (
	"aether/internal/errors"
	"fmt"
	"os"
)

// CreateDirectory Creates the directory if it doesn't exist
func CreateDirectory(dirName string) *errors.WriteError {
	_, err := os.Stat(dirName)
	if os.IsNotExist(err) {
		fmt.Printf("file directory %s does not exist, creating a new one", dirName)

		err = os.Mkdir(dirName, 0755)
		if err != nil {
			return errors.NewWriteError(fmt.Sprintf("failed to create directory %s", dirName), err)
		}

		fmt.Printf("created file directory %s", dirName)
	}

	return nil
}
