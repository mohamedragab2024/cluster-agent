package routers

import (
	controllers "github.com/kube-carbonara/cluster-agent/controllers"
	"github.com/kube-carbonara/cluster-agent/utils"
	"github.com/labstack/echo/v4"
)

type SeviceRouter struct{}

func (router SeviceRouter) Handle(e *echo.Echo) {
	serviceController := controllers.ServicesController{}
	e.GET("/:ns/services", func(context echo.Context) error {
		return serviceController.Get(context, context.Param("ns"))
	})

	e.GET("/:ns/services/:id", func(context echo.Context) error {
		return serviceController.GetOne(context, context.Param("ns"), context.Param("id"))
	})

	e.POST("/:ns/services", func(context echo.Context) error {
		deployment := utils.JsonBodyToMap(context.Request().Body)
		return serviceController.Create(context, context.Param("ns"), deployment)
	})

	e.DELETE("/:ns/services/:id", func(context echo.Context) error {
		return serviceController.Delete(context, context.Param("ns"), context.Param("id"))
	})

	e.PUT("/:ns/services", func(context echo.Context) error {
		deployment := utils.JsonBodyToMap(context.Request().Body)
		return serviceController.Update(context, context.Param("ns"), deployment)
	})
}
