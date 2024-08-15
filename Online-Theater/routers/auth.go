package routers

import (
	"Online-Theater/models"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"net/http"
	"regexp"
	"time"
)

var jwtSecret = []byte("your-secret-key")

func SetUserRoutes(r *gin.Engine, db *gorm.DB) {
	auth := r.Group("/api")
	{
		//注册功能
		auth.POST("/register", func(c *gin.Context) {
			var user models.User

			if err := c.ShouldBindJSON(&user); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": err.Error()})
				return
			}

			// 验证用户名长度
			if len(user.Username) < 5 || len(user.Username) > 15 {
				c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "用户名必须为5~15个字符"})
				return
			}

			// 验证密码长度
			if len(user.Password) < 5 || len(user.Password) > 15 {
				c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "密码必须为5~15个字符"})
				return
			}

			// 验证邮箱格式
			emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
			if matched, _ := regexp.MatchString(emailRegex, user.Email); !matched {
				c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "邮箱格式不正确"})
				return
			}

			// 检查用户名是否已存在
			var existingUser models.User
			if err := db.Where("username = ?", user.Username).First(&existingUser).Error; err == nil {
				c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "用户名已存在"})
				return
			}

			// 检查邮箱是否已存在
			if err := db.Where("email = ?", user.Email).First(&existingUser).Error; err == nil {
				c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "邮箱已被使用"})
				return
			}

			// 哈希密码
			hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": ""})
				return
			}
			user.Password = string(hashedPassword)

			// 保存用户到数据库
			if err := db.Create(&user).Error; err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "注册失败"})
				return
			}
			c.JSON(http.StatusOK, gin.H{"status": "success", "message": "注册成功"})
		})

		//登录功能
		auth.POST("/login", func(c *gin.Context) {
			var loginDetails struct {
				Email    string `json:"email"`
				Password string `json:"password"`
			}

			if err := c.ShouldBindJSON(&loginDetails); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": err.Error()})
				return
			}

			var user models.User
			if err := db.Where("email = ?", loginDetails.Email).First(&user).Error; err != nil {
				c.JSON(http.StatusUnauthorized, gin.H{"status": "error", "message": "邮箱或者密码错误"})
				return
			}

			if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginDetails.Password)); err != nil {
				c.JSON(http.StatusUnauthorized, gin.H{"status": "error", "message": "邮箱或者密码错误"})
				return
			}

			token, err := generateToken(user.ID, user.Username, user.Email)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "产生token失败"})
				return
			}

			c.JSON(http.StatusOK, gin.H{
				"status": "success",
				"token":  token,
				"user": gin.H{
					"id":       user.ID,
					"username": user.Username,
					"email":    user.Email,
				},
			})
		})

		//获取用户信息
		auth.GET("/user", func(c *gin.Context) {
			tokenStr := c.GetHeader("Authorization")
			token, claims, err := parseToken(tokenStr)
			if err != nil || !token.Valid {
				c.JSON(http.StatusUnauthorized, gin.H{"status": "error", "message": "无效的token"})
				return
			}

			var user models.User
			if err := db.First(&user, claims.Id).Error; err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "不存在该用户"})
				return
			}

			c.JSON(http.StatusOK, gin.H{
				"status": "success",
				"user": gin.H{
					"id":       user.ID,
					"username": user.Username,
					"email":    user.Email,
				},
			})
		})

		//更新用户信息
		auth.PUT("/user", func(c *gin.Context) {
			tokenStr := c.GetHeader("Authorization")
			token, claims, err := parseToken(tokenStr)
			if err != nil || !token.Valid {
				c.JSON(http.StatusUnauthorized, gin.H{"status": "error", "message": "Invalid token"})
				return
			}

			var updateDetails struct {
				Username string `json:"username"`
				Email    string `json:"email"`
				Password string `json:"password"`
				Nickname string `json:"nickname"`
				Phone    string `json:"phone"`
			}

			if err := c.ShouldBindJSON(&updateDetails); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": err.Error()})
				return
			}

			// 验证新的邮箱格式
			emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
			if matched, _ := regexp.MatchString(emailRegex, updateDetails.Email); !matched {
				c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "邮箱格式错误"})
				return
			}

			var user models.User
			if err := db.First(&user, claims.Id).Error; err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "用户获取失败"})
				return
			}

			// 更新用户名和邮箱
			user.Username = updateDetails.Username
			user.Email = updateDetails.Email
			user.Password = updateDetails.Password
			user.Nickname = updateDetails.Nickname
			user.Phone = updateDetails.Phone

			if err := db.Save(&user).Error; err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "用户更新失败"})
				return
			}

			c.JSON(http.StatusOK, gin.H{"status": "success", "message": "更新成功"})
		})
	}
}

// 产生登录的token
func generateToken(userID int, username, email string) (string, error) {
	claims := &jwt.StandardClaims{
		Id:        fmt.Sprintf("%d", userID),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

// 解析token
func parseToken(tokenStr string) (*jwt.Token, *jwt.StandardClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if err != nil {
		return nil, nil, err
	}
	claims, ok := token.Claims.(*jwt.StandardClaims)
	if !ok {
		return nil, nil, fmt.Errorf("invalid token claims")
	}
	return token, claims, nil
}
