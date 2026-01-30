package db

import "microservice/internal/models"

// OrderRepository описывает интерфейс работы с заказами
type OrderRepository interface {
	GetAllOrders() ([]*models.Order, error)
	GetOrderByUID(orderUID string) (*models.Order, error)
	InsertOrder(order *models.Order) error
}
