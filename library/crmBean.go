package library

import (
	"database/sql"
	"strings"
)

type CrmBean struct {
	db *sql.DB
}

func (t *CrmBean) GetCollection(beanName string, where []string, order string) *sql.Rows {

	var conn MysqlConnector

	t.db = conn.GetConnect("root", "me", "crmgo")
	defer t.db.Close()

	results, err := t.db.Query("SELECT main.*, attr.*  " +
		"FROM " + beanName + " as main " +
		"INNER JOIN " + beanName + "_attr attr ON attr.id_users = main.id AND main.deleted != 1 " +
		"WHERE " + strings.Join(where, " AND ") + " " +
		order)

	if err != nil {
		panic(err.Error())
	}

	return results
}
