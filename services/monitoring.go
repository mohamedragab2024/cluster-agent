package services

import (
	"encoding/json"
	"log"

	"github.com/gorilla/websocket"
	utils "github.com/kube-carbonara/cluster-agent/utils"
)

type MonitoringService struct {
	NameSpace string
	Resource  string
	EventName string
	PayLoad   interface{}
	ClusterId string
}

func (m MonitoringService) PushEvent(wsConn *websocket.Conn) {
	m.ClusterId = utils.NewConfig().RemoteProxy
	msg, _ := json.Marshal(m)
	err := wsConn.WriteMessage(websocket.TextMessage, msg)
	if err != nil {
		log.Println("write:", err)
		return
	}

}
