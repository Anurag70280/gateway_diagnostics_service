package test

import (
	"encoding/json"
	"golang_service_template/app"
	"golang_service_template/models"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetDevicesNoErrors(t *testing.T) {

	assert := assert.New(t)

	app.SetupCognitoRoutes()

	funcRestClientMockResponse = func() (statusCode int, responseBody []byte, err error) {

		var responseBodyJson models.GetDevicesResponse

		responseBodyJson = models.GetDevicesResponse{
			Type:    "success",
			Message: models.Devices{[]models.Device{models.Device{SerialNumber: "1001A0111234", Name: "malcolm"}}},
		}

		responseBody, _ = json.Marshal(responseBodyJson)

		return 200, responseBody, nil
	}

	resWriter := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/devices", nil)
	req.Header.Set("Content-Type", "application/json")
	app.Router.ServeHTTP(resWriter, req)

	getDevicesResponse := models.GetDevicesResponse{}
	json.Unmarshal(resWriter.Body.Bytes(), &getDevicesResponse)

	assert.Equal(200, resWriter.Code)
	assert.Equal("success", getDevicesResponse.Type)
	assert.Equal("1001A0111234", getDevicesResponse.Message.Devices[0].SerialNumber)
	assert.Equal("malcolm", getDevicesResponse.Message.Devices[0].Name)

}
