package kafka

import (
	"context"
	"log"
	"time"

	"microservice/internal/cache"
	"microservice/internal/db"
	"microservice/internal/models"
	"microservice/internal/validation"

	"github.com/segmentio/kafka-go"
)

type Consumer struct {
	reader     *kafka.Reader
	db         db.OrderRepository
	orderCache *cache.Cache
	topic      string
}

func NewConsumer(brokers []string, topic string, groupID string, dbRepo db.OrderRepository, orderCache *cache.Cache) ConsumerInterface {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:     brokers,
		Topic:       topic,
		GroupID:     groupID,
		MinBytes:    1,
		MaxBytes:    10e6,
		StartOffset: kafka.FirstOffset,
	})

	return &Consumer{
		reader:     r,
		db:         dbRepo,
		orderCache: orderCache,
		topic:      topic,
	}
}

// Start запускает бесконечный цикл обработки сообщений
func (c *Consumer) Start(ctx context.Context) {
	defer c.reader.Close()
	log.Printf("Kafka consumer подключен. Топик: %s", c.topic)

	for {
		select {
		case <-ctx.Done():
			log.Println("Kafka consumer остановлен по сигналу")
			return
		default:
		}

		m, err := c.reader.FetchMessage(ctx)
		if err != nil {
			if ctx.Err() != nil {
				return
			}
			log.Printf("Ошибка получения сообщения Kafka: %v", err)
			time.Sleep(time.Second)
			continue
		}

		log.Printf("Получено сообщение Kafka: key=%s", string(m.Key))

		order := &models.Order{}
		if err := models.UnmarshalOrder(m.Value, order); err != nil {
			log.Printf("Ошибка парсинга JSON заказа: %v", err)
			continue
		}

		if err := validation.ValidateOrder(order); err != nil {
			log.Printf("Заказ %s не прошел валидацию: %v", order.OrderUID, err)
			continue
		}

		if err := c.db.InsertOrder(order); err != nil {
			log.Printf("Ошибка вставки заказа %s в БД: %v", order.OrderUID, err)
			continue
		}

		c.orderCache.Set(order.OrderUID, order)
		log.Printf("Заказ %s успешно обработан: вставлен в БД и добавлен в кэш", order.OrderUID)

		if err := c.reader.CommitMessages(ctx, m); err != nil {
			log.Printf("Ошибка подтверждения сообщения Kafka: %v", err)
		} else {
			log.Printf("Сообщение Kafka с ключом %s подтверждено", string(m.Key))
		}
	}
}

// Close закрывает reader
func (c *Consumer) Close() error {
	return c.reader.Close()
}
