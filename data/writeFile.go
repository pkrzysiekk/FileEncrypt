package data

import (
	"FileEncrypt/utils"
	"log"
	"os"
)

func WriteFile(data []byte, filename string) {
	filenameToSave := utils.TrimExtension(filename) + ".enc"
	err := os.WriteFile(filenameToSave, data, 666)
	if err != nil {
		log.Fatal(err)
	}
}
