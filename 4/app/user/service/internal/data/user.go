package data

import (
	"context"
	"go_crouse/4/app/user/service/internal/biz"
	"go_crouse/4/app/user/service/internal/data/ent/user"
	"log"
)

func (ud *UserData) AddUser(ctx context.Context, info *biz.UserInfo) error {
	log.Println("data-Adduser")
	log.Println(info)
	ud.db.User.Create().
		SetAge(info.Age).
		SetName(info.Name).
		SetGender(user.Gender0).
		Save(ctx)
	return nil
}

func (ud *UserData) GetUser(ctx context.Context, id int) (*biz.UserInfo, error) {
	userInfo := &biz.UserInfo{Name: "xxx", Age: 0, Id: id}
	//ud.db.User.Query().Where(user.IDEQ(id)).Only(ctx)
	return userInfo, nil
}
