package controllers

import (
	"fmt"
	"golang_service_template/models"
	"golang_service_template/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateTask(c *gin.Context) {

	var input models.CreateTaskRequest

	if err := c.ShouldBindJSON(&input); err != nil {

		apiError := models.ApiError{
			StatusCode: http.StatusBadRequest,
			ApplicationError: models.ApplicationError{
				Type: "error",
				Message: models.ApplicationErrorMessage{
					ErrorCode:    1006,
					ErrorMessage: fmt.Sprintf("%s", err),
				},
			},
		}

		c.JSON(apiError.StatusCode, apiError.ApplicationError)
		return

	}

	msg, apiError := services.CreateTask(input)

	if apiError != nil {
		c.JSON(apiError.StatusCode, apiError.ApplicationError)
		return
	}

	c.JSON(http.StatusOK, models.SuccessResponse{Type: "success", Message: msg})
	return

}

func CreateMessage(c *gin.Context) {

	var input models.CreateMessageRequest

	if err := c.ShouldBindJSON(&input); err != nil {

		apiError := models.ApiError{
			StatusCode: http.StatusBadRequest,
			ApplicationError: models.ApplicationError{
				Type: "error",
				Message: models.ApplicationErrorMessage{
					ErrorCode:    1006,
					ErrorMessage: fmt.Sprintf("%s", err),
				},
			},
		}

		c.JSON(apiError.StatusCode, apiError.ApplicationError)
		return

	}

	msg, apiError := services.CreateMessage(input)

	if apiError != nil {
		c.JSON(apiError.StatusCode, apiError.ApplicationError)
		return
	}

	c.JSON(http.StatusOK, models.SuccessResponse{Type: "success", Message: msg})
	return

}

func CreateAck(c *gin.Context) {

	var input models.AckEvent

	if err := c.ShouldBindJSON(&input); err != nil {

		apiError := models.ApiError{
			StatusCode: http.StatusBadRequest,
			ApplicationError: models.ApplicationError{
				Type: "error",
				Message: models.ApplicationErrorMessage{
					ErrorCode:    1006,
					ErrorMessage: fmt.Sprintf("%s", err),
				},
			},
		}

		c.JSON(apiError.StatusCode, apiError.ApplicationError)
		return

	}

	msg, apiError := services.CreateAck(input)

	if apiError != nil {
		c.JSON(apiError.StatusCode, apiError.ApplicationError)
		return
	}

	c.JSON(http.StatusOK, models.SuccessResponse{Type: "success", Message: msg})
	return

}

func CreateDcm(c *gin.Context) {

	var input models.DcmEvent

	if err := c.ShouldBindJSON(&input); err != nil {

		apiError := models.ApiError{
			StatusCode: http.StatusBadRequest,
			ApplicationError: models.ApplicationError{
				Type: "error",
				Message: models.ApplicationErrorMessage{
					ErrorCode:    1006,
					ErrorMessage: fmt.Sprintf("%s", err),
				},
			},
		}

		c.JSON(apiError.StatusCode, apiError.ApplicationError)
		return

	}

	msg, apiError := services.CreateDcm(input)

	if apiError != nil {
		c.JSON(apiError.StatusCode, apiError.ApplicationError)
		return
	}

	c.JSON(http.StatusOK, models.SuccessResponse{Type: "success", Message: msg})
	return

}
