package operation

import (
	"encoding/json"
	"fmt"
	"net/http"

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
