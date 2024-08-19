package models

import "time"

type UserRoom struct {
	ID       uint      `gorm:"primaryKey"`
	UserID   uint      `gorm:"not null"`
	RoomID   uint      `gorm:"not null"`
	JoinedAt time.Time `gorm:"default:current_timestamp"`
	User     User      `gorm:"foreignKey:UserID"`
	Room     Room      `gorm:"foreignKey:RoomID"`
}
