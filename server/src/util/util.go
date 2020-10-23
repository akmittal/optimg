package util

import "path/filepath"

func GetPathWithoutExtension(path string) string {
	extension := filepath.Ext(path)
	return path[0 : len(path)-len(extension)]
}
