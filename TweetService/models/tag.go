package models


type Tag struct {
    ID        int     `json:"id"`
    Name      string    `json:"name"`
    Count     int     `json:"count,omitempty"` // optional: used when fetching top tags
    CreatedAt string    `json:"created_at,omitempty"`
    UpdatedAt string    `json:"updated_at,omitempty"`
}
type TweetTag struct {
    TweetID int `json:"tweet_id"`
    TagID   int `json:"tag_id"`
}