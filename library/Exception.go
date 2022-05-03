package library

import (
	"encoding/json"
	"strconv"
)

func Exception(errorCode int, errorMessage string, stack string) {

	message1 := map[string]string{
		"code":  strconv.Itoa(errorCode),
		"error": errorMessage,
		"stack": string(stack),
	}
	jsonStr, _ := json.Marshal(message1)
	panic(string(jsonStr))
}
