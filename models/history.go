package models

// 观看历史模型

type History struct {
	ID      uint `gorm:"primaryKey"`
	UserID  uint
	MovieID uint
	Movie   Movie `gorm:"foreignKey:MovieID"`
}
