package image

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/h2non/bimg"
)

type ImageType bimg.ImageType

const (
	UNKNOWN ImageType = iota
	// JPEG represents the JPEG image type.
	JPEG
	// WEBP represents the WEBP image type.
	WEBP
	// PNG represents the PNG image type.
	PNG
	// TIFF represents the TIFF image type.
	TIFF
	// GIF represents the GIF image type.
	GIF
	// PDF represents the PDF type.
	PDF
	// SVG represents the SVG image type.
	SVG
	// MAGICK represents the libmagick compatible genetic image type.
	MAGICK
	// HEIF represents the HEIC/HEIF/HVEC image type
	HEIF
	// AVIF NEW FILE FORMAT
	AVIF
)

var ImageTypes = map[ImageType]string{
	UNKNOWN: "",
	// JPEG represents the JPEG image type.
	JPEG: ".jpg",
	// WEBP represents the WEBP image type.
	WEBP: ".webp",
	// PNG represents the PNG image type.
	PNG: ".png",
	// TIFF represents the TIFF image type.
	TIFF: ".tiff",
	// GIF represents the GIF image type.
	GIF: ".giff",
	// PDF represents the PDF type.
	PDF: ".pdf",
	// SVG represents the SVG image type.
	SVG: ".svg",
	// MAGICK represents the libmagick compatible genetic image type.
	MAGICK: ".magick",
	// HEIF represents the HEIC/HEIF/HVEC image type
	HEIF: ".heif",
	// AVIF NEW FILE FORMAT
	AVIF: ".avif",
}

type Operation struct {
	Format  ImageType
	Scale   int
	Quality int
	Height  int
	Width   int
}

type Image struct {
	bimg.Image
	Path string
}

func GetExtensionForFormat(imageType ImageType) string {
	return ImageTypes[imageType]

}

func (img *Image) Convert(operation Operation) ([]byte, error) {
	if operation.Format == AVIF {
		return img.ConvertAVIF(operation.Quality)
	}

	options := bimg.Options{
		Height:  operation.Height,
		Width:   operation.Width,
		Quality: operation.Quality,
		Zoom:    operation.Scale,
		Type:    bimg.ImageType(operation.Format),
	}
	return img.Process(options)

}

func (img *Image) ConvertAVIF(quality int) ([]byte, error) {
	outPath := strings.Replace(img.Path, filepath.Ext(img.Path), ".avif", 0)
	return RunCavif(img.Path, outPath, quality)
}

func RunCavif(inputPath string, outputPath string, quality int) ([]byte, error) {
	qualityFlag := fmt.Sprintf("--quality=%d", quality)
	cmd := exec.Command("cavif", qualityFlag, inputPath, outputPath, "--overwrite")
	stderr, err := cmd.StderrPipe()
	if err != nil {
		log.Fatal(err)
	}
	err = cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	slurp, _ := ioutil.ReadAll(stderr)
	fmt.Printf("%s\n", slurp)

	log.Printf("Waiting for command to finish...")
	err = cmd.Wait()

	if err != nil {
		log.Printf("Command finished with error: %v", err)
		return nil, errors.New(string(slurp))
	}
	return ioutil.ReadFile(outputPath)

}

func FromPath(path string) (Image, error) {
	data, err := bimg.Read(path)
	if err != nil {
		return Image{}, err
	}
	img := bimg.NewImage(data)
	return Image{*img, path}, nil
}
