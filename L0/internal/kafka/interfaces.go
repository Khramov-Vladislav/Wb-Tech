package kafka

import (
	"context"
	"microservice/internal/models"
)

// ProducerInterface описывает методы Kafka-продюсера
type ProducerInterface interface {
	SendOrder(ctx context.Context, order *models.Order) error
	Close() error
}

// ConsumerInterface описывает методы Kafka-консьюмера
type ConsumerInterface interface {
	Start(ctx context.Context)
	Close() error
}
