package web

import (
	"encoding/json"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"time"

	"github.com/akmittal/optimg/server/src/util"
)

var targetExtensions = []string{".avif", ".webp", ".jpg", ".png"}

type GalleryImages struct {
	Main     GalleryImage
	Varients []GalleryImage
}

type GalleryImage struct {
	Path     string
	Height   int
	Width    int
	Size     int64
	Modified time.Time
}

const (
	sourcePATH = "/Users/amittal/projects/images"
	targetPath = "/Users/amittal/images"
)

func Gallery(rw http.ResponseWriter, req *http.Request) {
	var result []GalleryImages
	imgMapping, err := GetAllFilesAtPath(sourcePATH)
	query := req.URL.Query()
	page := query.Get("page")
	limit, err := strconv.Atoi(query.Get("limit"))
	totalResults := 0
	if err != nil {
		limit = 10
	}
	pageNo := 0
	if page != "" {
		pageNo, err = strconv.Atoi(page)
	}

	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}
	for path, imageList := range imgMapping {
		for _, file := range imageList {
			images := GalleryImages{}

			imagePath := filepath.Join(path, file.Name())
			image, err := getGalleryImage(imagePath, sourcePATH)
			if err != nil {
				http.Error(rw, err.Error(), http.StatusInternalServerError)
			}
			images.Main = image
			images.Varients, err = getImageVarients(image)
			if err != nil {
				http.Error(rw, err.Error(), http.StatusInternalServerError)
				return
			}
			result = append(result, images)
			totalResults = len(result)

		}
	}
	sort.Slice(result, func(i, j int) bool {
		return result[i].Main.Modified.Before(result[j].Main.Modified)
	})

	result = result[(pageNo)*limit : (pageNo+1)*limit]
	rw.Header().Set("Content-Type", "application/json")

	json.NewEncoder(rw).Encode(map[string]interface{}{"Data": result, "TotalPages": totalResults / limit})

}

func getImageVarients(image GalleryImage) ([]GalleryImage, error) {
	var result []GalleryImage

	targetFilename := filepath.Join(targetPath, image.Path)
	pathWithoutExt := util.GetPathWithoutExtension(targetFilename)
	for _, ext := range targetExtensions {
		newFilePath := pathWithoutExt + ext
		if _, err := os.Stat(newFilePath); os.IsNotExist(err) {
			continue
		} else {
			image, err := getGalleryImage(newFilePath, targetPath)
			if err != nil {
				return nil, err
			}

			result = append(result, image)
		}
	}
	return result, nil

}
