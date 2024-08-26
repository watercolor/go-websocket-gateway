package handlers

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true // 可以根据实际情况调整此策略
	},
}

const (
	pingInterval = 30 * time.Second // 心跳间隔
	pingMessage  = "ping"           // 心跳消息
)

func WebSocketHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Upgrade:", err)
		return
	}
	defer conn.Close()

	// 设置读取截止时间为心跳间隔加一些额外的时间，以便有时间接收心跳响应
	conn.SetReadDeadline(time.Now().Add(pingInterval * 2))
	conn.SetPongHandler(func(string) error {
		conn.SetReadDeadline(time.Now().Add(pingInterval * 2))
		return nil
	})

	ticker := time.NewTicker(pingInterval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			if err := conn.WriteMessage(websocket.PingMessage, []byte(pingMessage)); err != nil {
				log.Println("write ping:", err)
				return
			}
		default:
			// 这里可以添加其他逻辑，例如处理从客户端接收到的消息
		}
	}
}
