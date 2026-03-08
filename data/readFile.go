package data

import (
	"fmt"
	"log"
	"os"
)

func ReadFile(filename string) []byte {
	fmt.Println(filename)
	file, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	return file
}
