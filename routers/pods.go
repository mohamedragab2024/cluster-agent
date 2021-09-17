package routers

import (
	controllers "github.com/kube-carbonara/cluster-agent/controllers"
	"github.com/kube-carbonara/cluster-agent/utils"
	"github.com/labstack/echo/v4"
)

type PodsRouter struct {
}

func (router PodsRouter) Handle(e *echo.Echo) {
	podsController := controllers.PodsController{}
	e.GET("/:ns/pods", func(context echo.Context) error {
		var ns string
		if context.Param("ns") == "all" {
			ns = ""
		} else {
			ns = context.Param("ns")
		}
		return podsController.Get(context, ns)
	})

	e.GET("/:ns/pods/:id", func(context echo.Context) error {
		return podsController.GetOne(context, context.Param("ns"), context.Param("id"))
	})

	e.POST("/:ns/pods", func(context echo.Context) error {
		deployment := utils.JsonBodyToMap(context.Request().Body)
		return podsController.Create(context, context.Param("ns"), deployment)
	})

	e.DELETE("/:ns/pods/:id", func(context echo.Context) error {
		return podsController.Delete(context, context.Param("ns"), context.Param("id"))
	})

	e.PUT("/:ns/pods", func(context echo.Context) error {
		deployment := utils.JsonBodyToMap(context.Request().Body)
		return podsController.Update(context, context.Param("ns"), deployment)
	})
}
