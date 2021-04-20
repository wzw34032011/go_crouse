package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
)

const (
	USERNAME = "db_user_name"
	PASSWORD = "db_password"
	NETWORK  = "tcp"
	SERVER   = "db_address"
	PORT     = 3306
	DATABASE = "db_database"
)

var DB = &sql.DB{}

func init() {
	dsn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s", USERNAME, PASSWORD, NETWORK, SERVER, PORT, DATABASE)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic("mysql connect fail:" + err.Error())
	}
	DB = db
}

func DaoGetName(id int) (string, error) {
	var name sql.NullString
	err := DB.QueryRow("SELECT name FROM user WHERE id=?", id).Scan(&name)

	var de *DaoError
	switch {
	case err == sql.ErrNoRows:
		de = &DaoError{msg: "query not find", emptyRow: true, err: err}
	case err != nil:
		de = &DaoError{msg: "query error", emptyRow: false, err: err}
	default:
		de = nil
	}

	return name.String, de
}

func ServiceUser(id int) (string, error) {
	name, err := DaoGetName(id)
	if err != nil {
		var se *ServiceError
		if de, ok := err.(*DaoError); ok && de.IsEmptyRow() {
			se = &ServiceError{code: 1, msg: "no row find", emptyRow: true, err: de}
		} else {
			//做一些重试查询，数据兜底相关逻辑
			se = &ServiceError{code: 2, msg: "db error", emptyRow: false, err: de}
		}
		return name, se
	}
	return name, err
}

func ApiController(request *http.Request) string {
	id := 1
	name, err := ServiceUser(id)
	if err != nil {
		if e, ok := err.(*ServiceError); ok {
			return structToJson(Response{code: e.code, msg: e.Error(), data: ResponseData{name: name}})
		}
	}

	return structToJson(Response{code: 0, msg: "", data: ResponseData{name: name}})
}

type ServiceError struct {
	code     int
	msg      string
	emptyRow bool
	err      error
}

func (se *ServiceError) IsEmptyRow() bool {
	return se.emptyRow
}

func (se *ServiceError) Error() string {
	return se.msg
}

func (se *ServiceError) Unwrap() error {
	return se.err
}

func (se *ServiceError) Is(target error) bool {
	_, ok := target.(*ServiceError)
	return ok
}

type DaoError struct {
	msg      string
	emptyRow bool
	err      error
}

func (de *DaoError) Unwrap() error {
	return de.err
}

func (de *DaoError) Error() string {
	return de.msg
}

func (de *DaoError) IsEmptyRow() bool {
	return de.emptyRow
}

type Response struct {
	code int
	msg  string
	data ResponseData
}

type ResponseData struct {
	name string
}

func structToJson(s interface{}) string {
	jsonBytes, err := json.Marshal(s)
	if err != nil {
		return ""
	}
	return string(jsonBytes)
}

func main() {
	ApiController(new(http.Request))
}
