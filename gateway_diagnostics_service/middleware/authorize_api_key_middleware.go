package middleware

import (
	"golang_service_template/models"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

var (
	API_KEY string
)

func init() {

	API_KEY = os.Getenv("API_KEY")

}

func AuthorizeApiKey() gin.HandlerFunc {

	return func(c *gin.Context) {

		xApiKey := c.GetHeader("x-api-key")

		if xApiKey == "" {

			apiError := models.ApiError{
				StatusCode: http.StatusUnauthorized,
				ApplicationError: models.ApplicationError{
					Type: "error",
					Message: models.ApplicationErrorMessage{
						ErrorCode:    1001,
						ErrorMessage: "X-API-KEY missing in header!",
					},
				},
			}
			c.Abort()
			c.JSON(apiError.StatusCode, apiError.ApplicationError)
			return

		}

		if xApiKey != API_KEY {

			apiError := models.ApiError{
				StatusCode: http.StatusUnauthorized,
				ApplicationError: models.ApplicationError{
					Type: "error",
					Message: models.ApplicationErrorMessage{
						ErrorCode:    1002,
						ErrorMessage: "API key mismatch!",
					},
				},
			}
			c.Abort()
			c.JSON(apiError.StatusCode, apiError.ApplicationError)
			return
		}

		c.Next()

	}

}
