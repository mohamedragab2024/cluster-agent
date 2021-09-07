package controllers

import (
	ctx "context"
	"encoding/json"
	"net/http"

	"github.com/kube-carbonara/cluster-agent/models"
	utils "github.com/kube-carbonara/cluster-agent/utils"
	"github.com/labstack/echo/v4"
	"k8s.io/api/extensions/v1beta1"
	"k8s.io/client-go/1.5/pkg/runtime"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

type IngressController struct {
}

func (c IngressController) Create(context echo.Context, nameSpaceName string, ingressConfig map[string]interface{}) error {
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	ingress := &v1beta1.Ingress{}
	UnmarshalErr := json.Unmarshal(utils.MapToJson(ingressConfig), ingress)
	if UnmarshalErr != nil {
		return context.JSON(http.StatusBadRequest, models.Response{
			Message: UnmarshalErr.Error(),
		})
	}

	var result runtime.Object
	createErr := clientset.RESTClient().Post().
		NamespaceIfScoped(nameSpaceName, true).
		Resource("ingress").
		Body(ingress).
		Do(ctx.TODO()).
		Into(result)

	if createErr != nil {
		return context.JSON(http.StatusInternalServerError, models.Response{
			Message:      createErr.Error(),
			ResourceType: utils.RESOUCETYPE_INGRESS,
		})
	}
	return context.JSON(http.StatusOK, models.Response{
		Data:         utils.StructToMap(result),
		ResourceType: utils.RESOUCETYPE_INGRESS,
	})
}
