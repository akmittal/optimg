package operation

import (
	"encoding/json"
	"fmt"
	"net/http"

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

		newOperation, _ := Get(opr.SourcePath, opr.TargetPath, opr.CopyUnknown, opr.Monitor, opr.Transformations)
		tx := db.Create(&newOperation)
		tx.Commit()
		err = newOperation.Process(db)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusBadRequest)
			return
		}
		fmt.Fprint(rw, "done")
	}
}
