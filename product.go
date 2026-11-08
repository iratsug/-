package repository

import (
	"database/sql"
	"fmt"
	"strings"

	"shop/internal/models"

	"github.com/jmoiron/sqlx"
)

type ProductRepo struct{ db *sqlx.DB }

func NewProductRepo(db *sqlx.DB) *ProductRepo { return &ProductRepo{db: db} }

func (r *ProductRepo) Create(p *models.Product) error {
	_, err := r.db.NamedExec(`
		INSERT INTO products (id, name, type, gender, size, color, description, brand, season, price)
		VALUES (:id, :name, :type, :gender, :size, :color, :description, :brand, :season, :price)`, p)
	return err
}

func (r *ProductRepo) GetByID(id string) (*models.Product, error) {
	var p models.Product
	err := r.db.Get(&p, `SELECT * FROM products WHERE id = ?`, id)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return &p, err
}

func (r *ProductRepo) List(filter map[string]string, sortBy string, limit, offset int) ([]models.Product, error) {
	where := []string{"1=1"}
	args := []interface{}{}

	allowed := map[string]bool{
		"type": true, "gender": true, "size": true,
		"color": true, "brand": true, "season": true,
	}
	for k, v := range filter {
		if allowed[k] && v != "" {
			where = append(where, fmt.Sprintf("%s = ?", k))
			args = append(args, v)
		}
	}

	order := "name ASC"
	if sortBy == "price_asc" {
		order = "price ASC"
	} else if sortBy == "price_desc" {
		order = "price DESC"
	}

	query := fmt.Sprintf(
		`SELECT * FROM products WHERE %s ORDER BY %s LIMIT ? OFFSET ?`,
		strings.Join(where, " AND "), order,
	)
	args = append(args, limit, offset)

	var out []models.Product
	err := r.db.Select(&out, query, args...)
	return out, err
}

func (r *ProductRepo) Update(p *models.Product) error {
	_, err := r.db.NamedExec(`
		UPDATE products SET
			name=:name, type=:type, gender=:gender, size=:size, color=:color,
			description=:description, brand=:brand, season=:season, price=:price
		WHERE id=:id`, p)
	return err
}

func (r *ProductRepo) Delete(id string) error {
	_, err := r.db.Exec(`DELETE FROM products WHERE id = ?`, id)
	return err
}