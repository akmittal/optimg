package web

import (
	"encoding/json"
	"fmt"

	"net/http"
	"os"
	"path/filepath"
	"strings"

	img "github.com/akmittal/optimg/server/src/image"
)

func ImageHandler(rw http.ResponseWriter, req *http.Request) {
	var result GalleryImages
	query := req.URL.Query()
	path := query.Get("path")

	imagePath := filepath.Join(sourcePATH, path)
	image, err := getGalleryImage(imagePath, sourcePATH)
	if err != nil {
		fmt.Println(err)
	}

	result.Main = image
	result.Varients = getImageVarients(image)
	json.NewEncoder(rw).Encode(result)
}

func getGalleryImage(path string, rootPath string) (GalleryImage, error) {
	fmt.Println(path, "path")
	file, err := os.Stat(path)
	if err != nil {

		return GalleryImage{}, err
	}
	image := GalleryImage{}
	bimgImage, err := img.FromPath(path)
	if err != nil {

		return GalleryImage{}, err
	}
	meta, err := bimgImage.Metadata()
	if err != nil {

		return GalleryImage{}, err
	}
	image.Width = meta.Size.Height
	image.Height = meta.Size.Height
	image.Path = strings.Replace(path, rootPath, "", 1)
	image.Size = file.Size()
	image.Modified = file.ModTime()
	return image, nil
}
