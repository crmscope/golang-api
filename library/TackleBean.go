package library

import (
	"database/sql"
	"reflect"
	"strings"
)

type TackleBean struct {
	db *sql.DB
}

func (t *TackleBean) GetCollection(beanName string, where []string, order string) *sql.Rows {

	var conn MysqlConnector

	t.db = conn.GetConnect("root", "me", "tackle_api")
	defer t.db.Close()

	results, err := t.db.Query("SELECT main.*, attr.*  " +
		"FROM " + beanName + " as main " +
		"INNER JOIN " + beanName + "_attr attr ON attr.id_user = main.id AND main.deleted != 1 " +
		"WHERE " + strings.Join(where, " AND ") + " " +
		order)

	if err != nil {
		Exception(400, "CrmBean: SQL query error.", string(err.Error()))
	}

	return results
}

func (t *TackleBean) GetBean(beanName string, id string) *sql.Rows {

	var conn MysqlConnector

	t.db = conn.GetConnect("root", "me", "tackle_api")
	defer t.db.Close()

	results, err := t.db.Query("SELECT main.*, attr.*  " +
		"FROM " + beanName + " as main " +
		"INNER JOIN " + beanName + "_attr attr ON attr.id_user = main.id AND main.deleted != 1 " +
		"WHERE main.id = \"" + id + "\"")

	if err != nil {
		Exception(400, "CrmBean: SQL query error.", string(err.Error()))
	}

	return results
}

func GetFields(bean interface{}) []string {
	var fields []string
	fieldReflect := reflect.Indirect(reflect.ValueOf(bean))
	if fieldReflect.Kind() == reflect.Struct {
		for i := 0; i < fieldReflect.NumField(); i++ {
			fields = append(fields, fieldReflect.Type().Field(i).Name)
		}
	}
	return fields
}
