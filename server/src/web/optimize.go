package web

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path"
	"path/filepath"

	"github.com/akmittal/bimg"
	findImages "github.com/akmittal/optimg/server/src/findimages"
	"github.com/akmittal/optimg/server/src/image"
)

type Transformation struct {
	Operations []image.Operation
	SourcePath string
	TargetPath string
}

func Optimize(rw http.ResponseWriter, req *http.Request) {

	var transformations Transformation
	err := json.NewDecoder(req.Body).Decode(&transformations)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}
	// check if path exists
	if info, err := os.Stat(transformations.SourcePath); os.IsNotExist(err) || !info.IsDir() {
		http.Error(rw, "Path not valid", http.StatusBadRequest)
		return
	}
	processAllDirectries(transformations)

}

func GetAllFilesAtPath(sourcePath string) (map[string][]os.FileInfo, error) {
	var result map[string][]os.FileInfo = make(map[string][]os.FileInfo)

	imgs, err := findImages.GetImageFiles(sourcePath)
	if err != nil {
		return nil, err
	}
	result[sourcePath] = imgs
	dirs, err := findImages.GetSubDirectories(sourcePath)

	if err != nil {
		return nil, err
	}
	for _, dir := range dirs {
		dirPath := path.Join(sourcePath, dir.Name())
		images, _ := GetAllFilesAtPath(dirPath)
		for k, v := range images {
			result[k] = v
		}
	}

	return result, nil
}

func processAllDirectries(transformations Transformation) {

	imgMapping, err := GetAllFilesAtPath(transformations.SourcePath)
	for sourcepath, _ := range imgMapping {

		fmt.Println(sourcepath)

	}

	if err != nil {
		fmt.Print(err)
	}
	for sourcepath, images := range imgMapping {
		for _, file := range images {
			path.Join(sourcepath, file.Name())
		}
	}

	for _, operation := range transformations.Operations {
		for sourcepath, images := range imgMapping {
			for _, file := range images {
				targetPath, err := filepath.Rel(transformations.SourcePath, sourcepath)
				targetPath = path.Join(transformations.TargetPath, targetPath)
				if err != nil {
					fmt.Println(err)
				}

				Convert(sourcepath, operation, file, targetPath)

			}
		}
	}

}

func Convert(imageSrc string, operation image.Operation, file os.FileInfo, targetPath string) {
	imagepath := path.Join(imageSrc, file.Name())

	img, err := image.FromPath(imagepath)
	if err != nil {
		// http.Error(rw, imagepath+err.Error(), http.StatusBadRequest)
		fmt.Println("Error", err)
		return
	}
	data, err := img.Convert(operation)
	if err != nil {
		// http.Error(rw, imagepath+err.Error(), http.StatusBadRequest)
		fmt.Println("Error2", err)
		return
	}
	if _, err := os.Stat(targetPath); os.IsNotExist(err) {
		fmt.Println("MAKING", targetPath)
		os.MkdirAll(targetPath, os.ModePerm)
	}
	var extension = filepath.Ext(file.Name())
	var newName = file.Name()[0:len(file.Name())-len(extension)] + image.GetExtensionForFormat(operation.Format)
	targetimagepath := path.Join(targetPath, newName)
	err = bimg.Write(targetimagepath, data)
	if err != nil {
		// http.Error(rw, imagepath+err.Error(), http.StatusBadRequest)
		fmt.Println("Error3", err)
		return
	}

}
