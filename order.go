package repository

import (
	"database/sql"

	"shop/internal/models"

	"github.com/jmoiron/sqlx"
)

type OrderRepo struct{ db *sqlx.DB }

func NewOrderRepo(db *sqlx.DB) *OrderRepo { return &OrderRepo{db: db} }

func (r *OrderRepo) Create(o *models.Order) error {
	_, err := r.db.NamedExec(`
		INSERT INTO orders (id, user_id, date_time, status)
		VALUES (:id, :user_id, :date_time, :status)`, o)
	return err
}

func (r *OrderRepo) GetByID(id string) (*models.Order, error) {
	var o models.Order
	err := r.db.Get(&o, `SELECT * FROM orders WHERE id = ?`, id)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return &o, err
}

func (r *OrderRepo) List(status, from, to string, limit, offset int) ([]models.Order, error) {
	query := `SELECT * FROM orders WHERE 1=1`
	args := []interface{}{}
	if status != "" {
		query += ` AND status = ?`
		args = append(args, status)
	}
	if from != "" {
		query += ` AND date_time >= ?`
		args = append(args, from)
	}
	if to != "" {
		query += ` AND date_time <= ?`
		args = append(args, to)
	}
	query += ` ORDER BY date_time DESC LIMIT ? OFFSET ?`
	args = append(args, limit, offset)

	var out []models.Order
	err := r.db.Select(&out, query, args...)
	return out, err
}

func (r *OrderRepo) Update(o *models.Order) error {
	_, err := r.db.NamedExec(`
		UPDATE orders SET user_id=:user_id, date_time=:date_time, status=:status
		WHERE id=:id`, o)
	return err
}

func (r *OrderRepo) UpdateStatus(id, status string) error {
	_, err := r.db.Exec(`UPDATE orders SET status = ? WHERE id = ?`, status, id)
	return err
}