package hash

import (
	"crypto/sha512"
	"encoding/hex"
	"io"
	"os"
)

func Sha512sum(filePath string) (string, error) {
	var returnSHA512String string
	file, err := os.Open(filePath)
	if err != nil {
		return returnSHA512String, err
	}
	defer file.Close()
	hash := sha512.New()
	if _, err := io.Copy(hash, file); err != nil {
		return returnSHA512String, err
	}
	hashInBytes := hash.Sum(nil)[:64]
	returnSHA512String = hex.EncodeToString(hashInBytes)
	return returnSHA512String, nil
}
