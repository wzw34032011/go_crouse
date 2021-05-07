package biz

import "go_crouse/4/app/user/service/internal/data"

type UserBiz struct {
	data *data.UserData
}

func NewUserBiz(data *data.UserData) *UserBiz {
	return &UserBiz{
		data: data,
	}
}
