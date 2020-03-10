package hash

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"os"
)

func Md5sum(filePath string) (string) {
	var returnMD5String string
	file, err := os.Open(filePath)
	if err != nil {
		return returnMD5String
	}
	defer file.Close()
	hash := md5.New()
	if _, err := io.Copy(hash, file); err != nil {
		return returnMD5String
	}
	hashInBytes := hash.Sum(nil)[:16]
	returnMD5String = hex.EncodeToString(hashInBytes)
	return returnMD5String
}