package user

import (
	//"encoding/json"
	"tackle-api/library"

	"github.com/google/uuid"
)

type Controller struct{}

func (t *Controller) Run(rData []library.Data) (errorCode int, errorMessage string, resultData []library.Data) {

	var lbr library.TackleBean
	user := UserBean{}
	userCollection := []UserBean{}

	errorCode = 200
	errorMessage = ""

	beanId := library.GetVal("id", rData)

	results := lbr.GetBean("user", beanId)

	for results.Next() {
		err := results.Scan(user.GetPointers()...)
		userCollection = append(userCollection, user)
		//jsonStr, _ := json.Marshal(user)

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
