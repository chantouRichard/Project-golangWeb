package models

import "time"

type UserFavoriteTag struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	TagID     int       `json:"tag_id"`
	CreatedAt time.Time `json:"created_at"`
}
