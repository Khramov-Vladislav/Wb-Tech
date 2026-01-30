package kafka

import (
	"context"
	"testing"
	"time"

	"microservice/internal/models"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestProducer_SendOrder(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockProd := NewMockProducerInterface(ctrl)
	order := &models.Order{OrderUID: "order-123"}

	mockProd.EXPECT().SendOrder(gomock.Any(), order).Return(nil).Times(1)
	mockProd.EXPECT().Close().Return(nil).Times(1)

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	err := mockProd.SendOrder(ctx, order)
	assert.NoError(t, err)

	err = mockProd.Close()
	assert.NoError(t, err)
}
