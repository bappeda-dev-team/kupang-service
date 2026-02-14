package repository

import (
	"context"
	"database/sql"
	"errors"
	"kupang-service/model/domain"
)

type IndikatorPokinOpdRepositoryImpl struct {
}

func NewIndikatorPokinOpdRepositoryImpl() *IndikatorPokinOpdRepositoryImpl {
	return &IndikatorPokinOpdRepositoryImpl{}
}

func (repository *IndikatorPokinOpdRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, indikatorPokinOpd domain.IndikatorPokinOpd) (domain.IndikatorPokinOpd, error) {
	query := "INSERT INTO indikator_pokin_opd (tujuan_pokin_opd_id, nama_indikator) VALUES ($1, $2) RETURNING id"
	err := tx.QueryRowContext(ctx, query, indikatorPokinOpd.TujuanPokinOpdId, indikatorPokinOpd.NamaIndikator).Scan(&indikatorPokinOpd.Id)
	if err != nil {
		return domain.IndikatorPokinOpd{}, err
	}

	return indikatorPokinOpd, nil
}

func (repository *IndikatorPokinOpdRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, indikatorPokinOpd domain.IndikatorPokinOpd) (domain.IndikatorPokinOpd, error) {
	query := "UPDATE indikator_pokin_opd SET nama_indikator = $1, last_modified_date = NOW() WHERE id = $2"
	_, err := tx.ExecContext(ctx, query, indikatorPokinOpd.NamaIndikator, indikatorPokinOpd.Id)
	if err != nil {
		return domain.IndikatorPokinOpd{}, err
	}

	return indikatorPokinOpd, nil
}

func (repository *IndikatorPokinOpdRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, id int) (domain.IndikatorPokinOpd, error) {
	query := "SELECT id, tujuan_pokin_opd_id, nama_indikator FROM indikator_pokin_opd WHERE id = $1"
	row := tx.QueryRowContext(ctx, query, id)

	var indikatorPokinOpd domain.IndikatorPokinOpd
	err := row.Scan(&indikatorPokinOpd.Id, &indikatorPokinOpd.TujuanPokinOpdId, &indikatorPokinOpd.NamaIndikator)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.IndikatorPokinOpd{}, errors.New("id tidak ditemukan")
		}
		return domain.IndikatorPokinOpd{}, err
	}

	return indikatorPokinOpd, nil
}

func (repository *IndikatorPokinOpdRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) ([]domain.IndikatorPokinOpd, error) {
	query := "SELECT id, tujuan_pokin_opd_id, nama_indikator FROM indikator_pokin_opd ORDER BY id ASC"
	rows, err := tx.QueryContext(ctx, query)
	if err != nil {
		return []domain.IndikatorPokinOpd{}, err
	}
	defer rows.Close()

	var indikatorPokinOpdList []domain.IndikatorPokinOpd
	for rows.Next() {
		var indikatorPokinOpd domain.IndikatorPokinOpd
		err := rows.Scan(&indikatorPokinOpd.Id, &indikatorPokinOpd.TujuanPokinOpdId, &indikatorPokinOpd.NamaIndikator)
		if err != nil {
			return []domain.IndikatorPokinOpd{}, err
		}

		indikatorPokinOpdList = append(indikatorPokinOpdList, indikatorPokinOpd)
	}

	return indikatorPokinOpdList, nil
}

func (repository *IndikatorPokinOpdRepositoryImpl) FindByTujuanPokinOpdId(ctx context.Context, tx *sql.Tx, tujuanPokinOpdId int) ([]domain.IndikatorPokinOpd, error) {
	query := "SELECT id, tujuan_pokin_opd_id, nama_indikator FROM indikator_pokin_opd WHERE tujuan_pokin_opd_id = $1 ORDER BY id ASC"
	rows, err := tx.QueryContext(ctx, query, tujuanPokinOpdId)
	if err != nil {
		return []domain.IndikatorPokinOpd{}, err
	}
	defer rows.Close()

	var indikatorPokinOpdList []domain.IndikatorPokinOpd
	for rows.Next() {
		var indikatorPokinOpd domain.IndikatorPokinOpd
		err := rows.Scan(&indikatorPokinOpd.Id, &indikatorPokinOpd.TujuanPokinOpdId, &indikatorPokinOpd.NamaIndikator)
		if err != nil {
			return []domain.IndikatorPokinOpd{}, err
		}

		indikatorPokinOpdList = append(indikatorPokinOpdList, indikatorPokinOpd)
	}

	return indikatorPokinOpdList, nil
}

func (repository *IndikatorPokinOpdRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, id int) error {
	query := "DELETE FROM indikator_pokin_opd WHERE id = $1"
	_, err := tx.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}
