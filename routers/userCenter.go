package routers

import (
	"Online-Theater/models"
	middleware "Online-Theater/package"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"net/http"
	"strconv"
	"strings"
)

// 个人中心路由

func SetupUserCenterRoutes(r *gin.Engine, db *gorm.DB) {
	userCenter := r.Group("/api/user")
	userCenter.Use(middleware.TokenAuthMiddleware())

	// 修改密码
	userCenter.POST("/updatepwd", func(c *gin.Context) {
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

		// 获取用户 ID
		userID := claims.Id
		var user models.User
		if err := db.First(&user, userID).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"status": "error", "message": "User not found"})
			return
		}

		// 处理密码更改逻辑
		var requestData struct {
			OldPassword string `json:"old_pwd"`
			NewPassword string `json:"new_pwd"`
			RePassword  string `json:"re_pwd"`
		}
		if err := c.ShouldBindJSON(&requestData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": err.Error()})
			return
		}

		// 验证旧密码
		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(requestData.OldPassword)); err != nil {
			c.JSON(400, gin.H{"status": "error", "message": "Old password is incorrect"})
			return
		}

		// 哈希处理新密码
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(requestData.NewPassword), bcrypt.DefaultCost)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "Failed to hash new password"})
			return
		}
		user.Password = string(hashedPassword)

		if requestData.RePassword != requestData.NewPassword {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "error", "message": "New password is incorrect"})
			return
		}

		// 更新用户密码
		if err := db.Save(&user).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "Failed to update password"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "success", "message": "Password changed successfully"})

	})

	// 我的收藏
	userCenter.GET("/collections", func(c *gin.Context) {
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

		// 获取用户 ID
		userID := claims.Id
		var user models.User
		if err := db.First(&user, userID).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"status": "error", "message": err.Error()})
			return
		}

		var collections []models.Collection
		if err := db.Preload("Movie").Where("user_id = ?", userID).Find(&collections).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "Failed to fetch collections"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "success", "collections": collections})
	})

	// 取消收藏
	userCenter.DELETE("/collections/:id", func(c *gin.Context) {
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

		// 获取用户 ID
		userID := claims.Id
		var user models.User
		if err := db.First(&user, userID).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"status": "error", "message": err.Error()})
			return
		}

		// 获取收藏项的 ID
		collectionID := c.Param("id")

		// 查找收藏记录
		var collection models.Collection
		if err := db.Where("id = ? AND user_id = ?", collectionID, userID).First(&collection).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"status": "error", "message": "Collection not found"})
			return
		}

		// 删除收藏记录
		if err := db.Delete(&collection).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "Failed to delete collection"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "success", "message": "Collection deleted successfully"})
	})

	// 我的观看历史
	userCenter.GET("/history", func(c *gin.Context) {
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

		// 获取用户 ID
		userID := claims.Id
		var histories []models.History
		if err := db.Preload("Movie").Where("user_id = ?", userID).Find(&histories).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "Failed to fetch history"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "success", "histories": histories})
	})

	// 删除所有观看历史记录
	userCenter.DELETE("/history", func(c *gin.Context) {
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

		// 获取用户 ID
		userID := claims.Id
		var user models.User
		if err := db.First(&user, userID).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"status": "error", "message": err.Error()})
			return
		}

		// 删除用户的所有观看历史记录
		if err := db.Where("user_id = ?", userID).Delete(&models.History{}).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "Failed to delete history"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "success", "message": "All history records deleted successfully"})
	})

	// 添加观看历史记录
	userCenter.POST("/history", func(c *gin.Context) {
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

		// 将 userID 从字符串转换为 uint
		userIDStr := claims.Id
		userID, err := strconv.ParseUint(userIDStr, 10, 32)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "Failed to parse user ID"})
			return
		}

		// 获取传入的 movie_id 参数
		var input struct {
			MovieID uint `json:"movie_id" binding:"required"`
		}
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Movie ID is required"})
			return
		}

		// 检查电影是否存在
		var movie models.Movie
		if err := db.First(&movie, input.MovieID).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"status": "error", "message": "Movie not found"})
			return
		}

		// 创建新的历史记录
		history := models.History{
			UserID:  uint(userID), // 这里将 uint64 转换为 uint
			MovieID: input.MovieID,
		}

		// 保存历史记录到数据库
		if err := db.Create(&history).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "Failed to add history"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "success", "message": "History added successfully"})
	})
}
