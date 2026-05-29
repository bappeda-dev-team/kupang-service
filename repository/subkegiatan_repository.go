package repository

import (
	"context"
	"database/sql"
	"kupang-service/model/domain"
)

type SubkegiatanRepository interface {
	Create(ctx context.Context, tx *sql.Tx, subkegiatan domain.Subkegiatan) (domain.Subkegiatan, error)
	Update(ctx context.Context, tx *sql.Tx, subkegiatan domain.Subkegiatan) (domain.Subkegiatan, error)
	Delete(ctx context.Context, tx *sql.Tx, id int) error
	FindById(ctx context.Context, tx *sql.Tx, id int) (domain.Subkegiatan, error)
	FindAll(ctx context.Context, tx *sql.Tx) ([]domain.Subkegiatan, error)
}
