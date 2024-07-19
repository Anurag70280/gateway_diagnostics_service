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
	AWS_KID1, AWS_PKEY1, AWS_KID2, AWS_PKEY2 string
)

func init() {

	AWS_KID1 = os.Getenv("AWS_KID1")
	AWS_PKEY1 = os.Getenv("AWS_PKEY1")
	AWS_KID2 = os.Getenv("AWS_KID2")
	AWS_PKEY2 = os.Getenv("AWS_PKEY2")

}

func AuthorizeAwsToken() gin.HandlerFunc {

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

		token, err := jwt.Parse(authHeader, func(token *jwt.Token) (interface{}, error) {

			_, ok := token.Method.(*jwt.SigningMethodRSA)
			if !ok {
				return nil, errors.New("Invalid token. Code 001!")
			}

			keyId, keyFound := token.Header["kid"]

			if !keyFound {
				return nil, errors.New("Invalid token. Code 002!")
			}

			keyIdStr, keyIdIsStr := keyId.(string)

			if !keyIdIsStr {
				return nil, errors.New("Invalid token. Code 003!")
			}

			var publicKey string

			if keyIdStr == AWS_KID1 {
				publicKey = AWS_PKEY1
			} else if keyIdStr == AWS_KID2 {
				publicKey = AWS_PKEY2
			} else {
				return nil, errors.New("Invalid token. Code 004!")
			}

			key, err := jwt.ParseRSAPublicKeyFromPEM([]byte(publicKey))

			return key, err

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

		claims, claimsIsOk := token.Claims.(jwt.MapClaims)

		//fmt.Printf("sub type=%T\n", claims["custom:userScopes"])

		if !claimsIsOk {
			apiError := models.ApiError{
				StatusCode: http.StatusUnauthorized,
				ApplicationError: models.ApplicationError{
					Type: "error",
					Message: models.ApplicationErrorMessage{
						ErrorCode:    1003,
						ErrorMessage: "Error: Not able to verfiy token claim",
					},
				},
			}

			c.Abort()
			c.JSON(apiError.StatusCode, apiError.ApplicationError)
			return
		}

		c.Set("sub", claims["sub"])
		c.Set("scopes", claims["custom:userScopes"])

	}

}
