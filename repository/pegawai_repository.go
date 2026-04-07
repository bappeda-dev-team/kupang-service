package repository

import (
	"context"
	"database/sql"
	"kupang-service/model/domain"
)

type PegawaiRepository interface {
	Create(ctx context.Context, tx *sql.Tx, pegawai domain.Pegawai) (domain.Pegawai, error)
	Update(ctx context.Context, tx *sql.Tx, pegawai domain.Pegawai) (domain.Pegawai, error)
	Delete(ctx context.Context, tx *sql.Tx, id int) error
	FindById(ctx context.Context, tx *sql.Tx, id int) (domain.Pegawai, error)
	FindAll(ctx context.Context, tx *sql.Tx) ([]domain.Pegawai, error)
	FindByKodeOpd(ctx context.Context, tx *sql.Tx, kodeOpd string) ([]domain.Pegawai, error)
}
