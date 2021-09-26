package controllers

import (
	ctx "context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/kube-carbonara/cluster-agent/models"
	services "github.com/kube-carbonara/cluster-agent/services"
	utils "github.com/kube-carbonara/cluster-agent/utils"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type SecretsController struct {
}

func (c SecretsController) runWatcherEventLoop() error {
	config := utils.NewConfig()
	var client utils.Client = *utils.NewClient()
	watch, err := client.Clientset.CoreV1().Secrets(v1.NamespaceAll).Watch(ctx.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Fatal(err.Error())
		return err
	}

	channel := watch.ResultChan()

	done := make(chan struct{})

	session := utils.Session{
		Host:    config.RemoteProxy,
		Channel: "monitoring",
	}
	session.NewSession()
	defer session.Conn.Close()
	defer close(done)
	for {
		select {
		case event, ok := <-channel:
			if !ok {
				log.Fatal("unexpected type")
				return nil
			}
			obj, ok := event.Object.(*v1.Secret)
			if !ok {
				log.Fatal("unexpected type")
				return nil
			}

			err := services.MonitoringService{
				NameSpace: obj.Namespace,
				EventName: string(event.Type),
				Resource:  utils.RESOUCETYPE_SECRETS,
				PayLoad:   obj,
			}.PushEvent(&session)

			if err != nil {
				logrus.Error(err)
				session.Conn.Close()
				session = *session.NewSession()
				services.MonitoringService{
					EventName: string(event.Type),
					Resource:  utils.RESOUCETYPE_SECRETS,
					PayLoad:   obj,
				}.PushEvent(&session)
			}

		case <-time.After(30 * time.Minute):
			logrus.Info("Timeout, restarting event watcher")
			return nil

		}
	}

}

func (c SecretsController) Watch() {

	for {
		if err := c.runWatcherEventLoop(); err != nil {
			logrus.Error(err)
		}

	}
}

func (c SecretsController) GetOne(context echo.Context, nameSpaceName string, name string) error {
	var client utils.Client = *utils.NewClient()
	result, err := client.Clientset.CoreV1().Secrets(nameSpaceName).Get(ctx.TODO(), name, metav1.GetOptions{})
	if err != nil {
		return context.JSON(http.StatusBadRequest, models.Response{
			Message: err.Error(),
		})
	}

	return context.JSON(http.StatusOK, models.Response{
		Data:         utils.StructToMap(result),
		ResourceType: utils.RESOUCETYPE_SECRETS,
	})
}

func (c SecretsController) Get(context echo.Context, nameSpaceName string) error {
	var client utils.Client = *utils.NewClient()
	result, err := client.Clientset.CoreV1().Secrets(nameSpaceName).List(ctx.TODO(), metav1.ListOptions{})
	if err != nil {
		return context.JSON(http.StatusBadRequest, models.Response{
			Message: err.Error(),
		})
	}

	return context.JSON(http.StatusOK, models.Response{
		Data:         utils.StructToMap(result),
		ResourceType: utils.RESOUCETYPE_SECRETS,
	})
}

func (c SecretsController) Create(context echo.Context, nameSpaceName string, secretConfig map[string]interface{}) error {
	secret := &v1.Secret{}
	UnmarshalErr := json.Unmarshal(utils.MapToJson(secretConfig), secret)
	if UnmarshalErr != nil {
		return context.JSON(http.StatusBadRequest, models.Response{
			Message: UnmarshalErr.Error(),
		})
	}

	var client utils.Client = *utils.NewClient()
	result, err := client.Clientset.CoreV1().Secrets(nameSpaceName).Create(ctx.TODO(), secret, metav1.CreateOptions{})
	if err != nil {
		return context.JSON(http.StatusBadRequest, models.Response{
			Message: err.Error(),
		})
	}

	return context.JSON(http.StatusOK, models.Response{
		Data:         utils.StructToMap(result),
		ResourceType: utils.RESOUCETYPE_SECRETS,
	})
}

func (c SecretsController) Update(context echo.Context, nameSpaceName string, secretConfig map[string]interface{}) error {
	secret := &v1.Secret{}
	UnmarshalErr := json.Unmarshal(utils.MapToJson(secretConfig), secret)
	if UnmarshalErr != nil {
		return context.JSON(http.StatusBadRequest, models.Response{
			Message: UnmarshalErr.Error(),
		})
	}

	var client utils.Client = *utils.NewClient()
	result, err := client.Clientset.CoreV1().Secrets(nameSpaceName).Update(ctx.TODO(), secret, metav1.UpdateOptions{})
	if err != nil {
		return context.JSON(http.StatusBadRequest, models.Response{
			Message: err.Error(),
		})
	}

	return context.JSON(http.StatusOK, models.Response{
		Data:         utils.StructToMap(result),
		ResourceType: utils.RESOUCETYPE_SECRETS,
	})
}

func (c SecretsController) Delete(context echo.Context, nameSpaceName string, name string) error {
	var client utils.Client = *utils.NewClient()
	err := client.Clientset.CoreV1().Secrets(nameSpaceName).Delete(ctx.TODO(), name, metav1.DeleteOptions{})
	if err != nil {
		return context.JSON(http.StatusBadRequest, models.Response{
			Message: err.Error(),
		})
	}

	return context.JSON(http.StatusNoContent, models.Response{
		Data:         nil,
		ResourceType: utils.RESOUCETYPE_SECRETS,
	})
}
