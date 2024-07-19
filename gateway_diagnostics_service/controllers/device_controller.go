package controllers

import (
	"golang_service_template/models"
	"golang_service_template/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetDevices(c *gin.Context) {

	devices, apiError := services.GetDevices()

	if apiError != nil {
		c.JSON(apiError.StatusCode, apiError.ApplicationError)
		return
	}

	c.JSON(http.StatusOK, models.SuccessResponse{Type: "success", Message: devices})
	return

}
