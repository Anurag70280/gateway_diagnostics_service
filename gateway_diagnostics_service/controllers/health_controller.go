package controllers

import (
	"golang_service_template/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetHealth(c *gin.Context) {

	c.JSON(http.StatusOK, models.SuccessResponse{Type: "success", Message: "healthy"})
	return

}
