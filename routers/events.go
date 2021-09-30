package routers

import (
	controllers "github.com/kube-carbonara/cluster-agent/controllers"
	"github.com/labstack/echo/v4"
)

type EventsRouter struct{}

func (router EventsRouter) Handle(e *echo.Echo) {
	eventController := controllers.EventsController{}
	e.GET("/:ns/events", func(context echo.Context) error {
		var ns string
		if context.Param("ns") == "all" {
			ns = ""
		} else {
			ns = context.Param("ns")
		}
		return eventController.Get(context, ns)
	})

	e.GET("/:ns/events/:id", func(context echo.Context) error {
		return eventController.GetOne(context, context.Param("ns"), context.Param("id"))
	})

}
