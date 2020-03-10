package hash

import (
	"crypto/sha256"
	"encoding/hex"
	"io"
	"os"
)

func Sha256sum(filePath string) string {
	var returnSHA256String string
	file, err := os.Open(filePath)
	if err != nil {
		return returnSHA256String
	}
	defer file.Close()
	hash := sha256.New()
	if _, err := io.Copy(hash, file); err != nil {
		return returnSHA256String
	}
	hashInBytes := hash.Sum(nil)[:64]
	returnSHA256String = hex.EncodeToString(hashInBytes)
	return returnSHA256String
}