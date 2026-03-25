package main

import (
	"userservice/app"
	config "userservice/config/env"
)

func main() {
	config.Load()
	cfg := app.NewConfig()
	app := app.NewApplication(cfg)
	app.Run()
}
