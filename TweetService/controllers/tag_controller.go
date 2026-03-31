package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"TweetService/services"
	"TweetService/utils"
)
type TagController interface{
    GetAllTags(w http.ResponseWriter, r *http.Request)
    GetTagWithTweets(w http.ResponseWriter, r *http.Request)
    GetTopTags(w http.ResponseWriter, r *http.Request)
    DeleteTag(w http.ResponseWriter, r *http.Request)
}
 
type TagControllerImpl struct{
    TagService services.TagService
}
 
func NewTagController(tagService services.TagService) TagController {
    return &TagControllerImpl{
        TagService: tagService,
    }
}

func (c *TagControllerImpl) GetAllTags(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all tags called in tag controller")

	limit,_ := strconv.Atoi(r.URL.Query().Get("limit"))
	offset,_ := strconv.Atoi(r.URL.Query().Get("offset"))

	resp,err := c.TagService.GetAllTags(limit,offset)

	if err !=nil{
		utils.WriteJsonErrorResponse(w,http.StatusInternalServerError,"Failed to fetch tag",err)
		return
	}

	utils.WriteJsonSuccessResponse(w,http.StatusOK,"Tags fetched successfully",resp)
}

func (c *TagControllerImpl) GetTagWithTweets(w http.ResponseWriter, r *http.Request){
	fmt.Println("GetTag With tweets called in tag controller")

	name := r.URL.Query().Get("name")

	if name == ""{
		utils.WriteJsonErrorResponse(w,http.StatusBadRequest,"Tag name is required",fmt.Errorf("missing name"))
		return
	}
	
	resp, err := c.TagService.GetTagWithTweets(name)
	
	if err != nil {
		utils.WriteJsonErrorResponse(w,http.StatusInternalServerError,"Failed to fetch tag with tweets",err)
		return
	}
	
	utils.WriteJsonSuccessResponse(w,http.StatusOK,"Tag with tweets fetched successfully",resp)
}

func (c *TagControllerImpl) GetTopTags(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GetTopTags called in tag controller")
	
	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))
	
	resp, err := c.TagService.GetTopTags(limit)
	
	if err != nil {
		utils.WriteJsonErrorResponse(w, http.StatusInternalServerError, "Failed to fetch top tags", err)
		return
	}
	
	utils.WriteJsonSuccessResponse(w, http.StatusOK, "Top tags fetched successfully", resp)
}

func (c *TagControllerImpl) DeleteTag(w http.ResponseWriter, r *http.Request) {
	fmt.Println("DeleteTag called in tag controller")
	
	idStr := r.URL.Query().Get("id")
	id, _ := strconv.Atoi(idStr)
	
	err := c.TagService.DeleteTag(id)
	
	if err != nil {
		utils.WriteJsonErrorResponse(w, http.StatusInternalServerError, "Failed to delete tag", err)
		return
	}
	
	utils.WriteJsonSuccessResponse(w, http.StatusOK, "Tag deleted successfully", nil)
}
