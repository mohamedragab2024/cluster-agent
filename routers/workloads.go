package routers

import (
	controllers "github.com/kube-carbonara/cluster-agent/controllers"
	"github.com/labstack/echo/v4"
)

type WorkLoadsRouter struct {
}

func (router WorkLoadsRouter) Handle(e *echo.Echo) {
	workLoadController := controllers.WorkLoadController{}
	e.GET("/:ns/workloads", func(context echo.Context) error {
		var ns string
		if context.Param("ns") == "all" {
			ns = ""
		} else {
			ns = context.Param("ns")
		}

		selector := context.QueryParam("selector")
		if selector != "" {
			return workLoadController.GetBySelector(context, ns, selector)
		} else {
			return workLoadController.Get(context, ns)
		}

	})
}
