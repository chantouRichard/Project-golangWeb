package models

// 中间表MovieTag用于表示电影与标签的多对多关系

type MovieTag struct {
	ID      int `json:"id"`
	MovieID int `json:"movie_id"`
	TagID   int `json:"tag_id"`
}
