package main

import (
	"crypto/md5"
	"encoding/hex"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func md5sum(filePath string) (string, error) {
	var returnMD5String string
	file, err := os.Open(filePath)
	if err != nil {
		return returnMD5String, err
	}
	defer file.Close()
	hash := md5.New()
	if _, err := io.Copy(hash, file); err != nil {
		return returnMD5String, err
	}
	hashInBytes := hash.Sum(nil)[:16]
	returnMD5String = hex.EncodeToString(hashInBytes)
	return returnMD5String, nil
}

func visit(path string, f os.FileInfo, err error) error {
	hash, err := md5sum(path)
	fmt.Printf("md5sum for %s:\t%s\n", path, hash)
	return nil
}

func main() {
	//filecrawl
	flag.Parse()
	root := flag.Arg(0)
	err := filepath.Walk(root, visit)
	//md5sum
	hash, err := md5sum(os.Args[1])
	if err == nil {
		fmt.Printf("md5sum for %s\n", hash)
	}
}
