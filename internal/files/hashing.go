package files

import (
	"aether/internal/errors"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"mime/multipart"
)

// HashFile Creates a hex encoded SHA-256 hash of the file
func HashFile(f multipart.File, filename string) (string, *errors.HashError) {
	h := sha256.New()

	if _, err := io.Copy(h, f); err != nil {
		return "", errors.NewHashError(fmt.Sprintf("unable to hash file %s", filename), err)
	}

	return hex.EncodeToString(h.Sum(nil)), nil
}