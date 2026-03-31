package router

import (
	"net/http"
	"TweetService/controllers"
	"TweetService/middlewares"
	"github.com/go-chi/chi/v5"
)

type TweetRouter struct {
	TweetController controllers.TweetController
}

func NewTweetRouter(_tweetController controllers.TweetController) Router {
	return &TweetRouter{
		TweetController: _tweetController,
	}
}
func (tr *TweetRouter) Register(r chi.Router){

	//health check
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	//tweet routes
	r.With(middlewares.PayloadMiddleware).Post("/tweet", tr.TweetController.CreateTweet)
	r.Get("/tweet", tr.TweetController.GetAllTweets)
	r.Get("/tweet/{id}", tr.TweetController.GetTweetByID)
	r.With(middlewares.PayloadMiddleware).Put("/tweet/{id}", tr.TweetController.UpdateTweet)
	r.Delete("/tweet/{id}", tr.TweetController.DeleteTweet)
}
