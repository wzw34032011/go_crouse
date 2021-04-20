package main

import (
	"database/sql"
	"errors"
	xerrors "github.com/pkg/errors"
	"log"
)

func Dao() (string, error) {
	var name sql.NullString
	var DB sql.DB
	err := DB.QueryRow("SELECT name FROM user WHERE id=1").Scan(&name)

	var de *DaoError
	switch {
	case err == sql.ErrNoRows:
		de = &DaoError{msg: "query not find", emptyRow: true, err: err}
	case err != nil:
		de = &DaoError{msg: "query error", emptyRow: false, err: err}
	default:
		de = nil
	}
	return name.String, xerrors.Wrap(de, "dao error")
}

func Service() (string, error) {
	name, err := Dao()
	if err != nil {
		var se *ServiceError
		var de *DaoError
		if errors.As(err, &de) && de.IsEmptyRow() {
			//空数据错误
			se = &ServiceError{code: 1, msg: "未查询到数据", err: err}
		} else {
			//其他错误
			se = &ServiceError{code: 2, msg: "服务异常", err: err}
		}
		return name, xerrors.Wrap(se, "service error")
	}
	return name, nil
}

func Handler() Response {
	name, err := Service()
	if err != nil {
		log.Printf("stack trace:\n%+v\n", err)

		var se *ServiceError
		if errors.As(err, &se) {
			return Response{code: se.code, msg: se.Error(), data: ResponseData{name: name}}
		} else {
			panic("error invalid")
		}
	}

	return Response{code: 0, msg: "", data: ResponseData{name: name}}
}

// service层错误
type ServiceError struct {
	code int
	msg  string
	err  error
}

func (se *ServiceError) Error() string {
	return se.msg
}

func (se *ServiceError) Unwrap() error {
	return se.err
}

// Dao层错误
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

func main() {
	Handler()
}
