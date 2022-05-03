package library

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type MysqlConnector struct{}

func (t *MysqlConnector) GetConnect(user string, password string, dbName string) (db *sql.DB) {

	db, err := sql.Open("mysql", user+":"+password+"@tcp(localhost:3306)/"+dbName)
	if err != nil {
		Exception(400, "MysqlConnector: MySQL connection error.", string(err.Error()))
	}

	return db
}
