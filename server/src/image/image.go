package image

import (
	"fmt"

	"github.com/akmittal/bimg"
)

// type ImageType bimg.ImageType

// const (
// 	UNKNOWN ImageType = iota
// 	// JPEG represents the JPEG image type.
// 	JPEG
// 	// WEBP represents the WEBP image type.
// 	WEBP
// 	// PNG represents the PNG image type.
// 	PNG
// 	// TIFF represents the TIFF image type.
// 	TIFF
// 	// GIF represents the GIF image type.
// 	GIF
// 	// PDF represents the PDF type.
// 	PDF
// 	// SVG represents the SVG image type.
// 	SVG
// 	// MAGICK represents the libmagick compatible genetic image type.
// 	MAGICK
// 	// HEIF represents the HEIC/HEIF/HVEC image type
// 	HEIF
// 	// AVIF NEW FILE FORMAT
// 	AVIF
// )

// var ImageTypes = map[ImageType]string{
// 	UNKNOWN: "",
// 	// JPEG represents the JPEG image type.
// 	JPEG: ".jpg",
// 	// WEBP represents the WEBP image type.
// 	WEBP: ".webp",
// 	// PNG represents the PNG image type.
// 	PNG: ".png",
// 	// TIFF represents the TIFF image type.
// 	TIFF: ".tiff",
// 	// GIF represents the GIF image type.
// 	GIF: ".giff",
// 	// PDF represents the PDF type.
// 	PDF: ".pdf",
// 	// SVG represents the SVG image type.
// 	SVG: ".svg",
// 	// MAGICK represents the libmagick compatible genetic image type.
// 	MAGICK: ".magick",
// 	// HEIF represents the HEIC/HEIF/HVEC image type
// 	HEIF: ".heif",
// 	// AVIF NEW FILE FORMAT
// 	AVIF: ".avif",
// }

type Operation struct {
	Format  bimg.ImageType
	Scale   int
	Quality int
	Height  int
	Width   int
}

type Image struct {
	bimg.Image
	Path string
}

func GetExtensionForFormat(imageType bimg.ImageType) string {
	return fmt.Sprintf("%s%s", ".", bimg.ImageTypes[imageType])

}

func (img *Image) Convert(operation Operation) ([]byte, error) {

	options := bimg.Options{
		Height:  operation.Height,
		Width:   operation.Width,
		Quality: operation.Quality,
		Zoom:    operation.Scale,
		Type:    bimg.ImageType(operation.Format),
		Speed:   8,
	}
	return img.Process(options)

}

func FromPath(path string) (Image, error) {
	data, err := bimg.Read(path)
	if err != nil {
		return Image{}, err
	}
	img := bimg.NewImage(data)
	return Image{*img, path}, nil
}
