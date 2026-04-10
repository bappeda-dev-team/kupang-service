package repository

import (
	"context"
	"database/sql"
	"kupang-service/model/domain"
)

type MusrenbangRepository interface {
	Create(ctx context.Context, tx *sql.Tx, musrenbang domain.Musrenbang) (domain.Musrenbang, error)
	Update(ctx context.Context, tx *sql.Tx, musrenbang domain.Musrenbang) (domain.Musrenbang, error)
	Delete(ctx context.Context, tx *sql.Tx, id int) error
	FindById(ctx context.Context, tx *sql.Tx, id int) (domain.Musrenbang, error)
	FindAll(ctx context.Context, tx *sql.Tx) ([]domain.Musrenbang, error)
	FindByKodeOpd(ctx context.Context, tx *sql.Tx, kodeOpd string) ([]domain.Musrenbang, error)
}
