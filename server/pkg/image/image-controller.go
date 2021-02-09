package image

import (
	"encoding/json"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/go-chi/chi"
	"gorm.io/gorm"
)

type Images struct {
	Image    Image   `json:"image"`
	Varients []Image `json:"varients"`
}

var optimalformats = [2]string{".avif", ".webp"}

func ImageServer(basePath string) http.HandlerFunc {
	return func(rw http.ResponseWriter, req *http.Request) {
		// query := req.URL.Query()

		imagePath := filepath.Join(basePath, chi.URLParam(req, "path"))

		supportedFormats := getSupportedFormats(req)
		imagePath = checkIfBetterFile(imagePath, supportedFormats)
		extension := filepath.Ext(imagePath)

		if extension == ".avif" {
			rw.Header().Set("content-type", "image/avif")
		}

		http.ServeFile(rw, req, imagePath)
	}
}
func getSupportedFormats(req *http.Request) []string {
	result := []string{}
	accepts := req.Header.Get("Accept")
	if strings.Contains(accepts, "image/avif") {
		result = append(result, ".avif")
	}
	if strings.Contains(accepts, "image/webp") {
		result = append(result, ".webp")
	}
	return result
}

func checkIfBetterFile(path string, supportedFormats []string) string {
	extension := filepath.Ext(path)
	filePathWithoutExt := path[0 : len(path)-len(extension)]
	for _, format := range supportedFormats {
		optimalPath := filePathWithoutExt + format

		if _, err := os.Stat(optimalPath); os.IsNotExist(err) {

			continue
		} else {
			return optimalPath
		}
	}
	return path

}

func ImageController(db *gorm.DB) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		path := r.URL.Query().Get("path")
		filename := r.URL.Query().Get("name")
		var image Image
		var varients []Image
		var images Images
		if path == "/" {
			path = "."
		} else if strings.HasPrefix(path, "/") {
			path = path[1:]
		}
		db.First(&image, "path = ? AND name = ?", path, filename)
		images.Image = image
		db.Where("parent_id = ? ", image.ID).Find(&varients)
		images.Varients = varients
		json.NewEncoder(rw).Encode(images)
	}
}
