package dto

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

type TagWithTweetsResponseDTO struct {
    ID     int                 `json:"id"`
    Name   string              `json:"name"`
    Tweets []TweetResponseDTO  `json:"tweets"`
}

