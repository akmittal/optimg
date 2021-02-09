package image

import (
	"os"
	"path/filepath"

	"github.com/h2non/bimg"
	"gorm.io/gorm"
)

type Image struct {
	*gorm.Model
	*bimg.Image       `gorm:"-"`
	Path              string `json:"path"`
	Name              string `json:"name"`
	Size              int64  `json:"size"`
	ParentID          uint   `json:"parentID"`
	OperationRef      uint
	TransformationRef uint
}

func Get(path string, sourcePath string) (*Image, error) {
	imageBytes, err := bimg.Read(path)
	if err != nil {
		return nil, err
	}
	stats, err := os.Stat(path)
	if err != nil {
		return nil, err
	}
	relativeFilePath, _ := filepath.Rel(sourcePath, filepath.Dir(path))
	return &Image{
		Image: bimg.NewImage(imageBytes),
		Path:  relativeFilePath,
		Name:  filepath.Base(path),
		Size:  stats.Size(),
	}, nil
}
