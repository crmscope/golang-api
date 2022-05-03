package main

import (
	"crmgo/library"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func response(errorCode int, errorMessage string, resultData []library.Data) string {

	rData, err := json.Marshal(resultData)

	if err != nil {
		library.Exception(400, "Main: Json parsing error.", string(err.Error()))
	}

	message1 := map[string]string{
		"code":        strconv.Itoa(errorCode),
		"result_data": string(rData),
		"error":       errorMessage,
	}

	jsonStr, _ := json.Marshal(message1)
	return string(jsonStr)
}

func apiAction(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		module := GetModule(r.FormValue("module"))
		fmt.Fprint(w, response(module.Init(r.FormValue("rest_data"))))
	}
}

func main() {

	http.HandleFunc("/api", apiAction)
	http.ListenAndServe(":8001", nil)

}
