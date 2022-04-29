package auth

import (
	"crmgo/library"
	"crmgo/module/user"
)

type Controller struct{}

func (t *Controller) Run(rData []library.Data) string {

	var lbr library.CrmBean
	userCollection := []user.UserBean{}

	whare := []string{
		"main.login = '" + library.GetVal("name", rData) + "'",
		"main.password = MD5('" + library.GetVal("password", rData) + "')",
	}

	user := user.UserBean{}
	results := lbr.GetCollection("sys_users", whare, "ORDER BY main.id DESC")

	for results.Next() {
		err := results.Scan(user.GetScans()...)
		userCollection = append(userCollection, user)

		if err != nil {
			panic(err.Error())
		}

	}

	return userCollection[0].Login.String
}
