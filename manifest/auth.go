package manifest

import (
	"crmgo/library"
	"crmgo/module/auth"
	"encoding/json"
)

type ManifestAuth struct{}

func (t *ManifestAuth) Init(restData string) (int, string, []library.Data) {

	var rData []library.Data
	var contr auth.Controller

	err := json.Unmarshal([]byte(restData), &rData)
	if err != nil {
		library.Exception(400, "ManifestAuth: Json parsing error.", string(err.Error()))
	}

	return contr.Run(rData)
}
