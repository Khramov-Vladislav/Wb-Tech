package cache

import (
	"testing"
	"time"

	"microservice/internal/models"

	"github.com/stretchr/testify/assert"
)

func validOrder(uid string) *models.Order {
	return &models.Order{
		OrderUID: uid,
		Delivery: models.Delivery{
			Name:    "Иван Иванов",
			Phone:   "+79161234567",
			Zip:     "101000",
			City:    "Москва",
			Address: "ул. Тверская, 1",
			Region:  "Москва",
			Email:   "ivan@example.com",
		},
		Payment: models.Payment{
			Transaction:  "txn-001",
			RequestID:    "req-001",
			Currency:     "RUB",
			Provider:     "SBERBANK",
			Amount:       1000,
			PaymentDt:    1673953200,
			Bank:         "Sberbank",
			DeliveryCost: 200,
			GoodsTotal:   800,
			CustomFee:    0,
		},
		Items: []models.Item{
			{
				ChrtID:      1,
				TrackNumber: "TRK-001",
				Price:       800,
				Rid:         "RID-001",
				Name:        "Товар 1",
				Sale:        0,
				Size:        "M",
				TotalPrice:  800,
				NmID:        101,
				Brand:       "BrandA",
				Status:      1,
			},
		},
	}
}

func TestCache_SetAndGet(t *testing.T) {
	c := NewCache(2 * time.Second)

	order := &models.Order{OrderUID: "order-1"}
	c.Set("order-1", order)

	got, ok := c.Get("order-1")
	assert.True(t, ok)
	assert.Equal(t, order, got)
}

func TestCache_SetValidated_Valid(t *testing.T) {
	c := NewCache(2 * time.Second)

	order := validOrder("order-2")
	err := c.SetValidated("order-2", order)
	assert.NoError(t, err)

	got, ok := c.Get("order-2")
	assert.True(t, ok)
	assert.Equal(t, order, got)
}

func TestCache_SetValidated_Invalid(t *testing.T) {
	c := NewCache(2 * time.Second)

	order := &models.Order{} 
	err := c.SetValidated("", order)
	assert.Error(t, err)
}

func TestCache_TTLExpiration(t *testing.T) {
	c := NewCache(1 * time.Second)

	order := &models.Order{OrderUID: "order-ttl"}
	c.Set("order-ttl", order)

	got, ok := c.Get("order-ttl")
	assert.True(t, ok)
	assert.Equal(t, order, got)

	// Ждем пока TTL истечет
	time.Sleep(1100 * time.Millisecond)
	got, ok = c.Get("order-ttl")
	assert.False(t, ok)
	assert.Nil(t, got)
}
