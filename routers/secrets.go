package routers

import (
	controllers "github.com/kube-carbonara/cluster-agent/controllers"
	"github.com/kube-carbonara/cluster-agent/utils"
	"github.com/labstack/echo/v4"
)

type SecretRouter struct{}

func (router SecretRouter) Handle(e *echo.Echo) {
	secretController := controllers.SecretsController{}
	e.GET("/:ns/secrets", func(context echo.Context) error {
		var ns string
		if context.Param("ns") == "all" {
			ns = ""
		} else {
			ns = context.Param("ns")
		}
		return secretController.Get(context, ns)
	})

	e.GET("/:ns/secrets/:id", func(context echo.Context) error {
		return secretController.GetOne(context, context.Param("ns"), context.Param("id"))
	})

	e.POST("/:ns/secrets", func(context echo.Context) error {
		deployment := utils.JsonBodyToMap(context.Request().Body)
		return secretController.Create(context, context.Param("ns"), deployment)
	})

	e.DELETE("/:ns/secrets/:id", func(context echo.Context) error {
		return secretController.Delete(context, context.Param("ns"), context.Param("id"))
	})

	e.PUT("/:ns/secrets", func(context echo.Context) error {
		deployment := utils.JsonBodyToMap(context.Request().Body)
		return secretController.Update(context, context.Param("ns"), deployment)
	})
}
