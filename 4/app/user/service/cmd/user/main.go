package main

import (
	"go_crouse/4/app"
	"go_crouse/4/app/user/service/internal/server"
)

// go build -ldflags "-X main.Version=x.y.z"
var (
	// Name is the name of the compiled software.
	Name string
)

func NewApp(name string, server1 *server.Server1, server2 *server.Server_2) *app.App {
	return app.NewApp(name, app.AddServer(server1, server2))
}

func main() {
	app := initApp("my app")
	app.Run()
}
