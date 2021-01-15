package operation

import (
	"path/filepath"
	"strings"

	"github.com/akmittal/optimg/server/pkg/image"
	"github.com/h2non/bimg"
	"gorm.io/gorm"
)

type Transformation struct {
	Quality int            `json:"quality"`
	Format  bimg.ImageType `json:"format"`
}

type Operation struct {
	*gorm.Model
	SourcePath      string           `json:"sourcePath"`
	TargetPath      string           `json:"targetPath"`
	CopyUnknown     bool             `json:"copyUnknown"`
	Monitor         bool             `json:"monitor"`
	Transformations []Transformation `json:"transformations"`
}

// Get Returns instance of Operation
func Get(sourcePath string, targetPath string, copyUnknown bool, monitor bool, transformations []Transformation) (Operation, error) {
	return Operation{
		SourcePath:      sourcePath,
		TargetPath:      targetPath,
		CopyUnknown:     copyUnknown,
		Monitor:         monitor,
		Transformations: transformations,
	}, nil
}

func (o *Operation) Process(db *gorm.DB) error {

	data, err := image.GetImagesByPath(o.SourcePath)
	for _, img := range data {
		absPath := filepath.Join(o.SourcePath, img)
		sourceImage, err := image.Get(absPath, o.SourcePath)
		if err != nil {
			return err
		}
		err = PerformTransformations(sourceImage, db, o.SourcePath, o.TargetPath, o.Transformations)

	}
	return err
}

func PerformTransformations(i *image.Image, db *gorm.DB, sourcePath string, targetPath string, transformations []Transformation) error {
	tx := db.Create(i)
	for _, transformation := range transformations {

		convertedImg, err := i.Process(bimg.Options{
			Quality: (transformation.Quality),
			Type:    transformation.Format,
		})
		targetFilepath := getTargetFilePath(i.Name, sourcePath, targetPath, transformation.Format)

		err = bimg.Write(targetFilepath, convertedImg)
		if err != nil {
			return err
		}
		targetImage, err := image.Get(targetFilepath, targetFilepath)
		targetImage.ParentID = i.ID
		tx = db.Create(&targetImage)
	}

	tx.Commit()
	return nil
}
func getTargetFilePath(fileName string, sourcePath string, targetPath string, format bimg.ImageType) string {
	targetExtenstion := bimg.ImageTypes[format]
	targetFilename := strings.TrimSuffix(fileName, filepath.Ext(fileName)) + "." + targetExtenstion
	return filepath.Join(targetPath, targetFilename)
}
