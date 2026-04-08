package repository

import (
	"context"
	"database/sql"
	"kupang-service/model/domain"
)

type RoleRepository interface {
	Create(ctx context.Context, tx *sql.Tx, role domain.Role) (domain.Role, error)
	Update(ctx context.Context, tx *sql.Tx, role domain.Role) (domain.Role, error)
	Delete(ctx context.Context, tx *sql.Tx, id int) error
	FindById(ctx context.Context, tx *sql.Tx, id int) (domain.Role, error)
	FindAll(ctx context.Context, tx *sql.Tx) ([]domain.Role, error)
}
