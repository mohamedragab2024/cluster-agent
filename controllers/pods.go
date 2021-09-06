package controllers

import (
	"fmt"
	"net/http"

	"github.com/kube-carbonara/cluster-agent/models"
	"github.com/labstack/echo/v4"
)

type PodsController struct {
}

func (c PodsController) Get(context echo.Context) error {
	fmt.Print("Getting from pods controller \n")
	return context.JSON(http.StatusOK, models.Response{
		ResourceType: "pods",
		Message:      "Test Data",
	})
}
