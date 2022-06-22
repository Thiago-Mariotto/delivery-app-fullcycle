package kafka

import (
	"log"
	"os"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
)

// consome os dados da nossa fila

type KafkaConsumer struct {
	MsgChannel chan *ckafka.Message
}

func (k *KafkaConsumer) Consume() {
	configmap := &ckafka.ConfigMap{
		"bootstrap.servers": os.Getenv("KafkaBootstrapServers"),
		"group.id":          os.Getenv("KafkaConsumerGroupId"),
	}

	c, err := ckafka.NewConsumer(configmap)
	if err != nil {
		log.Fatalf("error consuming kafka message: " + err.Error())
	}

	topics := []string{os.Getenv("KafkaReadTopic")}
	c.SubscribeTopics(topics, nil)
}
