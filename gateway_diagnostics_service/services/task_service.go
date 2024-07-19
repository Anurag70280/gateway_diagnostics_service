package services

import (
	"encoding/json"
	"golang_service_template/clients"
	"golang_service_template/logger"
	"golang_service_template/models"

	"go.uber.org/zap"
)

func CreateTask(task models.CreateTaskRequest) (*string, *models.ApiError) {

	logger.Log.Debug("Task details:", zap.Any("task:", task))

	txByteSlice, err := json.Marshal(task)
	if err != nil {
		logger.Log.Debug("Error:" + err.Error())
	} else {

		taskKafkaMsg := clients.ToKafkaMessage{
			Topic: "MyTopic5",
			Key:   "001",
			Value: txByteSlice,
		}

		clients.ToKafkaCh1 <- taskKafkaMsg

	}

	return nil, nil
}

func CreateMessage(msg models.CreateMessageRequest) (*string, *models.ApiError) {

	txByteSlice, err := json.Marshal(msg)
	if err != nil {
		logger.Log.Debug("Error:" + err.Error())
	} else {

		kafkaMsg := clients.ToKafkaMessage{
			Topic: "infraServiceAckEvent",
			Key:   "001",
			Value: txByteSlice,
		}

		clients.ToKafkaCh1 <- kafkaMsg

	}

	return nil, nil
}

func CreateAck(msg models.AckEvent) (*string, *models.ApiError) {

	txByteSlice, err := json.Marshal(msg)
	if err != nil {
		logger.Log.Debug("Error:" + err.Error())
	} else {

		kafkaMsg := clients.ToKafkaMessage{
			Topic: "infraServiceResourcesAckEvent",
			Key:   "001",
			Value: txByteSlice,
		}

		clients.ToKafkaCh1 <- kafkaMsg

	}

	return nil, nil
}

func CreateDcm(msg models.DcmEvent) (*string, *models.ApiError) {

	txByteSlice, err := json.Marshal(msg)
	if err != nil {
		logger.Log.Debug("Error:" + err.Error())
	} else {

		kafkaMsg := clients.ToKafkaMessage{
			Topic: "deviceConfigurationManagementEvent",
			Key:   "001",
			Value: txByteSlice,
		}

		clients.ToKafkaCh1 <- kafkaMsg

	}

	return nil, nil
}
