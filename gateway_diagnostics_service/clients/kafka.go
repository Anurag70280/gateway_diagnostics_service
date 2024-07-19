package clients

import (
	"fmt"
	"golang_service_template/logger"
	"golang_service_template/models"
	"golang_service_template/utils"
	"os"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

type callbackFunctionWithMsg func(*kafka.Message) error

type ToKafkaMessage struct {
	Topic string
	Key   string
	Value []byte
}

var (
	ToKafkaCh1 = make(chan ToKafkaMessage)
	ToKafkaCh2 = make(chan ToKafkaMessage)
)

func KafkaConsumer(kafkaConsumerGroup string, kafkaTopicName string, callbackFunction callbackFunctionWithMsg) {

	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers":  os.Getenv("KAFKA_BOOTSTRAP_SERVERS"),
		"group.id":           kafkaConsumerGroup,
		"auto.offset.reset":  "earliest",
		"enable.auto.commit": false,
		//"sasl.mechanisms":   "SCRAM-SHA-512",
		//"security.protocol": "sasl_ssl",
		//"sasl.username":     KafkaUsername,
		//"sasl.password":     KafkaPassword,
	})

	if err != nil {
		logger.Log.Error(err.Error())
		utils.ProcessError(models.ProcessErrorMessage{Priority: 1, Error: err}, nil)
		panic(err)
	}

	err = c.SubscribeTopics([]string{kafkaTopicName}, nil)

	if err != nil {
		logger.Log.Error(err.Error())
		utils.ProcessError(models.ProcessErrorMessage{Priority: 1, Error: err}, nil)
		panic(err)
	}

	for {
		msg, err := c.ReadMessage(-1)
		if err == nil {
			err = callbackFunction(msg)
			if err == nil {
				c.CommitMessage(msg)
			} else {
				logger.Log.Error(err.Error())
				c.Close()
				utils.ProcessError(models.ProcessErrorMessage{Priority: 1, Error: err}, msg)
				return
			}
		} else {
			// The client will automatically try to recover from all errors.
			errorMsg := fmt.Sprintf("Consumer error:%v,Consumer error:%v\n", err, msg)
			logger.Log.Error(errorMsg)
			utils.ProcessError(models.ProcessErrorMessage{Priority: 2, Error: err}, msg)
		}
	}

}

func KafkaProducer(producerChannel <-chan ToKafkaMessage) {

	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": os.Getenv("KAFKA_BOOTSTRAP_SERVERS"),
		//"sasl.mechanisms":   "SCRAM-SHA-512",
		//"security.protocol": "sasl_ssl",
		//"sasl.username":     KafkaUsername,
		//"sasl.password":     KafkaPassword,
	})

	if err != nil {
		logger.Log.Error(err.Error())
		panic(err)
	}

	deliveryChan := make(chan kafka.Event)

	for {

		message := <-producerChannel

		err := p.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &(message.Topic), Partition: kafka.PartitionAny},
			Key:            []byte(message.Key),
			Value:          []byte(message.Value),
		}, deliveryChan)

		kafkaEvent := <-deliveryChan
		kafkaMessage := kafkaEvent.(*kafka.Message)

		if kafkaMessage.TopicPartition.Error != nil {
			fmt.Printf("Failed to send message: %v\n", kafkaMessage.TopicPartition.Error)
		} else {
			fmt.Printf("Sent message to topic %s [%d] at offset %v\n", *kafkaMessage.TopicPartition.Topic, kafkaMessage.TopicPartition.Partition, kafkaMessage.TopicPartition.Offset)
		}

		if err != nil {
			errorMsg := fmt.Sprintf("Producer error: %v\n", err)
			logger.Log.Error(errorMsg)
		}

	}

}
