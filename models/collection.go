package models

// 收藏模型

type Collection struct {
	ID      uint `gorm:"primaryKey"`
	UserID  uint
	MovieID uint
	Movie   Movie `gorm:"foreignKey:MovieID"`
}
