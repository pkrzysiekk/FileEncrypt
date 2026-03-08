package data

import (
	"log"
	"os"
	"path/filepath"
)

func WriteFile(data []byte, filename string) {
	filenameToSave := trimExtension(filename) + ".enc"
	err := os.WriteFile(filenameToSave, data, 666)
	if err != nil {
		log.Fatal(err)
	}
}

func trimExtension(filename string) string {
	extension := filepath.Ext(filename)
	return filename[:len(filename)-len(extension)]
}
