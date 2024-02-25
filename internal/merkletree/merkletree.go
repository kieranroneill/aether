package merkletree

import (
	"aether/internal/constants"
	"aether/internal/types"
	"crypto/sha256"
	"encoding/hex"
	"math"
)

/**
  private functions
*/

func buildMerkleTree(hashes []string, tree [][]string) [][]string {
	var newTreeLevel []string

	// if the hashes only contain one hash, we are at the root, so return the tree as-is
	if len(hashes) == 1 {
		return tree
	}

	// make sure there is an even number of hashes
	preparedHashes := prepareHashes(hashes)

	// for each left and right hash, concatenate and make a new hash and add it to the new level of the tree
	for i := 0; i < len(preparedHashes); i += 2 {
		newTreeLevel = append(newTreeLevel, createHashPair(preparedHashes[i], preparedHashes[i+1]))
	}

	tree = append(tree, newTreeLevel)

	return buildMerkleTree(newTreeLevel, tree)
}

func createHashPair(leftHash string, rightHash string) string {
	hashPair := leftHash + rightHash
	hash := sha256.Sum256([]byte(hashPair))

	return hex.EncodeToString(hash[:])
}

// findLeafIndexInTreeLevel Simply gets the index of the leaf in the level. If the leaf doesn't exist in the level, -1 is returned.
func findLeafIndexInTreeLevel(leaf string, level []string) int {
	for index, value := range level {
		if value == leaf {
			return index
		}
	}

	// if the hash is not in the level
	return -1
}

// determineLeafPosition Convenience function that determines if a leaf index is a left or right leaf. Even indexes are
// left leafs whereas odd indexed equate to a right leaf.
func determineLeafPosition(index int) int {
	if index%2 == 0 { // if the index is even, this will be a left node, otherwise (odd) will be a right leaf
		return constants.Left
	}

	return constants.Right
}

// prepareHashes A Merkle Tree must have an even number of nodes, if the number of hashes is odd, we append the last hash to the end
// to make it even
func prepareHashes(hashes []string) []string {
	if len(hashes)%2 != 0 {
		return append(hashes, hashes[len(hashes)-1])
	}

	return hashes
}

// generateMerkleTree Generates a Merkle Tree from a list of hashes. This function takes several steps:
// 1. Ensures the list of hashes is even by duplicating the last hash and appending it to the end if it is an odd length
// 2. Concatenates each pair and hashes this pair to create a new parent node.
// 3. Recursively repeats the above steps until a root node (a level with one node) is achieved.
func generateMerkleTree(hashes []string) [][]string {
	// for empty hashes, return nil, we cannot make a merkle tree with no hashes
	if len(hashes) <= 0 {
		return nil
	}

	// start the tree with the hashes at index 0 - the bottom of the tree
	tree := [][]string{hashes}

	// recursively build the merkle tree
	return buildMerkleTree(hashes, tree)
}

/**
  public functions
*/

func GenerateMerkleTreeProof(hash string, hashes []string) []*types.MerkleTreeProofItem {
	var siblingIndex int

	if len(hashes) <= 0 {
		return nil
	}

	// generate the merkle tree
	tree := generateMerkleTree(hashes)

	// get the index of the leaf
	leafIndex := findLeafIndexInTreeLevel(hash, tree[0])

	// if the leaf is not in the tree, the proof is invalid
	if leafIndex < 0 {
		return nil
	}

	// add the current hash to the proof
	proof := []*types.MerkleTreeProofItem{{
		Hash:     hash,
		Position: determineLeafPosition(leafIndex),
	}}

	// traverse the tree adding the necessary nodes to the proof
	for i := 0; i < len(tree)-1; i++ {
		// find out the position of this leaf
		leafPosition := determineLeafPosition(leafIndex)

		// get the index of the leaf's sibling, for left leafs, the sibling is to the right of the slice (index + 1), or right leafs, the sibling is to the left of the slice (index -1)
		if leafPosition == constants.Left {
			siblingIndex = leafIndex + 1
		} else {
			siblingIndex = leafIndex - 1
		}

		// add the sibling to the proof
		proof = append(proof, &types.MerkleTreeProofItem{
			Hash:     tree[i][siblingIndex],
			Position: determineLeafPosition(siblingIndex),
		})

		leafIndex = int(math.Floor(float64(leafIndex / 2)))
	}

	return proof
}

// GenerateMerkleTreeRoot Generates a Merkle Tree from a list of hashes and simply returns the root.
func GenerateMerkleTreeRoot(hashes []string) string {
	// if there are no hashes, return an empty string
	if len(hashes) <= 0 {
		return ""
	}

	// generate the merkle tree
	tree := generateMerkleTree(hashes)

	// the root not will be the last element
	rootNode := tree[len(tree)-1]

	// if there is more than one node in the last element something went wrong
	if len(rootNode) > 1 {
		return ""
	}

	return rootNode[0]
}

func VerifyMerkleTreeProof(root string, proof []*types.MerkleTreeProofItem) bool {
	var leafHash string

	// if the merkle tree proof is empty, the root is not in there :P
	if len(proof) <= 0 {
		return false
	}

	for index, value := range proof {
		// if we are at the first element just use the first hash
		if index == 0 {
			leafHash = value.Hash
			continue
		}

		// if the next leaf is a left leaf, create the parent
		if value.Position == constants.Left {
			leafHash = createHashPair(value.Hash, leafHash)
			continue
		}

		// if the next leaf is a right leaf, create the parent
		leafHash = createHashPair(leafHash, value.Hash)
	}

	// check if the traversed hash matches the supplied root
	return leafHash == root
}
