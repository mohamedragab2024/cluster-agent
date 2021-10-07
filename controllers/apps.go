package controllers

import (
	ctx "context"
	"net/http"

	"github.com/kube-carbonara/cluster-agent/models"
	utils "github.com/kube-carbonara/cluster-agent/utils"
	"github.com/labstack/echo/v4"
)

type AppsController struct{}

func (c AppsController) Get(context echo.Context) error {
	var client utils.Client = *utils.NewClient()
	result := client.Clientset.AppsV1().RESTClient().Get().Do(ctx.TODO())
	if result.Error() != nil {
		return context.JSON(http.StatusBadRequest, models.Response{
			Message: result.Error().Error(),
		})
	}

	return context.JSON(http.StatusOK, models.Response{
		Data:         utils.StructToMap(result),
		ResourceType: utils.RESOUCETYPE_NODES,
	})
}
