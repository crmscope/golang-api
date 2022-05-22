package auth

import (
	"tackle-api/library"
	"tackle-api/module/user"

	"github.com/google/uuid"
)

type Controller struct{}

func (t *Controller) Run(rData []library.Data) (errorCode int, errorMessage string, resultData []library.Data) {

	var lbr library.TackleBean
	userCollection := []user.UserBean{}

	errorCode = 200
	errorMessage = ""

	whare := []string{
		"main.login = '" + library.GetVal("name", rData) + "'",
		"main.password = MD5('" + library.GetVal("password", rData) + "')",
	}

	user := user.UserBean{}

	results := lbr.GetCollection("user", whare, "ORDER BY main.id DESC")

	for results.Next() {
		err := results.Scan(user.GetPointers()...)
		userCollection = append(userCollection, user)

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
