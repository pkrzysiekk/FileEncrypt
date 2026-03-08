package utils

import "path/filepath"

func TrimExtension(filename string) string {
	extension := filepath.Ext(filename)
	return filename[:len(filename)-len(extension)]
}
