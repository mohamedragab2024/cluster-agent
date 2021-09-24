package main

import (
	"context"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/joho/godotenv"
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
	addr   string
	id     string
	debug  bool
	appKey string
)

func handleRouting(e *echo.Echo) {
	namespacesRouter := routers.NameSpacesRouter{}
	podsRouter := routers.PodsRouter{}
	deplymentRouter := routers.DeploymentsRouter{}
	serviceRouter := routers.SeviceRouter{}
	nodeRouter := routers.NodesRouter{}
	ingressRouter := routers.IngresRouter{}
	metricsRouter := routers.MetricsRouter{}
	secretRouter := routers.SecretRouter{}
	eventRouter := routers.EventsRouter{}
	workloadRouter := routers.WorkLoadsRouter{}
	namespacesRouter.Handle(e)
	podsRouter.Handle(e)
	deplymentRouter.Handle(e)
	serviceRouter.Handle(e)
	nodeRouter.Handle(e)
	ingressRouter.Handle(e)
	metricsRouter.Handle(e)
	secretRouter.Handle(e)
	eventRouter.Handle(e)
	workloadRouter.Handle(e)
}

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
	time.AfterFunc(5*time.Second, func() {
		remotedialer.ClientConnect(context.Background(), addr, headers, nil, func(string, string) bool { return true }, nil)

	})

	go controllers.ServicesController{}.Watch()
	go controllers.PodsController{}.Watch()
	go controllers.DeploymentsController{}.Watch()
	go controllers.NameSpacesController{}.Watch()
	go controllers.NodesController{}.Watch()
	go controllers.IngressController{}.Watch()
	go controllers.SecretsController{}.Watch()
	go controllers.EventsController{}.Watch()

	e := echo.New()
	e.GET("/", func(context echo.Context) error {
		return context.String(http.StatusOK, "Hello, World!")
	})

	e.GET("/health", func(context echo.Context) error {
		resp, err := http.Get(fmt.Sprintf("%s://%s", config.RemoteSchema, config.RemoteProxy))
		if err != nil {
			log.Fatalln(err)
			return context.String(http.StatusGatewayTimeout, "error connecting to gateway")
		}
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln(err)
			return context.String(http.StatusGatewayTimeout, "error connecting to gateway")
		}

		fmt.Print(string(body) + "\n")

		return context.String(http.StatusOK, "App is running")

	})
	handleRouting(e)

	e.Logger.Fatal(e.Start(":1323"))
}
