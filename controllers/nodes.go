package controllers

import (
	ctx "context"
	"encoding/json"
	"net/http"

	"github.com/kube-carbonara/cluster-agent/models"
	utils "github.com/kube-carbonara/cluster-agent/utils"
	"github.com/labstack/echo/v4"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

type NodesController struct{}

func (c NodesController) GetOne(context echo.Context, name string) error {
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	result, err := clientset.CoreV1().Nodes().Get(ctx.TODO(), name, metav1.GetOptions{})
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
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	result, err := clientset.CoreV1().Nodes().List(ctx.TODO(), metav1.ListOptions{})
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
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return context.JSON(http.StatusBadRequest, models.Response{
			Message: err.Error(),
		})
	}
	clientSetErr := clientset.CoreV1().Nodes().Delete(ctx.TODO(), name, metav1.DeleteOptions{})
	if clientSetErr != nil {
		return context.JSON(http.StatusBadRequest, models.Response{
			Message: err.Error(),
		})
	}
	return context.JSON(http.StatusNoContent, nil)
}

func (c NodesController) Create(context echo.Context, nodeConfig map[string]interface{}) error {
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return context.JSON(http.StatusBadRequest, models.Response{
			Message: err.Error(),
		})
	}

	node := &v1.Node{}
	UnmarshalErr := json.Unmarshal(utils.MapToJson(nodeConfig), node)
	if UnmarshalErr != nil {
		return context.JSON(http.StatusBadRequest, models.Response{
			Message: UnmarshalErr.Error(),
		})
	}

	result, err := clientset.CoreV1().Nodes().Create(ctx.TODO(), node, metav1.CreateOptions{})
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
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	node := &v1.Node{}
	UnmarshalErr := json.Unmarshal(utils.MapToJson(nodeConfig), node)
	if UnmarshalErr != nil {
		return context.JSON(http.StatusBadRequest, models.Response{
			Message: UnmarshalErr.Error(),
		})
	}
	result, err := clientset.CoreV1().Nodes().Update(ctx.TODO(), node, metav1.UpdateOptions{})
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
