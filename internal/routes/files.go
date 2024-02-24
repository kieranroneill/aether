package routes

import (
	"aether/internal/errors"
	internalfiles "aether/internal/files"
	"aether/internal/types"
	"fmt"
	"github.com/labstack/echo/v4"
	"io"
	"net/http"
	"os"
)

func NewFilesUploadRoute() echo.HandlerFunc {
	return func(c echo.Context) error {
		var readError *errors.ReadError
		var writeError *errors.WriteError

		rootFilesDirectory, writeError := internalfiles.GetRootFilesDirectory()
		if writeError != nil {
			fmt.Println(writeError)

			return c.JSON(http.StatusInternalServerError, writeError)
		}

		form, err := c.MultipartForm()
		if err != nil {
			fmt.Println(err)

			return c.JSON(http.StatusInternalServerError, err)
		}
		formFiles := form.File["files"]

		for _, fileHeader := range formFiles {
			// read file
			fileFromUpload, err := fileHeader.Open()
			if err != nil {
				readError = errors.NewReadError(fmt.Sprintf("failed to read file %s", fileHeader.Filename), err)

				fmt.Println(readError)

				return c.JSON(http.StatusInternalServerError, readError)
			}
			defer fileFromUpload.Close()

			// get a hash of the file
			fileHash, hashError := internalfiles.HashFile(fileFromUpload, fileHeader.Filename)
			if hashError != nil {
				fmt.Println(hashError)

				return c.JSON(http.StatusInternalServerError, hashError)
			}

			fmt.Println(fmt.Sprintf("file %s has a hash of %s", fileHeader.Filename, fileHash))

			// create the file with the hash as its file name
			fileOnStorage, err := os.Create(fmt.Sprintf("%s/%s", rootFilesDirectory, fileHeader.Filename))
			if err != nil {
				writeError = errors.NewWriteError(fmt.Sprintf("failed to write file %s", fileHeader.Filename), err)

				fmt.Println(writeError)

				return c.JSON(http.StatusInternalServerError, writeError)
			}
			defer fileOnStorage.Close()

			// copy file contents to storage
			if _, err = io.Copy(fileOnStorage, fileFromUpload); err != nil {
				writeError = errors.NewWriteError(fmt.Sprintf("failed to write contents of file %s", fileHeader.Filename), err)

				fmt.Println(writeError)

				return c.JSON(http.StatusInternalServerError, writeError)
			}
		}

		return c.JSON(http.StatusOK, types.FilesUploadResponse{
			Root: "",
		})
	}
}
