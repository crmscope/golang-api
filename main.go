package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func response(code int, message string, err string) string {
	message1 := map[string]string{
		"code":    strconv.Itoa(code),
		"message": message,
		"error":   err,
	}

	jsonStr, _ := json.Marshal(message1)
	return string(jsonStr)
}

func apiAction(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		module := GetModule(r.FormValue("module"))
		fmt.Fprint(w, response(200, module.Init(r.FormValue("rest_data")), ""))
	}
}

func main() {
	http.HandleFunc("/api", apiAction)
	http.ListenAndServe(":8001", nil)

}
