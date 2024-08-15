package services

import (
	"io"
	"net/http"
	"os"
)

func StreamVideo(w http.ResponseWriter, videoPath string) {
	file, err := os.Open(videoPath)
	if err != nil {
		http.Error(w, "Video not found", http.StatusNotFound)
		return
	}
	defer file.Close()

	w.Header().Set("Content-Type", "video/mp4")
	_, err = io.Copy(w, file)
	if err != nil {
		http.Error(w, "Error streaming video", http.StatusInternalServerError)
	}
}
