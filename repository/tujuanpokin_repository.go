package repository

import (
	"context"
	"database/sql"
	"kupang-service/model/domain"
)

type TujuanPokinOpdRepository interface {
	Create(ctx context.Context, tx *sql.Tx, tujuanPokinOpd domain.TujuanPokinOpd) (domain.TujuanPokinOpd, error)
	Update(ctx context.Context, tx *sql.Tx, tujuanPokinOpd domain.TujuanPokinOpd) (domain.TujuanPokinOpd, error)
	Delete(ctx context.Context, tx *sql.Tx, id int) error
	FindById(ctx context.Context, tx *sql.Tx, id int) (domain.TujuanPokinOpd, error)
	FindAll(ctx context.Context, tx *sql.Tx) ([]domain.TujuanPokinOpd, error)
}
