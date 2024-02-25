package routes

import (
	"aether/internal/constants"
	"aether/internal/errors"
	"aether/internal/merkletree"
	"aether/internal/types"
	"aether/internal/utils"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"sort"
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
			sort.Slice(hashes, func(i int, j int) bool {
				return hashes[i] < hashes[j]
			})

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
		var files []*types.FileReadData
		var readError *errors.ReadError
		var writeError *errors.WriteError

		form, err := c.MultipartForm()
		if err != nil {
			fmt.Println(err)

			return c.JSON(http.StatusInternalServerError, err)
		}

		formFiles := form.File["file"]

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

			// add files so they can be written to storage when we have the merkle root
			files = append(files, &types.FileReadData{
				File: uploadFile,
				Name: fileHeader.Filename,
			})

			// add to the directory
			fileDirectory = append(fileDirectory, &types.FileDirectoryItem{
				Hash: fileHash,
				Name: fileHeader.Filename,
			})
		}

		// sort the files by hash
		sort.Slice(fileDirectory, func(i int, j int) bool {
			return fileDirectory[i].Hash < fileDirectory[j].Hash
		})

		// create a merkle root
		merkleRoot := merkletree.GenerateMerkleTreeRoot(utils.ExtractHashesFromFileDirectory(fileDirectory))

		fmt.Println(fmt.Sprintf("created merkle root %s", merkleRoot))

		// if we have an empty merkle root, no files were uploaded
		if merkleRoot == "" {
			fmt.Println("empty merkle root there was no files added to directory")

			return c.NoContent(http.StatusBadRequest)
		}

		// create a subdirectory to store the files using the merkle root hash as the directory name
		dirName := fmt.Sprintf("%s/%s", constants.RootFileDirectory, merkleRoot)
		err = utils.CreateDir(dirName)
		if err != nil {
			writeError = errors.NewWriteError(fmt.Sprintf("failed create the %s directory", dirName), err)

			fmt.Println(writeError)

			return c.JSON(http.StatusInternalServerError, writeError)
		}

		// add the files to the directory
		err = utils.SaveFilesToDir(dirName, files)
		if err != nil {
			writeError = errors.NewWriteError(fmt.Sprintf("failed write files to %s", dirName), err)

			fmt.Println(writeError)

			return c.JSON(http.StatusInternalServerError, writeError)
		}

		fmt.Println(fmt.Sprintf("added files to directory %s", dirName))

		// create a file directory json in the directory
		err = utils.WriteFileDirectoryJSONToStorage(dirName, fileDirectory)
		if err != nil {
			fmt.Println(fmt.Sprintf("failed to create file directory json at %s", dirName))
		}

		fmt.Println(fmt.Sprintf("created file directory json at %s, with %d entries", dirName, len(fileDirectory)))

		// finally return the merkle root and the directory
		return c.JSON(http.StatusOK, types.FilesUploadResponse{
			Directory: fileDirectory,
			Root:      merkleRoot,
		})
	}
}
