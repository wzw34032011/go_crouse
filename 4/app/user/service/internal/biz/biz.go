package biz

import (
	"context"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(NewUserBiz)

type UserBiz struct {
	data UserRepo
}

type UserRepo interface {
	AddUser(ctx context.Context, info *UserInfo) error
	GetUser(ctx context.Context, id int) (*UserInfo, error)
}

func NewUserBiz(data UserRepo) *UserBiz {
	return &UserBiz{
		data: data,
	}
}
