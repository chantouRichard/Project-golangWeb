package routers

import (
	"Online-Theater/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func SetRoomRoutes(r *gin.Engine, db *gorm.DB) {
	rooms := r.Group("/api/rooms")
	{
		// 创建房间
		rooms.POST("", func(c *gin.Context) {
			// 从 Authorization 头中提取 Token
			tokenStr := c.GetHeader("Authorization")
			if tokenStr == "" {
				c.JSON(http.StatusUnauthorized, gin.H{"status": "error", "message": "Token is required"})
				return
			}

			// 去掉 'Bearer ' 前缀
			tokenStr = strings.TrimPrefix(tokenStr, "Bearer ")

			// 解析 Token 并获取用户 ID
			token, claims, err := parseToken(tokenStr)
			if err != nil || !token.Valid {
				c.JSON(http.StatusUnauthorized, gin.H{"status": "error", "message": "Invalid token"})
				return
			}

			// 获取用户 ID，假设 claims.Id 是字符串，需要转换成 uint 类型
			userID, err := strconv.ParseUint(claims.Id, 10, 32)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "Failed to parse user ID"})
				return
			}
			// 将解析到的 userID 赋值给 req.UserID

			var req struct {
				MovieID  uint   `json:"movie_id"`
				RoomName string `json:"room_name"`
			}

			if err := c.BindJSON(&req); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": err.Error()})
				return
			}

			creatorID := uint(userID) // 自定义CreatorID，这里假设为123

			room := models.Room{
				MovieID:   req.MovieID,
				CreatorID: creatorID, // 使用自定义的CreatorID
				RoomName:  req.RoomName,
				CreatedAt: time.Now(),
			}

			if err := db.Create(&room).Error; err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": err.Error()})
				return
			}

			c.JSON(http.StatusOK, gin.H{"status": "success", "room": room})
		})

		// 获取电影的房间列表
		rooms.GET("/movie/:movie_id", func(c *gin.Context) {
			movieID := c.Param("movie_id")
			var rooms []models.Room

			// 预加载 Movie 和 Creator
			if err := db.Preload("Movie").Preload("Creator").Where("movie_id = ?", movieID).Find(&rooms).Error; err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "Failed to fetch rooms"})
				return
			}

			// 转换 Room 数据为 RoomDTO
			var roomsWithDTO []models.RoomDTO
			for _, room := range rooms {
				roomDTO := models.RoomDTO{
					ID:        room.ID,
					MovieID:   room.MovieID,
					CreatorID: room.CreatorID,
					RoomName:  room.RoomName,
					CreatedAt: room.CreatedAt,
					Movie:     room.Movie,
					Creator: models.UserDTO{
						ID:   uint(room.Creator.ID),
						Name: room.Creator.Username,
					},
				}
				roomsWithDTO = append(roomsWithDTO, roomDTO)
			}

			c.JSON(http.StatusOK, gin.H{"status": "success", "rooms": roomsWithDTO})
		})

		// 用户加入房间
		rooms.POST("/join", func(c *gin.Context) {
			var req struct {
				UserID uint `json:"user_id"`
				RoomID uint `json:"room_id"`
			}
			if err := c.BindJSON(&req); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": err.Error()})
				return
			}

			userRoom := models.UserRoom{
				UserID: req.UserID,
				RoomID: req.RoomID,
			}

			// Check if the user already joined the room
			if err := db.Where("user_id = ? AND room_id = ?", req.UserID, req.RoomID).First(&userRoom).Error; err == nil {
				c.JSON(http.StatusConflict, gin.H{"status": "error", "message": "User already in room"})
				return
			}

			if err := db.Create(&userRoom).Error; err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": err.Error()})
				return
			}

			c.JSON(http.StatusOK, gin.H{"status": "success", "message": "Joined room successfully"})
		})

		// 关闭房间
		rooms.DELETE("/:room_id", func(c *gin.Context) {
			// 从 Authorization 头中提取 Token
			tokenStr := c.GetHeader("Authorization")
			if tokenStr == "" {
				c.JSON(http.StatusUnauthorized, gin.H{"status": "error", "message": "Token is required"})
				return
			}

			// 去掉 'Bearer ' 前缀
			tokenStr = strings.TrimPrefix(tokenStr, "Bearer ")

			// 解析 Token 并获取用户 ID
			token, claims, err := parseToken(tokenStr)
			if err != nil || !token.Valid {
				c.JSON(http.StatusUnauthorized, gin.H{"status": "error", "message": "Invalid token"})
				return
			}

			// 获取用户 ID，假设 claims.Id 是字符串，需要转换成 uint 类型
			userID, err := strconv.ParseUint(claims.Id, 10, 32)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "Failed to parse user ID"})
				return
			}

			// 定义请求体结构体
			var req struct {
				UserID uint `json:"user_id"`
			}

			// 将解析到的 userID 赋值给 req.UserID
			req.UserID = uint(userID)

			if err := c.BindJSON(&req); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": err.Error()})
				return
			}

			roomID := c.Param("room_id")

			// 开启事务，确保数据一致性
			tx := db.Begin()

			// 删除 user_rooms 表中与该房间相关的所有记录
			if err := tx.Where("room_id = ?", roomID).Delete(&models.UserRoom{}).Error; err != nil {
				tx.Rollback()
				c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": err.Error()})
				return
			}

			// 删除 rooms 表中的房间记录
			if err := tx.Where("id = ?", roomID).Delete(&models.Room{}).Error; err != nil {
				tx.Rollback()
				c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": err.Error()})
				return
			}

			// 如果有其他关联表，也在这里删除相应的记录
			// 例如删除聊天记录:
			// if err := tx.Where("room_id = ?", roomID).Delete(&models.ChatMessage{}).Error; err != nil {
			//     tx.Rollback()
			//     c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "Failed to delete associated chat messages"})
			//     return
			// }

			// 提交事务
			tx.Commit()

			c.JSON(http.StatusOK, gin.H{"status": "success", "message": "关闭房间成功"})
		})

	}
}
