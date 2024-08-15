package models

import "time"

type Room struct {
	ID        uint      `gorm:"primaryKey"`
	MovieID   uint      `gorm:"not null"`
	CreatorID uint      `gorm:"not null"`
	RoomName  string    `gorm:"size:255"`
	CreatedAt time.Time `gorm:"default:current_timestamp"`
	Movie     Movie     `gorm:"foreignKey:MovieID"`
	Creator   User      `gorm:"foreignKey:CreatorID"`
}
