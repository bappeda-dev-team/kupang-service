package repository

import (
	"context"
	"database/sql"
	"errors"
	"kupang-service/model/domain"
)

type IndikatorPokinOpdOperationalRepositoryImpl struct {
}

func NewIndikatorPokinOpdOperationalRepositoryImpl() *IndikatorPokinOpdOperationalRepositoryImpl {
	return &IndikatorPokinOpdOperationalRepositoryImpl{}
}

func (repository *IndikatorPokinOpdOperationalRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, indikator domain.IndikatorPokinOpdOperational) (domain.IndikatorPokinOpdOperational, error) {
	query := "INSERT INTO indikator_pokin_opd_operational (pokin_opd_operational_id, nama_indikator) VALUES ($1, $2) RETURNING id"
	err := tx.QueryRowContext(ctx, query, indikator.PokinOpdOperationalId, indikator.NamaIndikator).Scan(&indikator.Id)
	if err != nil {
		return domain.IndikatorPokinOpdOperational{}, err
	}

	return indikator, nil
}

func (repository *IndikatorPokinOpdOperationalRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, indikator domain.IndikatorPokinOpdOperational) (domain.IndikatorPokinOpdOperational, error) {
	query := "UPDATE indikator_pokin_opd_operational SET pokin_opd_operational_id = $1, nama_indikator = $2, last_modified_date = NOW() WHERE id = $3"
	_, err := tx.ExecContext(ctx, query, indikator.PokinOpdOperationalId, indikator.NamaIndikator, indikator.Id)
	if err != nil {
		return domain.IndikatorPokinOpdOperational{}, err
	}

	return indikator, nil
}

func (repository *IndikatorPokinOpdOperationalRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, id int) error {
	query := "DELETE FROM indikator_pokin_opd_operational WHERE id = $1"
	_, err := tx.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}

func (repository *IndikatorPokinOpdOperationalRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, id int) (domain.IndikatorPokinOpdOperational, error) {
	query := "SELECT id, pokin_opd_operational_id, nama_indikator FROM indikator_pokin_opd_operational WHERE id = $1"
	row := tx.QueryRowContext(ctx, query, id)

	var indikator domain.IndikatorPokinOpdOperational
	err := row.Scan(&indikator.Id, &indikator.PokinOpdOperationalId, &indikator.NamaIndikator)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.IndikatorPokinOpdOperational{}, errors.New("id tidak ditemukan")
		}
		return domain.IndikatorPokinOpdOperational{}, err
	}

	return indikator, nil
}

func (repository *IndikatorPokinOpdOperationalRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) ([]domain.IndikatorPokinOpdOperational, error) {
	query := "SELECT id, pokin_opd_operational_id, nama_indikator FROM indikator_pokin_opd_operational ORDER BY id ASC"
	rows, err := tx.QueryContext(ctx, query)
	if err != nil {
		return []domain.IndikatorPokinOpdOperational{}, err
	}
	defer rows.Close()

	var indikators []domain.IndikatorPokinOpdOperational
	for rows.Next() {
		var indikator domain.IndikatorPokinOpdOperational
		err := rows.Scan(&indikator.Id, &indikator.PokinOpdOperationalId, &indikator.NamaIndikator)
		if err != nil {
			return []domain.IndikatorPokinOpdOperational{}, err
		}
		indikators = append(indikators, indikator)
	}

	return indikators, nil
}

func (repository *IndikatorPokinOpdOperationalRepositoryImpl) FindByPokinOpdOperationalId(ctx context.Context, tx *sql.Tx, pokinOpdOperationalId int) ([]domain.IndikatorPokinOpdOperational, error) {
	query := "SELECT id, pokin_opd_operational_id, nama_indikator FROM indikator_pokin_opd_operational WHERE pokin_opd_operational_id = $1 ORDER BY id ASC"
	rows, err := tx.QueryContext(ctx, query, pokinOpdOperationalId)
	if err != nil {
		return []domain.IndikatorPokinOpdOperational{}, err
	}
	defer rows.Close()

	var indikators []domain.IndikatorPokinOpdOperational
	for rows.Next() {
		var indikator domain.IndikatorPokinOpdOperational
		err := rows.Scan(&indikator.Id, &indikator.PokinOpdOperationalId, &indikator.NamaIndikator)
		if err != nil {
			return []domain.IndikatorPokinOpdOperational{}, err
		}
		indikators = append(indikators, indikator)
	}

	return indikators, nil
}
