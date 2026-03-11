package utils

import (
	"log"
	"os"
	"path"
)

func GetAbsolutePath(filename string) string {
	currentDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	absPath := path.Join(currentDir, filename)
	return absPath
}
