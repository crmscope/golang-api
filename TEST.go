package main

import (
	"database/sql"
	"fmt"
	"reflect"
	"time"

	"github.com/go-sql-driver/mysql"
)

type BaseBean struct {
}

type UserBean struct {
	id               int32
	id_module        sql.NullInt32
	login            sql.NullString
	password         sql.NullString
	date_entered     mysql.NullTime
	date_modified    mysql.NullTime
	deleted          sql.NullInt32
	created_by       sql.NullInt32
	assigned_user_id sql.NullInt32
	modified_user_id sql.NullInt32
	id_users         sql.NullInt32
	id_language      sql.NullInt32
	name             sql.NullString
	description      sql.NullString
	phone            sql.NullInt32
}

//------------------------------------------------------------------------------
func (t *UserBean) SetId(id int32) {
	t.id = id
}
func (t *UserBean) GetId() int32 {
	return t.id
}

//------------------------------------------------------------------------------
func (t *UserBean) SetIdModule(id int32) {
	idInt32 := sql.NullInt32{id, true}
	t.id_module = idInt32
}

func (t *UserBean) GetIdModule() int32 {
	return t.id_module.Int32
}

//------------------------------------------------------------------------------

func (t *UserBean) SetLogin(login string) {
	loginString := sql.NullString{login, true}
	t.login = loginString
}

func (t *UserBean) GetLogin() string {
	return t.login.String
}

//------------------------------------------------------------------------------
func (t *UserBean) SetPassword(password string) {
	passwordString := sql.NullString{password, true}
	t.password = passwordString
}

func (t *UserBean) GetPassword() string {
	return t.password.String
}

//------------------------------------------------------------------------------
func (t *UserBean) SetDateEntered(dateEntered time.Time) {
	dateEnteredTime := mysql.NullTime{dateEntered, true}
	t.date_entered = dateEnteredTime
}

func (t *UserBean) GetDateEntered() time.Time {
	return t.date_entered.Time
}

//------------------------------------------------------------------------------
func (t *UserBean) SetDateModified(dateModified time.Time) {
	dateModifiedTime := mysql.NullTime{dateModified, true}
	t.date_modified = dateModifiedTime
}

func (t *UserBean) GetDateModified() time.Time {
	return t.date_modified.Time
}

//------------------------------------------------------------------------------
func (t *UserBean) SetDeleted(deleted int32) {
	deletedInt32 := sql.NullInt32{deleted, true}
	t.deleted = deletedInt32
}

func (t *UserBean) GetDeleted() int32 {
	return t.deleted.Int32
}

//------------------------------------------------------------------------------
func (t *UserBean) SetCreatedBy(createdBy int32) {
	createdByInt32 := sql.NullInt32{createdBy, true}
	t.created_by = createdByInt32
}

func (t *UserBean) GetCreatedBy() int32 {
	return t.created_by.Int32
}

//------------------------------------------------------------------------------
func (t *UserBean) SetAssignedUserId(assignedUserId int32) {
	idInt32 := sql.NullInt32{assignedUserId, true}
	t.assigned_user_id = idInt32
}

func (t *UserBean) GetAssignedUserId() int32 {
	return t.assigned_user_id.Int32
}

//------------------------------------------------------------------------------
func (t *UserBean) SetIdUsers(idUsers int32) {
	idUsersInt32 := sql.NullInt32{idUsers, true}
	t.id_users = idUsersInt32
}

func (t *UserBean) GetIdUsers() int32 {
	return t.id_users.Int32
}

//------------------------------------------------------------------------------
func (t *UserBean) SetIdLanguage(idLanguage int32) {
	idInt32 := sql.NullInt32{idLanguage, true}
	t.id_language = idInt32
}

func (t *UserBean) GetIdLanguage() int32 {
	return t.id_language.Int32
}

//------------------------------------------------------------------------------
func (t *UserBean) SetName(name string) {
	idString := sql.NullString{name, true}
	t.name = idString
}

func (t *UserBean) GetName() string {
	return t.name.String
}

//------------------------------------------------------------------------------
func (t *UserBean) SetDescription(description string) {
	descriptionString := sql.NullString{description, true}
	t.description = descriptionString
}

func (t *UserBean) GetDescription() string {
	return t.description.String
}

//------------------------------------------------------------------------------

func (t *BaseBean) GetPointers() []interface{} {
	val := reflect.ValueOf(t).Elem()
	v := make([]interface{}, val.NumField())
	for i := 0; i < val.NumField(); i++ {
		valueField := val.Field(i)
		v[i] = valueField.Addr().Pointer()
	}
	return v
}

func (t *BaseBean) GetFields() []string {
	var fields []string
	fieldReflect := reflect.TypeOf(*t)
	if fieldReflect.Kind() == reflect.Struct {
		for i := 0; i < fieldReflect.NumField(); i++ {
			fields = append(fields, fieldReflect.Field(i).Name)
		}
	}
	return fields
}

func main() {

	var varTest UserBean

	varTest.SetId(11)
	varTest.SetIdModule(22)
	varTest.SetLogin("Login")
	varTest.SetPassword("Password")
	varTest.SetDateEntered(time.Now())
	varTest.SetDateModified(time.Now())
	varTest.SetDeleted(77)
	varTest.SetCreatedBy(88)
	varTest.SetAssignedUserId(99)
	varTest.SetIdUsers(110)
	varTest.SetIdLanguage(120)
	varTest.SetName("name")
	varTest.SetDescription("Description")

	fmt.Println(varTest.GetId())
	fmt.Println(varTest.GetIdModule())
	fmt.Println(varTest.GetLogin())
	fmt.Println(varTest.GetPassword())
	fmt.Println(varTest.GetDateEntered())
	fmt.Println(varTest.GetDateModified())
	fmt.Println(varTest.GetDeleted())
	fmt.Println(varTest.GetCreatedBy())
	fmt.Println(varTest.GetAssignedUserId())
	fmt.Println(varTest.GetIdUsers())
	fmt.Println(varTest.GetIdLanguage())
	fmt.Println(varTest.GetName())
	fmt.Println(varTest.GetDescription())

	fmt.Println(varTest.GetPointers())
	fmt.Println(varTest.GetFields())

}
