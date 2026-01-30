package cache

import "microservice/internal/models"

// CacheIface описывает поведение кэша заказов
type CacheIface interface {
	// Get возвращает заказ по ключу
	Get(key string) (*models.Order, bool)

	// Set сохраняет заказ в кэш без проверки
	Set(key string, order *models.Order)

	// SetValidated сохраняет заказ с проверкой
	SetValidated(key string, order *models.Order) error
}
