package routers

import (
	controllers "github.com/kube-carbonara/cluster-agent/controllers"
	"github.com/kube-carbonara/cluster-agent/utils"
	"github.com/labstack/echo/v4"
)

type MetricsRouter struct{}

func (router MetricsRouter) Handle(e *echo.Echo) {
	metricsController := controllers.MetricsController{}

	e.GET("/metrics/:resource", func(context echo.Context) error {
		switch context.Param("resource") {
		case utils.RESOUCETYPE_NODES:
			{
				return metricsController.NodeMetrics(context)

			}
		default:
			{
				return metricsController.ClusterMetrics(context)

			}

		}

	})

}
