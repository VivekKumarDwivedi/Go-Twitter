package router

import (
	"net/http"
	"TweetService/controllers"

	"github.com/go-chi/chi/v5"
)

type TagRouter struct {
	tagController controllers.TagController
}

func NewTagRouter(_tagController controllers.TagController) Router {
	return &TagRouter{
		tagController: _tagController,
	}
}

func (tr *TagRouter) Register(r chi.Router) {

	//health check
	r.Get("/tag/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	//tag routes
	r.Get("/tag", tr.tagController.GetAllTags)
	r.Get("/tag/tweets", tr.tagController.GetTagWithTweets)
	r.Get("/tag/top", tr.tagController.GetTopTags)
	r.Delete("/tag", tr.tagController.DeleteTag)
}
