package utils

import (
	"fmt"
	"golang_service_template/models"
	"os"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/joho/godotenv"
)

func init() {

	if _, err := os.Stat(".env"); err == nil {
		err = godotenv.Load(".env")
		if err != nil {
			panic(err)
		}
	}
}

func ShowServiceInfo() {

	serviceName := os.Getenv("SERVICE_NAME")
	fmt.Printf("Starting service %s!\n", serviceName)

	envName := os.Getenv("ENV")
	fmt.Printf("Environment is %s!\n", envName)

}

func ProcessError(errMsg models.ProcessErrorMessage, data interface{}) {

	switch value := data.(type) {

	case *kafka.Message:
		fmt.Printf("Kafka message error,Error:%s,Data:%v\n", errMsg.Error.Error(), data)

	default:
		fmt.Printf("Unsupported message type:%v,Error:%s,Data:%v\n", value, errMsg.Error.Error(), data)

	}

}
