package merkletree

import (
	"aether/internal/constants"
	"crypto/sha256"
	"encoding/hex"
)

// determineLeafNodeDirectionFromHash Given the hash and the merkletree, determine whether the leaf node is a left or
// right node
func determineLeafNodePositionFromHash(merkletree [][]string, leafHash string, level int) int {
	for nodeIndex, value := range merkletree[level] { // we are only concerned with the bottom level
		if value == leafHash {
			if nodeIndex%2 == 0 { // if the index is even, this will be a left node, otherwise (odd) will be a right leaf
				return constants.Left
			}

			return constants.Right
		}
	}

	// if the hash is not in the level
	return -1
}

// prepareHashes A Merkle Tree must have an even number of nodes, if the number of hashes is odd, we append the last hash to the end
// to make it even
func prepareHashes(hashes []string) []string {
	if len(hashes)%2 != 0 {
		return append(hashes, hashes[len(hashes)-1])
	}

	return hashes
}

// GenerateMerkleTreeRoot Generates a Merkle Tree root from a list of hashes. This function takes several steps:
// 1. Ensures the list of hashes is even by duplicating the last hash and appending it to the end if it is an odd length
// 2. Concatenates each pair and hashes this pair to create a new parent node
// 3. Recursively repeats the above steps until a root node (a level with one node) is achieved. This root is returned.
func GenerateMerkleTreeRoot(hashes []string) string {
	var newTreeLevel []string

	if len(hashes) <= 0 {
		return ""
	}

	// make sure there is an even number of hashes
	preparedHashes := prepareHashes(hashes)

	// for each left and right hash, concatenate and make a new hash and add it to the new level of the tree
	for i := 0; i < len(preparedHashes); i += 2 {
		hashPair := preparedHashes[i] + preparedHashes[i+1]
		hash := sha256.Sum256([]byte(hashPair))
		newTreeLevel = append(newTreeLevel, hex.EncodeToString(hash[:]))
	}

	// if the new tree level has only one node, we have the root
	if len(newTreeLevel) == 1 {
		return newTreeLevel[0]
	}

	// if there is more than one level, we have more levels to go to get to the root
	return GenerateMerkleTreeRoot(newTreeLevel)
}
