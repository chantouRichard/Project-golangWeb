package models

import "time"

type ChatMessage struct {
	UserID    string `json:"user_id"`
	Message   string `json:"message"`
	RoomID    string `json:"room_id"`
	CreatedAt time.Time
}

// 模拟保存聊天记录

func SaveChatMessage(message ChatMessage) {
	// TODO: 将消息保存到数据库
}

// 模拟广播消息给所有用户

func BroadcastMessage(message ChatMessage) {
	// TODO: 广播消息给聊天室里的所有用户
}
