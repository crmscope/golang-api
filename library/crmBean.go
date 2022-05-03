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
		Exception(400, "CrmBean: SQL query error.", string(err.Error()))
	}

	return results
}

func (t *CrmBean) GetBean(beanName string, id int) *sql.Rows {

	var conn MysqlConnector

	t.db = conn.GetConnect("root", "me", "crmgo")
	defer t.db.Close()

	results, err := t.db.Query("SELECT main.*, attr.*  " +
		"FROM " + beanName + " as main " +
		"INNER JOIN " + beanName + "_attr attr ON attr.id_users = main.id AND main.deleted != 1 " +
		"WHERE main.id = ".id)

	if err != nil {
		Exception(400, "CrmBean: SQL query error.", string(err.Error()))
	}

	return results
}
