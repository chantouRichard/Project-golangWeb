package utils

import (
	"io"
	"net/http"
	"os"
	"path/filepath"
)

// StreamPicture 处理图片流媒体请求
func StreamPicture(w http.ResponseWriter, pictureID string) {
	// 假设图片文件保存在 ./pictures 目录下，并且图片格式为 jpg
	filePath := "./pictures/" + pictureID + ".jpg"

	// 打开图片文件
	file, err := os.Open(filePath)
	if err != nil {
		http.Error(w, "找不到图片", http.StatusNotFound)
		return
	}
	defer file.Close()

	// 获取图片的 MIME 类型
	fileExt := filepath.Ext(filePath)
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

	// 设置响应头信息
	w.Header().Set("Content-Type", mimeType)

	// 使用流方式读取图片并写入响应
	buffer := make([]byte, 1024*64) // 64KB 缓冲区
	for {
		bytesRead, err := file.Read(buffer)
		if err != nil {
			if err == io.EOF {
				break
			}
			http.Error(w, "Error reading file", http.StatusInternalServerError)
			return
		}
		w.Write(buffer[:bytesRead])
		w.(http.Flusher).Flush()
	}
}
