package utils

import (
	"flag"
	"fmt"
	"log"
	"net/url"
	"os"
	"os/signal"
	"sync"

	"github.com/gorilla/websocket"
)

type Session struct {
	Host    string
	Channel string
	Conn    *websocket.Conn
	mu      sync.Mutex
}

func (s *Session) newSession() {
	var addr = flag.String("addr", s.Host, "http service address")
	flag.Parse()
	log.SetFlags(0)
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	u := url.URL{Scheme: "ws", Host: *addr, Path: fmt.Sprintf("/%s", s.Channel)}
	log.Printf("connecting to %s", u.String())
	conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
		return
	}
	s.Conn = conn
}

func (s *Session) Send(message []byte) error {
	if s.Conn == nil {
		s.newSession()
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.Conn.WriteMessage(websocket.TextMessage, message)
}
