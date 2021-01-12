package entity

import "time"

type Score struct {
	Id        int       `json:"id"`
	UserName  string    `json:"username"`
	Score     int64     `json:"score"`
	CreatedAt time.Time `json:"created_at"`
}
