package routers

import (
	"Online-Theater/models"
	middleware "Online-Theater/package"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"net/http"
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
}
