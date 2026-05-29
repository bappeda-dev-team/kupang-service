package repository

import (
	"context"
	"database/sql"
	"kupang-service/model/domain"
)

type KegiatanRepository interface {
	Create(ctx context.Context, tx *sql.Tx, kegiatan domain.Kegiatan) (domain.Kegiatan, error)
	Update(ctx context.Context, tx *sql.Tx, kegiatan domain.Kegiatan) (domain.Kegiatan, error)
	Delete(ctx context.Context, tx *sql.Tx, id int) error
	FindById(ctx context.Context, tx *sql.Tx, id int) (domain.Kegiatan, error)
	FindAll(ctx context.Context, tx *sql.Tx) ([]domain.Kegiatan, error)
}
