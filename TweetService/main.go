package main

import (
	"TweetService/app"
	config "TweetService/config/env"
)

func main(){

	config.Load()
	cfg := app.NewConfig()
    app := app.NewApplication(cfg)

	app.Run()
}