package main

import (
	"database/sql"
	"fmt"
	"github.com/pkg/errors"
	"log"
)

// 我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？

// sql.ErrNoRows 是 Scan(dest ...interface{}) 方法在无查询记录时返回的错误，需要抛给上层。
// 第一种情况，能查询到记录，但是这条记录的字段值可能是空值，这时Scan方法给dest字段赋值为空值
// 第二种情况，没有查询到记录，Scan方法不会给dest字段赋值，dest变量为初始零值（空值）
// 情况一和二都返回了一样的空值，如果不上抛错误，上层无法得知数据是否存在

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

type ServiceError struct {
	code int
	msg  string
	err  error
}

func (se *ServiceError) Cause() error {
	return se.err
}

func (se *ServiceError) Unwrap() error {
	return se.err
}

func (se *ServiceError) Error() string {
	return fmt.Sprintf("code:%d,msg:%s", se.code, se.msg)
}

func (se *ServiceError) Is(tag error) bool {
	_, ok := tag.(*ServiceError)
	return ok
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

func Dao() (string, error) {
	var data string
	var err = sql.ErrNoRows
	//var err = sql.ErrConnDone

	var de *DaoError
	var we error = nil
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			de = &DaoError{msg: "query not find", emptyRow: true, err: err}
			we = WrapStackOnce(de, "dao error")
		} else {
			de = &DaoError{msg: "connect err", emptyRow: false, err: err}
			we = WrapStackOnce(de, "dao error")
		}
	} else {
		data = "正常数据"
	}
	return data, we
}

func Service() (string, error) {
	data, err := Dao()
	if err != nil {
		var de *DaoError
		if errors.As(err, &de) && de.IsEmptyRow() {
			//空数据错误
			if returnError := true; returnError {
				//错误上抛
				se := &ServiceError{
					code: 1,
					msg:  "未查询到数据",
					err:  err,
				}
				return data, WrapStackOnce(se, "service error")
			} else {
				//降级处理
				return "兜底数据", nil
			}
		} else {
			//无法处理错误，错误上抛
			return data, WrapStackOnce(err, "service error")
		}
	}

	//正常流程
	return data, nil
}

func main() {
	data, err := Service()
	if err != nil {
		if !errors.Is(err, new(ServiceError)) {
			//错误无法处理，写log
			log.Printf("stack trace:\n%+v\n", err)
			log.Printf("original error: %v", errors.Cause(err))

			err = &ServiceError{
				code: 2,
				msg:  "unknown error",
				err:  err,
			}
		}
	}

	fmt.Printf("data=%s\nerr=%v", data, err)
}
