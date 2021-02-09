package operation

import (
	"encoding/json"
	"net/http"

	"github.com/akmittal/optimg/server/pkg/image"
	"github.com/akmittal/optimg/server/pkg/util"
	"gorm.io/gorm"
)

func HandleStartOperation(db *gorm.DB) http.HandlerFunc {
	return func(rw http.ResponseWriter, req *http.Request) {
		var opr Operation
		err := json.NewDecoder(req.Body).Decode(&opr)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusBadRequest)
			return
		}

		newOperation, err := Get(opr.SourcePath, opr.TargetPath, opr.CopyUnknown, opr.Monitor, opr.Transformations)

		if err != nil {

			util.HTTPError(rw, util.Error{Msg: err.Error()}, http.StatusBadRequest)
			return
		}
		// tx := db.Create(&newOperation)
		// tx.Commit()

		err = newOperation.Process(db)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusBadRequest)
			return
		}
		util.HTTPJson(rw, map[string]string{"msg": "Done"})
	}
}

func HandleStats(db *gorm.DB) http.HandlerFunc {
	return func(rw http.ResponseWriter, req *http.Request) {
		var images []image.Image
		subQuery := db.Select("MAX(id)").Table("operations")
		db.Table("images").Select("images.*").Where("operation_ref = (?)", subQuery).Find(&images)
		var sourceImages []image.Image
		var totalSourceSize int
		var totalDestSize map[*Transformation]int
		for _, image := range images {
			if image.ParentID == 0 {
				sourceImages = append(sourceImages, image)
				totalSourceSize += int(image.Size)

			} else {
				if transformationNotExist(totalDestSize, image.TransformationRef) {
					var transform Transformation
					db.First(&transform, image.TransformationRef)
					totalDestSize[&transform] = int(image.Size)
				} else {
					transform := getTransformByID(totalDestSize, image.TransformationRef)
					totalDestSize[transform] = totalDestSize[transform] + int(image.Size)
				}

			}
		}
		json.NewEncoder(rw).Encode(map[string]interface{}{"sourceLength": len(images), "totalDestSize": totalDestSize})

	}
}

func transformationNotExist(mapping map[*Transformation]int, transformID uint) bool {
	for transform := range mapping {
		if transform.ID == transformID {
			return false
		}

	}
	return true
}
func getTransformByID(mapping map[*Transformation]int, transformID uint) *Transformation {
	for transform := range mapping {
		if transform.ID == transformID {
			return transform
		}

	}
	return nil
}
