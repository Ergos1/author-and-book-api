package kafkarequestlogger

import (
	"context"
	"encoding/json"
	"log"

	"github.com/Shopify/sarama"
)

const (
	topic = "requestLogger"
)

type KafkaProducer interface {
	SendSyncMessage(message *sarama.ProducerMessage) (partition int32, offset int64, err error)
}

type KafkaConsumer interface {
	Subscribe(ctx context.Context, topic string, handler func(message *sarama.ConsumerMessage)) error
}

type KafkaRequestLogger struct {
	producer KafkaProducer
	consumer KafkaConsumer
}

func NewKafkaRequestLogger(producer KafkaProducer, consumer KafkaConsumer) *KafkaRequestLogger {
	return &KafkaRequestLogger{
		producer: producer,
		consumer: consumer,
	}
}

type KafkaMessageValue struct {
	Method string `json:"method"`
	Url    string `json:"url"`
	Body   any    `json:"body"`
}

func (krl *KafkaRequestLogger) buildMessage(method string, url string, body any) (*sarama.ProducerMessage, error) {
	messageValue := &KafkaMessageValue{
		Method: method,
		Body:   body,
		Url:    url,
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

func (krl *KafkaRequestLogger) Log(method string, url string, body any) error {
	message, err := krl.buildMessage(method, url, body)
	if err != nil {
		return err
	}

	if _, _, err := krl.producer.SendSyncMessage(message); err != nil {
		return err
	}

	return nil
}

func (krl *KafkaRequestLogger) Subscribe(ctx context.Context) error {
	log.SetPrefix("[KafkaRequestLogger.Subscribe] ")
	krl.consumer.Subscribe(ctx, topic, func(message *sarama.ConsumerMessage) {

		log.Printf("\nRequest timestamp: %v\nRequest metadata: %v\n", message.Timestamp, string(message.Value))
	})

	return nil
}
