package repository

import (
	"context"
	"database/sql"
	"kupang-service/model/domain"
)

type PegawaiRepository interface {
	Create(ctx context.Context, tx *sql.Tx, pegawai domain.Pegawai) (domain.Pegawai, error)
	Update(ctx context.Context, tx *sql.Tx, pegawai domain.Pegawai) (domain.Pegawai, error)
	UpdateJabatan(ctx context.Context, tx *sql.Tx, jabatan domain.Jabatan) (domain.Jabatan, error)
	Delete(ctx context.Context, tx *sql.Tx, id int) error
	FindById(ctx context.Context, tx *sql.Tx, id int) (domain.Pegawai, error)
	FindJabatanById(ctx context.Context, tx *sql.Tx, id int) (domain.Jabatan, error)
	FindAll(ctx context.Context, tx *sql.Tx) ([]domain.Pegawai, error)
	FindByKodeOpd(ctx context.Context, tx *sql.Tx, kodeOpd string) ([]domain.Pegawai, error)
	UpdatePegawaiNamaJabatanByJabatanId(ctx context.Context, tx *sql.Tx, jabatanId int, namaJabatan string) error
}
