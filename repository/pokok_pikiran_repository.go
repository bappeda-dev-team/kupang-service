package repository

import (
	"context"
	"database/sql"
	"kupang-service/model/domain"
)

type PokokPikiranRepository interface {
	Create(ctx context.Context, tx *sql.Tx, pokokPikiran domain.PokokPikiran) (domain.PokokPikiran, error)
	Update(ctx context.Context, tx *sql.Tx, pokokPikiran domain.PokokPikiran) (domain.PokokPikiran, error)
	Delete(ctx context.Context, tx *sql.Tx, id int) error
	FindById(ctx context.Context, tx *sql.Tx, id int) (domain.PokokPikiran, error)
	FindAll(ctx context.Context, tx *sql.Tx) ([]domain.PokokPikiran, error)
	FindByKodeOpd(ctx context.Context, tx *sql.Tx, kodeOpd string) ([]domain.PokokPikiran, error)
}
