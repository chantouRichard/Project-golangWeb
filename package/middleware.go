package middleware

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// TokenAuthMiddleware 验证请求头中的 token
func TokenAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")

		// 检查 token 是否存在
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status":  "error",
				"message": "Authorization token is required",
			})
			c.Abort()
			return
		}

		// 在这里添加你的 token 验证逻辑
		if !isValidToken(token) {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status":  "error",
				"message": "Invalid authorization token",
			})
			c.Abort()
			return
		}

		c.Next() // 继续处理请求
	}
}

var jwtSecret = []byte("your-secret-key") // 替换为实际的 JWT 密钥

// isValidToken 验证 token 是否有效
func isValidToken(tokenStr string) bool {
	// 解析 token
	_, claims, err := parseToken(tokenStr)
	if err != nil {
		// 解析失败，token 无效
		return false
	}

	// 检查 token 是否过期
	if time.Now().Unix() > claims.ExpiresAt {
		return false
	}

	return true
}

// parseToken 解析 token 并返回 claims
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
