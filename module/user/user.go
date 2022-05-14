package user

import (
	"database/sql"
	"crmgo/library"

	"github.com/go-sql-driver/mysql"
)

type UserBean struct {
	Id               int32
	Id_module        sql.NullInt32
	Login            sql.NullString
	Password         sql.NullString
	Date_entered     mysql.NullTime
	Date_modified    mysql.NullTime
	Deleted          sql.NullInt32
	Created_by       sql.NullInt32
	Assigned_user_id sql.NullInt32
	Modified_user_id sql.NullInt32
	Id_users         sql.NullInt32
	Id_language      sql.NullInt32
	Name             sql.NullString
	Description      sql.NullString
	Phone            sql.NullInt32
}

func (t *UserBean) GetScans() []interface{} {
	return []interface{}{
		&t.Id,
		&t.Id_module,
		&t.Login,
		&t.Password,
		&t.Date_entered,
		&t.Date_modified,
		&t.Deleted,
		&t.Created_by,
		&t.Assigned_user_id,
		&t.Modified_user_id,
		&t.Id_users,
		&t.Id_language,
		&t.Name,
		&t.Description,
		&t.Phone,
	}
}

func (t *UserBean) GetJson() jsonMap []library.Data {
	
	jsonMap[0].Name
	
	
	jsonMap[0]["namr"]
	
	
	jsonMap["Id"] = t.Id;
	jsonMap["Id_module"] = t.Id_module,
	jsonMap["Login"] = t.Login,
	jsonMap["Password"] = t.Password,
	jsonMap["Date_entered"] = t.Date_entered,
	jsonMap["Date_modified"] = t.Date_modified,
	jsonMap["Deleted"] = t.Deleted,
	jsonMap["Created_by"] = t.Created_by,
	jsonMap["Assigned_user_id"] = t.Assigned_user_id,
	jsonMap["Modified_user_id"] = t.Modified_user_id,
	jsonMap["Id_users"] = t.Id_users,
	jsonMap["Id_language"] = t.Id_language,
	jsonMap["Name"] = t.Name,
	jsonMap["Description"] = t.Description,
	jsonMap["Phone"] = t.Phone,
}

