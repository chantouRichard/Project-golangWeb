//package models
//
//import "gorm.io/gorm"
//
//type Movie struct {
//	gorm.Model
//	Title       string  `json:"title"`
//	Description string  `json:"description"`
//	ReleaseDate string  `json:"release_date"`
//	Genre       string  `json:"genre"`
//	Rating      float64 `json:"rating"`
//	PosterURL   string  `json:"poster_url"`
//}

package models

import (
	"gorm.io/gorm"
	"time"
)

type Movie struct {
	gorm.Model
	ID           int     `gorm:"primaryKey;autoIncrement:false"`
	Title        string  `json:"title"`
	Description  string  `json:"description"`
	ThumbnailURL string  `json:"thumbnail_url"`
	VideoURL     string  `json:"video_url"`
	Genre        string  `json:"genre"`
	Rating       float64 `json:"rating"`
	DeletedAt    time.Time
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
