package services

import (
	"encoding/json"
	"log"

	"github.com/gorilla/websocket"
)

type MonitoringService struct {
	Resource  string
	EventName string
	PayLoad   interface{}
}

func (m MonitoringService) PushEvent(wsConn *websocket.Conn) {

	msg, _ := json.Marshal(m)
	err := wsConn.WriteMessage(websocket.TextMessage, msg)
	if err != nil {
		log.Println("write:", err)
		return
	}

}
