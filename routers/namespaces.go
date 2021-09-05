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
}
