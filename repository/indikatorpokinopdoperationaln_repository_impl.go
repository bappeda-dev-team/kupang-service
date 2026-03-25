package repository

import (
	"context"
	"database/sql"
	"errors"
	"kupang-service/model/domain"
)

type IndikatorPokinOpdOperationalNRepositoryImpl struct {}

func NewIndikatorPokinOpdOperationalNRepositoryImpl() *IndikatorPokinOpdOperationalNRepositoryImpl {
	return &IndikatorPokinOpdOperationalNRepositoryImpl{}
}

func (repository *IndikatorPokinOpdOperationalNRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, indikator domain.IndikatorPokinOpdOperationalN) (domain.IndikatorPokinOpdOperationalN, error) {
	query := "INSERT INTO indikator_pokin_opd_operationalN (pokin_opd_operationalN_id, nama_indikator) VALUES ($1, $2) RETURNING id"
	err := tx.QueryRowContext(ctx, query, indikator.PokinOpdOperationalNId, indikator.NamaIndikator).Scan(&indikator.Id)
	if err != nil {
		return domain.IndikatorPokinOpdOperationalN{}, err
	}

	return indikator, nil
}

func (repository *IndikatorPokinOpdOperationalNRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, indikator domain.IndikatorPokinOpdOperationalN) (domain.IndikatorPokinOpdOperationalN, error) {
	query := "UPDATE indikator_pokin_opd_operationalN SET pokin_opd_operationalN_id = $1, nama_indikator = $2, last_modified_date = NOW() WHERE id = $3"
	_, err := tx.ExecContext(ctx, query, indikator.PokinOpdOperationalNId, indikator.NamaIndikator, indikator.Id)
	if err != nil {
		return domain.IndikatorPokinOpdOperationalN{}, err
	}

	return indikator, nil
}

func (repository *IndikatorPokinOpdOperationalNRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, id int) error {
	query := "DELETE FROM indikator_pokin_opd_operationalN WHERE id = $1"
	_, err := tx.ExecContext(ctx, query, id)
	return err
}

func (repository *IndikatorPokinOpdOperationalNRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, id int) (domain.IndikatorPokinOpdOperationalN, error) {
	query := "SELECT id, pokin_opd_operationalN_id, nama_indikator FROM indikator_pokin_opd_operationalN WHERE id = $1"
	row := tx.QueryRowContext(ctx, query, id)

	var indikator domain.IndikatorPokinOpdOperationalN
	err := row.Scan(&indikator.Id, &indikator.PokinOpdOperationalNId, &indikator.NamaIndikator)
	if err != nil {
		if err == sql.ErrNoRows {
			return domain.IndikatorPokinOpdOperationalN{}, errors.New("id tidak ditemukan")
		}
		return domain.IndikatorPokinOpdOperationalN{}, err
	}

	return indikator, nil
}

func (repository *IndikatorPokinOpdOperationalNRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) ([]domain.IndikatorPokinOpdOperationalN, error) {
	query := "SELECT id, pokin_opd_operationalN_id, nama_indikator FROM indikator_pokin_opd_operationalN ORDER BY id ASC"
	rows, err := tx.QueryContext(ctx, query)
	if err != nil {
		return []domain.IndikatorPokinOpdOperationalN{}, err
	}
	defer rows.Close()

	var indikators []domain.IndikatorPokinOpdOperationalN
	for rows.Next() {
		var indikator domain.IndikatorPokinOpdOperationalN
		err := rows.Scan(&indikator.Id, &indikator.PokinOpdOperationalNId, &indikator.NamaIndikator)
		if err != nil {
			return []domain.IndikatorPokinOpdOperationalN{}, err
		}
		indikators = append(indikators, indikator)
	}

	return indikators, nil
}

func (repository *IndikatorPokinOpdOperationalNRepositoryImpl) FindByPokinOpdOperationalNId(ctx context.Context, tx *sql.Tx, pokinOpdOperationalNId int) ([]domain.IndikatorPokinOpdOperationalN, error) {
	query := "SELECT id, pokin_opd_operationalN_id, nama_indikator FROM indikator_pokin_opd_operationalN WHERE pokin_opd_operationalN_id = $1 ORDER BY id ASC"
	rows, err := tx.QueryContext(ctx, query, pokinOpdOperationalNId)
	if err != nil {
		return []domain.IndikatorPokinOpdOperationalN{}, err
	}
	defer rows.Close()

	var indikators []domain.IndikatorPokinOpdOperationalN
	for rows.Next() {
		var indikator domain.IndikatorPokinOpdOperationalN
		err := rows.Scan(&indikator.Id, &indikator.PokinOpdOperationalNId, &indikator.NamaIndikator)
		if err != nil {
			return []domain.IndikatorPokinOpdOperationalN{}, err
		}
		indikators = append(indikators, indikator)
	}

	return indikators, nil
}
