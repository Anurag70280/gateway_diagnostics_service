package services

import (
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

func KafkaConsumerCallback2(msg *kafka.Message) (err error) {
	fmt.Printf("Message:%s\n", msg.Value)
	return nil
}
