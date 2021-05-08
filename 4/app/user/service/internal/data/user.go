package data

import (
	"context"
	"go_crouse/4/app/user/service/internal/biz"
	"log"
)

func (ud *UserData) AddUser(ctx context.Context, info *biz.UserInfo) error {
	//ud.db.Exec("")
	log.Println("data-Adduser")
	log.Println(info)
	return nil
}

func (ud *UserData) GetUser(ctx context.Context, id int) (*biz.UserInfo, error) {
	userInfo := &biz.UserInfo{Name: "xxx", Age: 0, Id: id}
	return userInfo, nil
}
