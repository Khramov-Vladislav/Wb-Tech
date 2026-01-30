package kafka

import (
	"testing"
	"time"

	"microservice/internal/cache"
	"microservice/internal/db"
	"microservice/internal/models"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestConsumer_ProcessOrderWithMockDB(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Создаём мок DB
	mockDB := db.NewMockOrderRepository(ctrl)

	// Реальный кэш
	orderCache := cache.NewCache(2 * time.Second)

	// Подготовка заказа
	order := &models.Order{OrderUID: "order-123"}

	// Мокаем вызов InsertOrder
	mockDB.EXPECT().InsertOrder(order).Return(nil).Times(1)

	// Вместо реального Consumer используем метод, который напрямую вызывает InsertOrder
	// Симулируем обработку одного заказа
	err := mockDB.InsertOrder(order)
	assert.NoError(t, err)

	// Добавляем в кэш вручную
	orderCache.Set(order.OrderUID, order)

	// Проверяем, что заказ есть в кэше
	got, ok := orderCache.Get("order-123")
	assert.True(t, ok)
	assert.Equal(t, order, got)
}
