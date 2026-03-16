package repository

import (
	"context"
	"database/sql"
	"errors"
	"kupang-service/model/domain"
)

type IndikatorPokinOpdStrategicRepositoryImpl struct {
}

func NewIndikatorPokinOpdStrategicRepositoryImpl() *IndikatorPokinOpdStrategicRepositoryImpl {
	return &IndikatorPokinOpdStrategicRepositoryImpl{}
}

func (repository *IndikatorPokinOpdStrategicRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, indikator domain.IndikatorPokinOpdStrategic) (domain.IndikatorPokinOpdStrategic, error) {
	query := "INSERT INTO indikator_pokin_opd_strategic (pokin_opd_strategic_id, nama_indikator) VALUES ($1, $2) RETURNING id"
	err := tx.QueryRowContext(ctx, query, indikator.PokinOpdStrategicId, indikator.NamaIndikator).Scan(&indikator.Id)
	if err != nil {
		return domain.IndikatorPokinOpdStrategic{}, err
	}

	return indikator, nil
}

func (repository *IndikatorPokinOpdStrategicRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, indikator domain.IndikatorPokinOpdStrategic) (domain.IndikatorPokinOpdStrategic, error) {
	query := "UPDATE indikator_pokin_opd_strategic SET pokin_opd_strategic_id = $1, nama_indikator = $2, last_modified_date = NOW() WHERE id = $3"
	_, err := tx.ExecContext(ctx, query, indikator.PokinOpdStrategicId, indikator.NamaIndikator, indikator.Id)
	if err != nil {
		return domain.IndikatorPokinOpdStrategic{}, err
	}

	return indikator, nil
}

func (repository *IndikatorPokinOpdStrategicRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, id int) error {
	query := "DELETE FROM indikator_pokin_opd_strategic WHERE id = $1"
	_, err := tx.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}
	return nil
}

func (repository *IndikatorPokinOpdStrategicRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, id int) (domain.IndikatorPokinOpdStrategic, error) {
	query := "SELECT id, pokin_opd_strategic_id, nama_indikator FROM indikator_pokin_opd_strategic WHERE id = $1"
	row := tx.QueryRowContext(ctx, query, id)

	var indikator domain.IndikatorPokinOpdStrategic
	err := row.Scan(&indikator.Id, &indikator.PokinOpdStrategicId, &indikator.NamaIndikator)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.IndikatorPokinOpdStrategic{}, errors.New("id tidak ditemukan")
		}
		return domain.IndikatorPokinOpdStrategic{}, err
	}

	return indikator, nil
}

func (repository *IndikatorPokinOpdStrategicRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) ([]domain.IndikatorPokinOpdStrategic, error) {
	query := "SELECT id, pokin_opd_strategic_id, nama_indikator FROM indikator_pokin_opd_strategic ORDER BY id ASC"
	rows, err := tx.QueryContext(ctx, query)
	if err != nil {
		return []domain.IndikatorPokinOpdStrategic{}, err
	}
	defer rows.Close()

	var indikators []domain.IndikatorPokinOpdStrategic
	for rows.Next() {
		var indikator domain.IndikatorPokinOpdStrategic
		err := rows.Scan(&indikator.Id, &indikator.PokinOpdStrategicId, &indikator.NamaIndikator)
		if err != nil {
			return []domain.IndikatorPokinOpdStrategic{}, err
		}
		indikators = append(indikators, indikator)
	}

	return indikators, nil
}

func (repository *IndikatorPokinOpdStrategicRepositoryImpl) FindByPokinOpdStrategicId(ctx context.Context, tx *sql.Tx, pokinOpdStrategicId int) ([]domain.IndikatorPokinOpdStrategic, error) {
	query := "SELECT id, pokin_opd_strategic_id, nama_indikator FROM indikator_pokin_opd_strategic WHERE pokin_opd_strategic_id = $1 ORDER BY id ASC"
	rows, err := tx.QueryContext(ctx, query, pokinOpdStrategicId)
	if err != nil {
		return []domain.IndikatorPokinOpdStrategic{}, err
	}
	defer rows.Close()

	var indikators []domain.IndikatorPokinOpdStrategic
	for rows.Next() {
		var indikator domain.IndikatorPokinOpdStrategic
		err := rows.Scan(&indikator.Id, &indikator.PokinOpdStrategicId, &indikator.NamaIndikator)
		if err != nil {
			return []domain.IndikatorPokinOpdStrategic{}, err
		}
		indikators = append(indikators, indikator)
	}

	return indikators, nil
}
