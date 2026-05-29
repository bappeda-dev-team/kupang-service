package repository

import (
	"context"
	"database/sql"
	"kupang-service/model/domain"
)

type UrusanRepository interface {
	Create(ctx context.Context, tx *sql.Tx, urusan domain.Urusan) (domain.Urusan, error)
	Update(ctx context.Context, tx *sql.Tx, urusan domain.Urusan) (domain.Urusan, error)
	Delete(ctx context.Context, tx *sql.Tx, id int) error
	FindById(ctx context.Context, tx *sql.Tx, id int) (domain.Urusan, error)
	FindAll(ctx context.Context, tx *sql.Tx) ([]domain.Urusan, error)
}
