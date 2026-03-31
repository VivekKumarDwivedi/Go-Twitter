package router

import (
	"TweetService/controllers"

	"github.com/go-chi/chi/v5"
)

type Router interface {
	Register(r chi.Router)
}

// setup Router wires up all feature routers

func SetupRouter(tweetRouter Router, tagRouter Router) *chi.Mux {
	chiRouter := chi.NewRouter()

	chiRouter.Get("/ping", controllers.PingHandler)
	
	// Register all feature routers
	tweetRouter.Register(chiRouter)
	tagRouter.Register(chiRouter)
	
	return chiRouter
}