package service

import (
	"go_crouse/4/app/user/service/internal/biz"
)

type UserService struct {
	Biz *biz.UserBiz
}

func NewUserService(biz *biz.UserBiz) *UserService {
	return &UserService{
		Biz: biz,
	}
}
