package routers

import (
	"strconv"

	controllers "github.com/kube-carbonara/cluster-agent/controllers"
	"github.com/kube-carbonara/cluster-agent/utils"
	"github.com/labstack/echo/v4"
)

type DeploymentsRouter struct {
}

func (router DeploymentsRouter) Handle(e *echo.Echo) {
	deploymentController := controllers.DeploymentsController{}
	e.GET("/:ns/deployments", func(context echo.Context) error {
		var ns string
		if context.Param("ns") == "all" {
			ns = ""
		} else {
			ns = context.Param("ns")
		}

		selector := context.QueryParam("selector")
		if selector != "" {
			return deploymentController.GetBySelector(context, ns, selector)
		} else {
			return deploymentController.Get(context, ns)
		}

	})

	e.GET("/:ns/deployments/:id", func(context echo.Context) error {
		return deploymentController.GetOne(context, context.Param("ns"), context.Param("id"))
	})

	e.POST("/:ns/deployments", func(context echo.Context) error {
		deployment := utils.JsonBodyToMap(context.Request().Body)
		return deploymentController.Create(context, context.Param("ns"), deployment)
	})

	e.DELETE("/:ns/deployments/:id", func(context echo.Context) error {
		return deploymentController.Delete(context, context.Param("ns"), context.Param("id"))
	})

	e.PUT("/:ns/deployments", func(context echo.Context) error {
		deployment := utils.JsonBodyToMap(context.Request().Body)
		reDeployParam := context.QueryParam("restart")
		scaleParam := context.QueryParam("scale")
		if reDeployParam == "1" {
			return deploymentController.Restart(context, context.Param("ns"), deployment)
		}
		if scaleParam != "" {
			scale, err := strconv.ParseInt(scaleParam, 0, 32)
			if err == nil {
				return deploymentController.ReScale(context, context.Param("ns"), int32(scale), deployment)
			}
		}
		return deploymentController.Update(context, context.Param("ns"), deployment)
	})
}
