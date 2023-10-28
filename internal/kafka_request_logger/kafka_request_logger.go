package kafkarequestlogger

import (
	"encoding/json"

	"github.com/Shopify/sarama"
)

const (
	topic = "requestLogger"
)

type KafkaProducer interface {
	SendSyncMessage(message *sarama.ProducerMessage) (partition int32, offset int64, err error)
}

type KafkaRequestLogger struct {
	producer KafkaProducer
}

func NewKafkaRequestLogger(producer KafkaProducer) *KafkaRequestLogger {
	return &KafkaRequestLogger{
		producer: producer,
	}
}

type KafkaMessageValue struct {
	Method string `json:"method"`
	Body   any    `json:"body"`
}

func (krl *KafkaRequestLogger) buildMessage(method string, body any) (*sarama.ProducerMessage, error) {
	messageValue := &KafkaMessageValue{
		Method: method,
		Body:   body,
	}

	messageValueStr, err := json.Marshal(messageValue)
	if err != nil {
		return nil, err
	}

	return &sarama.ProducerMessage{
		Topic:     topic,
		Value:     sarama.ByteEncoder(messageValueStr),
		Partition: -1,
	}, nil
}

func (krl *KafkaRequestLogger) Log(method string, body any) error {
	message, err := krl.buildMessage(method, body)
	if err != nil {
		return err
	}

	if _, _, err := krl.producer.SendSyncMessage(message); err != nil {
		return err
	}

	return nil
}
