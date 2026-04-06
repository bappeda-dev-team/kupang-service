package repository

import (
	"context"
	"database/sql"
	"errors"
	"kupang-service/model/domain"
)

type OpdRepositoryImpl struct {
}

func NewOpdRepositoryImpl() *OpdRepositoryImpl {
	return &OpdRepositoryImpl{}
}

func (repository *OpdRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, opd domain.Opd) (domain.Opd, error) {
	query := "INSERT INTO opd (kode_opd, nama_opd, kode_lembaga, nama_kepala_perangkat_daerah, nip_kepala_perangkat_daerah, pangkat_kepala_perangkat_daerah) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id"
	err := tx.QueryRowContext(ctx, query, opd.KodeOpd, opd.NamaOpd, opd.KodeLembaga, opd.NamaKepalaPerangkatDaerah, opd.NipKepalaPerangkatDaerah, opd.PangkatKepalaPerangkatDaerah).Scan(&opd.Id)
	if err != nil {
		return domain.Opd{}, err
	}

	return opd, nil
}

func (repository *OpdRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, opd domain.Opd) (domain.Opd, error) {
	query := "UPDATE opd SET kode_opd = $1, nama_opd = $2, kode_lembaga = $3, nama_kepala_perangkat_daerah = $4, nip_kepala_perangkat_daerah = $5, pangkat_kepala_perangkat_daerah = $6, last_modified_date = NOW() WHERE id = $7"
	_, err := tx.ExecContext(ctx, query, opd.KodeOpd, opd.NamaOpd, opd.KodeLembaga, opd.NamaKepalaPerangkatDaerah, opd.NipKepalaPerangkatDaerah, opd.PangkatKepalaPerangkatDaerah, opd.Id)
	if err != nil {
		return domain.Opd{}, err
	}

	return opd, nil
}

func (repository *OpdRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, id int) (domain.Opd, error) {
	query := "SELECT id, kode_opd, nama_opd, COALESCE(kode_lembaga, '') AS kode_lembaga, nama_kepala_perangkat_daerah, nip_kepala_perangkat_daerah, pangkat_kepala_perangkat_daerah FROM opd WHERE id = $1"
	row := tx.QueryRowContext(ctx, query, id)

	var opd domain.Opd
	err := row.Scan(&opd.Id, &opd.KodeOpd, &opd.NamaOpd, &opd.KodeLembaga, &opd.NamaKepalaPerangkatDaerah, &opd.NipKepalaPerangkatDaerah, &opd.PangkatKepalaPerangkatDaerah)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.Opd{}, errors.New("id tidak ditemukan")
		}
		return domain.Opd{}, err
	}

	return opd, nil
}

func (repository *OpdRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) ([]domain.Opd, error) {
	query := "SELECT id, kode_opd, nama_opd, COALESCE(kode_lembaga, '') AS kode_lembaga, nama_kepala_perangkat_daerah, nip_kepala_perangkat_daerah, pangkat_kepala_perangkat_daerah FROM opd ORDER BY id ASC"
	rows, err := tx.QueryContext(ctx, query)
	if err != nil {
		return []domain.Opd{}, err
	}
	defer rows.Close()

	var opdList []domain.Opd
	for rows.Next() {
		var opd domain.Opd
		err := rows.Scan(&opd.Id, &opd.KodeOpd, &opd.NamaOpd, &opd.KodeLembaga, &opd.NamaKepalaPerangkatDaerah, &opd.NipKepalaPerangkatDaerah, &opd.PangkatKepalaPerangkatDaerah)
		if err != nil {
			return []domain.Opd{}, err
		}

		opdList = append(opdList, opd)
	}

	return opdList, nil
}

func (repository *OpdRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, id int) error {
	query := "DELETE FROM opd WHERE id = $1"
	_, err := tx.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}
