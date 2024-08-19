package handlers

import (
	"Online-Theater/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

// Database model for Room (assuming you have this)

type Room struct {
	ID        uint        `gorm:"primaryKey"`
	MovieID   uint        `gorm:"not null"`
	CreatorID uint        `gorm:"not null"`
	RoomName  string      `gorm:"size:255"`
	CreatedAt time.Time   `gorm:"default:current_timestamp"`
	Movie     Movie       `gorm:"foreignKey:MovieID"`
	Creator   models.User `gorm:"foreignKey:CreatorID"`
}

// Database model for Movie (assuming you have this)

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

func StreamMovie(c *gin.Context, roomID string, db *gorm.DB) {
	// 根据 roomID 查询对应的 room 和 movie
	var room Room
	if err := db.Where("id = ?", roomID).First(&room).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "房间未找到"})
		return
	}

	var movie Movie
	if err := db.Where("id = ?", room.MovieID).First(&movie).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "无法找到电影"})
		return
	}

	// 打开电影文件
	file, err := os.Open(movie.VideoURL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "无法打开电影"})
		return
	}
	defer file.Close()

	// 获取文件信息
	fileInfo, err := file.Stat()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "无法获取电影信息"})
		return
	}

	// 设置响应头
	c.Writer.Header().Set("Content-Type", "video/mp4")
	c.Writer.Header().Set("Content-Length", fmt.Sprintf("%d", fileInfo.Size()))
	c.Writer.Header().Set("Accept-Ranges", "bytes")

	// 处理 HTTP Range 请求
	http.ServeContent(c.Writer, c.Request, filepath.Base(movie.VideoURL), fileInfo.ModTime(), file)
}
