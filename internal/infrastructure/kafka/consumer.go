package kafka

import (
	"context"
	"fmt"
	"time"

	"github.com/Shopify/sarama"
)

type Consumer struct {
	brokers        []string
	SingleConsumer sarama.Consumer
}

// type ConsumerMessageHandler = func(message *sarama.ConsumerMessage)

func NewConsumer(brokers []string) (*Consumer, error) {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = false
	config.Consumer.Offsets.AutoCommit.Enable = true
	config.Consumer.Offsets.AutoCommit.Interval = 5 * time.Second
	/*
		sarama.OffsetNewest - получаем только новые сообщений, те, которые уже были игнорируются
		sarama.OffsetOldest - читаем все с самого начала
	*/
	config.Consumer.Offsets.Initial = sarama.OffsetOldest

	consumer, err := sarama.NewConsumer(brokers, config)

	if err != nil {
		return nil, err
	}

	/*
		consumer.Topics() - список топиков
		consumer.Partitions("test_topic") - партиции топика
		consumer.ConsumePartition("test_topic", 1, 12) - чтение конкретного топика с 12 сдвига в первой партиции
		consumer.Pause() - останавливаем чтение определенных топиков
		consumer.Resume() - восстанавливаем чтение определенных топиков
		consumer.PauseAll() - останавливаем чтение всех топиков
		consumer.ResumeAll() - восстанавливаем чтение всех топиков
	*/

	return &Consumer{
		brokers:        brokers,
		SingleConsumer: consumer,
	}, err
}

func (c *Consumer) Subscribe(ctx context.Context, topic string, handler func(message *sarama.ConsumerMessage)) error {
	partitionList, err := c.SingleConsumer.Partitions(topic)

	if err != nil {
		return err
	}

	initialOffset := sarama.OffsetOldest

	for _, partition := range partitionList {
		pc, err := c.SingleConsumer.ConsumePartition(topic, partition, initialOffset)

		if err != nil {
			return err
		}

		go func(pc sarama.PartitionConsumer, partition int32) {
			for {
				select {
				case message, ok := <-pc.Messages():
					if !ok {
						return
					}

					fmt.Println("Read Topic: ", topic, " Partition: ", partition, " Offset: ", message.Offset)
					fmt.Println("Received Key: ", string(message.Key), " Value: ", string(message.Value))
					handler(message)

				case <-ctx.Done():
					fmt.Println("[Consumer] stop listenning")
					return
				}
			}
		}(pc, partition)
	}

	return nil

}
