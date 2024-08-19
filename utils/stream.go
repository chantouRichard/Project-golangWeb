package utils

import (
	"io"
	"net/http"
	"os"
)

// StreamMovieMultithreaded 使用多线程处理电影流媒体播放
func StreamMovieMultithreaded(w http.ResponseWriter, movieID string) {
	filePath := "./movies/" + movieID + ".mp4"
	file, err := os.Open(filePath)
	if err != nil {
		http.Error(w, "找不到电影", http.StatusNotFound)
		return
	}
	defer file.Close()

	// 设置响应头信息
	w.Header().Set("Content-Type", "video/mp4")

	// 使用多线程处理文件流
	buffer := make([]byte, 1024*1024) // 1MB 缓冲区
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
