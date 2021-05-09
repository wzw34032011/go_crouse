package service

import (
	"github.com/google/wire"
	"go_crouse/4/app/user/service/internal/biz"
)

var ProviderSet = wire.NewSet(NewUserService)

type UserService struct {
	Biz *biz.UserBiz
}

func NewUserService(biz *biz.UserBiz) *UserService {
	return &UserService{
		Biz: biz,
	}
}
