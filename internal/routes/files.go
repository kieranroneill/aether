package routes

import (
	"aether/internal/constants"
	"aether/internal/errors"
	internalfiles "aether/internal/files"
	"aether/internal/merkletree"
	"aether/internal/types"
	"aether/internal/utils"
	"fmt"
	"github.com/labstack/echo/v4"
	"io"
	"net/http"
	"os"
)

func copyFiles(srcDir string, destDir string, fileMetadatas []*types.FileMetadata) error {
	// create a new directory at the root .files directory with the merkle tree root as the name
	writeError := internalfiles.CreateDirectory(destDir)
	if writeError != nil {
		fmt.Println(writeError)

		// attempt to clean up
		err := removeTempFiles(srcDir)
		if err != nil {
			return err
		}

		return writeError.Error
	}

	for _, fileMetadata := range fileMetadatas {
		// read the source file
		srcFile, err := os.Open(fmt.Sprintf("%s/%s", srcDir, fileMetadata.FileName))
		if err != nil {
			return err
		}
		defer srcFile.Close()

		// create the destination file
		destFile, err := os.Create(fmt.Sprintf("%s/%s", destDir, fileMetadata.FileName))
		if err != nil {
			return err
		}
		defer destFile.Close()

		// copy the contents over
		_, err = io.Copy(destFile, srcFile)
		if err != nil {
			return err
		}
	}

	return nil
}

func extractFileHashesFromFileMetadata(fileMetadatas []*types.FileMetadata) []string {
	hashes := make([]string, len(fileMetadatas))

	for i := range fileMetadatas {
		hashes[i] = fileMetadatas[i].Hash
	}

	return hashes
}

func removeTempFiles(tempDir string) error {
	err := os.RemoveAll(fmt.Sprintf("%s", tempDir)) // the entire .temp/ subdirectory
	if err != nil {
		return err
	}

	return nil
}

func NewFilesUploadRoute() echo.HandlerFunc {
	return func(c echo.Context) error {
		var fileMetadatas []*types.FileMetadata
		var hashError *errors.HashError
		var readError *errors.ReadError
		var writeError *errors.WriteError

		form, err := c.MultipartForm()
		if err != nil {
			fmt.Println(err)

			return c.JSON(http.StatusInternalServerError, err)
		}

		formFiles := form.File["files"]
		tempSubDir, err := utils.CreateRandomSha256Hash()
		if err != nil {
			hashError = errors.NewHashError("failed to create temp sub-directory hash", err)

			fmt.Println(hashError)

			return c.JSON(http.StatusInternalServerError, hashError)
		}

		tempDir := fmt.Sprintf("%s/%s", constants.TempFileDirectory, tempSubDir)

		// create the temp subdirectory where the files will be stored before the merkle root is calculated
		writeError = internalfiles.CreateDirectory(tempDir)
		if writeError != nil {
			fmt.Println(writeError)

			return c.JSON(http.StatusInternalServerError, writeError)
		}

		// iterate through each file, hash it and copy it to the .temp/ directory
		for _, fileHeader := range formFiles {
			// read file
			uploadFile, err := fileHeader.Open()
			if err != nil {
				readError = errors.NewReadError(fmt.Sprintf("failed to read file %s", fileHeader.Filename), err)

				fmt.Println(readError)

				return c.JSON(http.StatusInternalServerError, readError)
			}
			defer uploadFile.Close()

			// get a hash of the file
			fileHash, hashError := internalfiles.HashFile(uploadFile, fileHeader.Filename)
			if hashError != nil {
				fmt.Println(hashError)

				return c.JSON(http.StatusInternalServerError, hashError)
			}

			fmt.Println(fmt.Sprintf("file %s has a hash of %s", fileHeader.Filename, fileHash))

			fileMetadatas = append(fileMetadatas, &types.FileMetadata{
				FileName: fileHeader.Filename,
				Hash:     fileHash,
			})

			// store the file in the .temp/ directory
			tempFile, err := os.Create(fmt.Sprintf("%s/%s", tempDir, fileHeader.Filename))
			if err != nil {
				writeError = errors.NewWriteError(fmt.Sprintf("failed to write file %s", fileHeader.Filename), err)

				fmt.Println(writeError)

				return c.JSON(http.StatusInternalServerError, writeError)
			}
			defer tempFile.Close()

			// copy file contents to of the files
			if _, err = io.Copy(tempFile, uploadFile); err != nil {
				writeError = errors.NewWriteError(fmt.Sprintf("failed to write contents of file %s", fileHeader.Filename), err)

				fmt.Println(writeError)

				return c.JSON(http.StatusInternalServerError, writeError)
			}
		}

		// create a merkle tree root
		merkleTreeRoot := merkletree.GenerateMerkleTreeRoot(extractFileHashesFromFileMetadata(fileMetadatas))
		merkleTreeRootDir := fmt.Sprintf("%s/%s", constants.RootFileDirectory, merkleTreeRoot)

		fmt.Println(fmt.Sprintf("created merkle tree root %s", merkleTreeRoot))
		fmt.Println(fmt.Sprintf("copying files from %s to %s", tempDir, merkleTreeRootDir))

		// copy the files from the .temp/ directory to the .files/ merkle tree directory
		err = copyFiles(
			tempDir,
			merkleTreeRootDir,
			fileMetadatas,
		)
		if err != nil {
			writeError = errors.NewWriteError(fmt.Sprintf("failed to copy files from %s to %s", tempDir, merkleTreeRootDir), err)

			fmt.Println(writeError)

			return c.JSON(http.StatusInternalServerError, writeError)
		}

		fmt.Println(fmt.Sprintf("cleaning up files from %s", tempDir))

		// remove the temp directory files
		err = removeTempFiles(tempDir)
		if err != nil {
			fmt.Println(err)
		}

		// finally return the merkle tree root
		return c.JSON(http.StatusOK, types.FilesUploadResponse{
			Root: merkleTreeRoot,
		})
	}
}
