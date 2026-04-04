package repository

import (
	"context"
	"database/sql"
	"errors"
	"kupang-service/model/domain"
)

type PemdaRepositoryImpl struct {
}

func NewPemdaRepositoryImpl() *PemdaRepositoryImpl {
	return &PemdaRepositoryImpl{}
}

func (repository *PemdaRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, pemda domain.Pemda) (domain.Pemda, error) {
	query := "INSERT INTO pemda (kode_pemda, nama_pemda) VALUES ($1, $2) RETURNING id"
	err := tx.QueryRowContext(ctx, query, pemda.KodePemda, pemda.NamaPemda).Scan(&pemda.Id)
	if err != nil {
		return domain.Pemda{}, err
	}

	return pemda, nil
}

func (repository *PemdaRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, pemda domain.Pemda) (domain.Pemda, error) {
	query := "UPDATE pemda SET kode_pemda = $1, nama_pemda = $2, last_modified_date = NOW() WHERE id = $3"
	_, err := tx.ExecContext(ctx, query, pemda.KodePemda, pemda.NamaPemda, pemda.Id)
	if err != nil {
		return domain.Pemda{}, err
	}

	return pemda, nil
}

func (repository *PemdaRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, id int) (domain.Pemda, error) {
	query := "SELECT id, kode_pemda, nama_pemda FROM pemda WHERE id = $1"
	row := tx.QueryRowContext(ctx, query, id)

	var pemda domain.Pemda
	err := row.Scan(&pemda.Id, &pemda.KodePemda, &pemda.NamaPemda)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.Pemda{}, errors.New("id tidak ditemukan")
		}
		return domain.Pemda{}, err
	}

	return pemda, nil
}

func (repository *PemdaRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) ([]domain.Pemda, error) {
	query := "SELECT id, kode_pemda, nama_pemda FROM pemda ORDER BY id ASC"
	rows, err := tx.QueryContext(ctx, query)
	if err != nil {
		return []domain.Pemda{}, err
	}
	defer rows.Close()

	var pemdaList []domain.Pemda
	for rows.Next() {
		var pemda domain.Pemda
		err := rows.Scan(&pemda.Id, &pemda.KodePemda, &pemda.NamaPemda)
		if err != nil {
			return []domain.Pemda{}, err
		}

		pemdaList = append(pemdaList, pemda)
	}

	return pemdaList, nil
}

func (repository *PemdaRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, id int) error {
	query := "DELETE FROM pemda WHERE id = $1"
	_, err := tx.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}
