package repository

import (
	"context"
	"database/sql"
	"errors"
	"kupang-service/model/domain"
)

type IndikatorPokinOpdTacticalRepositoryImpl struct {
}

func NewIndikatorPokinOpdTacticalRepositoryImpl() *IndikatorPokinOpdTacticalRepositoryImpl {
	return &IndikatorPokinOpdTacticalRepositoryImpl{}
}

func (repository *IndikatorPokinOpdTacticalRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, indikator domain.IndikatorPokinOpdTactical) (domain.IndikatorPokinOpdTactical, error) {
	query := "INSERT INTO indikator_pokin_opd_tactical (pokin_opd_tactical_id, nama_indikator) VALUES ($1, $2) RETURNING id"
	err := tx.QueryRowContext(ctx, query, indikator.PokinOpdTacticalId, indikator.NamaIndikator).Scan(&indikator.Id)
	if err != nil {
		return domain.IndikatorPokinOpdTactical{}, err
	}

	return indikator, nil
}

func (repository *IndikatorPokinOpdTacticalRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, indikator domain.IndikatorPokinOpdTactical) (domain.IndikatorPokinOpdTactical, error) {
	query := "UPDATE indikator_pokin_opd_tactical SET pokin_opd_tactical_id = $1, nama_indikator = $2, last_modified_date = NOW() WHERE id = $3"
	_, err := tx.ExecContext(ctx, query, indikator.PokinOpdTacticalId, indikator.NamaIndikator, indikator.Id)
	if err != nil {
		return domain.IndikatorPokinOpdTactical{}, err
	}

	return indikator, nil
}

func (repository *IndikatorPokinOpdTacticalRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, id int) error {
	query := "DELETE FROM indikator_pokin_opd_tactical WHERE id = $1"
	_, err := tx.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}

func (repository *IndikatorPokinOpdTacticalRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, id int) (domain.IndikatorPokinOpdTactical, error) {
	query := "SELECT id, pokin_opd_tactical_id, nama_indikator FROM indikator_pokin_opd_tactical WHERE id = $1"
	row := tx.QueryRowContext(ctx, query, id)

	var indikator domain.IndikatorPokinOpdTactical
	err := row.Scan(&indikator.Id, &indikator.PokinOpdTacticalId, &indikator.NamaIndikator)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.IndikatorPokinOpdTactical{}, errors.New("id tidak ditemukan")
		}
		return domain.IndikatorPokinOpdTactical{}, err
	}

	return indikator, nil
}

func (repository *IndikatorPokinOpdTacticalRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) ([]domain.IndikatorPokinOpdTactical, error) {
	query := "SELECT id, pokin_opd_tactical_id, nama_indikator FROM indikator_pokin_opd_tactical ORDER BY id ASC"
	rows, err := tx.QueryContext(ctx, query)
	if err != nil {
		return []domain.IndikatorPokinOpdTactical{}, err
	}
	defer rows.Close()

	var indikators []domain.IndikatorPokinOpdTactical
	for rows.Next() {
		var indikator domain.IndikatorPokinOpdTactical
		err := rows.Scan(&indikator.Id, &indikator.PokinOpdTacticalId, &indikator.NamaIndikator)
		if err != nil {
			return []domain.IndikatorPokinOpdTactical{}, err
		}
		indikators = append(indikators, indikator)
	}

	return indikators, nil
}

func (repository *IndikatorPokinOpdTacticalRepositoryImpl) FindByPokinOpdTacticalId(ctx context.Context, tx *sql.Tx, pokinOpdTacticalId int) ([]domain.IndikatorPokinOpdTactical, error) {
	query := "SELECT id, pokin_opd_tactical_id, nama_indikator FROM indikator_pokin_opd_tactical WHERE pokin_opd_tactical_id = $1 ORDER BY id ASC"
	rows, err := tx.QueryContext(ctx, query, pokinOpdTacticalId)
	if err != nil {
		return []domain.IndikatorPokinOpdTactical{}, err
	}
	defer rows.Close()

	var indikators []domain.IndikatorPokinOpdTactical
	for rows.Next() {
		var indikator domain.IndikatorPokinOpdTactical
		err := rows.Scan(&indikator.Id, &indikator.PokinOpdTacticalId, &indikator.NamaIndikator)
		if err != nil {
			return []domain.IndikatorPokinOpdTactical{}, err
		}
		indikators = append(indikators, indikator)
	}

	return indikators, nil
}
