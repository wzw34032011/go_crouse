package service

import (
	"context"
	v1 "go_crouse/4/api/user/v1"
	"go_crouse/4/app/user/service/internal/biz"
	"log"
)

func (us *UserService) AddUser(ctx context.Context, req *v1.AddUserReq) (*v1.AddUserReply, error) {
	log.Println("service-Adduser")
	user := &biz.UserInfo{Name: req.Name}
	err := us.Biz.AddUser(ctx, user)
	return &v1.AddUserReply{Code: 0, Msg: "ok"}, err
}
