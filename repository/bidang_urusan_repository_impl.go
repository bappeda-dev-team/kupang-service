package repository

import (
	"context"
	"database/sql"
	"errors"
	"kupang-service/model/domain"
)

type BidangUrusanRepositoryImpl struct {
}

func NewBidangUrusanRepositoryImpl() *BidangUrusanRepositoryImpl {
	return &BidangUrusanRepositoryImpl{}
}

func (repository *BidangUrusanRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, bidangUrusan domain.BidangUrusan) (domain.BidangUrusan, error) {
	query := "INSERT INTO bidang_urusan (kode_urusan, nama_urusan) VALUES ($1, $2) RETURNING id"
	err := tx.QueryRowContext(ctx, query, bidangUrusan.KodeUrusan, bidangUrusan.NamaUrusan).Scan(&bidangUrusan.Id)
	if err != nil {
		return domain.BidangUrusan{}, err
	}

	return bidangUrusan, nil
}

func (repository *BidangUrusanRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, bidangUrusan domain.BidangUrusan) (domain.BidangUrusan, error) {
	query := "UPDATE bidang_urusan SET kode_urusan = $1, nama_urusan = $2, last_modified_date = NOW() WHERE id = $3"
	_, err := tx.ExecContext(ctx, query, bidangUrusan.KodeUrusan, bidangUrusan.NamaUrusan, bidangUrusan.Id)
	if err != nil {
		return domain.BidangUrusan{}, err
	}

	return bidangUrusan, nil
}

func (repository *BidangUrusanRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, id int) (domain.BidangUrusan, error) {
	query := "SELECT id, kode_urusan, nama_urusan FROM bidang_urusan WHERE id = $1"
	row := tx.QueryRowContext(ctx, query, id)

	var bidangUrusan domain.BidangUrusan
	err := row.Scan(&bidangUrusan.Id, &bidangUrusan.KodeUrusan, &bidangUrusan.NamaUrusan)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.BidangUrusan{}, errors.New("id tidak ditemukan")
		}
		return domain.BidangUrusan{}, err
	}

	return bidangUrusan, nil
}

func (repository *BidangUrusanRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) ([]domain.BidangUrusan, error) {
	query := "SELECT id, kode_urusan, nama_urusan FROM bidang_urusan ORDER BY id ASC"
	rows, err := tx.QueryContext(ctx, query)
	if err != nil {
		return []domain.BidangUrusan{}, err
	}
	defer rows.Close()

	var bidangUrusanList []domain.BidangUrusan
	for rows.Next() {
		var bidangUrusan domain.BidangUrusan
		err := rows.Scan(&bidangUrusan.Id, &bidangUrusan.KodeUrusan, &bidangUrusan.NamaUrusan)
		if err != nil {
			return []domain.BidangUrusan{}, err
		}

		bidangUrusanList = append(bidangUrusanList, bidangUrusan)
	}

	return bidangUrusanList, nil
}

func (repository *BidangUrusanRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, id int) error {
	query := "DELETE FROM bidang_urusan WHERE id = $1"
	_, err := tx.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}
