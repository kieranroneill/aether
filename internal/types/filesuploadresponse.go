package types

type FilesUploadResponse struct {
	Directory []*FileDirectoryItem `json:"directory"`
	Root      string               `json:"root"`
}
