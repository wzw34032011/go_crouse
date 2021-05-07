package service

import (
	"context"
	v1 "go_crouse/4/api/user/v1"
	"log"
)

func (us *UserService) AddUser(ctx context.Context, req *v1.AddUserReq) (*v1.AddUserReply, error) {
	log.Println("service-Adduser")
	us.Biz.AddUser(ctx)
	return &v1.AddUserReply{Code: 0, Msg: "ok"}, nil
}
