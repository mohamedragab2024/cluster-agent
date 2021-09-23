package utils

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"sync"

	"github.com/gorilla/websocket"
)

type Session struct {
	Host    string
	Channel string
	Conn    *websocket.Conn
	mu      sync.Mutex
}

func (s *Session) NewSession() {
	u := url.URL{Scheme: "ws", Host: s.Host, Path: fmt.Sprintf("/%s", s.Channel)}
	log.Printf("connecting to %s", u.String())
	conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
		os.Exit(0)
	}
	s.Conn = conn
}

func (s *Session) Send(message []byte) error {

	s.mu.Lock()
	defer s.mu.Unlock()
	return s.Conn.WriteMessage(websocket.TextMessage, message)
}
