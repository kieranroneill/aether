package routes

import (
	"aether/internal/constants"
	"aether/internal/errors"
	"aether/internal/merkletree"
	"aether/internal/types"
	"aether/internal/utils"
	"fmt"
	"github.com/labstack/echo/v4"
	"io"
	"net/http"
	"os"
	"sort"
)

func NewGetFilesRoute() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, c)
	}
}

func NewPostFilesUploadRoute() echo.HandlerFunc {
	return func(c echo.Context) error {
		var fileDirectory []*types.FileDirectoryItem
		var hashError *errors.HashError
		var readError *errors.ReadError
		var writeError *errors.WriteError

		form, err := c.MultipartForm()
		if err != nil {
			fmt.Println(err)

			return c.JSON(http.StatusInternalServerError, err)
		}

		formFiles := form.File["files"]
		subDirName, err := utils.CreateRandomSha256Hash()
		if err != nil {
			hashError = errors.NewHashError("failed to create .files/ subdirectory hash", err)

			fmt.Println(hashError)

			return c.JSON(http.StatusInternalServerError, hashError)
		}

		dirName := fmt.Sprintf("%s/%s", constants.RootFileDirectory, subDirName)

		// create the subdirectory to store these files
		err = utils.CreateDirectory(dirName)
		if err != nil {
			writeError = errors.NewWriteError(fmt.Sprintf("failed create the %s directory", dirName), err)

			fmt.Println(writeError)

			return c.JSON(http.StatusInternalServerError, writeError)
		}

		// iterate through each file, hash it and copy it to the .files/<subDirName> directory
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
			fileHash, err := utils.HashFile(uploadFile, fileHeader.Filename)
			if err != nil {
				hashError := errors.NewHashError(fmt.Sprintf("unable to hash file %s", fileHeader.Filename), err)

				fmt.Println(hashError)

				return c.JSON(http.StatusInternalServerError, hashError)
			}

			fmt.Println(fmt.Sprintf("file %s has a hash of %s", fileHeader.Filename, fileHash))

			// add to the directory
			fileDirectory = append(fileDirectory, &types.FileDirectoryItem{
				Hash: fileHash,
				Name: fileHeader.Filename,
			})

			// store the file in the .files/<subDirName> directory
			destFile, err := os.Create(fmt.Sprintf("%s/%s", dirName, fileHeader.Filename))
			if err != nil {
				writeError = errors.NewWriteError(fmt.Sprintf("failed to write file to %s/%s", dirName, fileHeader.Filename), err)

				fmt.Println(writeError)

				return c.JSON(http.StatusInternalServerError, writeError)
			}
			defer destFile.Close()

			// copy file contents to of the files from the upload to the source
			if _, err = io.Copy(destFile, uploadFile); err != nil {
				writeError = errors.NewWriteError(fmt.Sprintf("failed to write contents of file %s", fileHeader.Filename), err)

				fmt.Println(writeError)

				return c.JSON(http.StatusInternalServerError, writeError)
			}
		}

		// sort the directory by hash
		sort.Slice(fileDirectory, func(i int, j int) bool {
			return fileDirectory[i].Hash < fileDirectory[j].Hash
		})

		// create a merkle tree root
		merkleTreeRoot := merkletree.GenerateMerkleTreeRoot(utils.ExtractHashesFromFileDirectory(fileDirectory))

		fmt.Println(fmt.Sprintf("created merkle tree root %s", merkleTreeRoot))

		// create a directory file
		err = utils.WriteFileDirectoryToStorage(dirName, fileDirectory)
		if err != nil {
			fmt.Println(fmt.Sprintf("failed to create directory file at %s", dirName))
		}

		// finally return the merkle tree root
		return c.JSON(http.StatusOK, types.FilesUploadResponse{
			Directory: fileDirectory,
			Root:      merkleTreeRoot,
		})
	}
}
