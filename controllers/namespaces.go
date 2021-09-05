package controllers

import (
	ctx "context"
	"net/http"

	"github.com/kube-carbonara/cluster-agent/models"
	"github.com/labstack/echo/v4"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

type NameSpacesController struct {
}

func (c NameSpacesController) Get(context echo.Context) error {
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	result, err := clientset.CoreV1().Namespaces().List(ctx.TODO(), metav1.ListOptions{})
	if err != nil {
		return context.JSON(http.StatusBadRequest, models.Response{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
	}
	return context.JSON(http.StatusOK, result)
}
