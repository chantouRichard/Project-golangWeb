package routers

import (
	"Online-Theater/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

func SetMovieRoutes(r *gin.Engine, db *gorm.DB) {
	movies := r.Group("/api/movies")
	{
		// 获取电影列表
		movies.GET("", func(c *gin.Context) {
			var movies []models.Movie
			if err := db.Find(&movies).Error; err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"status":  "error",
					"message": err.Error(),
				})
				return
			}
			c.JSON(http.StatusOK, gin.H{"status": "success", "movies": movies})
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

			var movies []models.Movie
			query := db

			if genre != "" {
				query = query.Where("genre = ?", genre)
			}
			if minRating != "" {
				rating, err := strconv.ParseFloat(minRating, 64)
				if err == nil {
					query = query.Where("rating >= ?", rating)
				}
			}

			if err := query.Find(&movies).Error; err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "Failed to filter movies"})
				return
			}
			if len(movies) == 0 {
				c.JSON(http.StatusOK, gin.H{"status": "error", "movies": "没有符合条件的电影"})
			} else {
				c.JSON(http.StatusOK, gin.H{"status": "success", "movies": movies})
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

	}
}
