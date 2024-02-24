package types

type VersionsResponse struct {
	Environment string `json:"environment"`
	Name        string `json:"name"`
	Version     string `json:"version"`
}
