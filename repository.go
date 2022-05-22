package main

import (
	"tackle-api/library"
	"tackle-api/manifest"
)

type ModuleInterface interface {
	Init(restData string) (int, string, []library.Data)
}

func GetModule(moduleName string) ModuleInterface {

	modules := make(map[string]interface{})
	modules["auth"] = &manifest.ManifestAuth{}
	modules["user"] = &manifest.ManifestUser{}

	module := modules[moduleName]
	return module.(ModuleInterface)
}
