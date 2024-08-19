package routers

import (
	"Online-Theater/handlers"
	"Online-Theater/models"
	middleware "Online-Theater/package"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

func SetMovieRoutes(r *gin.Engine, db *gorm.DB) {
	movies := r.Group("/api/movies")
	movies.Use(middleware.TokenAuthMiddleware())
	{
		// 获取电影列表并增加分页功能
		movies.GET("", func(c *gin.Context) {
			// 获取分页参数，默认为第一页，每页10条数据
			page := c.DefaultQuery("page", "1")
			pageSize := c.DefaultQuery("page_size", "10")

			// 将分页参数转换为整数类型
			pageInt, err := strconv.Atoi(page)
			if err != nil || pageInt < 1 {
				pageInt = 1
			}

			pageSizeInt, err := strconv.Atoi(pageSize)
			if err != nil || pageSizeInt < 1 {
				pageSizeInt = 10
			}

			var movies []models.Movie
			var total int64

			// 先获取总记录数
			db.Model(&models.Movie{}).Count(&total)

			// 使用 Offset 和 Limit 实现分页
			db.Offset((pageInt - 1) * pageSizeInt).Limit(pageSizeInt).Find(&movies)

			// 返回分页后的数据以及总记录数
			c.JSON(http.StatusOK, gin.H{
				"status": "success",
				"movies": movies,
				"total":  total,
			})
		})

		// 获取单个电影的详细信息
		movies.GET("/:id", func(c *gin.Context) {
			id := c.Param("id")
			var movie models.Movie
			if err := db.First(&movie, id).Error; err != nil {
				c.JSON(http.StatusNotFound, gin.H{"status": "error", "message": "Movie not found"})
				return
			}
			c.JSON(http.StatusOK, gin.H{"status": "success", "movie": movie})
		})

		//获取电影封面图片
		movies.GET("/picture/:id", func(c *gin.Context) {
			id := c.Param("id")

			// 查询电影记录
			var movie models.Movie
			if err := db.Where("id = ?", id).First(&movie).Error; err != nil {
				c.JSON(http.StatusNotFound, gin.H{"error": "电影未找到"})
				return
			}

			// 调用 StreamPicture 处理图片流
			picturePath := movie.ThumbnailURL // 从电影记录中获取图片路径
			handlers.StreamPicture(c, picturePath)
		})

		// 搜索电影
		movies.GET("/search", func(c *gin.Context) {
			query := c.Query("query")
			var movies []models.Movie
			if err := db.Where("title LIKE ?", "%"+query+"%").Find(&movies).Error; err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "Failed to search movies"})
				return
			}
			if len(movies) == 0 {
				c.JSON(http.StatusOK, gin.H{"status": "error", "movies": "没有符合条件的电影"})
			} else {
				c.JSON(http.StatusOK, gin.H{"status": "success", "movies": movies})
			}
		})

		// 筛选电影
		movies.GET("/filter", func(c *gin.Context) {
			genre := c.Query("genre")
			minRating := c.Query("min_rating")

			page := c.DefaultQuery("page", "1")
			pageSize := c.DefaultQuery("page_size", "10")

			// 将 page 和 pageSize 转换为 int
			pageInt, err := strconv.Atoi(page)
			if err != nil || pageInt < 1 {
				pageInt = 1
			}

			pageSizeInt, err := strconv.Atoi(pageSize)
			if err != nil || pageSizeInt < 1 {
				pageSizeInt = 10
			}

			// 计算 Offset
			offset := (pageInt - 1) * pageSizeInt

			var movies []models.Movie
			query := db

			// 根据 genre 和 minRating 进行筛选
			if genre != "" {
				query = query.Where("genre = ?", genre)
			}
			if minRating != "" {
				rating, err := strconv.ParseFloat(minRating, 64)
				if err == nil {
					query = query.Where("rating >= ?", rating)
				}
			}

			// 获取总记录数
			var total int64
			if err := query.Model(&models.Movie{}).Count(&total).Error; err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "Failed to count movies"})
				return
			}

			// 添加分页
			if err := query.Limit(pageSizeInt).Offset(offset).Find(&movies).Error; err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "Failed to filter movies"})
				return
			}

			if len(movies) == 0 {
				c.JSON(http.StatusOK, gin.H{"status": "error", "movies": "没有符合条件的电影"})
			} else {
				c.JSON(http.StatusOK, gin.H{
					"status":    "success",
					"movies":    movies,
					"total":     total,       // 总记录数
					"page":      pageInt,     // 当前页码
					"page_size": pageSizeInt, // 每页的条数
				})
			}
		})

		//加入收藏
		movies.POST("/collections", func(c *gin.Context) {
			userIDStr := c.PostForm("user_id")
			movieIDStr := c.PostForm("movie_id")

			if userIDStr == "" || movieIDStr == "" {
				c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "User ID and Movie ID are required"})
				return
			}

			userID, err := strconv.ParseUint(userIDStr, 10, 32)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Invalid user ID"})
				return
			}

			movieID, err := strconv.ParseUint(movieIDStr, 10, 32)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Invalid movie ID"})
				return
			}

			// Check if the collection already exists
			var collection models.Collection
			if err := db.Where("user_id = ? AND movie_id = ?", userID, movieID).First(&collection).Error; err == nil {
				c.JSON(http.StatusConflict, gin.H{"status": "error", "message": "Movie already in collection"})
				return
			}

			// Create new collection entry
			newCollection := models.Collection{
				UserID:  uint(userID),
				MovieID: uint(movieID),
			}
			if err := db.Create(&newCollection).Error; err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "Failed to add movie to collection"})
				return
			}
			c.JSON(http.StatusOK, gin.H{"status": "success", "message": "Movie added to collection"})
		})

		//搜索电影
		movies.GET("/search", func(c *gin.Context) {
			query := c.Query("query")
			mode := c.Query("mode")

			if query == "" {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Query parameter is required"})
				return
			}

			var movies []models.Movie
			if mode == "exact" {
				// 精确搜索
				db.Where("name = ?", query).Find(&movies)
			} else {
				// 模糊搜索
				db.Where("name LIKE ?", "%"+query+"%").Find(&movies)
			}

			c.JSON(http.StatusOK, gin.H{"movies": movies})
		})
	}

}
