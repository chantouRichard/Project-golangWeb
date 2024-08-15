package handlers

import (
	"Online-Theater/models"
	"Online-Theater/services"
	"encoding/json"
	"net/http"
)

func ChatHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		var chat *models.ChatMessage
		if err := json.NewDecoder(r.Body).Decode(&chat); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		services.SaveChatMessage(chat)
		w.WriteHeader(http.StatusCreated)
	case "GET":
		messages := services.GetChatMessages()
		json.NewEncoder(w).Encode(messages)
	}
}
