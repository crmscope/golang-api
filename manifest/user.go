package manifest

import (
	"crmgo/library"
	"crmgo/module/user"
	"encoding/json"
	"fmt"
)

type ManifestUser struct{}

func (t *ManifestUser) Init(restData string) (int, string, []library.Data) {

	var rData []library.Data
	var contr user.Controller

	err := json.Unmarshal([]byte(restData), &rData)
	fmt.Println("sss: " + rData[0].Value)

	if err != nil {
		library.Exception(400, "ManifestUser: Json parsing error.", string(err.Error()))
	}

	return contr.Run(rData)
}
