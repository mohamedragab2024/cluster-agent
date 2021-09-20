package controllers

import (
	ctx "context"
	"log"
	"net/http"
	"time"

	"github.com/kube-carbonara/cluster-agent/models"
	services "github.com/kube-carbonara/cluster-agent/services"
	utils "github.com/kube-carbonara/cluster-agent/utils"
	"github.com/labstack/echo/v4"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type NameSpacesController struct {
}

func (c NameSpacesController) WatchTest(session *utils.Session) {
	go func() {

		for {
			services.MonitoringService{}.PushEvent(session)
			time.Sleep(30 * time.Second)
		}

	}()

}

func (c NameSpacesController) Watch(session *utils.Session) {
	var client utils.Client = *utils.NewClient()
	watch, err := client.Clientset.CoreV1().Namespaces().Watch(ctx.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Fatal(err.Error())
	}
	go func() {
		for event := range watch.ResultChan() {

			obj, ok := event.Object.(*v1.Namespace)
			if !ok {
				log.Fatal("unexpected type")
			}
			if err != nil {
				log.Println("write:", err)
				return
			}
			services.MonitoringService{
				EventName: string(event.Type),
				Resource:  utils.RESOUCETYPE_NAMESPACES,
				PayLoad:   obj,
			}.PushEvent(session)
		}

	}()
}

func (c NameSpacesController) GetOne(context echo.Context, name string) error {
	var client utils.Client = *utils.NewClient()
	result, err := client.Clientset.CoreV1().Namespaces().Get(ctx.TODO(), name, metav1.GetOptions{})
	if err != nil {
		return context.JSON(http.StatusBadRequest, models.Response{
			Message: err.Error(),
		})
	}

	return context.JSON(http.StatusOK, models.Response{
		Data:         utils.StructToMap(result),
		ResourceType: utils.RESOUCETYPE_NAMESPACES,
	})
}

func (c NameSpacesController) Get(context echo.Context) error {
	var client utils.Client = *utils.NewClient()
	result, err := client.Clientset.CoreV1().Namespaces().List(ctx.TODO(), metav1.ListOptions{})
	if err != nil {
		return context.JSON(http.StatusBadRequest, models.Response{
			Message: err.Error(),
		})
	}

	return context.JSON(http.StatusOK, models.Response{
		Data:         utils.StructToMap(result),
		ResourceType: utils.RESOUCETYPE_NAMESPACES,
	})
}

func (c NameSpacesController) Delete(context echo.Context, name string) error {
	var client utils.Client = *utils.NewClient()
	err := client.Clientset.CoreV1().Namespaces().Delete(ctx.TODO(), name, metav1.DeleteOptions{})
	if err != nil {
		return context.JSON(http.StatusBadRequest, models.Response{
			Message: err.Error(),
		})
	}
	return context.JSON(http.StatusNoContent, nil)
}

func (c NameSpacesController) Create(context echo.Context, name string) error {
	ns := &v1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name: name,
		},
	}
	var client utils.Client = *utils.NewClient()
	result, err := client.Clientset.CoreV1().Namespaces().Create(ctx.TODO(), ns, metav1.CreateOptions{})
	if err != nil {
		return context.JSON(http.StatusBadRequest, models.Response{
			Message: err.Error(),
		})
	}
	return context.JSON(http.StatusCreated, models.Response{
		ResourceType: utils.RESOUCETYPE_NAMESPACES,
		Data:         utils.StructToMap(result),
	})
}
