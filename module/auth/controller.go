package auth

import (
	"crmgo/library"
)

type Controller struct{}

func (t *Controller) Run(rData []library.Data) string {

	var lbr library.CrmBean

	whare := []string{
		"main.login = '" + library.GetVal("name", rData) + "'",
		"main.password = MD5('" + library.GetVal("password", rData) + "')",
	}

	tmp := lbr.GetCollection("sys_users", whare, "ORDER BY main.id DESC")

	return string(tmp)
}
