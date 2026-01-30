# Microservice

Сервис обработки заказов с PostgreSQL, Kafka и кэшем.

## Запуск

1. Настройте `.env` (см. `.env.example`).  
2. Запустите PostgreSQL и Kafka (Docker).  
3. Сервер: ``go run cmd/server/main.go``

4. Продюсер Kafka (отправка заказов): ``go run cmd/producer/main.go``

## API
Получение заказа по UID: ``http://localhost:8081/order/<order_uid>``

## Тестирование
 ``go test ./internal/... -v``