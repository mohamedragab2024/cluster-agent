package services

import (
	"encoding/json"
	"flag"
	"log"
	"net/url"
	"os"
	"os/signal"

	"github.com/gorilla/websocket"
)

type MonitoringService struct{}

func (m MonitoringService) PushEvent(conn *websocket.Conn, data interface{}) {
	var addr = flag.String("addr", os.Getenv("SERVER_ADDRESS"), "http service address")
	flag.Parse()
	log.SetFlags(0)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	u := url.URL{Scheme: "ws", Host: *addr, Path: "/services"}
	log.Printf("connecting to %s", u.String())

	defer conn.Close()
	msg, _ := json.Marshal(data)
	err := conn.WriteMessage(websocket.TextMessage, msg)
	//closeErr := conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
	if err != nil {
		log.Println("write:", err)
		return
	}

}
