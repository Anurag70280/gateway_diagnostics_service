package controllers

import (
	"fmt"
	"golang_service_template/models"
	"golang_service_template/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {

	scopes, exists := c.Get("scopes")

	if !exists {
		apiError := models.ApiError{
			StatusCode: http.StatusForbidden,
			ApplicationError: models.ApplicationError{
				Type: "error",
				Message: models.ApplicationErrorMessage{
					ErrorCode:    1007,
					ErrorMessage: "Error: scopes is not present in token!",
				},
			},
		}
		c.JSON(apiError.StatusCode, apiError.ApplicationError)
		return
	}

	_, ok := scopes.(int)

	if !ok {
		apiError := models.ApiError{
			StatusCode: http.StatusForbidden,
			ApplicationError: models.ApplicationError{
				Type: "error",
				Message: models.ApplicationErrorMessage{
					ErrorCode:    1007,
					ErrorMessage: "Error: scopes in token is not a string!",
				},
			},
		}
		c.JSON(apiError.StatusCode, apiError.ApplicationError)
		return
	}

	fmt.Printf("scopes value=%v\n", scopes)

	users, apiError := services.GetUsers()

	if apiError != nil {
		c.JSON(apiError.StatusCode, apiError.ApplicationError)
		return
	}

	c.JSON(http.StatusOK, models.SuccessResponse{Type: "success", Message: users})
	return

}

func GetUser(c *gin.Context) {

	userId, err := strconv.ParseInt(c.Param("user_id"), 10, 32)

	if err != nil {

		apiError := models.ApiError{
			StatusCode: http.StatusBadRequest,
			ApplicationError: models.ApplicationError{
				Type: "error",
				Message: models.ApplicationErrorMessage{
					ErrorCode:    1005,
					ErrorMessage: fmt.Sprintf("%s", err),
				},
			},
		}

		c.JSON(apiError.StatusCode, apiError.ApplicationError)
		return

	}

	user, apiError := services.GetUser(int(userId))

	if apiError != nil {
		c.JSON(apiError.StatusCode, apiError.ApplicationError)
		return
	}

	c.JSON(http.StatusOK, models.SuccessResponse{Type: "success", Message: user})
	return

}

func DeleteUser(c *gin.Context) {

	userId, err := strconv.ParseInt(c.Param("user_id"), 10, 32)

	if err != nil {

		apiError := models.ApiError{
			StatusCode: http.StatusBadRequest,
			ApplicationError: models.ApplicationError{
				Type: "error",
				Message: models.ApplicationErrorMessage{
					ErrorCode:    1005,
					ErrorMessage: fmt.Sprintf("%s", err),
				},
			},
		}

		c.JSON(apiError.StatusCode, apiError.ApplicationError)
		return

	}

	msg, apiError := services.DeleteUser(int(userId))

	if apiError != nil {
		c.JSON(apiError.StatusCode, apiError.ApplicationError)
		return
	}

	c.JSON(http.StatusOK, models.SuccessResponse{Type: "success", Message: msg})
	return

}

func CreateUser(c *gin.Context) {

	var input models.CreateUserRequest

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

	msg, apiError := services.CreateUser(input)

	if apiError != nil {
		c.JSON(apiError.StatusCode, apiError.ApplicationError)
		return
	}

	c.JSON(http.StatusOK, models.SuccessResponse{Type: "success", Message: msg})
	return

}
