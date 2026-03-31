package controllers

import(
	"fmt"
	"net/http"
	"strconv"
	"TweetService/services"
	"TweetService/dto"
	"TweetService/middlewares"
	"TweetService/utils"
	"github.com/go-chi/chi/v5"
)

type TweetController interface{
	CreateTweet(w http.ResponseWriter, r *http.Request)
	GetTweetByID(w http.ResponseWriter, r *http.Request)
	GetAllTweets(w http.ResponseWriter, r *http.Request)
	UpdateTweet(w http.ResponseWriter, r *http.Request)
	DeleteTweet(w http.ResponseWriter, r *http.Request)
}

type TweetControllerImpl struct {
	TweetService services.TweetService
}

func NewTweetController(_tweetService services.TweetService) TweetController {
	return &TweetControllerImpl{
		TweetService: _tweetService,
	}
}

func (c *TweetControllerImpl) CreateTweet(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create tweet called in tweet controller")

	payload, ok := r.Context().Value(middlewares.PayloadKey).(dto.CreateTweetRequestDTO)

	if !ok {
		utils.WriteJsonErrorResponse(w, http.StatusBadRequest, "Invalid request payload", fmt.Errorf("payload missing"))
		return
	}

	resp, err := c.TweetService.CreateTweet(&payload)
	if err != nil {
		utils.WriteJsonErrorResponse(w,http.StatusInternalServerError,"Failed created tweet",err)
		return
	}
	
	utils.WriteJsonSuccessResponse(w, http.StatusCreated, "Tweet created successfully",resp)
	
}

func (c *TweetControllerImpl) GetTweetByID(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get tweet by id called in tweet controller")

	idStr := chi.URLParam(r, "id")
	if idStr == "" {
		utils.WriteJsonErrorResponse(w, http.StatusBadRequest, "Tweet ID is required", fmt.Errorf("id missing"))
		return
	}
	
	id, _ :=strconv.Atoi(idStr)

	resp, err := c.TweetService.GetTweetByID(id)

	if err != nil {
		utils.WriteJsonErrorResponse(w,http.StatusNotFound,"Tweet not found",err)
		return
	}
	utils.WriteJsonSuccessResponse(w, http.StatusOK, "Tweet fetched successfully",resp)
}

func (c *TweetControllerImpl) GetAllTweets(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all tweets called in tweet controller")

	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))
	offset,_ := strconv.Atoi(r.URL.Query().Get("offset"))

	userIDstr := r.URL.Query().Get("user_id")
	hashtag := r.URL.Query().Get("hashtag")
	search := r.URL.Query().Get("q")

	var userID *int 
	if userIDstr != ""{
		id,_ := strconv.Atoi(userIDstr)
		userID = &id
	}

	resp, err := c.TweetService.GetAllTweets(limit, offset, userID, &hashtag, &search)
	if err != nil {
		utils.WriteJsonErrorResponse(w,http.StatusInternalServerError,"Failed to get tweets",err)
		return
	}
	
	utils.WriteJsonSuccessResponse(w, http.StatusOK, "Tweets fetched successfully",resp)
	
}
func (c *TweetControllerImpl) UpdateTweet(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update tweet called in tweet controller")

	idStr := chi.URLParam(r, "id")
	id,_ := strconv.Atoi(idStr)

	payload, ok := r.Context().Value(middlewares.PayloadKey).(dto.UpdateTweetRequestDTO)
	if !ok {
		utils.WriteJsonErrorResponse(w, http.StatusBadRequest, "Invalid request payload", fmt.Errorf("payload missing"))
		return
	}
userID := 0  // Default value
if uid := r.Context().Value(middlewares.UserIDKey); uid != nil {
    if uid, ok := uid.(int); ok {
        userID = uid
        fmt.Println("User ID:", userID)
    }
}
	err := c.TweetService.UpdateTweet(id, &payload)
	if err != nil {
		utils.WriteJsonErrorResponse(w, http.StatusInternalServerError, "Failed to update tweet", err)
		return
	}
	
	utils.WriteJsonSuccessResponse(w, http.StatusOK, "Tweet updated successfully", nil)
	
}
func (c *TweetControllerImpl) DeleteTweet(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete tweet called in tweet controller")

	idStr := chi.URLParam(r, "id")
	id,_ := strconv.Atoi(idStr)

	err := c.TweetService.DeleteTweet(id)
	if err != nil {
		utils.WriteJsonErrorResponse(w, http.StatusInternalServerError, "Failed to delete tweet", err)
		return
	}
	
	utils.WriteJsonSuccessResponse(w, http.StatusOK, "Tweet deleted successfully", nil)
	
}