package routers

import (
	controllers "github.com/kube-carbonara/cluster-agent/controllers"
	"github.com/labstack/echo/v4"
)

type NameSpacesRouter struct {
}

func (router NameSpacesRouter) Handle(e *echo.Echo) {
	nameSpacesController := controllers.NameSpacesController{}
	e.GET("/namespaces", func(context echo.Context) error {
		return nameSpacesController.Get(context)
	})

	e.GET("/namespaces/:id", func(context echo.Context) error {
		return nameSpacesController.GetOne(context, context.Param("id"))
	})
	e.POST("/namespaces/:id", func(context echo.Context) error {
		return nameSpacesController.Create(context, context.Param("id"))
	})

	e.DELETE("/namespaces/:id", func(context echo.Context) error {
		return nameSpacesController.Delete(context, context.Param("id"))
	})
}
