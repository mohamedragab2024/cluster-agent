package routers

import (
	controllers "github.com/kube-carbonara/cluster-agent/controllers"
	"github.com/labstack/echo/v4"
)

type MetricsRouter struct{}

func (router MetricsRouter) Handle(e *echo.Echo) {
	metricsController := controllers.MetricsController{}

	e.GET("/metrics/:resource", func(context echo.Context) error {
		return metricsController.NodeMetrics(context)

	})

}
