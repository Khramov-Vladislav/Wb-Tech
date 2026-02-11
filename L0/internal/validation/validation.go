package validation

import (
	"fmt"

	"microservice/internal/models"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

// ValidateOrder проверяет данные заказа
func ValidateOrder(order *models.Order) error {
	// Добавляем теги в модели или проверяем вручную
	if err := validate.Struct(order); err != nil {
		return fmt.Errorf("структурная ошибка: %w", err)
	}

	// Дополнительные проверки
	if order.OrderUID == "" {
		return fmt.Errorf("OrderUID не может быть пустым")
	}
	if order.Payment.Amount < 0 {
		return fmt.Errorf("Payment.Amount не может быть отрицательным")
	}
	if order.Delivery.Phone == "" {
		return fmt.Errorf("Delivery.Phone обязательное поле")
	}
	for i, item := range order.Items {
		if item.Price < 0 {
			return fmt.Errorf("Item[%d].Price не может быть отрицательным", i)
		}
		if item.NmID <= 0 {
			return fmt.Errorf("Item[%d].NmID должно быть > 0", i)
		}
	}
	return nil
}
