package repository

import (
	"context"
	"database/sql"
	"kupang-service/model/domain"
)

type OpdRepository interface {
	Create(ctx context.Context, tx *sql.Tx, opd domain.Opd) (domain.Opd, error)
	Update(ctx context.Context, tx *sql.Tx, opd domain.Opd) (domain.Opd, error)
	Delete(ctx context.Context, tx *sql.Tx, id int) error
	FindById(ctx context.Context, tx *sql.Tx, id int) (domain.Opd, error)
	FindAll(ctx context.Context, tx *sql.Tx) ([]domain.Opd, error)
}
