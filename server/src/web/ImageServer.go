package web

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/akmittal/optimg/server/src/util"
)

const basePath = "/Users/amittal/images"

var optimalformats = [2]string{".avif", ".webp"}

func ImageServer(rw http.ResponseWriter, req *http.Request) {
	query := req.URL.Query()

	format := query["format"]
	width := query["w"]
	height := query["h"]
	quality := query["q"]
	scale := query["s"]
	fmt.Println(format, width, height, quality, scale, req.URL.Path)
	imagePath := filepath.Join(basePath, req.URL.Path)
	supportedFormats := getSupportedFormats(req)
	imagePath = checkIfBetterFile(imagePath, supportedFormats)
	extension := filepath.Ext(imagePath)

	if extension == ".avif" {
		rw.Header().Set("content-type", "image/avif")
	}

	http.ServeFile(rw, req, imagePath)
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
	filePathWithoutExt := util.GetPathWithoutExtension(path)

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
