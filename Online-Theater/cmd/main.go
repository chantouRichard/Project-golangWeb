package main

import (
	"Online-Theater/models"
	"Online-Theater/routers"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func main() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/theater?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database:", err)
	}

	// 自动迁移数据库模式
	db.AutoMigrate(&models.User{})

	// 初始化Gin
	r := gin.Default()

	// 设置路由
	routers.SetUserRoutes(r, db)
	routers.SetMovieRoutes(r, db)
	routers.SetupUserCenterRoutes(r, db)
	routers.SetRoomRoutes(r, db)

	// 启动服务器
	r.Run(":8080")
}
