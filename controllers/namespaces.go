package controllers

import (
	"net/http"

	"github.com/kube-carbonara/cluster-agent/models"
	"github.com/labstack/echo/v4"
)

type NameSpacesController struct {
}

func (c NameSpacesController) Get(context echo.Context) error {

	return context.JSON(http.StatusOK, models.Response{
		ResourceType: "namespace",
		Status:       200,
		Message:      "Test Data",
	})
}
