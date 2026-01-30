package services

import (
	"database/sql"
	"errors"
	"testing"
	"time"

	"microservice/internal/cache"
	"microservice/internal/db"
	"microservice/internal/models"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestOrderService_GetOrder(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := db.NewMockOrderRepository(ctrl)
	orderCache := cache.NewCache(2 * time.Second)
	orderService := NewOrderService(mockRepo, orderCache)

	order := &models.Order{OrderUID: "order-123"}

	mockRepo.
		EXPECT().
		GetOrderByUID("order-123").
		Return(order, nil).
		Times(1)

	got, err := orderService.GetOrder("order-123")
	assert.NoError(t, err)
	assert.Equal(t, order, got)

	// Проверяем, что кэш работает
	got2, err2 := orderService.GetOrder("order-123")
	assert.NoError(t, err2)
	assert.Equal(t, order, got2)
}

func TestOrderService_GetOrder_NotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := db.NewMockOrderRepository(ctrl)
	orderCache := cache.NewCache(2 * time.Second)
	orderService := NewOrderService(mockRepo, orderCache)

	// Возвращаем sql.ErrNoRows вместо несуществующей переменной
	mockRepo.
		EXPECT().
		GetOrderByUID("order-missing").
		Return(nil, sql.ErrNoRows).
		Times(1)

	got, err := orderService.GetOrder("order-missing")
	assert.Nil(t, got)
	assert.Error(t, err)
	assert.True(t, errors.Is(err, sql.ErrNoRows))
}
