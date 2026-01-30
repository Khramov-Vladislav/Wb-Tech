package services

import "microservice/internal/models"

// OrderService описывает бизнес-логику работы с заказами
type OrderService interface {
	GetOrder(orderUID string) (*models.Order, error)
}
