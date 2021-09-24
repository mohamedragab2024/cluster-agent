package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/kube-carbonara/cluster-agent/utils"
	"github.com/rancher/remotedialer"
	"github.com/sirupsen/logrus"
)

func init() {
}

var (
	addr   string
	id     string
	debug  bool
	appKey string
)

func main() {
	godotenv.Load()

	config := utils.NewConfig()
	flag.StringVar(&addr, "connect", fmt.Sprintf("ws://%s/connect", config.RemoteProxy), "Address to connect to")
	flag.StringVar(&id, "id", config.ClientId, "Client ID")
	flag.StringVar(&appKey, "appKey", config.AppKey, "App Key")
	flag.BoolVar(&debug, "debug", true, "Debug logging")
	flag.Parse()

	if debug {
		logrus.SetLevel(logrus.DebugLevel)
	}

	headers := http.Header{
		"X-Tunnel-ID":     []string{id},
		"x-agent":         []string{id},
		"x-agent-app-key": []string{appKey},
	}

	remotedialer.ClientConnect(context.Background(), addr, headers, nil, func(string, string) bool { return true }, nil)

}
