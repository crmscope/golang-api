package main

import (
	"crmgo/manifest"
)

type ModuleInterface interface {
	Init(restData string) string
}

func GetModule(moduleName string) ModuleInterface {

	modules := make(map[string]interface{})
	modules["auth"] = &manifest.ManifestAuth{}
	modules["user"] = &manifest.ManifestUser{}

	module := modules[moduleName]
	return module.(ModuleInterface)
}
