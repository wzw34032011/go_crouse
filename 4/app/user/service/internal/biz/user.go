package biz

import (
	"context"
	"log"
)

type UserInfo struct {
	Id   int
	Name string
	Age  int8
}

func (ub *UserBiz) AddUser(ctx context.Context, user *UserInfo) error {
	log.Println(user)
	return ub.data.AddUser(ctx, user)
}
