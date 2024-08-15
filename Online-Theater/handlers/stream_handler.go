package handlers

import (
	"Online-Theater/services"
	"net/http"
)

func StreamHandler(w http.ResponseWriter, r *http.Request) {
	// Example: Streaming a video
	videoPath := r.URL.Query().Get("video")
	if videoPath == "" {
		http.Error(w, "Video not specified", http.StatusBadRequest)
		return
	}
	services.StreamVideo(w, videoPath)
}
