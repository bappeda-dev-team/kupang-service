package repository

import (
	"context"
	"database/sql"
	"errors"
	"kupang-service/model/domain"
)

type LembagaRepositoryImpl struct {
}

func NewLembagaRepositoryImpl() *LembagaRepositoryImpl {
	return &LembagaRepositoryImpl{}
}

func (repository *LembagaRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, lembaga domain.Lembaga) (domain.Lembaga, error) {
	query := "INSERT INTO lembaga (kode_lembaga, nama_lembaga) VALUES ($1, $2) RETURNING id"
	err := tx.QueryRowContext(ctx, query, lembaga.KodeLembaga, lembaga.NamaLembaga).Scan(&lembaga.Id)
	if err != nil {
		return domain.Lembaga{}, err
	}

	return lembaga, nil
}

func (repository *LembagaRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, lembaga domain.Lembaga) (domain.Lembaga, error) {
	query := "UPDATE lembaga SET kode_lembaga = $1, nama_lembaga = $2, last_modified_date = NOW() WHERE id = $3"
	_, err := tx.ExecContext(ctx, query, lembaga.KodeLembaga, lembaga.NamaLembaga, lembaga.Id)
	if err != nil {
		return domain.Lembaga{}, err
	}

	return lembaga, nil
}

func (repository *LembagaRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, id int) (domain.Lembaga, error) {
	query := "SELECT id, kode_lembaga, nama_lembaga FROM lembaga WHERE id = $1"
	row := tx.QueryRowContext(ctx, query, id)

	var lembaga domain.Lembaga
	err := row.Scan(&lembaga.Id, &lembaga.KodeLembaga, &lembaga.NamaLembaga)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.Lembaga{}, errors.New("id tidak ditemukan")
		}
		return domain.Lembaga{}, err
	}

	return lembaga, nil
}

func (repository *LembagaRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) ([]domain.Lembaga, error) {
	query := "SELECT id, kode_lembaga, nama_lembaga FROM lembaga ORDER BY id ASC"
	rows, err := tx.QueryContext(ctx, query)
	if err != nil {
		return []domain.Lembaga{}, err
	}
	defer rows.Close()

	var lembagaList []domain.Lembaga
	for rows.Next() {
		var lembaga domain.Lembaga
		err := rows.Scan(&lembaga.Id, &lembaga.KodeLembaga, &lembaga.NamaLembaga)
		if err != nil {
			return []domain.Lembaga{}, err
		}

		lembagaList = append(lembagaList, lembaga)
	}

	return lembagaList, nil
}

func (repository *LembagaRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, id int) error {
	query := "DELETE FROM lembaga WHERE id = $1"
	_, err := tx.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}
