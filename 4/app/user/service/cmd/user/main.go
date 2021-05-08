package main

import (
	"flag"
	"go_crouse/4/app"
	"go_crouse/4/app/user/service/internal/server"
)

// go build -ldflags "-X main.Version=x.y.z"
var (
	// Name is the name of the compiled software.
	Name string
	// Version is the version of the compiled software.
	Version string
	// flagconf is the config flag.
	flagconf string
)

func init() {
	flag.StringVar(&flagconf, "conf", "../../configs", "config path, eg: -conf config.yaml")
}

func NewApp(version string, server1 *server.Server1, server2 *server.Server_2) *app.App {
	return app.NewApp(version, app.AddServer(server1, server2))
}

func main() {
	//ud := data.NewUserData()
	//ub := biz.NewUserBiz(ud)
	//us := service.NewUserService(ub)

	//app := app.NewApp(
	//	Name,
	//	Version,
	//	app.AddMeta("env", "test"),
	//	app.AddServer(
	//		server.NewService1(us),
	//		server.NewService2(),
	//	),
	//)
	app := initApp(Version)
	app.Run()
}
