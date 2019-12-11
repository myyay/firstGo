package main

import (
	"firstGo/superstar/bootstrap"
	"firstGo/superstar/web/middleware/identity"
	"firstGo/superstar/web/routes"
)

func newApp() *bootstrap.Bootstrapper {
	app := bootstrap.New("Superstar database", "Me")
	app.Bootstrap()
	app.Configure(identity.Configure, routes.Configure)
	return app
}

func main() {
	app := newApp()
	app.Listen(":8080")
}
