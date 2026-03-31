package dto

//This is what you return when listing tags or fetching tag details.
type TagResponseDTO struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Count int    `json:"count,omitempty"` // usage count, optional
}

type CreateTagRequestDTO struct {
	Name string `json:"name"`
}

type UpdateTagRequestDTO struct {
	Name string `json:"name"`
}

//Combines tag info with a list of tweets that use it
type TagWithTweetsResponseDTO struct {
	ID     int                `json:"id"`
	Name   string             `json:"name"`
	Tweets []TweetResponseDTO `json:"tweets"`
}
