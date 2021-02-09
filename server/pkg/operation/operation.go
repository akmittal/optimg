package operation

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/akmittal/optimg/server/pkg/image"
	"github.com/akmittal/optimg/server/pkg/util"
	"github.com/h2non/bimg"
	"gorm.io/gorm"
)

var (
	FormatMapping = map[string]bimg.ImageType{
		"jpg":  bimg.JPEG,
		"jpeg": bimg.JPEG,
		"avif": bimg.AVIF,
		"webp": bimg.WEBP,
	}
)

type Transformation struct {
	*gorm.Model
	Quality      int            `json:"quality"`
	Format       bimg.ImageType `json:"format"`
	OperationRef uint
	Images       []image.Image `gorm:"foreignKey:TransformationRef"`
}

type Operation struct {
	*gorm.Model
	SourcePath      string           `json:"sourcePath"`
	TargetPath      string           `json:"targetPath"`
	CopyUnknown     bool             `json:"copyUnknown"`
	Monitor         bool             `json:"monitor"`
	Transformations []Transformation `json:"transformations" gorm:"foreignKey:OperationRef"`
	Images          []image.Image    `gorm:"foreignKey:OperationRef"`
}

// Get Returns instance of Operation
func Get(sourcePath string, targetPath string, copyUnknown bool, monitor bool, transformations []Transformation) (Operation, error) {
	opr := Operation{
		SourcePath:      sourcePath,
		TargetPath:      targetPath,
		CopyUnknown:     copyUnknown,
		Monitor:         monitor,
		Transformations: transformations,
	}
	err := opr.Validate()

	return opr, err
}

func (o *Operation) Process(db *gorm.DB) error {
	tx := db.Create(o)

	images, err := image.GetImagesByPath(o.SourcePath)
	if err != nil {
		return err
	}

	err = processImages(images, *o, ".", db)
	tx.Commit()

	return err
}
func processImages(images []string, operation Operation, relPath string, db *gorm.DB) error {
	destDir := filepath.Join(operation.TargetPath, relPath)
	if _, err := os.Stat(destDir); os.IsNotExist(err) {
		os.MkdirAll(destDir, 0755)
	}
	for _, img := range images {
		absPath := filepath.Join(operation.SourcePath, relPath, img)
		sourceImage, err := image.Get(absPath, operation.SourcePath)
		sourceImage.OperationRef = operation.ID

		if err != nil {
			return err
		}
		db.Create(sourceImage)
		varientImages, err := PerformTransformations(sourceImage, db, &operation)
		db.Create(varientImages)

	}
	db.Commit()
	dirs, err := image.GetSubDirectoriesByPath(filepath.Join(operation.SourcePath, relPath))
	if err != nil {
		return err
	}

	for _, dir := range dirs {
		dirPath := filepath.Join(operation.SourcePath, relPath, dir)
		// targetPath := filepath.Join(targetPath, dir)
		images, err := image.GetImagesByPath(dirPath)
		if err != nil {
			return err
		}
		relPath := filepath.Join(relPath, dir)
		processImages(images, operation, relPath, db)
	}
	return nil
}

func (o *Operation) Validate() error {

	if stat, err := os.Stat(o.SourcePath); os.IsNotExist(err) || !stat.IsDir() {
		return &util.Error{Msg: "Sourcepath not exist", Field: "sourcePath"}
	}
	if path, _ := filepath.Rel(o.SourcePath, o.TargetPath); path == "." {

		return &util.Error{Msg: "Sourcepath and Target path cant be same", Field: "sourcePath"}
	}
	if path, _ := filepath.Rel(o.SourcePath, o.TargetPath); !strings.HasPrefix(path, "..") {

		return &util.Error{Msg: "Target path cant be subdirectory of Source", Field: "sourcePath"}
	}
	formats := make(map[bimg.ImageType]bool)
	for _, tranform := range o.Transformations {
		if formats[tranform.Format] {
			return &util.Error{Msg: "Formats must be unique", Field: "sourcePath"}
		}
		formats[tranform.Format] = true
	}
	return nil
}

func PerformTransformations(i *image.Image, db *gorm.DB, o *Operation) ([]image.Image, error) {
	var varientImages []image.Image
	for _, transformation := range o.Transformations {

		convertedImg, err := i.Process(bimg.Options{
			Quality: (transformation.Quality),
			Type:    transformation.Format,
		})
		targetFilepath := getTargetFilePath(i.Name, filepath.Join(o.TargetPath, i.Path), transformation.Format)

		err = bimg.Write(targetFilepath, convertedImg)
		if err != nil {
			return varientImages, err
		}
		targetImage, err := image.Get(targetFilepath, o.TargetPath)
		targetImage.ParentID = i.ID
		targetImage.OperationRef = o.ID
		varientImages = append(varientImages, *targetImage)
		targetImage.TransformationRef = transformation.ID

	}

	return varientImages, nil
}
func getTargetFilePath(fileName string, targetPath string, format bimg.ImageType) string {
	targetExtenstion := bimg.ImageTypes[format]
	targetFilename := strings.TrimSuffix(fileName, filepath.Ext(fileName)) + "." + targetExtenstion
	fmt.Println("path", filepath.Join(targetPath, targetFilename))
	return filepath.Join(targetPath, targetFilename)
}
