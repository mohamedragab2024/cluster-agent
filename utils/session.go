package utils

import (
	"flag"
	"fmt"
	"log"
	"net/url"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/websocket"
)

type Session struct {
	Host    string
	Channel string
	Conn    *websocket.Conn
}

func (s Session) NewSession() (*Session, error) {
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
	}
	return &Session{
		Host:    s.Host,
		Conn:    conn,
		Channel: s.Channel,
	}, err
}

func (s Session) Serv(timeout time.Duration) {
	lastResponse := time.Now()
	s.Conn.SetPongHandler(func(msg string) error {
		lastResponse = time.Now()
		return nil
	})

	go func() {
		for {
			err := s.Conn.WriteMessage(websocket.PingMessage, []byte("keepalive"))
			if err != nil {
				return
			}
			time.Sleep(timeout / 2)
			if time.Since(lastResponse) > timeout {
				s.Conn.Close()
				os.Exit(3)
				return
			}
		}
	}()
}
