package repository

import (
	"shop/internal/models"

	"github.com/jmoiron/sqlx"
)

type OrderItemRepo struct{ db *sqlx.DB }

func NewOrderItemRepo(db *sqlx.DB) *OrderItemRepo { return &OrderItemRepo{db: db} }

func (r *OrderItemRepo) Create(oi *models.OrderItem) error {
	_, err := r.db.NamedExec(`
		INSERT INTO order_items (id, order_id, product_id, price)
		VALUES (:id, :order_id, :product_id, :price)`, oi)
	return err
}

func (r *OrderItemRepo) GetByOrder(orderID string) ([]models.OrderItem, error) {
	var out []models.OrderItem
	err := r.db.Select(&out, `SELECT * FROM order_items WHERE order_id = ?`, orderID)
	return out, err
}