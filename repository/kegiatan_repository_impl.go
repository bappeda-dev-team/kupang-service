package repository

import (
	"context"
	"database/sql"
	"errors"
	"kupang-service/model/domain"
)

type KegiatanRepositoryImpl struct {
}

func NewKegiatanRepositoryImpl() *KegiatanRepositoryImpl {
	return &KegiatanRepositoryImpl{}
}

func (repository *KegiatanRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, kegiatan domain.Kegiatan) (domain.Kegiatan, error) {
	query := "INSERT INTO kegiatan (kode_kegiatan, nama_kegiatan, tahun) VALUES ($1, $2, $3) RETURNING id"
	err := tx.QueryRowContext(ctx, query, kegiatan.KodeKegiatan, kegiatan.NamaKegiatan, kegiatan.Tahun).Scan(&kegiatan.Id)
	if err != nil {
		return domain.Kegiatan{}, err
	}

	return kegiatan, nil
}

func (repository *KegiatanRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, kegiatan domain.Kegiatan) (domain.Kegiatan, error) {
	query := "UPDATE kegiatan SET kode_kegiatan = $1, nama_kegiatan = $2, tahun = $3 WHERE id = $4 RETURNING id"
	err := tx.QueryRowContext(ctx, query, kegiatan.KodeKegiatan, kegiatan.NamaKegiatan, kegiatan.Tahun, kegiatan.Id).Scan(&kegiatan.Id)
	if err != nil {
		return domain.Kegiatan{}, err
	}

	return kegiatan, nil
}

func (repository *KegiatanRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, id int) (domain.Kegiatan, error) {
	query := "SELECT id, kode_kegiatan, nama_kegiatan, tahun FROM kegiatan WHERE id = $1"
	row := tx.QueryRowContext(ctx, query, id)

	var kegiatan domain.Kegiatan
	err := row.Scan(&kegiatan.Id, &kegiatan.KodeKegiatan, &kegiatan.NamaKegiatan, &kegiatan.Tahun)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.Kegiatan{}, errors.New("id tidak ditemukan")
		}
		return domain.Kegiatan{}, err
	}

	return kegiatan, nil
}

func (repository *KegiatanRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) ([]domain.Kegiatan, error) {
	query := "SELECT id, kode_kegiatan, nama_kegiatan, tahun FROM kegiatan ORDER BY id ASC"
	rows, err := tx.QueryContext(ctx, query)
	if err != nil {
		return []domain.Kegiatan{}, err
	}
	defer rows.Close()

	var kegiatanList []domain.Kegiatan
	for rows.Next() {
		var kegiatan domain.Kegiatan
		err := rows.Scan(&kegiatan.Id, &kegiatan.KodeKegiatan, &kegiatan.NamaKegiatan, &kegiatan.Tahun)
		if err != nil {
			return []domain.Kegiatan{}, err
		}

		kegiatanList = append(kegiatanList, kegiatan)
	}

	return kegiatanList, nil
}

func (repository *KegiatanRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, id int) error {
	query := "DELETE FROM kegiatan WHERE id = $1"
	_, err := tx.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}
