package routers

import (
	controllers "github.com/kube-carbonara/cluster-agent/controllers"
	"github.com/labstack/echo/v4"
)

type AppsRouter struct{}

func (router AppsRouter) Handle(e *echo.Echo) {
	appsController := controllers.AppsController{}
	e.GET("/apps", func(context echo.Context) error {
		return appsController.Get(context)
	})

}
