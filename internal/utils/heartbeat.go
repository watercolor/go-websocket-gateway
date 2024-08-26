package utils

import (
	"github.com/gorilla/websocket"
)

func SendHeartbeat(conn *websocket.Conn) error {
	return conn.WriteMessage(websocket.TextMessage, []byte("heartbeat"))
}
