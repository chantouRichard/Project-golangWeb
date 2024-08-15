package services

import (
	"Online-Theater/models"
	"sync"
)

var chatMessages []models.ChatMessage
var chatMutex sync.Mutex

func SaveChatMessage(chat *models.ChatMessage) {
	chatMutex.Lock()
	chatMessages = append(chatMessages, *chat)
	chatMutex.Unlock()
}

func GetChatMessages() []models.ChatMessage {
	chatMutex.Lock()
	messages := chatMessages
	chatMutex.Unlock()
	return messages
}
