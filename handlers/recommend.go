package handlers

import (
	"Online-Theater/models"
	"gorm.io/gorm"
	"sort"
)

type Tag struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// 获取用户喜欢的标签
func getUserFavoriteTags(db *gorm.DB, userID int) ([]Tag, error) {
	var tags []Tag
	err := db.Table("tags").
		Select("tags.id, tags.name").
		Joins("join user_favorite_tags uft on uft.tag_id = tags.id").
		Where("uft.user_id = ?", userID).
		Scan(&tags).Error

	return tags, err
}

// 获取所有电影
func getAllMovies(db *gorm.DB) ([]models.Movie, error) {
	var movies []models.Movie
	err := db.Find(&movies).Error
	return movies, err
}

// 获取电影对应的标签
func getTagsForMovie(db *gorm.DB, movieID int) ([]Tag, error) {
	var tags []Tag
	err := db.Table("tags").
		Select("tags.id, tags.name").
		Joins("join movie_tags mt on mt.tag_id = tags.id").
		Where("mt.movie_id = ?", movieID).
		Scan(&tags).Error

	return tags, err
}

// 基于内容的推荐算法实现
func contentBasedRecommendation(db *gorm.DB, userTags []Tag) ([]models.Movie, error) {
	var recommendedMovies []models.Movie
	movieTagCounts := make(map[int]int)

	// 将用户喜欢的标签转换为一个map，方便查找
	userTagMap := make(map[int]bool)
	for _, tag := range userTags {
		userTagMap[tag.ID] = true
	}

	// 获取所有电影
	movies, err := getAllMovies(db)
	if err != nil {
		return nil, err
	}

	// 计算每部电影与用户标签的匹配程度
	for _, movie := range movies {
		tags, err := getTagsForMovie(db, movie.ID)
		if err != nil {
			continue
		}
		var matchCount int
		for _, tag := range tags {
			if userTagMap[tag.ID] {
				matchCount++
			}
		}
		if matchCount > 0 {
			movieTagCounts[movie.ID] = matchCount
			recommendedMovies = append(recommendedMovies, movie)
		}
	}

	// 按照匹配度排序
	sort.Slice(recommendedMovies, func(i, j int) bool {
		return movieTagCounts[recommendedMovies[i].ID] > movieTagCounts[recommendedMovies[j].ID]
	})

	return recommendedMovies, nil
}

// 基于协同过滤的推荐算法实现
func collaborativeFilteringRecommendation(db *gorm.DB, userID int) ([]models.Movie, error) {
	var recommendedMovies []models.Movie

	query := `
        SELECT DISTINCT m.id, m.title, m.description, m.genre, m.rating, m.thumbnail_url, m.video_url
        FROM user_favorite_tags uft1
        JOIN user_favorite_tags uft2 ON uft1.tag_id = uft2.tag_id AND uft1.user_id != uft2.user_id
        JOIN movie_tags mt ON uft2.tag_id = mt.tag_id
        JOIN movies m ON mt.movie_id = m.id
        WHERE uft1.user_id = ?
    `

	err := db.Raw(query, userID).Scan(&recommendedMovies).Error
	return recommendedMovies, err
}

// 合并推荐结果并选出前8个
func mergeAndPickTopMovies(contentBased, collaborative []models.Movie, limit int) []models.Movie {
	movieMap := make(map[int]models.Movie)

	// 添加基于内容的推荐电影
	for _, movie := range contentBased {
		movieMap[movie.ID] = movie
	}

	// 添加基于协同过滤的推荐电影
	for _, movie := range collaborative {
		if len(movieMap) < limit {
			movieMap[movie.ID] = movie
		}
	}

	// 选出前8个电影
	var finalMovies []models.Movie
	for _, movie := range movieMap {
		finalMovies = append(finalMovies, movie)
		if len(finalMovies) >= limit {
			break
		}
	}

	return finalMovies
}

// 综合推荐算法
func RecommendMovies(db *gorm.DB, userID int) ([]models.Movie, error) {
	// Step 1: 获取用户的喜欢标签
	userTags, err := getUserFavoriteTags(db, userID)
	if err != nil {
		return nil, err
	}

	// Step 2: 基于内容的推荐
	contentBasedMovies, err := contentBasedRecommendation(db, userTags)
	if err != nil {
		return nil, err
	}

	// Step 3: 基于协同过滤的推荐
	collaborativeMovies, err := collaborativeFilteringRecommendation(db, userID)
	if err != nil {
		return nil, err
	}

	// 合并结果并返回前8个电影
	return mergeAndPickTopMovies(contentBasedMovies, collaborativeMovies, 8), nil
}
