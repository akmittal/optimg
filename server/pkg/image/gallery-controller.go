package image

import (
	"encoding/json"
	"math"
	"net/http"
	"strconv"

	"gorm.io/gorm"
)

type ImageData struct {
	Image    Image   `json:"image"`
	Varients []Image `json:"varients"`
}

type Directory struct {
	Name string `json:"name"`
	Path string `json:"path"`
}

type Gallery struct {
	CurrentPage uint        `json:"currentPage"`
	TotalPages  uint        `json:"totalPages"`
	Images      []ImageData `json:"images"`
	Directories []Directory `json:"directories"`
}

const (
	PageLimit = 15
)

// Todo : check path and return subdirectories

func HandleGallery(db *gorm.DB) http.HandlerFunc {
	return func(rw http.ResponseWriter, req *http.Request) {
		var gallery Gallery
		var images []Image

		pageQuery := req.URL.Query().Get("page")
		// path := req.URL.Query().Get("path")
		page := 1
		if pageQuery != "" {
			page, _ = strconv.Atoi(pageQuery)
		}

		// path := "/"
		var totalCount int64
		db.Model(&Image{}).Where(&Image{}, "parent_id").Count(&totalCount)
		gallery.TotalPages = uint(math.Ceil(float64(totalCount) / float64(PageLimit)))
		gallery.CurrentPage = uint(page)
		db.Where(&Image{}, "parent_id").Limit(PageLimit).Offset(PageLimit * (page - 1)).Find(&images)
		rw.Header().Add("Content-Type", "application/json")
		for _, image := range images {
			var varientImages []Image
			db.Where(&Image{ParentID: image.ID}).Find(&varientImages)
			var imageData ImageData
			imageData.Image = image
			imageData.Varients = varientImages
			gallery.Images = append(gallery.Images, imageData)
		}
		json.NewEncoder(rw).Encode(gallery)

	}
}
