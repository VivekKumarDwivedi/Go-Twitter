package models

type Tweet struct {
	ID        int    `json:"id"`
	UserID    int    `json:"user_id"`
	Tweet     string `json:"tweet"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

