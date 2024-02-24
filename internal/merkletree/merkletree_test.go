package merkletree

import (
	"fmt"
	"testing"
)

var hashes []string

func TestGenerateMerkleTreeRootForEmptyHashes(t *testing.T) {
	const expectedMerkleRoot string = ""

	hashes = []string{}
	actualMerkleRoot := GenerateMerkleTreeRoot(hashes)

	if expectedMerkleRoot != actualMerkleRoot {
		t.Errorf("expect result: %s, actual result: %s", expectedMerkleRoot, actualMerkleRoot)
	}
}

func TestGenerateMerkleTreeRootForHash(t *testing.T) {
	const expectedMerkleRoot string = "d9f893bed0c563e78a5b225dbfd642b3957a56879ee7afed605479025847af50"

	hashes = []string{
		"e9616a8f682133fe550840eedecb0492a8c209044b6644dc999738b64a6a11aa",
	}
	actualMerkleRoot := GenerateMerkleTreeRoot(hashes)

	if expectedMerkleRoot != actualMerkleRoot {
		t.Errorf("expect result: %s, actual result: %s", expectedMerkleRoot, actualMerkleRoot)
	}
}

func TestGenerateMerkleTreeRootForAPairOfHashes(t *testing.T) {
	const expectedMerkleRoot string = "e818f4f035a41a36a574e42dc6986e730b70b7c8473715c0e6171c3ee6e50f26"

	hashes = []string{
		"e9616a8f682133fe550840eedecb0492a8c209044b6644dc999738b64a6a11aa",
		"08e8378e98dd1b8c81992a113c73e3b50a42aa24f744f984adc3b5b28fc690ed",
	}
	actualMerkleRoot := GenerateMerkleTreeRoot(hashes)

	if expectedMerkleRoot != actualMerkleRoot {
		t.Errorf("expect result: %s, actual result: %s", expectedMerkleRoot, actualMerkleRoot)
	}
}

func TestGenerateMerkleTreeRootForOddNumberOfHashes(t *testing.T) {
	const expectedMerkleRoot string = "1ddc0cfe8640ab380b5dbb186b1e84011c3b6bcf7a79be2b9257fec31ca606c3"

	hashes = []string{
		"e9616a8f682133fe550840eedecb0492a8c209044b6644dc999738b64a6a11aa",
		"08e8378e98dd1b8c81992a113c73e3b50a42aa24f744f984adc3b5b28fc690ed",
		"f17a98f8dcb95ea1fb3a7016ef08301f3482eb89ec7d1e43164bdcf1cfac323e",
		"522b2aa04d3541d8d67d382d659c992314620d34807b9439090708c2519fb232",
		"59fa57badd1a68045672bf90360e8a7fb401709149878e837a91ac85ccf5031d",
		"79a61238173fb912c6a9e251081768b8f13bf80d81afdafe8572269bd352c58e",
		"66da398be63468af7d410cd03df2f0c6def65ef78542f91a3c02dc955c7c0205",
		"3a38a370a8fd060daf623e7985c55c994d1387a0a7b5ef2d740e89a9b3cb73d3",
		"d681f09b4e03ee9e887e30dc3fb7307df143ed7c213272f87192714ecbab2f63",
		"d90ac811fa0da57444d997a770e0d7fa296c6b7978f384ba82bdd70fa3f50776",
		"9291313a9f9f9cff6760d868726135a6af2a82d70cf549de65d33f4362230a98",
	}
	actualMerkleRoot := GenerateMerkleTreeRoot(hashes)

	if expectedMerkleRoot != actualMerkleRoot {
		t.Errorf("expect result: %s, actual result: %s", expectedMerkleRoot, actualMerkleRoot)
	}
}

func TestGenerateMerkleTreeRootForEvenNumberOfHashes(t *testing.T) {
	const expectedMerkleRoot string = "0c2fdf0d51ab5a8f9577ba8909c9815d73c4845fab332af9c671e98d8a3a3971"

	hashes = []string{
		"e9616a8f682133fe550840eedecb0492a8c209044b6644dc999738b64a6a11aa",
		"08e8378e98dd1b8c81992a113c73e3b50a42aa24f744f984adc3b5b28fc690ed",
		"f17a98f8dcb95ea1fb3a7016ef08301f3482eb89ec7d1e43164bdcf1cfac323e",
		"522b2aa04d3541d8d67d382d659c992314620d34807b9439090708c2519fb232",
		"59fa57badd1a68045672bf90360e8a7fb401709149878e837a91ac85ccf5031d",
		"79a61238173fb912c6a9e251081768b8f13bf80d81afdafe8572269bd352c58e",
		"66da398be63468af7d410cd03df2f0c6def65ef78542f91a3c02dc955c7c0205",
		"3a38a370a8fd060daf623e7985c55c994d1387a0a7b5ef2d740e89a9b3cb73d3",
		"d681f09b4e03ee9e887e30dc3fb7307df143ed7c213272f87192714ecbab2f63",
		"d90ac811fa0da57444d997a770e0d7fa296c6b7978f384ba82bdd70fa3f50776",
		"9291313a9f9f9cff6760d868726135a6af2a82d70cf549de65d33f4362230a98",
		"68e6cdf0cae7fb8eef39cc899c8882e34dd1727a2d08f2303811886949c539e6",
	}
	actualMerkleRoot := GenerateMerkleTreeRoot(hashes)

	fmt.Println(actualMerkleRoot)

	if expectedMerkleRoot != actualMerkleRoot {
		t.Errorf("expect result: %s, actual result: %s", expectedMerkleRoot, actualMerkleRoot)
	}
}
