package websocket

import (
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

type Room struct {
	ID         string
	Host       *websocket.Conn
	Members    map[*websocket.Conn]bool
	VideoState VideoState
}

type VideoState struct {
	Progress  int64
	IsPlaying bool
}

type Message struct {
	Type     string `json:"type"`
	RoomID   string `json:"roomID"`
	Progress int64  `json:"progress"`
}

var rooms = make(map[string]*Room)
var mu sync.Mutex
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func HandleConnections(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer ws.Close()

	var currentRoom *Room

	for {
		var msg Message
		err := ws.ReadJSON(&msg)
		if err != nil {
			log.Printf("错误: %v", err)
			if currentRoom != nil {
				mu.Lock()
				delete(currentRoom.Members, ws)
				mu.Unlock()
			}
			break
		}

		switch msg.Type {
		case "create_room":
			mu.Lock()
			roomID := generateRoomID()
			currentRoom = &Room{
				ID:      roomID,
				Host:    ws,
				Members: make(map[*websocket.Conn]bool),
			}
			rooms[roomID] = currentRoom
			mu.Unlock()
			ws.WriteJSON(Message{Type: "room_created", RoomID: roomID})

		case "join_room":
			mu.Lock()
			currentRoom = rooms[msg.RoomID]
			if currentRoom != nil {
				currentRoom.Members[ws] = true
				ws.WriteJSON(Message{Type: "sync", Progress: currentRoom.VideoState.Progress})
			}
			mu.Unlock()

		case "play", "pause", "seek":
			if currentRoom != nil && currentRoom.Host == ws {
				mu.Lock()
				if msg.Type == "play" {
					currentRoom.VideoState.IsPlaying = true
				} else if msg.Type == "pause" {
					currentRoom.VideoState.IsPlaying = false
				} else if msg.Type == "seek" {
					currentRoom.VideoState.Progress = msg.Progress
				}
				broadcastToRoom(currentRoom, currentRoom.VideoState)
				mu.Unlock()
			}
		}
	}
}

func broadcastToRoom(room *Room, state VideoState) {
	for client := range room.Members {
		err := client.WriteJSON(Message{Type: "sync", Progress: state.Progress, RoomID: room.ID})
		if err != nil {
			log.Printf("广播错误: %v", err)
			client.Close()
			mu.Lock()
			delete(room.Members, client)
			mu.Unlock()
		}
	}
}

func generateRoomID() string {
	// 生成唯一Room ID的逻辑，可使用随机数或UUID
	return "room123" // 示例，实际应用中应生成唯一的ID
}
