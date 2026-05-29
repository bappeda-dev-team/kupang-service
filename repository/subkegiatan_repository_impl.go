package repository

import (
	"context"
	"database/sql"
	"errors"
	"kupang-service/model/domain"
)

type SubkegiatanRepositoryImpl struct {
}

func NewSubkegiatanRepositoryImpl() *SubkegiatanRepositoryImpl {
	return &SubkegiatanRepositoryImpl{}
}

func (repository *SubkegiatanRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, subkegiatan domain.Subkegiatan) (domain.Subkegiatan, error) {
	query := "INSERT INTO subkegiatan (kode_subkegiatan, nama_subkegiatan, tahun) VALUES ($1, $2, $3) RETURNING id"
	err := tx.QueryRowContext(ctx, query, subkegiatan.KodeSubkegiatan, subkegiatan.NamaSubkegiatan, subkegiatan.Tahun).Scan(&subkegiatan.Id)
	if err != nil {
		return domain.Subkegiatan{}, err
	}

	return subkegiatan, nil
}

func (repository *SubkegiatanRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, subkegiatan domain.Subkegiatan) (domain.Subkegiatan, error) {
	query := "UPDATE subkegiatan SET kode_subkegiatan = $1, nama_subkegiatan = $2, tahun = $3 WHERE id = $4 RETURNING id"
	err := tx.QueryRowContext(ctx, query, subkegiatan.KodeSubkegiatan, subkegiatan.NamaSubkegiatan, subkegiatan.Tahun, subkegiatan.Id).Scan(&subkegiatan.Id)
	if err != nil {
		return domain.Subkegiatan{}, err
	}

	return subkegiatan, nil
}

func (repository *SubkegiatanRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, id int) (domain.Subkegiatan, error) {
	query := "SELECT id, kode_subkegiatan, nama_subkegiatan, tahun FROM subkegiatan WHERE id = $1"
	row := tx.QueryRowContext(ctx, query, id)

	var subkegiatan domain.Subkegiatan
	err := row.Scan(&subkegiatan.Id, &subkegiatan.KodeSubkegiatan, &subkegiatan.NamaSubkegiatan, &subkegiatan.Tahun)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.Subkegiatan{}, errors.New("id tidak ditemukan")
		}
		return domain.Subkegiatan{}, err
	}

	return subkegiatan, nil
}

func (repository *SubkegiatanRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) ([]domain.Subkegiatan, error) {
	query := "SELECT id, kode_subkegiatan, nama_subkegiatan, tahun FROM subkegiatan ORDER BY id ASC"
	rows, err := tx.QueryContext(ctx, query)
	if err != nil {
		return []domain.Subkegiatan{}, err
	}
	defer rows.Close()

	var subkegiatanList []domain.Subkegiatan
	for rows.Next() {
		var subkegiatan domain.Subkegiatan
		err := rows.Scan(&subkegiatan.Id, &subkegiatan.KodeSubkegiatan, &subkegiatan.NamaSubkegiatan, &subkegiatan.Tahun)
		if err != nil {
			return []domain.Subkegiatan{}, err
		}

		subkegiatanList = append(subkegiatanList, subkegiatan)
	}

	return subkegiatanList, nil
}

func (repository *SubkegiatanRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, id int) error {
	query := "DELETE FROM subkegiatan WHERE id = $1"
	_, err := tx.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}
