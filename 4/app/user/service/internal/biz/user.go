package biz

import (
	"context"
	"go_crouse/4/app/user/service/internal/data"
	"log"
)

func (ub *UserBiz) AddUser(ctx context.Context) {
	log.Println("biz-Adduser")
	ub.data.AddUser(ctx, data.UserModel{})
}
