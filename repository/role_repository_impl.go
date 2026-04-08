package repository

import (
	"context"
	"database/sql"
	"errors"
	"kupang-service/model/domain"
)

type RoleRepositoryImpl struct {
}

func NewRoleRepositoryImpl() *RoleRepositoryImpl {
	return &RoleRepositoryImpl{}
}

func (repository *RoleRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, role domain.Role) (domain.Role, error) {
	query := `INSERT INTO "role" (role) VALUES ($1) RETURNING id`
	err := tx.QueryRowContext(ctx, query, role.Role).Scan(&role.Id)
	if err != nil {
		return domain.Role{}, err
	}

	return role, nil
}

func (repository *RoleRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, role domain.Role) (domain.Role, error) {
	query := `UPDATE "role" SET role = $1, last_modified_date = NOW() WHERE id = $2`
	_, err := tx.ExecContext(ctx, query, role.Role, role.Id)
	if err != nil {
		return domain.Role{}, err
	}

	return role, nil
}

func (repository *RoleRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, id int) error {
	query := `DELETE FROM "role" WHERE id = $1`
	_, err := tx.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}

func (repository *RoleRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, id int) (domain.Role, error) {
	query := `SELECT id, role FROM "role" WHERE id = $1`
	row := tx.QueryRowContext(ctx, query, id)

	var role domain.Role
	err := row.Scan(&role.Id, &role.Role)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.Role{}, errors.New("id tidak ditemukan")
		}
		return domain.Role{}, err
	}

	return role, nil
}

func (repository *RoleRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) ([]domain.Role, error) {
	query := `SELECT id, role FROM "role" ORDER BY id ASC`
	rows, err := tx.QueryContext(ctx, query)
	if err != nil {
		return []domain.Role{}, err
	}
	defer rows.Close()

	var roles []domain.Role
	for rows.Next() {
		var role domain.Role
		err := rows.Scan(&role.Id, &role.Role)
		if err != nil {
			return nil, err
		}
		roles = append(roles, role)
	}

	return roles, nil
}
