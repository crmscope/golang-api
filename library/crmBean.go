package library

import (
	"database/sql"
	"strings"
)

type CrmBean struct {
	db *sql.DB
}

func (t *CrmBean) GetCollection(beanName string, where []string, order string) string {

	var conn MysqlConnector
	var TMP string
	t.db = conn.GetConnect("root", "me", "crmgo")
	defer t.db.Close()

	results, err := t.db.Query("SELECT main.login " +
		"FROM " + beanName + " as main " +
		"WHERE " + strings.Join(where, " AND ") + " " +
		order)

	if err != nil {
		panic(err.Error())
	}

	for results.Next() {

		err = results.Scan(&TMP)
		if err != nil {
			panic(err.Error())
		}

	}
	return TMP

}
