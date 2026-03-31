package app 


import (
	dbConfig "TweetService/config/db"
    config "TweetService/config/env"
    "TweetService/controllers"
    repo "TweetService/db/repositories"
    "TweetService/router"
    "TweetService/services"
    "fmt"
    "net/http"
    "strings"
    "time"
)

// Config hold server configuration

type Config struct {
	Addr string
}

// Application struct

type Application struct {
	Config Config
}

// Constructor For Config

func NewConfig() Config {
	port := config.GetString("PORT",":8081")

	if !strings.HasPrefix(port, ":") {
		port = ":" + port
	}
	return Config{
		Addr: port,
	}
}

// Constructor for Application 
func NewApplication(cfg Config) *Application {
	return &Application{
		Config:cfg,
	}
}

// Run starts the server
func (app *Application) Run() error{
	// step 1: Setup DB

	db, err := dbConfig.SetupDB()
	if err != nil {
		fmt.Println("Error setting up database:",err)
		return err
	}

	// step 2: setup repositories
	tweetRepo := repo.NewTweetRepository(db)
	tagRepo := repo.NewTagRepository(db)
	
	// step 3 : setup services
	tweetService := services.NewTweetService(tweetRepo, tagRepo)
	tagService := services.NewTagService(tagRepo)

	// step 4: Setup controllers
	tweetController := controllers.NewTweetController(tweetService)
	tagController := controllers.NewTagController(tagService)

	// step 5: Setup router
	tweetRouter := router.NewTweetRouter(tweetController)
	tagRouter := router.NewTagRouter(tagController)
	router := router.SetupRouter(tweetRouter, tagRouter)

	// step 6: Start server
	server := &http.Server{
		Addr: app.Config.Addr,
		Handler: router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	fmt.Println("Server starting on", app.Config.Addr)
	return server.ListenAndServe()
}
