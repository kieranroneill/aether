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
)

func NewGetFilesRoute() echo.HandlerFunc {
	return func(c echo.Context) error {
		var filesResponse map[string][]*types.FileResponse
		var hashes []string
		var readError *errors.ReadError

		fileDirectories, err := utils.GetAllFileDirectories()
		if err != nil {
			readError = errors.NewReadError("failed to get file directories", err)

			fmt.Println(readError)

			return c.JSON(http.StatusInternalServerError, readError)
		}

		filesResponse = map[string][]*types.FileResponse{}

		// add the merkle proofs to the file directories
		for key, fileDirectory := range fileDirectories {
			hashes = utils.ExtractHashesFromFileDirectory(fileDirectory)

			for _, value := range fileDirectory {
				filesResponse[key] = append(filesResponse[key], &types.FileResponse{
					Hash:  value.Hash,
					Name:  value.Name,
					Proof: merkletree.GenerateMerkleTreeProof(value.Hash, hashes),
				})
			}
		}

		return c.JSON(http.StatusOK, filesResponse)
	}
}

func NewPostFilesUploadRoute() echo.HandlerFunc {
	return func(c echo.Context) error {
		var fileDirectory []*types.FileDirectoryItem
		var readError *errors.ReadError
		var writeError *errors.WriteError

		form, err := c.MultipartForm()
		if err != nil {
			fmt.Println(err)

			return c.JSON(http.StatusInternalServerError, err)
		}

		formFiles := form.File["file"]

		// create a directory name with a random hash
		subDirName, err := utils.CreateRandomSha256Hash()
		if err != nil {
			hashError := errors.NewHashError("failed to create hash", err)

			fmt.Println(hashError)

			return c.JSON(http.StatusInternalServerError, hashError)
		}

		dirName := fmt.Sprintf("%s/%s", constants.RootFileDirectory, subDirName)

		// create the new directory
		err = utils.CreateDir(dirName)
		if err != nil {
			writeError = errors.NewWriteError(fmt.Sprintf("failed create the %s directory", dirName), err)

			fmt.Println(writeError)

			return c.JSON(http.StatusInternalServerError, writeError)
		}

		// iterate through each file, hash it and copy it to the .files/<dirName> directory
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

			// create the file
			destFile, err := os.Create(fmt.Sprintf("%s/%s", dirName, fileHeader.Filename))
			if err != nil {
				writeError = errors.NewWriteError(fmt.Sprintf("failed to create file %s/%s", dirName, fileHeader.Filename), err)

				fmt.Println(writeError)

				return c.JSON(http.StatusInternalServerError, writeError)
			}
			defer destFile.Close()

			// copy the contents of the uploaded file
			_, err = io.Copy(destFile, uploadFile)
			if err != nil {
				writeError = errors.NewWriteError(fmt.Sprintf("failed to write file %s/%s", dirName, fileHeader.Filename), err)

				fmt.Println(writeError)

				return c.JSON(http.StatusInternalServerError, writeError)
			}
		}

		fmt.Println(fmt.Sprintf("added files to directory %s", dirName))

		// create a merkle root
		merkleRoot := merkletree.GenerateMerkleTreeRoot(utils.ExtractHashesFromFileDirectory(fileDirectory))

		fmt.Println(fmt.Sprintf("created merkle root %s", merkleRoot))

		// if we have an empty merkle root, remove the new directory
		if merkleRoot == "" {
			// clean up
			utils.RemoveDir(dirName)

			fmt.Println("empty merkle root")

			return c.NoContent(http.StatusBadRequest)
		}

		merkleRootDirName := fmt.Sprintf("%s/%s", constants.RootFileDirectory, merkleRoot)

		// check if the merkle root directory exists
		if utils.DoesDirExist(merkleRootDirName) {
			// remove the newly created directory, we don't need it
			utils.RemoveDir(dirName)

			// return the previous directory details
			return c.JSON(http.StatusOK, types.FilesUploadResponse{
				Directory: fileDirectory,
				Root:      merkleRoot,
			})
		}

		// create a file directory json in the directory
		err = utils.WriteFileDirectoryJSONToStorage(dirName, fileDirectory)
		if err != nil {
			// clean up
			utils.RemoveDir(dirName)

			writeError = errors.NewWriteError(fmt.Sprintf("failed to create file directory json at %s", dirName), err)

			fmt.Println(writeError)

			return c.JSON(http.StatusInternalServerError, writeError)
		}

		fmt.Println(fmt.Sprintf("created file directory json at %s, with %d entries", dirName, len(fileDirectory)))

		// rename the directory to the merkle root
		err = os.Rename(dirName, merkleRootDirName)
		if err != nil {
			// clean up
			utils.RemoveDir(dirName)

			writeError = errors.NewWriteError(fmt.Sprintf("failed to rename directory %s to %s", dirName, merkleRootDirName), err)

			fmt.Println(writeError)

			return c.JSON(http.StatusInternalServerError, writeError)
		}

		// finally return the merkle root and the directory
		return c.JSON(http.StatusOK, types.FilesUploadResponse{
			Directory: fileDirectory,
			Root:      merkleRoot,
		})
	}
}
