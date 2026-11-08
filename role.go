package repository

import (
	"database/sql"

	"shop/internal/models"

	"github.com/jmoiron/sqlx"
)

type RoleRepo struct{ db *sqlx.DB }

func NewRoleRepo(db *sqlx.DB) *RoleRepo { return &RoleRepo{db: db} }

func (r *RoleRepo) Create(rl *models.Role) error {
	_, err := r.db.NamedExec(`INSERT INTO roles (id, name) VALUES (:id, :name)`, rl)
	return err
}

func (r *RoleRepo) GetByID(id string) (*models.Role, error) {
	var rl models.Role
	err := r.db.Get(&rl, `SELECT * FROM roles WHERE id = ?`, id)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return &rl, err
}

func (r *RoleRepo) List() ([]models.Role, error) {
	var out []models.Role
	err := r.db.Select(&out, `SELECT * FROM roles ORDER BY name`)
	return out, err
}

func (r *RoleRepo) Update(rl *models.Role) error {
	_, err := r.db.NamedExec(`UPDATE roles SET name=:name WHERE id=:id`, rl)
	return err
}

func (r *RoleRepo) Delete(id string) error {
	_, err := r.db.Exec(`DELETE FROM roles WHERE id = ?`, id)
	return err
}