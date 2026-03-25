package app

import (
	"fmt"
	"net/http"
	dbConfig "userservice/config/db"

	config "userservice/config/env"
)

// config holds the configuration for the server

type Config struct {
	Addr string
}

// create server details
type Application struct {
	Config Config
}

// constructor for config

func NewConfig() Config {

	port := config.GetString("PORT", ":8080")

	return Config{
		Addr: port,
	}
}

func NewApplication(cfg Config) *Application {
	return &Application{
		Config: cfg,
	}
}

//member function

func (app *Application) Run() error {
	db, err := dbConfig.SetupDB()

	if err != nil {
		fmt.Println("Error setting up database:", err)
		return err
	}
	fmt.Println("db:", db)
	fmt.Println("Starting server on", app.Config.Addr)
	server := &http.Server{
		Addr: app.Config.Addr,
	}
	return server.ListenAndServe()
}
