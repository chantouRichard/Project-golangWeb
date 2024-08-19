// models/room_dto.go

package models

import "time"

type RoomDTO struct {
	ID        uint      `json:"id"`
	MovieID   uint      `json:"movie_id"`
	CreatorID uint      `json:"creator_id"`
	RoomName  string    `json:"room_name"`
	CreatedAt time.Time `json:"created_at"`
	Movie     Movie     `json:"movie"`
	Creator   UserDTO   `json:"creator"`
}
