package main

import (
	"crypto/md5"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func md5sum(filePath string) string {
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

func sha256sum(filePath string) string {
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
	hashInBytes := hash.Sum(nil)[:32]
	returnSHA256String = hex.EncodeToString(hashInBytes)
	return returnSHA256String
}

func sha512sum(filePath string) string {
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
	md5Cmd := flag.NewFlagSet("md5", flag.ExitOnError)
	sha256Cmd := flag.NewFlagSet("sha256", flag.ExitOnError)
	sha512Cmd := flag.NewFlagSet("sha512", flag.ExitOnError)

	//flags
	switch os.Args[1] {
	case "md5":
		md5Cmd.Parse(os.Args[2:])
		hash := md5sum(path)
		fmt.Printf("md5 for %s:\t\t%s\n", path, hash)
	case "sha256":
		sha256Cmd.Parse(os.Args[2:])
		hash := sha256sum(path)
		fmt.Printf("sha256 for %s:\t%s\n", path, hash)
	case "sha512":
		sha512Cmd.Parse(os.Args[2:])
		hash := sha512sum(path)
		fmt.Printf("sha512 for %s:\t%s\n", path, hash)
	default:
		fmt.Println("expected 'md5', 'sha256','sha512' flags")
		os.Exit(1)
	}
	return nil
}

func main() {
	if len(os.Args) > 1 {
		flag.Parse()
		root := flag.Arg(1)
		filepath.Walk(root, visit)
	} else {
		fmt.Println("No arguments given\n\n Usage: ", os.Args[0], " <flags> <path>")
	}
}
