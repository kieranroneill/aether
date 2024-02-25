package utils

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
)

// CreateRandomSha256Hash Creates a hex encoded SHA-256 hash of 32 random bytes
func CreateRandomSha256Hash() (string, error) {
	bytes := make([]byte, 32)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", sha256.Sum256(bytes)), nil
}

// HashFile Creates a hex encoded SHA-256 hash of the file
func HashFile(f io.Reader, filename string) (string, error) {
	h := sha256.New()

	if _, err := io.Copy(h, f); err != nil {
		return "", err
	}

	return hex.EncodeToString(h.Sum(nil)), nil
}
