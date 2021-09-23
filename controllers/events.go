package controllers

import (
	ctx "context"
	"log"
	"net/http"

	"github.com/kube-carbonara/cluster-agent/models"
	services "github.com/kube-carbonara/cluster-agent/services"
	utils "github.com/kube-carbonara/cluster-agent/utils"
	"github.com/labstack/echo/v4"
	CoreV1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type EventsController struct {
}

func (c EventsController) Watch() {
	config := utils.NewConfig()
	var client utils.Client = *utils.NewClient()
	watch, err := client.Clientset.CoreV1().Events(CoreV1.NamespaceAll).Watch(ctx.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Fatal(err.Error())
	}
	go func() {
		session := utils.Session{
			Host:    config.RemoteProxy,
			Channel: "monitoring",
		}
		session.NewSession()
		for event := range watch.ResultChan() {

			obj, ok := event.Object.(*CoreV1.Event)
			if !ok {
				log.Fatal("unexpected type")
			}

			services.MonitoringService{
				NameSpace: obj.Namespace,
				EventName: string(event.Type),
				Resource:  utils.EVENTS,
				PayLoad:   obj,
			}.PushEvent(&session)
		}
	}()

}

func (c EventsController) GetOne(context echo.Context, name string, nameSpace string) error {
	var client utils.Client = *utils.NewClient()
	result, err := client.Clientset.CoreV1().Events(nameSpace).Get(ctx.TODO(), name, metav1.GetOptions{})
	if err != nil {
		return context.JSON(http.StatusBadRequest, models.Response{
			Message: err.Error(),
		})
	}

	return context.JSON(http.StatusOK, models.Response{
		Data:         utils.StructToMap(result),
		ResourceType: utils.EVENTS,
	})
}

func (c EventsController) Get(context echo.Context, nameSpace string) error {
	var client utils.Client = *utils.NewClient()
	result, err := client.Clientset.CoreV1().Events(nameSpace).List(ctx.TODO(), metav1.ListOptions{})
	if err != nil {
		return context.JSON(http.StatusBadRequest, models.Response{
			Message: err.Error(),
		})
	}

	return context.JSON(http.StatusOK, models.Response{
		Data:         utils.StructToMap(result),
		ResourceType: utils.EVENTS,
	})
}
