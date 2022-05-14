package user

import (
	"crmgo/library"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/google/uuid"
)

type Controller struct{}

func (t *Controller) Run(rData []library.Data) (errorCode int, errorMessage string, resultData []library.Data) {

	var lbr library.CrmBean
	user := UserBean{}
	userCollection := []UserBean{}

	errorCode = 200
	errorMessage = ""

	beanId := library.GetVal("id", rData)

	bId, err := strconv.Atoi(beanId)

	if err != nil {
		fmt.Println("is not an integer.")
	}

	results := lbr.GetBean("sys_users", bId)

	for results.Next() {
		err := results.Scan(user.GetScans()...)
		userCollection = append(userCollection, user)

		jsonStr, _ := json.Marshal(user)
		fmt.Println(string(jsonStr))

		if err != nil {
			library.Exception(400, "Main: Json parsing error.", string(err.Error()))
		}
	}

	if userCollection[0].Login.String == "" {
		errorCode = 401
		errorMessage = "Invalid login or password"
	}

	resultData = append(resultData, library.Data{Name: "token", Value: uuid.New().String()})

	return errorCode, errorMessage, resultData

}
