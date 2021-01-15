package image

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/h2non/filetype"
)

func GetImagesByPath(path string) ([]string, error) {
	if file, err := os.Stat(path); os.IsNotExist(err) || !file.IsDir() {
		return nil, err
	}
	result := []string{}
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}

	for _, f := range files {
		absPath := filepath.Join(path, f.Name())
		buf, _ := ioutil.ReadFile(absPath)
		if !f.IsDir() && filetype.IsImage(buf) {
			result = append(result, f.Name())

		}
	}
	return result, err
}

func GetSubDirectoriesByPath(path string) ([]string, error) {
	if file, err := os.Stat(path); os.IsNotExist(err) || !file.IsDir() {
		return nil, err
	}
	result := []string{}
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}

	for _, f := range files {

		if f.IsDir() {
			result = append(result, f.Name())

		}
	}
	return result, err
}
