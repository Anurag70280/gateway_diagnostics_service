package services

import (
	"encoding/json"
	"fmt"
	"golang_service_template/clients"
	"golang_service_template/logger"
	"golang_service_template/models"
	"net/http"
	"os"
	"time"
)

func GetDevices() (models.Devices, *models.ApiError) {

	statusCode := 0

	headers := map[string]string{"Content-Type": "application/json", "x-api-key": os.Getenv("DEVICES_SERVICE_X_API_KEY")}

	url := fmt.Sprintf("%s/mainService/v1/data", os.Getenv("DEVICES_SERVICE_BASE_URL"))

	statusCode, responseBody, err := clients.RestClient.Get(url, headers, time.Duration(time.Duration.Seconds(10)))

	if err != nil {
		logger.Log.Error(err.Error())
		apiError := models.ApiError{
			StatusCode: http.StatusInternalServerError,
			ApplicationError: models.ApplicationError{
				Type: "error",
				Message: models.ApplicationErrorMessage{
					ErrorCode:    11008,
					ErrorMessage: fmt.Sprintf("%s", err),
				},
			},
		}
		return models.Devices{}, &apiError
	}

	if statusCode != 200 {

		apiError := models.ApiError{
			StatusCode: http.StatusInternalServerError,
			ApplicationError: models.ApplicationError{
				Type: "error",
				Message: models.ApplicationErrorMessage{
					ErrorCode:          11009,
					ErrorMessage:       "Unexpected response from device service!",
					OriginStatusCode:   statusCode,
					OriginErrorMessage: string(responseBody),
				},
			},
		}

		return models.Devices{}, &apiError
	}

	getDevicesResponse := models.GetDevicesResponse{}

	err = json.Unmarshal(responseBody, &getDevicesResponse)

	if err != nil {
		formattedMessage := fmt.Sprintf("Could not convert expected api response message to json format.Error:%s", err)
		logger.Log.Error(formattedMessage)
		apiError := models.ApiError{
			StatusCode: http.StatusInternalServerError,
			ApplicationError: models.ApplicationError{
				Type: "error",
				Message: models.ApplicationErrorMessage{
					ErrorCode:    11010,
					ErrorMessage: formattedMessage,
				},
			},
		}
		return models.Devices{}, &apiError
	}

	devices := getDevicesResponse.Message

	return devices, nil
}
