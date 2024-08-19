// picture.go
package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"path/filepath"
)

// StreamPicture 处理图片流请求
func StreamPicture(c *gin.Context, picturePath string) {
	// 打开图片文件
	file, err := os.Open(picturePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "无法打开图片"})
		return
	}
	defer file.Close()

	// 获取文件信息
	fileInfo, err := file.Stat()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "无法获取图片信息"})
		return
	}

	// 获取图片的 MIME 类型
	fileExt := filepath.Ext(picturePath)
	mimeType := "image/jpeg" // 默认 MIME 类型
	switch fileExt {
	case ".png":
		mimeType = "image/png"
	case ".gif":
		mimeType = "image/gif"
	case ".bmp":
		mimeType = "image/bmp"
	case ".webp":
		mimeType = "image/webp"
	}

	// 设置响应头
	c.Writer.Header().Set("Content-Type", mimeType)
	c.Writer.Header().Set("Content-Length", fmt.Sprintf("%d", fileInfo.Size()))

	// 处理 HTTP Range 请求
	http.ServeContent(c.Writer, c.Request, filepath.Base(picturePath), fileInfo.ModTime(), file)
}
