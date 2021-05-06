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

func main() {

	app := app.NewApp(
		Name,
		Version,
		app.AddMeta("env", "test"),
		app.AddServer(
			server.NewService1(),
			server.NewService2(),
		),
	)

	app.Run()
}
