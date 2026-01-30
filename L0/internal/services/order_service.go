package services

import (
	"errors"

	"microservice/internal/cache"
	"microservice/internal/db"
	"microservice/internal/models"
)

// orderServiceImpl реализует OrderService
type orderServiceImpl struct {
	db    db.OrderRepository
	cache *cache.Cache
}

// NewOrderService создаёт экземпляр сервиса
func NewOrderService(dbRepo db.OrderRepository, orderCache *cache.Cache) OrderService {
	return &orderServiceImpl{
		db:    dbRepo,
		cache: orderCache,
	}
}

// GetOrder сначала проверяет кэш, затем БД
func (s *orderServiceImpl) GetOrder(orderUID string) (*models.Order, error) {
	if orderUID == "" {
		return nil, errors.New("order_uid не указан")
	}

	// Проверяем кэш
	if order, ok := s.cache.Get(orderUID); ok {
		return order, nil
	}

	// Если нет в кэше — ищем в БД
	order, err := s.db.GetOrderByUID(orderUID)
	if err != nil {
		return nil, err
	}

	// Сохраняем в кэш
	s.cache.Set(orderUID, order)
	return order, nil
}
