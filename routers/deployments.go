package routers

import (
	controllers "github.com/kube-carbonara/cluster-agent/controllers"
	"github.com/kube-carbonara/cluster-agent/utils"
	"github.com/labstack/echo/v4"
)

type DeploymentsRouter struct {
}

func (router DeploymentsRouter) Handle(e *echo.Echo) {
	deploymentController := controllers.DeploymentsControllers{}
	e.GET("/:ns/deployments", func(context echo.Context) error {
		return deploymentController.Get(context, context.Param("ns"))
	})

	e.POST("/:ns/deployments", func(context echo.Context) error {
		deployment := utils.JsonBodyToMap(context.Request().Body)
		return deploymentController.Create(context, context.Param("ns"), deployment)
	})

	e.DELETE("/:ns/deployments/:id", func(context echo.Context) error {
		return deploymentController.Delete(context, context.Param("ns"), context.Param("id"))
	})
}
