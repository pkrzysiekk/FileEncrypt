package data

import (
	"log"
	"os"
)

func ReadFile(filename string) []byte {
	file, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal("Failed to read the file")
	}
	return file
}
