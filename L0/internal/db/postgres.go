package db

import (
	"database/sql"
	"microservice/internal/models"
)

// PostgresOrderRepo реализует OrderRepository через PostgreSQL
type PostgresOrderRepo struct {
	Conn *sql.DB
}

// NewPostgresOrderRepo создает новый репозиторий
func NewPostgresOrderRepo(conn *sql.DB) *PostgresOrderRepo {
	return &PostgresOrderRepo{Conn: conn}
}

func (r *PostgresOrderRepo) GetAllOrders() ([]*models.Order, error) {
	return GetAllOrders(r.Conn)
}

func (r *PostgresOrderRepo) GetOrderByUID(orderUID string) (*models.Order, error) {
	return GetOrderByUID(r.Conn, orderUID)
}

func (r *PostgresOrderRepo) InsertOrder(order *models.Order) error {
	return InsertOrder(r.Conn, order)
}
