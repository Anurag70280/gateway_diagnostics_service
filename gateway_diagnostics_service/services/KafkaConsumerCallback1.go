package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"golang_service_template/clients"
	"golang_service_template/logger"
	"golang_service_template/models"
	"golang_service_template/utils"
	"os"
	"time"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

func KafkaConsumerCallback1(msg *kafka.Message) (err error) {

	//fmt.Printf("Message Type:%T,Message Value:%v\n", msg.Value, msg.Value)

	accessHistoryMsg := models.AccessHistory{}
	err = json.Unmarshal(msg.Value, &accessHistoryMsg)

	if err != nil {
		logger.Log.Error(err.Error())
		utils.ProcessError(models.ProcessErrorMessage{Priority: 2, Error: err}, msg.Value)
		return nil
	}

	err = Validate(accessHistoryMsg)

	if err != nil {
		logger.Log.Error(err.Error())
		utils.ProcessError(models.ProcessErrorMessage{Priority: 2, Error: err}, msg.Value)
		return nil
	}

	headers := map[string]string{"Content-Type": "application/json", "x-api-key": os.Getenv("DEVICES_SERVICE_X_API_KEY")}

	url := fmt.Sprintf("%s/mainService/v1/data", os.Getenv("DEVICES_SERVICE_BASE_URL"))

	statusCode, responseBody, err := clients.RestClient.Post(url, headers, accessHistoryMsg, time.Duration(time.Duration.Seconds(10)))

	if err != nil {
		logger.Log.Error(err.Error())
		utils.ProcessError(models.ProcessErrorMessage{Priority: 2, Error: err}, msg.Value)
		return nil
	}

	if statusCode != 200 {

		errorMsg := fmt.Sprintf("API Error,Error Code:%v,API Response%v\n", statusCode, responseBody)
		logger.Log.Error(errorMsg)
		utils.ProcessError(models.ProcessErrorMessage{Priority: 1, Error: errors.New(errorMsg)}, msg.Value)
		return nil
	}

	strResponseBody := string(responseBody)

	logger.Log.Debug(strResponseBody)

	return nil
}

func Validate(ah models.AccessHistory) error {

	if ah.OrgId == 0 {
		return fmt.Errorf("missing 'orgId' key")
	}
	if ah.MsgType == "" {
		return fmt.Errorf("missing 'msgType' key")
	}
	if ah.AccessPointId == 0 {
		return fmt.Errorf("missing 'AccessPointId' key")
	}
	if ah.UserId == 0 {
		return fmt.Errorf("missing 'UserId' key")
	}
	return nil

}
