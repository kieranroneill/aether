package types

type FileResponse struct {
	Hash  string                 `json:"hash"`
	Name  string                 `json:"name"`
	Proof []*MerkleTreeProofItem `json:"proof"`
}
