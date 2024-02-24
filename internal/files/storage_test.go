package files

import (
	"aether/internal/constants"
	"log"
	"os"
	"testing"
)

func TestGetRootFilesDirectory(t *testing.T) {
	rootFileDir, writeError := GetRootFilesDirectory()
	if writeError != nil {
		log.Fatal(writeError)
	}

	if rootFileDir != constants.RootFileDirectory {
		t.Errorf("expect result: %s, actual result: %s", rootFileDir, constants.RootFileDirectory)
	}

	// clean up
	err := os.RemoveAll(rootFileDir)
	if err != nil {
		log.Fatal(err)
	}
}
