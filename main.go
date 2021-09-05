package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	handlers "github.com/kube-carbonara/cluster-agent/handlers"
	"github.com/kube-carbonara/cluster-agent/models"
	routers "github.com/kube-carbonara/cluster-agent/routers"
	"github.com/labstack/echo/v4"
)

func init() {
}

var onMessage mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	clusterGuid := "afc014b73ac645b096ba49b261f0227b"
	fetch := msg.Payload()
	if string(fetch) == "" {
		log.Print("Empty request ..")
		return
	}
	var request models.ServerRequest
	err := json.Unmarshal([]byte(fetch), &request)
	if err != nil {
		log.Printf("Failed to unmarshal ..: %s\n", err.Error())
		return
	}
	if request.Prefix == "X-" {
		return
	}
	log.Printf("Recevied new message")

	requestHandler := handlers.RequestHandler{}
	res := requestHandler.Handle(request)
	res.Prefix = "X-"
	resStr, err := json.Marshal(res)
	token := client.Publish(fmt.Sprintf("clients/%s", clusterGuid), 0, false, resStr)
	token.Wait()
	if token.Error() != nil {
		fmt.Print("Failed to publish message ", token.Error())
	}
}

var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	fmt.Println("Connected")
}

var connectionLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	fmt.Printf("Connection Lost: %s\n", err.Error())
}

func handleRouting(e *echo.Echo) {
	namespacesRouter := routers.NameSpacesRouter{}
	podsRouter := routers.PodsRouter{}
	namespacesRouter.Handle(e)
	podsRouter.Handle(e)
}

func main() {
	clusterGuid := "afc014b73ac645b096ba49b261f0227b"

	fmt.Print(clusterGuid)
	options := mqtt.NewClientOptions()
	options.AddBroker("tcp://localhost:1883")
	options.SetClientID(fmt.Sprintf("Client-%s", clusterGuid))
	options.OnConnect = connectHandler
	options.OnConnectionLost = connectionLostHandler
	client := mqtt.NewClient(options)
	token := client.Connect()
	if token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	topic := fmt.Sprintf("clients/%s", clusterGuid)
	if token := client.Subscribe(topic, 0, onMessage); token.Wait() && token.Error() != nil {
		fmt.Printf("Failed to subscribe to the MQ-topic:%s\n", token.Error())
		os.Exit(1)
	}
	e := echo.New()
	e.GET("/", func(context echo.Context) error {
		return context.String(http.StatusOK, "Hello, World!")
	})
	handleRouting(e)

	e.Logger.Fatal(e.Start(":1323"))
}
