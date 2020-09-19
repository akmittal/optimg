package findImages

import (
	"io/ioutil"
	"os"
	"strings"
)

var (
	//ImageExtensions all image extensions
	ImageExtensions = []string{".jpg", ".JPG", ".jpeg", ".JPEG", ".png", ".PNG", ".webp", ".WEBP"}
)

// IsImage check if a given path is of image or not
func IsImage(path string) bool {
	// TODO: maybe change the logic to check mime type
	for i := 0; i < len(ImageExtensions); i++ {
		isImage := strings.HasSuffix(path, ImageExtensions[i])
		if isImage {
			return true
		}
	}
	return false

}

//GetImageFiles Get all image files from path
func GetImageFiles(path string) ([]os.FileInfo, error) {
	var images []os.FileInfo

	files, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}
	for _, file := range files {
		if IsImage(file.Name()) {
			images = append(images, file)
		}
	}
	return images, nil
}

func GetSubDirectories(path string) ([]os.FileInfo, error) {
	var subDirs []os.FileInfo
	files, err := ioutil.ReadDir(path)
	if err == nil {
		for _, file := range files {
			if file.IsDir() {
				subDirs = append(subDirs, file)
			}
		}
	}
	return subDirs, err
}
