package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/kube-carbonara/cluster-agent/controllers"
	routers "github.com/kube-carbonara/cluster-agent/routers"
	"github.com/kube-carbonara/cluster-agent/utils"
	"github.com/labstack/echo/v4"
	"github.com/rancher/remotedialer"
	"github.com/sirupsen/logrus"
)

func init() {
}

var (
	addr  string
	id    string
	debug bool
)

func handleRouting(e *echo.Echo) {
	namespacesRouter := routers.NameSpacesRouter{}
	podsRouter := routers.PodsRouter{}
	deplymentRouter := routers.DeploymentsRouter{}
	serviceRouter := routers.SeviceRouter{}
	nodeRouter := routers.NodesRouter{}
	ingressRouter := routers.IngresRouter{}
	namespacesRouter.Handle(e)
	podsRouter.Handle(e)
	deplymentRouter.Handle(e)
	serviceRouter.Handle(e)
	nodeRouter.Handle(e)
	ingressRouter.Handle(e)
}

func main() {
	config := utils.NewConfig()
	flag.StringVar(&addr, "connect", fmt.Sprintf("ws://%s/connect", config.RemoteProxy), "Address to connect to")
	flag.StringVar(&id, "id", config.ClientId, "Client ID")
	flag.BoolVar(&debug, "debug", true, "Debug logging")
	flag.Parse()

	if debug {
		logrus.SetLevel(logrus.DebugLevel)
	}

	headers := http.Header{
		"X-Tunnel-ID": []string{id},
	}
	time.AfterFunc(5*time.Second, func() {
		remotedialer.ClientConnect(context.Background(), addr, headers, nil, func(string, string) bool { return true }, nil)

	})

	session, err := utils.Session{
		Host:    config.RemoteProxy,
		Channel: "monitoring",
	}.NewSession()

	if err != nil {
		logrus.Error(err)
		os.Exit(3)
	}

	defer session.Conn.Close()
	controllers.ServicesController{}.Watch(session.Conn)

	e := echo.New()
	e.GET("/", func(context echo.Context) error {
		return context.String(http.StatusOK, "Hello, World!")
	})
	handleRouting(e)

	e.Logger.Fatal(e.Start(":1323"))
}
