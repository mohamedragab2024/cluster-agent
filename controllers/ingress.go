package controllers

import (
	ctx "context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
	"github.com/kube-carbonara/cluster-agent/models"
	services "github.com/kube-carbonara/cluster-agent/services"
	utils "github.com/kube-carbonara/cluster-agent/utils"
	"github.com/labstack/echo/v4"
	CoreV1 "k8s.io/api/core/v1"
	networkingv1 "k8s.io/api/networking/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type IngressController struct {
}

func (c IngressController) Watch(wsConn *websocket.Conn) {
	var client utils.Client = *utils.NewClient()
	watch, err := client.Networkingv1client.Ingresses(CoreV1.NamespaceAll).Watch(ctx.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Fatal(err.Error())
	}
	go func() {
		for event := range watch.ResultChan() {

			obj, ok := event.Object.(*networkingv1.Ingress)
			if !ok {
				log.Fatal("unexpected type")
			}
			if err != nil {
				log.Println("write:", err)
				return
			}
			services.MonitoringService{
				NameSpace: obj.Namespace,
				EventName: string(event.Type),
				Resource:  utils.RESOUCETYPE_INGRESS,
				PayLoad:   obj,
			}.PushEvent(wsConn)
		}
		time.Sleep(30 * time.Second)
	}()
}

func (c IngressController) GetOne(context echo.Context, nameSpaceName string, name string) error {
	var client utils.Client = *utils.NewClient()
	result, err := client.Networkingv1client.Ingresses(nameSpaceName).Get(ctx.TODO(), name, metav1.GetOptions{})
	if err != nil {
		return context.JSON(http.StatusBadRequest, models.Response{
			Message: err.Error(),
		})
	}

	return context.JSON(http.StatusOK, models.Response{
		Data:         utils.StructToMap(result),
		ResourceType: utils.RESOUCETYPE_INGRESS,
	})
}

func (c IngressController) Get(context echo.Context, nameSpaceName string) error {
	var client utils.Client = *utils.NewClient()
	result, err := client.Networkingv1client.Ingresses(nameSpaceName).List(ctx.TODO(), metav1.ListOptions{})
	if err != nil {
		return context.JSON(http.StatusBadRequest, models.Response{
			Message: err.Error(),
		})
	}

	return context.JSON(http.StatusOK, models.Response{
		Data:         utils.StructToMap(result),
		ResourceType: utils.RESOUCETYPE_INGRESS,
	})
}

func (c IngressController) Create(context echo.Context, nameSpaceName string, ingressConfig map[string]interface{}) error {
	ingress := &networkingv1.Ingress{}
	UnmarshalErr := json.Unmarshal(utils.MapToJson(ingressConfig), ingress)
	if UnmarshalErr != nil {
		return context.JSON(http.StatusBadRequest, models.Response{
			Message: UnmarshalErr.Error(),
		})
	}

	var client utils.Client = *utils.NewClient()
	ingress, err := client.Networkingv1client.Ingresses(nameSpaceName).Create(ctx.TODO(), ingress, metav1.CreateOptions{})
	if err != nil {
		return context.JSON(http.StatusBadRequest, models.Response{
			Message: err.Error(),
		})
	}
	return context.JSON(http.StatusOK, models.Response{
		Data:         utils.StructToMap(ingress),
		ResourceType: utils.RESOUCETYPE_INGRESS,
	})
}

func (c IngressController) Update(context echo.Context, nameSpaceName string, ingressConfig map[string]interface{}) error {
	ingress := &networkingv1.Ingress{}
	UnmarshalErr := json.Unmarshal(utils.MapToJson(ingressConfig), ingress)
	if UnmarshalErr != nil {
		return context.JSON(http.StatusBadRequest, models.Response{
			Message: UnmarshalErr.Error(),
		})
	}
	var client utils.Client = *utils.NewClient()
	ingress, err := client.Networkingv1client.Ingresses(nameSpaceName).Update(ctx.TODO(), ingress, metav1.UpdateOptions{})
	if err != nil {
		return context.JSON(http.StatusBadRequest, models.Response{
			Message: err.Error(),
		})
	}
	return context.JSON(http.StatusOK, models.Response{
		Data:         utils.StructToMap(ingress),
		ResourceType: utils.RESOUCETYPE_INGRESS,
	})
}

func (c IngressController) Delete(context echo.Context, nameSpaceName string, name string) error {
	var client utils.Client = *utils.NewClient()
	err := client.Networkingv1client.Ingresses(nameSpaceName).Delete(ctx.TODO(), name, metav1.DeleteOptions{})
	if err != nil {
		return context.JSON(http.StatusBadRequest, models.Response{
			Message: err.Error(),
		})
	}
	return context.JSON(http.StatusNoContent, nil)
}
