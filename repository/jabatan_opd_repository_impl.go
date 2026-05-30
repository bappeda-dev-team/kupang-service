package repository

import (
	"context"
	"database/sql"
	"errors"
	"kupang-service/model/domain"
)

type JabatanOpdRepositoryImpl struct {
}

func NewJabatanOpdRepositoryImpl() *JabatanOpdRepositoryImpl {
	return &JabatanOpdRepositoryImpl{}
}

func (repository *JabatanOpdRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, jabatanOpd domain.JabatanOpd) (domain.JabatanOpd, error) {
	query := "INSERT INTO jabatan_opd (kode_jabatan, nama_jabatan, kode_opd, tahun) VALUES ($1, $2, $3, $4) RETURNING id"
	err := tx.QueryRowContext(ctx, query, jabatanOpd.KodeJabatan, jabatanOpd.NamaJabatan, jabatanOpd.KodeOpd, jabatanOpd.Tahun).Scan(&jabatanOpd.Id)
	if err != nil {
		return domain.JabatanOpd{}, err
	}

	return jabatanOpd, nil
}

func (repository *JabatanOpdRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, jabatanOpd domain.JabatanOpd) (domain.JabatanOpd, error) {
	query := "UPDATE jabatan_opd SET kode_jabatan = $1, nama_jabatan = $2, kode_opd = $3, tahun = $4 WHERE id = $5 RETURNING id"
	err := tx.QueryRowContext(ctx, query, jabatanOpd.KodeJabatan, jabatanOpd.NamaJabatan, jabatanOpd.KodeOpd, jabatanOpd.Tahun, jabatanOpd.Id).Scan(&jabatanOpd.Id)
	if err != nil {
		return domain.JabatanOpd{}, err
	}

	return jabatanOpd, nil
}

func (repository *JabatanOpdRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, id int) (domain.JabatanOpd, error) {
	query := "SELECT id, kode_jabatan, nama_jabatan, kode_opd, tahun FROM jabatan_opd WHERE id = $1"
	row := tx.QueryRowContext(ctx, query, id)

	var jabatanOpd domain.JabatanOpd
	err := row.Scan(&jabatanOpd.Id, &jabatanOpd.KodeJabatan, &jabatanOpd.NamaJabatan, &jabatanOpd.KodeOpd, &jabatanOpd.Tahun)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.JabatanOpd{}, errors.New("id tidak ditemukan")
		}
		return domain.JabatanOpd{}, err
	}

	return jabatanOpd, nil
}

func (repository *JabatanOpdRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) ([]domain.JabatanOpd, error) {
	query := "SELECT id, kode_jabatan, nama_jabatan, kode_opd, tahun FROM jabatan_opd ORDER BY id ASC"
	rows, err := tx.QueryContext(ctx, query)
	if err != nil {
		return []domain.JabatanOpd{}, err
	}
	defer rows.Close()

	var jabatanOpdList []domain.JabatanOpd
	for rows.Next() {
		var jabatanOpd domain.JabatanOpd
		err := rows.Scan(&jabatanOpd.Id, &jabatanOpd.KodeJabatan, &jabatanOpd.NamaJabatan, &jabatanOpd.KodeOpd, &jabatanOpd.Tahun)
		if err != nil {
			return []domain.JabatanOpd{}, err
		}

		jabatanOpdList = append(jabatanOpdList, jabatanOpd)
	}

	return jabatanOpdList, nil
}

func (repository *JabatanOpdRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, id int) error {
	query := "DELETE FROM jabatan_opd WHERE id = $1"
	_, err := tx.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}
