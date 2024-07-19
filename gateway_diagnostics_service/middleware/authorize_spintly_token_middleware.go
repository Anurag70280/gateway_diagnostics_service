package middleware

import (
	"errors"
	"golang_service_template/models"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var (
	SPINTLY_PKEY string
)

func init() {

	SPINTLY_PKEY = os.Getenv("SPINTLY_PKEY")

}

func AuthorizeSpintlyToken() gin.HandlerFunc {

	return func(c *gin.Context) {

		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {

			apiError := models.ApiError{
				StatusCode: http.StatusUnauthorized,
				ApplicationError: models.ApplicationError{
					Type: "error",
					Message: models.ApplicationErrorMessage{
						ErrorCode:    1001,
						ErrorMessage: "Authorization token missing in header!",
					},
				},
			}
			c.Abort()
			c.JSON(apiError.StatusCode, apiError.ApplicationError)
			return

		}

		_, err := jwt.Parse(authHeader, func(token *jwt.Token) (interface{}, error) {

			_, ok := token.Method.(*jwt.SigningMethodRSA)
			if !ok {
				return nil, errors.New("Invalid token. Code 001!")
			}

			publicKey := SPINTLY_PKEY

			key, _ := jwt.ParseRSAPublicKeyFromPEM([]byte(publicKey))

			return key, nil

		})

		if err != nil {

			apiError := models.ApiError{
				StatusCode: http.StatusUnauthorized,
				ApplicationError: models.ApplicationError{
					Type: "error",
					Message: models.ApplicationErrorMessage{
						ErrorCode:    1002,
						ErrorMessage: err.Error(),
					},
				},
			}
			c.Abort()
			c.JSON(apiError.StatusCode, apiError.ApplicationError)
			return
		}

		//c.Next()

	}

}
