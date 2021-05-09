// +build wireinject

package main

import (
	"github.com/google/wire"
	"go_crouse/4/app"
	"go_crouse/4/app/user/service/internal/biz"
	"go_crouse/4/app/user/service/internal/data"
	"go_crouse/4/app/user/service/internal/server"
	"go_crouse/4/app/user/service/internal/service"
)

func initApp(name string) *app.App {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, NewApp))
}
