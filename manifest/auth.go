package manifest

import (
	"crmgo/library"
	"crmgo/module/auth"
	"encoding/json"
)

type ManifestAuth struct{}

func (t *ManifestAuth) Init(restData string) string {

	var rData []library.Data
	var contr auth.Controller

	err := json.Unmarshal([]byte(restData), &rData)
	if err != nil {
		panic(err)
	}

	return contr.Run(rData)
}
