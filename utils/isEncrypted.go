package utils

import (
	"path/filepath"
)

const encryptedFileExt = ".enc"

func IsEncrypted(filename string) bool {
	extension := filepath.Ext(filename)
	return extension == encryptedFileExt
}
