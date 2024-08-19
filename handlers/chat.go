package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"gorm.io/gorm"
	"log"
	"net/http"
)

// Chat 处理在线聊天功能
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var roomConnections = make(map[string][]*websocket.Conn)

func Chat(c *gin.Context, roomID string, db *gorm.DB) {
	//roomID := c.Param("room_id")

	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("Failed to set WebSocket upgrade: ", err)
		return
	}
	defer conn.Close()

	roomConnections[roomID] = append(roomConnections[roomID], conn)

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println("ReadMessage Error:", err)
			break
		}

		for _, c := range roomConnections[roomID] {
			err = c.WriteMessage(websocket.TextMessage, msg)
			if err != nil {
				log.Println("WriteMessage Error:", err)
				c.Close()
				roomConnections[roomID] = removeConn(roomConnections[roomID], c)
			}
		}
	}
}

func removeConn(conns []*websocket.Conn, connToRemove *websocket.Conn) []*websocket.Conn {
	for i, conn := range conns {
		if conn == connToRemove {
			return append(conns[:i], conns[i+1:]...)
		}
	}
	return conns
}
