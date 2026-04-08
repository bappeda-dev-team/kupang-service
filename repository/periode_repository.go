package repository

import (
	"context"
	"database/sql"
	"kupang-service/model/domain"
)

type PeriodeRepository interface {
	Create(ctx context.Context, tx *sql.Tx, periode domain.Periode) (domain.Periode, error)
	Update(ctx context.Context, tx *sql.Tx, periode domain.Periode) (domain.Periode, error)
	Delete(ctx context.Context, tx *sql.Tx, id int) error
	FindById(ctx context.Context, tx *sql.Tx, id int) (domain.Periode, error)
	FindAll(ctx context.Context, tx *sql.Tx) ([]domain.Periode, error)
}
