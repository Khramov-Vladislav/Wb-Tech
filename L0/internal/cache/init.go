package cache

import (
	"database/sql"
	"log"

	"microservice/internal/db"
)

// InitCacheFromDB загружает все заказы из БД в кэш
func InitCacheFromDB(dbConn *sql.DB, orderCache CacheIface) {
	log.Println("Загрузка заказов из БД в кэш...")

	orders, err := db.GetAllOrders(dbConn)
	if err != nil {
		log.Printf("Ошибка загрузки заказов из БД: %v", err)
		return
	}

	if len(orders) == 0 {
		log.Println("В БД нет заказов — кэш инициализирован пустым")
		return
	}

	count := 0
	for _, order := range orders {
		if err := orderCache.SetValidated(order.OrderUID, order); err != nil {
			log.Printf("Заказ %s не прошел валидацию и не добавлен в кэш: %v", order.OrderUID, err)
			continue
		}
		count++
	}

	log.Printf("Загружено %d заказов в кэш", count)
}
