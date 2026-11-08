package repository

import (
	"shop/internal/models"

	"github.com/jmoiron/sqlx"
)

type CartRepo struct{ db *sqlx.DB }

func NewCartRepo(db *sqlx.DB) *CartRepo { return &CartRepo{db: db} }

func (r *CartRepo) Create(c *models.Cart) error {
	_, err := r.db.NamedExec(`
		INSERT INTO cart (id, user_id, product_id, count)
		VALUES (:id, :user_id, :product_id, :count)`, c)
	return err
}

func (r *CartRepo) GetByUser(userID string) ([]models.Cart, error) {
	var out []models.Cart
	err := r.db.Select(&out, `SELECT * FROM cart WHERE user_id = ?`, userID)
	return out, err
}

func (r *CartRepo) Update(c *models.Cart) error {
	_, err := r.db.NamedExec(`
		UPDATE cart SET user_id=:user_id, product_id=:product_id, count=:count
		WHERE id=:id`, c)
	return err
}

func (r *CartRepo) Delete(id string) error {
	_, err := r.db.Exec(`DELETE FROM cart WHERE id = ?`, id)
	return err
}