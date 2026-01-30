package kafka

import (
	"context"
	"encoding/json"
	"time"

	"microservice/internal/models"

	"github.com/segmentio/kafka-go"
)

// Producer оборачивает Kafka writer
type Producer struct {
	topic  string
	writer *kafka.Writer
}

// NewProducer создает новый Kafka продюсер 
func NewProducer(brokers []string, topic string) ProducerInterface {
	writer := &kafka.Writer{
		Addr:     kafka.TCP(brokers...),
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
		Async:    false,
	}
	return &Producer{
		topic:  topic,
		writer: writer,
	}
}

// SendOrder отправляет заказ в Kafka
func (p *Producer) SendOrder(ctx context.Context, order *models.Order) error {
	data, err := json.Marshal(order)
	if err != nil {
		return err
	}

	msg := kafka.Message{
		Key:   []byte(order.OrderUID),
		Value: data,
		Time:  time.Now(),
	}

	return p.writer.WriteMessages(ctx, msg)
}

// Close закрывает Kafka writer
func (p *Producer) Close() error {
	return p.writer.Close()
}
