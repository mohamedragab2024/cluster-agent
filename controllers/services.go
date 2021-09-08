package controllers

import (
	ctx "context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/kube-carbonara/cluster-agent/models"
	utils "github.com/kube-carbonara/cluster-agent/utils"
	"github.com/labstack/echo/v4"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

type ServicesController struct {
}

func (c ServicesController) Watch() {
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	for {
		watcher, err := clientset.CoreV1().Services(v1.NamespaceAll).Watch(ctx.TODO(), metav1.ListOptions{})
		if err != nil {
			panic(err.Error())
		}
		for {
			for event := range watcher.ResultChan() {
				svc := event.Object.(*v1.Service)

				switch event.Type {
				case watch.Added:
					fmt.Printf("pod %s/%s added", svc.ObjectMeta.Namespace, svc.ObjectMeta.Name)
				case watch.Modified:
					fmt.Printf("pod %s/%s modified", svc.ObjectMeta.Namespace, svc.ObjectMeta.Name)
				case watch.Deleted:
					fmt.Printf("pod %s/%s deleted", svc.ObjectMeta.Namespace, svc.ObjectMeta.Name)
				}
			}
		}
	}
}

func (c ServicesController) GetOne(context echo.Context, nameSpaceName string, name string) error {
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	result, err := clientset.CoreV1().Services(nameSpaceName).Get(ctx.TODO(), name, metav1.GetOptions{})
	if err != nil {
		return context.JSON(http.StatusBadRequest, models.Response{
			Message: err.Error(),
		})
	}

	return context.JSON(http.StatusOK, models.Response{
		Data:         utils.StructToMap(result),
		ResourceType: utils.RESOUCETYPE_SERVICES,
	})
}

func (c ServicesController) Get(context echo.Context, nameSpaceName string) error {
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	result, err := clientset.CoreV1().Services(nameSpaceName).List(ctx.TODO(), metav1.ListOptions{})
	if err != nil {
		return context.JSON(http.StatusBadRequest, models.Response{
			Message: err.Error(),
		})
	}

	return context.JSON(http.StatusOK, models.Response{
		Data:         utils.StructToMap(result),
		ResourceType: utils.RESOUCETYPE_SERVICES,
	})
}

func (c ServicesController) Create(context echo.Context, nameSpaceName string, serviceConfig map[string]interface{}) error {
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	service := &v1.Service{}
	UnmarshalErr := json.Unmarshal(utils.MapToJson(serviceConfig), service)
	if UnmarshalErr != nil {
		return context.JSON(http.StatusBadRequest, models.Response{
			Message: UnmarshalErr.Error(),
		})
	}
	result, err := clientset.CoreV1().Services(nameSpaceName).Create(ctx.TODO(), service, metav1.CreateOptions{})
	if err != nil {
		return context.JSON(http.StatusBadRequest, models.Response{
			Message: err.Error(),
		})
	}

	return context.JSON(http.StatusOK, models.Response{
		Data:         utils.StructToMap(result),
		ResourceType: utils.RESOUCETYPE_SERVICES,
	})
}

func (c ServicesController) Update(context echo.Context, nameSpaceName string, serviceConfig map[string]interface{}) error {
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	service := &v1.Service{}
	UnmarshalErr := json.Unmarshal(utils.MapToJson(serviceConfig), service)
	if UnmarshalErr != nil {
		return context.JSON(http.StatusBadRequest, models.Response{
			Message: UnmarshalErr.Error(),
		})
	}
	result, err := clientset.CoreV1().Services(nameSpaceName).Update(ctx.TODO(), service, metav1.UpdateOptions{})
	if err != nil {
		return context.JSON(http.StatusBadRequest, models.Response{
			Message: err.Error(),
		})
	}

	return context.JSON(http.StatusOK, models.Response{
		Data:         utils.StructToMap(result),
		ResourceType: utils.RESOUCETYPE_SERVICES,
	})
}

func (c ServicesController) Delete(context echo.Context, nameSpaceName string, name string) error {
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	deleteErr := clientset.CoreV1().Services(nameSpaceName).Delete(ctx.TODO(), name, metav1.DeleteOptions{})
	if deleteErr != nil {
		return context.JSON(http.StatusBadRequest, models.Response{
			Message: err.Error(),
		})
	}

	return context.JSON(http.StatusNoContent, models.Response{
		Data:         nil,
		ResourceType: utils.RESOUCETYPE_SERVICES,
	})
}
