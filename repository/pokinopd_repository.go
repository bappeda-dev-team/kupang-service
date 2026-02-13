package repository

import (
	"context"
	"database/sql"
	"kupang-service/model/domain"
)

type PokinOpdRepository interface {
	Create(ctx context.Context, tx *sql.Tx, pokinOpd domain.PokinOpd) (domain.PokinOpd, error)
	Update(ctx context.Context, tx *sql.Tx, pokinOpd domain.PokinOpd) (domain.PokinOpd, error)
	Delete(ctx context.Context, tx *sql.Tx, id int) error
	FindById(ctx context.Context, tx *sql.Tx, id int) (domain.PokinOpd, error)
	FindAll(ctx context.Context, tx *sql.Tx) ([]domain.PokinOpd, error)
}
