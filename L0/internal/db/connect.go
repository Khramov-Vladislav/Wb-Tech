package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

// Connect устанавливает соединение с PostgreSQL и проверяет доступность
func Connect(dsn string) *sql.DB {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("Ошибка подключения к БД: %v", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("БД недоступна: %v", err)
	}

	log.Println("Подключение к БД успешно")

	var currentDB string
	err = db.QueryRow("SELECT current_database();").Scan(&currentDB)
	if err != nil {
		log.Fatal("Ошибка получения имени БД", err)
	}
	fmt.Println("Текущая БД:", currentDB)

	return db
}
