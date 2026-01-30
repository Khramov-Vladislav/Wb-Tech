package main

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"microservice/config"
	"microservice/internal/kafka"
	"microservice/internal/models"
)

func main() {
	cfg := config.LoadConfig()

	var producer kafka.ProducerInterface
	producer = kafka.NewProducer(cfg.KafkaBrokers, cfg.KafkaTopic)
	defer producer.Close()

	fmt.Println("Продюсер Kafka запущен!")
	fmt.Println("Введите путь к JSON-файлу для отправки заказа в Kafka или 'exit' для выхода:")

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("→ ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "" {
			continue
		}

		if strings.ToLower(input) == "exit" {
			fmt.Println("Завершение работы продюсера.")
			break
		}

		files := strings.Fields(input)
		for _, jsonFile := range files {
			data, err := os.ReadFile(jsonFile)
			if err != nil {
				log.Printf("Ошибка чтения файла %s: %v", jsonFile, err)
				continue
			}

			var order models.Order
			if err := json.Unmarshal(data, &order); err != nil {
				log.Printf("Ошибка парсинга JSON (%s): %v", jsonFile, err)
				continue
			}

			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			err = producer.SendOrder(ctx, &order)
			cancel()

			if err != nil {
				log.Printf("Ошибка отправки заказа в Kafka (%s): %v", jsonFile, err)
				continue
			}

			fmt.Printf("Заказ %s успешно отправлен в Kafka\n", order.OrderUID)
		}
	}
}
