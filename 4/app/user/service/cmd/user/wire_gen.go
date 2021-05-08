// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package main

import (
	"go_crouse/4/app"
	"go_crouse/4/app/user/service/internal/biz"
	"go_crouse/4/app/user/service/internal/data"
	"go_crouse/4/app/user/service/internal/server"
	"go_crouse/4/app/user/service/internal/service"
)

// Injectors from wire.go:

func initApp(version string) *app.App {
	userData := data.NewUserData()
	userBiz := biz.NewUserBiz(userData)
	userService := service.NewUserService(userBiz)
	server1 := server.NewService1(userService)
	server_2 := server.NewService2()
	appApp := NewApp(version, server1, server_2)
	return appApp
}