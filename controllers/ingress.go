package controllers

import (
	ctx "context"
	"encoding/json"
	"net/http"

	"github.com/kube-carbonara/cluster-agent/models"
	utils "github.com/kube-carbonara/cluster-agent/utils"
	"github.com/labstack/echo/v4"
	networkingv1 "k8s.io/api/networking/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	networkingv1client "k8s.io/client-go/kubernetes/typed/networking/v1"
	"k8s.io/client-go/rest"
)

type IngressController struct {
}

func (c IngressController) GetOne(context echo.Context, nameSpaceName string, name string) error {
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}
	client, err := networkingv1client.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	result, err := client.Ingresses(nameSpaceName).Get(ctx.TODO(), name, metav1.GetOptions{})
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
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}
	client, err := networkingv1client.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	result, err := client.Ingresses(nameSpaceName).List(ctx.TODO(), metav1.ListOptions{})
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
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}
	client, err := networkingv1client.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	ingress := &networkingv1.Ingress{}
	UnmarshalErr := json.Unmarshal(utils.MapToJson(ingressConfig), ingress)
	if UnmarshalErr != nil {
		return context.JSON(http.StatusBadRequest, models.Response{
			Message: UnmarshalErr.Error(),
		})
	}
	ingress, err = client.Ingresses(nameSpaceName).Create(ctx.TODO(), ingress, metav1.CreateOptions{})
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
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}
	client, err := networkingv1client.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	ingress := &networkingv1.Ingress{}
	UnmarshalErr := json.Unmarshal(utils.MapToJson(ingressConfig), ingress)
	if UnmarshalErr != nil {
		return context.JSON(http.StatusBadRequest, models.Response{
			Message: UnmarshalErr.Error(),
		})
	}
	ingress, err = client.Ingresses(nameSpaceName).Update(ctx.TODO(), ingress, metav1.UpdateOptions{})
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
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}
	client, err := networkingv1client.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	err = client.Ingresses(nameSpaceName).Delete(ctx.TODO(), name, metav1.DeleteOptions{})
	if err != nil {
		return context.JSON(http.StatusBadRequest, models.Response{
			Message: err.Error(),
		})
	}
	return context.JSON(http.StatusNoContent, nil)
}
