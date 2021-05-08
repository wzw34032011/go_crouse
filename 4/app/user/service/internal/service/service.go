package service

import (
	"github.com/google/wire"
	"go_crouse/4/app/user/service/internal/biz"
)

type UserService struct {
	Biz *biz.UserBiz
}

var ProviderSet = wire.NewSet(NewUserService)

func NewUserService(biz *biz.UserBiz) *UserService {
	return &UserService{
		Biz: biz,
	}
}
