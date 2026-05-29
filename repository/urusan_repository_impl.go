package repository

import (
	"context"
	"database/sql"
	"errors"
	"kupang-service/model/domain"
)

type UrusanRepositoryImpl struct {
}

func NewUrusanRepositoryImpl() *UrusanRepositoryImpl {
	return &UrusanRepositoryImpl{}
}

func (repository *UrusanRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, urusan domain.Urusan) (domain.Urusan, error) {
	query := "INSERT INTO urusan (kode_urusan, nama_urusan) VALUES ($1, $2) RETURNING id, created_date, last_modified_date"
	err := tx.QueryRowContext(ctx, query, urusan.KodeUrusan, urusan.NamaUrusan).Scan(&urusan.Id, &urusan.CreatedDate, &urusan.LastModifiedDate)
	if err != nil {
		return domain.Urusan{}, err
	}
	return urusan, nil
}

func (repository *UrusanRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, urusan domain.Urusan) (domain.Urusan, error) {
	query := "UPDATE urusan SET kode_urusan = $1, nama_urusan = $2, last_modified_date = NOW() WHERE id = $3 RETURNING created_date, last_modified_date"
	err := tx.QueryRowContext(ctx, query, urusan.KodeUrusan, urusan.NamaUrusan, urusan.Id).Scan(&urusan.CreatedDate, &urusan.LastModifiedDate)
	if err != nil {
		return domain.Urusan{}, err
	}
	return urusan, nil
}

func (repository *UrusanRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, id int) (domain.Urusan, error) {
	query := "SELECT id, kode_urusan, nama_urusan, created_date, last_modified_date FROM urusan WHERE id = $1"
	row := tx.QueryRowContext(ctx, query, id)
	var urusan domain.Urusan
	err := row.Scan(&urusan.Id, &urusan.KodeUrusan, &urusan.NamaUrusan, &urusan.CreatedDate, &urusan.LastModifiedDate)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.Urusan{}, errors.New("id tidak ditemukan")
		}
		return domain.Urusan{}, err
	}
	return urusan, nil
}

func (repository *UrusanRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) ([]domain.Urusan, error) {
	query := "SELECT id, kode_urusan, nama_urusan, created_date, last_modified_date FROM urusan ORDER BY id ASC"
	rows, err := tx.QueryContext(ctx, query)
	if err != nil {
		return []domain.Urusan{}, err
	}
	defer rows.Close()
	var urusanList []domain.Urusan
	for rows.Next() {
		var urusan domain.Urusan
		err := rows.Scan(&urusan.Id, &urusan.KodeUrusan, &urusan.NamaUrusan, &urusan.CreatedDate, &urusan.LastModifiedDate)
		if err != nil {
			return []domain.Urusan{}, err
		}
		urusanList = append(urusanList, urusan)
	}
	return urusanList, nil
}

func (repository *UrusanRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, id int) error {
	query := "DELETE FROM urusan WHERE id = $1"
	_, err := tx.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}
	return nil
}
