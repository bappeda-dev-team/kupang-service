package repository

import (
	"context"
	"database/sql"
	"errors"
	"kupang-service/model/domain"
)

type MusrenbangRepositoryImpl struct{}

func NewMusrenbangRepositoryImpl() *MusrenbangRepositoryImpl {
	return &MusrenbangRepositoryImpl{}
}

func (repository *MusrenbangRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, musrenbang domain.Musrenbang) (domain.Musrenbang, error) {
	query := `INSERT INTO "musrenbang" (usulan, alamat, uraian, tahun, kode_opd, nama_opd, status) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id, usulan, alamat, uraian, tahun, kode_opd, nama_opd, status, created_date, last_modified_date`
	row := tx.QueryRowContext(ctx, query, musrenbang.Usulan, musrenbang.Alamat, musrenbang.Uraian, musrenbang.Tahun, musrenbang.KodeOpd, musrenbang.NamaOpd, musrenbang.Status)
	err := row.Scan(&musrenbang.Id, &musrenbang.Usulan, &musrenbang.Alamat, &musrenbang.Uraian, &musrenbang.Tahun, &musrenbang.KodeOpd, &musrenbang.NamaOpd, &musrenbang.Status, &musrenbang.CreatedDate, &musrenbang.LastModifiedDate)
	if err != nil {
		return domain.Musrenbang{}, err
	}

	return musrenbang, nil
}

func (repository *MusrenbangRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, musrenbang domain.Musrenbang) (domain.Musrenbang, error) {
	query := `UPDATE "musrenbang" SET usulan = $1, alamat = $2, uraian = $3, tahun = $4, kode_opd = $5, nama_opd = $6, status = $7, last_modified_date = NOW() WHERE id = $8 RETURNING id, usulan, alamat, uraian, tahun, kode_opd, nama_opd, status, created_date, last_modified_date`
	row := tx.QueryRowContext(ctx, query, musrenbang.Usulan, musrenbang.Alamat, musrenbang.Uraian, musrenbang.Tahun, musrenbang.KodeOpd, musrenbang.NamaOpd, musrenbang.Status, musrenbang.Id)
	err := row.Scan(&musrenbang.Id, &musrenbang.Usulan, &musrenbang.Alamat, &musrenbang.Uraian, &musrenbang.Tahun, &musrenbang.KodeOpd, &musrenbang.NamaOpd, &musrenbang.Status, &musrenbang.CreatedDate, &musrenbang.LastModifiedDate)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.Musrenbang{}, errors.New("id tidak ditemukan")
		}
		return domain.Musrenbang{}, err
	}

	return musrenbang, nil
}

func (repository *MusrenbangRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, id int) error {
	query := `DELETE FROM "musrenbang" WHERE id = $1`
	_, err := tx.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}

func (repository *MusrenbangRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, id int) (domain.Musrenbang, error) {
	query := `SELECT id, usulan, alamat, uraian, tahun, kode_opd, nama_opd, status, created_date, last_modified_date FROM "musrenbang" WHERE id = $1`
	row := tx.QueryRowContext(ctx, query, id)

	var musrenbang domain.Musrenbang
	err := row.Scan(&musrenbang.Id, &musrenbang.Usulan, &musrenbang.Alamat, &musrenbang.Uraian, &musrenbang.Tahun, &musrenbang.KodeOpd, &musrenbang.NamaOpd, &musrenbang.Status, &musrenbang.CreatedDate, &musrenbang.LastModifiedDate)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.Musrenbang{}, errors.New("id tidak ditemukan")
		}
		return domain.Musrenbang{}, err
	}

	return musrenbang, nil
}

func (repository *MusrenbangRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) ([]domain.Musrenbang, error) {
	query := `SELECT id, usulan, alamat, uraian, tahun, kode_opd, nama_opd, status, created_date, last_modified_date FROM "musrenbang" ORDER BY id ASC`
	rows, err := tx.QueryContext(ctx, query)
	if err != nil {
		return []domain.Musrenbang{}, err
	}
	defer rows.Close()

	var musrenbangs []domain.Musrenbang
	for rows.Next() {
		var musrenbang domain.Musrenbang
		err := rows.Scan(&musrenbang.Id, &musrenbang.Usulan, &musrenbang.Alamat, &musrenbang.Uraian, &musrenbang.Tahun, &musrenbang.KodeOpd, &musrenbang.NamaOpd, &musrenbang.Status, &musrenbang.CreatedDate, &musrenbang.LastModifiedDate)
		if err != nil {
			return nil, err
		}
		musrenbangs = append(musrenbangs, musrenbang)
	}

	return musrenbangs, nil
}

func (repository *MusrenbangRepositoryImpl) FindByKodeOpd(ctx context.Context, tx *sql.Tx, kodeOpd string) ([]domain.Musrenbang, error) {
	query := `SELECT id, usulan, alamat, uraian, tahun, kode_opd, nama_opd, status, created_date, last_modified_date FROM "musrenbang" WHERE kode_opd = $1 ORDER BY id ASC`
	rows, err := tx.QueryContext(ctx, query, kodeOpd)
	if err != nil {
		return []domain.Musrenbang{}, err
	}
	defer rows.Close()

	var musrenbangs []domain.Musrenbang
	for rows.Next() {
		var musrenbang domain.Musrenbang
		err := rows.Scan(&musrenbang.Id, &musrenbang.Usulan, &musrenbang.Alamat, &musrenbang.Uraian, &musrenbang.Tahun, &musrenbang.KodeOpd, &musrenbang.NamaOpd, &musrenbang.Status, &musrenbang.CreatedDate, &musrenbang.LastModifiedDate)
		if err != nil {
			return nil, err
		}
		musrenbangs = append(musrenbangs, musrenbang)
	}

	return musrenbangs, nil
}
