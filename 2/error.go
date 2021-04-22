package main

import (
	"database/sql"
	"github.com/pkg/errors"
	"log"
)

type MyError struct {
	msg      string
	emptyRow bool
	err      error
}

func (de *MyError) Cause() error {
	return de.err
}

func (de *MyError) Unwrap() error {
	return de.err
}

func (de *MyError) Error() string {
	return de.msg
}

func (de *MyError) IsEmptyRow() bool {
	return de.emptyRow
}

func WrapStackOnce(err error, message string) error {
	if err == nil {
		return nil
	}
	//如果err包含过withStack,就不使用wrap重复记录stack了
	if e, ok := err.(interface{ StackTrace() errors.StackTrace }); ok {
		return errors.WithMessage(e.(error), message)
	} else {
		return errors.Wrap(err, message)
	}
}

func Dao() error {
	var err = sql.ErrNoRows
	//var err = sql.ErrConnDone

	var me *MyError
	var we error
	switch {
	case errors.Is(err, sql.ErrNoRows):
		me = &MyError{msg: "query not find", emptyRow: true, err: err}
		we = WrapStackOnce(me, "dao we")
	case err != nil:
		me = &MyError{msg: "query err", emptyRow: false, err: err}
		we = WrapStackOnce(me, "dao we")
	default:
		we = nil
	}
	return we
}

func Service() error {
	err := Dao()
	if err != nil {
		var me *MyError
		if errors.As(err, &me) && me.IsEmptyRow() {
			//空数据错误
			if canHandle := false; canHandle {
				//处理错误代码,降级业务流程
				return nil
			} else {
				//无法处理错误，返回给调用者
				return WrapStackOnce(err, "未查询到数据")
			}
		} else {
			//其他错误
			if canHandle := false; canHandle {
				//处理错误代码,降级业务流程
				return nil
			} else {
				//无法处理错误，返回给调用者
				return WrapStackOnce(err, "Dao 异常")
			}
		}
	}

	//正常业务流程
	return WrapStackOnce(err, "其他业务错误")
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
