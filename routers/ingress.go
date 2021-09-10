package routers

import (
	controllers "github.com/kube-carbonara/cluster-agent/controllers"
	"github.com/kube-carbonara/cluster-agent/utils"
	"github.com/labstack/echo/v4"
)

type IngresRouter struct {
}

func (router IngresRouter) Handle(e *echo.Echo) {
	ingressController := controllers.IngressController{}

	e.GET("/:ns/ingress", func(context echo.Context) error {
		var ns string
		if context.Param("ns") == "all" {
			ns = ""
		} else {
			ns = context.Param("ns")
		}
		return ingressController.Get(context, ns)
	})

	e.GET("/:ns/ingress/:id", func(context echo.Context) error {
		return ingressController.GetOne(context, context.Param("ns"), context.Param("id"))
	})

	e.POST("/:ns/ingress", func(context echo.Context) error {
		deployment := utils.JsonBodyToMap(context.Request().Body)
		return ingressController.Create(context, context.Param("ns"), deployment)
	})

	e.DELETE("/:ns/ingress/:id", func(context echo.Context) error {
		return ingressController.Delete(context, context.Param("ns"), context.Param("id"))
	})

	e.PUT("/:ns/ingress", func(context echo.Context) error {
		deployment := utils.JsonBodyToMap(context.Request().Body)
		return ingressController.Update(context, context.Param("ns"), deployment)
	})

}
