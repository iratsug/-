package repository

import (
	"database/sql"

	"shop/internal/models"

	"github.com/jmoiron/sqlx"
)

type PhotoRepo struct{ db *sqlx.DB }

func NewPhotoRepo(db *sqlx.DB) *PhotoRepo { return &PhotoRepo{db: db} }

func (r *PhotoRepo) Create(p *models.Photo) error {
	_, err := r.db.NamedExec(`
		INSERT INTO photos (id, filename, product_id)
		VALUES (:id, :filename, :product_id)`, p)
	return err
}

func (r *PhotoRepo) GetByID(id string) (*models.Photo, error) {
	var p models.Photo
	err := r.db.Get(&p, `SELECT * FROM photos WHERE id = ?`, id)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return &p, err
}

func (r *PhotoRepo) Delete(id string) error {
	_, err := r.db.Exec(`DELETE FROM photos WHERE id = ?`, id)
	return err
}