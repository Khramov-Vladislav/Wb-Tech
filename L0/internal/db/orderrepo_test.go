package db

import (
	"microservice/internal/models"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestOrderRepository_WithMock(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := NewMockOrderRepository(ctrl)

	order1 := &models.Order{OrderUID: "order-1"}
	order2 := &models.Order{OrderUID: "order-2"}

	// Тест GetAllOrders
	mockRepo.EXPECT().GetAllOrders().Return([]*models.Order{order1, order2}, nil)
	orders, err := mockRepo.GetAllOrders()
	assert.NoError(t, err)
	assert.Len(t, orders, 2)
	assert.Equal(t, order1.OrderUID, orders[0].OrderUID)
	assert.Equal(t, order2.OrderUID, orders[1].OrderUID)

	// Тест GetOrderByUID
	mockRepo.EXPECT().GetOrderByUID("order-1").Return(order1, nil)
	o, err := mockRepo.GetOrderByUID("order-1")
	assert.NoError(t, err)
	assert.Equal(t, "order-1", o.OrderUID)

	// Тест InsertOrder
	mockRepo.EXPECT().InsertOrder(order2).Return(nil)
	err = mockRepo.InsertOrder(order2)
	assert.NoError(t, err)
}
