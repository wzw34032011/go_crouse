package main

import (
	"database/sql"
	"github.com/pkg/errors"
	"log"
)

// Dao层错误
type DaoError struct {
	msg      string
	emptyRow bool
	err      error
}

func (de *DaoError) Cause() error {
	return de.err
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

func Dao() error {
	var err = sql.ErrNoRows
	//var err = sql.ErrConnDone

	var de *DaoError
	var err2 error
	switch {
	case errors.Is(err, sql.ErrNoRows):
		de = &DaoError{msg: "query not find", emptyRow: true, err: err}
		err2 = errors.Wrap(de, "dao error")
	case err != nil:
		de = &DaoError{msg: "query err", emptyRow: false, err: err}
		err2 = errors.Wrap(de, "dao error")
	default:
		err2 = nil
	}
	return err2
}

func Service() error {
	err := Dao()
	if err != nil {
		var de *DaoError
		if errors.As(err, &de) && de.IsEmptyRow() {
			//空数据错误
			if canHandle := false; canHandle {
				//处理错误代码,降级业务流程
				return nil
			} else {
				//无法处理错误，返回给调用者
				return errors.Wrap(err, "未查询到数据")
			}
		} else {
			//其他错误
			if canHandle := false; canHandle {
				//处理错误代码,降级业务流程
				return nil
			} else {
				//无法处理错误，返回给调用者
				return errors.Wrap(err, "Dao 异常")
			}
		}
	}
	//正常业务流程

	return err
}

func main() {
	err := Service()
	if err != nil {
		log.Printf("stack trace:\n%+v\n", err)
		log.Printf("original error: %v", errors.Cause(err))
	} else {
		log.Println("success")
	}
}
