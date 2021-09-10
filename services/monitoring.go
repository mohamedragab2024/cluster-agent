package services

import (
	"encoding/json"
	"log"

	"github.com/gorilla/websocket"
)

type MonitoringService struct{}

func (m MonitoringService) PushEvent(wsConn *websocket.Conn, data interface{}) {
	msg, _ := json.Marshal(data)
	err := wsConn.WriteMessage(websocket.TextMessage, msg)
	if err != nil {
		log.Println("write:", err)
		return
	}

}
