package dto

type TweetResponseDTO struct {
	ID int `json:"id"`
	UserID int `json:"user_id"`
	Tweet string `json:"tweet"`
	Tags []string `json:"tags,omitempty"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type CreateTweetRequestDTO struct {
	UserID int    `json:"user_id"`
	Tweet  string `json:"tweet"`
}

type UpdateTweetRequestDTO struct {
	Tweet string `json:"tweet"`
}
