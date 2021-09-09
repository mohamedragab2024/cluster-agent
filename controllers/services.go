package controllers

import (
	ctx "context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/kube-carbonara/cluster-agent/models"
	utils "github.com/kube-carbonara/cluster-agent/utils"
	"github.com/labstack/echo/v4"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

type ServicesController struct {
}

func (c ServicesController) Watch() {
	fmt.Print("Checking events ...")
	config, err := rest.InClusterConfig()
	if err != nil {
		fmt.Printf("error InClusterConfig %s", err.Error())
		panic(err.Error())
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		fmt.Printf("error create new k8s client %s", err.Error())
		panic(err.Error())
	}
	for {

		result, err := clientset.CoreV1().Services(v1.NamespaceAll).List(ctx.Background(), metav1.ListOptions{})
		if err != nil {
			fmt.Printf("failed to get services %s", err.Error())
		}
		fmt.Printf("Services count : %d", len(result.Items))
		time.Sleep(10 * time.Second)
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
