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

	e.POST("/:ns/ingress", func(context echo.Context) error {
		ingress := utils.JsonBodyToMap(context.Request().Body)
		return ingressController.Create(context, context.Param("ns"), ingress)
	})

}
