package main

import (
	"crypto/sha512"
	"crypto/md5"
	"encoding/hex"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func md5sum(filePath string) (string) {
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

func sha512sum(filePath string) (string) {
	var returnSHA512String string
	file, err := os.Open(filePath)
    if err != nil {
        return returnSHA512String
    }
    defer file.Close()
    hash := sha512.New()
    if _, err := io.Copy(hash, file); err != nil {
    	return returnSHA512String
    }
    hashInBytes := hash.Sum(nil)[:64]
    returnSHA512String = hex.EncodeToString(hashInBytes)
    return returnSHA512String
}

func visit(path string, f os.FileInfo, err error) error {
	sha512Cmd := flag.NewFlagSet("sha512", flag.ExitOnError)
	md5Cmd := flag.NewFlagSet("md5Cmd", flag.ExitOnError)

	//flags
	switch os.Args[1] {
		case "sha512":
			sha512Cmd.Parse(os.Args[2:])
			hash := sha512sum(path)
			fmt.Printf("sha512 for %s:\t%s\n",path, hash)
		case "md5":
			md5Cmd.Parse(os.Args[2:])
			hash := md5sum(path)
			fmt.Printf("md5 for %s:\t%s\n", path, hash)
		default:
			fmt.Println("expected 'md5' or 'sha512' flags\n")
			os.Exit(1)
	}
	return nil
}

func main() {
	flag.Parse()
	root := flag.Arg(1)
	err := filepath.Walk(root, visit)
	hash := md5sum(os.Args[2])
	if err != nil {
		fmt.Println("hash for %s\t%s\n", hash)
	}
}