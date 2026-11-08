package repository

import (
	"database/sql"

	"shop/internal/models"

	"github.com/jmoiron/sqlx"
)

type UserRepo struct{ db *sqlx.DB }

func NewUserRepo(db *sqlx.DB) *UserRepo { return &UserRepo{db: db} }

func (r *UserRepo) Create(u *models.User) error {
	_, err := r.db.NamedExec(`
		INSERT INTO users (id, name, email, phone, role_id, registration_date)
		VALUES (:id, :name, :email, :phone, :role_id, :registration_date)`, u)
	return err
}

func (r *UserRepo) GetByID(id string) (*models.User, error) {
	var u models.User
	err := r.db.Get(&u, `SELECT * FROM users WHERE id = ?`, id)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return &u, err
}

func (r *UserRepo) List(limit, offset int) ([]models.User, error) {
	var out []models.User
	err := r.db.Select(&out, `SELECT * FROM users ORDER BY registration_date DESC LIMIT ? OFFSET ?`, limit, offset)
	return out, err
}

func (r *UserRepo) ListByRole(roleID string) ([]models.User, error) {
	var out []models.User
	err := r.db.Select(&out, `SELECT * FROM users WHERE role_id = ?`, roleID)
	return out, err
}

func (r *UserRepo) Update(u *models.User) error {
	_, err := r.db.NamedExec(`
		UPDATE users SET name=:name, email=:email, phone=:phone, role_id=:role_id
		WHERE id=:id`, u)
	return err
}

func (r *UserRepo) Delete(id string) error {
	_, err := r.db.Exec(`DELETE FROM users WHERE id = ?`, id)
	return err
}