package controllers

import (
	"Online-Theater/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"regexp"
)

func RegisterRouter(c *gin.Context) {
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
}
