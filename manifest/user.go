package manifest

import (
	"encoding/json"
	"tackle-api/library"
	"tackle-api/module/user"
)

type ManifestUser struct{}

func (t *ManifestUser) Init(restData string) (int, string, []library.Data) {

	var rData []library.Data
	var contr user.Controller

	err := json.Unmarshal([]byte(restData), &rData)

	if err != nil {
		library.Exception(400, "ManifestUser: Json parsing error.", string(err.Error()))
	}

	return contr.Run(rData)
}
