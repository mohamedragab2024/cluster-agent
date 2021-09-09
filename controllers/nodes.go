package controllers

import (
	ctx "context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/kube-carbonara/cluster-agent/models"
	services "github.com/kube-carbonara/cluster-agent/services"
	utils "github.com/kube-carbonara/cluster-agent/utils"
	"github.com/labstack/echo/v4"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type NodesController struct{}

func (c NodesController) Watch() {
	var client utils.Client = *utils.NewClient()
	watch, err := client.Clientset.CoreV1().Nodes().Watch(ctx.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Fatal(err.Error())
	}
	conn := utils.SocketConnection{
		Host: os.Getenv("SERVER_ADDRESS"),
	}.EstablishNewConnection()
	go func() {
		for event := range watch.ResultChan() {

			obj, ok := event.Object.(*v1.Node)
			if !ok {
				log.Fatal("unexpected type")
			}
			if err != nil {
				log.Println("write:", err)
				return
			}
			services.MonitoringService{}.PushEvent(conn, obj)
		}
		time.Sleep(30 * time.Second)
	}()
}

func (c NodesController) GetOne(context echo.Context, name string) error {
	var client utils.Client = *utils.NewClient()
	result, err := client.Clientset.CoreV1().Nodes().Get(ctx.TODO(), name, metav1.GetOptions{})
	if err != nil {
		return context.JSON(http.StatusBadRequest, models.Response{
			Message: err.Error(),
		})
	}

	return context.JSON(http.StatusOK, models.Response{
		Data:         utils.StructToMap(result),
		ResourceType: utils.RESOUCETYPE_NODES,
	})
}

func (c NodesController) Get(context echo.Context) error {
	var client utils.Client = *utils.NewClient()
	result, err := client.Clientset.CoreV1().Nodes().List(ctx.TODO(), metav1.ListOptions{})
	if err != nil {
		return context.JSON(http.StatusBadRequest, models.Response{
			Message: err.Error(),
		})
	}

	return context.JSON(http.StatusOK, models.Response{
		Data:         utils.StructToMap(result),
		ResourceType: utils.RESOUCETYPE_NODES,
	})
}

func (c NodesController) Delete(context echo.Context, name string) error {
	var client utils.Client = *utils.NewClient()
	err := client.Clientset.CoreV1().Nodes().Delete(ctx.TODO(), name, metav1.DeleteOptions{})
	if err != nil {
		return context.JSON(http.StatusBadRequest, models.Response{
			Message: err.Error(),
		})
	}
	return context.JSON(http.StatusNoContent, nil)
}

func (c NodesController) Create(context echo.Context, nodeConfig map[string]interface{}) error {
	node := &v1.Node{}
	UnmarshalErr := json.Unmarshal(utils.MapToJson(nodeConfig), node)
	if UnmarshalErr != nil {
		return context.JSON(http.StatusBadRequest, models.Response{
			Message: UnmarshalErr.Error(),
		})
	}
	var client utils.Client = *utils.NewClient()
	result, err := client.Clientset.CoreV1().Nodes().Create(ctx.TODO(), node, metav1.CreateOptions{})
	if err != nil {
		return context.JSON(http.StatusBadRequest, models.Response{
			Message: err.Error(),
		})
	}
	return context.JSON(http.StatusCreated, models.Response{
		ResourceType: utils.RESOUCETYPE_NODES,
		Data:         utils.StructToMap(result),
	})
}

func (c NodesController) Update(context echo.Context, nodeConfig map[string]interface{}) error {
	node := &v1.Node{}
	UnmarshalErr := json.Unmarshal(utils.MapToJson(nodeConfig), node)
	if UnmarshalErr != nil {
		return context.JSON(http.StatusBadRequest, models.Response{
			Message: UnmarshalErr.Error(),
		})
	}
	var client utils.Client = *utils.NewClient()
	result, err := client.Clientset.CoreV1().Nodes().Update(ctx.TODO(), node, metav1.UpdateOptions{})
	if err != nil {
		return context.JSON(http.StatusBadRequest, models.Response{
			Message: err.Error(),
		})
	}

	return context.JSON(http.StatusOK, models.Response{
		Data:         utils.StructToMap(result),
		ResourceType: utils.RESOUCETYPE_NODES,
	})
}
