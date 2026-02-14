package repository

import (
	"context"
	"database/sql"
	"errors"
	"kupang-service/model/domain"
)

type PokinOpdRepositoryImpl struct {
}

func NewPokinOpdRepositoryImpl() *PokinOpdRepositoryImpl {
	return &PokinOpdRepositoryImpl{}
}

func (repository *PokinOpdRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, pokinOpd domain.PokinOpd) (domain.PokinOpd, error) {
	query := "INSERT INTO pokin_opd (kode_opd, nama_opd, tahun) VALUES ($1, $2, $3) RETURNING id"
	err := tx.QueryRowContext(ctx, query, pokinOpd.KodeOpd, pokinOpd.NamaOpd, pokinOpd.Tahun).Scan(&pokinOpd.Id)
	if err != nil {
		return domain.PokinOpd{}, err
	}

	return pokinOpd, nil
}

func (repository *PokinOpdRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, pokinOpd domain.PokinOpd) (domain.PokinOpd, error) {
	query := "UPDATE pokin_opd SET kode_opd = $1, nama_opd = $2, tahun = $3, last_modified_date = NOW() WHERE id = $4"
	_, err := tx.ExecContext(ctx, query, pokinOpd.KodeOpd, pokinOpd.NamaOpd, pokinOpd.Tahun, pokinOpd.Id)
	if err != nil {
		return domain.PokinOpd{}, err
	}

	return pokinOpd, nil
}

func (repository *PokinOpdRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, id int) (domain.PokinOpd, error) {
	query := "SELECT id, kode_opd, nama_opd, tahun FROM pokin_opd WHERE id = $1"
	row := tx.QueryRowContext(ctx, query, id)

	var pokinOpd domain.PokinOpd
	err := row.Scan(&pokinOpd.Id, &pokinOpd.KodeOpd, &pokinOpd.NamaOpd, &pokinOpd.Tahun)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.PokinOpd{}, errors.New("id tidak ditemukan")
		}
		return domain.PokinOpd{}, err
	}

	return pokinOpd, nil
}

func (repository *PokinOpdRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) ([]domain.PokinOpd, error) {
	query := "SELECT id, kode_opd, nama_opd, tahun FROM pokin_opd ORDER BY id ASC"
	rows, err := tx.QueryContext(ctx, query)
	if err != nil {
		return []domain.PokinOpd{}, err
	}
	defer rows.Close()

	var pokinOpdList []domain.PokinOpd
	for rows.Next() {
		var pokinOpd domain.PokinOpd
		err := rows.Scan(&pokinOpd.Id, &pokinOpd.KodeOpd, &pokinOpd.NamaOpd, &pokinOpd.Tahun)
		if err != nil {
			return []domain.PokinOpd{}, err
		}

		pokinOpdList = append(pokinOpdList, pokinOpd)
	}

	return pokinOpdList, nil
}

func (repository *PokinOpdRepositoryImpl) FindByKodeOpdAndTahun(ctx context.Context, tx *sql.Tx, kodeOpd string, tahun int) (domain.PokinOpd, error) {
	query := "SELECT id, kode_opd, nama_opd, tahun FROM pokin_opd WHERE kode_opd = $1 AND tahun = $2"
	row := tx.QueryRowContext(ctx, query, kodeOpd, tahun)

	var pokinOpd domain.PokinOpd
	err := row.Scan(&pokinOpd.Id, &pokinOpd.KodeOpd, &pokinOpd.NamaOpd, &pokinOpd.Tahun)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.PokinOpd{}, errors.New("data tidak ditemukan")
		}
		return domain.PokinOpd{}, err
	}

	return pokinOpd, nil
}

func (repository *PokinOpdRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, id int) error {
	query := "DELETE FROM pokin_opd WHERE id = $1"
	_, err := tx.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}
