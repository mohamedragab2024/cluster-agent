package routers

import (
	"fmt"

	controllers "github.com/kube-carbonara/cluster-agent/controllers"
	"github.com/labstack/echo/v4"
)

type PodsRouter struct {
}

func (router PodsRouter) Handle(e *echo.Echo) {
	podsController := controllers.PodsController{}
	e.GET("/pods", func(context echo.Context) error {
		fmt.Print("get pods \n")
		return podsController.Get(context)
	})
}
