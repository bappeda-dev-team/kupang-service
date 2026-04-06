package repository

import (
	"context"
	"database/sql"
	"kupang-service/model/domain"
)

type LembagaRepository interface {
	Create(ctx context.Context, tx *sql.Tx, lembaga domain.Lembaga) (domain.Lembaga, error)
	Update(ctx context.Context, tx *sql.Tx, lembaga domain.Lembaga) (domain.Lembaga, error)
	Delete(ctx context.Context, tx *sql.Tx, id int) error
	FindById(ctx context.Context, tx *sql.Tx, id int) (domain.Lembaga, error)
	FindAll(ctx context.Context, tx *sql.Tx) ([]domain.Lembaga, error)
}
