package repository

import (
	"context"
	"database/sql"
	"errors"
	"kupang-service/model/domain"
)

type PeriodeRepositoryImpl struct {
}

func NewPeriodeRepositoryImpl() *PeriodeRepositoryImpl {
	return &PeriodeRepositoryImpl{}
}

func (repository *PeriodeRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, periode domain.Periode) (domain.Periode, error) {
	query := "INSERT INTO periode (tahun_awal, tahun_akhir, jenis_periode) VALUES ($1, $2, $3) RETURNING id"
	err := tx.QueryRowContext(ctx, query, periode.TahunAwal, periode.TahunAkhir, periode.JenisPeriode).Scan(&periode.Id)
	if err != nil {
		return domain.Periode{}, err
	}

	return periode, nil
}

func (repository *PeriodeRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, periode domain.Periode) (domain.Periode, error) {
	query := "UPDATE periode SET tahun_awal = $1, tahun_akhir = $2, jenis_periode = $3, last_modified_date = NOW() WHERE id = $4"
	_, err := tx.ExecContext(ctx, query, periode.TahunAwal, periode.TahunAkhir, periode.JenisPeriode, periode.Id)
	if err != nil {
		return domain.Periode{}, err
	}

	return periode, nil
}

func (repository *PeriodeRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, id int) error {
	query := "DELETE FROM periode WHERE id = $1"
	_, err := tx.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}

func (repository *PeriodeRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, id int) (domain.Periode, error) {
	query := "SELECT id, tahun_awal, tahun_akhir, jenis_periode FROM periode WHERE id = $1"
	row := tx.QueryRowContext(ctx, query, id)

	var periode domain.Periode
	err := row.Scan(&periode.Id, &periode.TahunAwal, &periode.TahunAkhir, &periode.JenisPeriode)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.Periode{}, errors.New("id tidak ditemukan")
		}
		return domain.Periode{}, err
	}

	return periode, nil
}

func (repository *PeriodeRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) ([]domain.Periode, error) {
	query := "SELECT id, tahun_awal, tahun_akhir, jenis_periode FROM periode ORDER BY id ASC"
	rows, err := tx.QueryContext(ctx, query)
	if err != nil {
		return []domain.Periode{}, err
	}
	defer rows.Close()

	var periodes []domain.Periode
	for rows.Next() {
		var periode domain.Periode
		err := rows.Scan(&periode.Id, &periode.TahunAwal, &periode.TahunAkhir, &periode.JenisPeriode)
		if err != nil {
			return nil, err
		}
		periodes = append(periodes, periode)
	}

	return periodes, nil
}
