package repository

import (
	"context"
	"database/sql"
	"errors"
	"kupang-service/model/domain"
)

type PokokPikiranRepositoryImpl struct{}

func NewPokokPikiranRepositoryImpl() *PokokPikiranRepositoryImpl {
	return &PokokPikiranRepositoryImpl{}
}

func (repository *PokokPikiranRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, pokokPikiran domain.PokokPikiran) (domain.PokokPikiran, error) {
	query := `INSERT INTO "pokok_pikiran" (usulan, alamat, uraian, tahun, kode_opd, nama_opd, status) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id, usulan, alamat, uraian, tahun, kode_opd, nama_opd, status, created_date, last_modified_date`
	row := tx.QueryRowContext(ctx, query, pokokPikiran.Usulan, pokokPikiran.Alamat, pokokPikiran.Uraian, pokokPikiran.Tahun, pokokPikiran.KodeOpd, pokokPikiran.NamaOpd, pokokPikiran.Status)
	err := row.Scan(&pokokPikiran.Id, &pokokPikiran.Usulan, &pokokPikiran.Alamat, &pokokPikiran.Uraian, &pokokPikiran.Tahun, &pokokPikiran.KodeOpd, &pokokPikiran.NamaOpd, &pokokPikiran.Status, &pokokPikiran.CreatedDate, &pokokPikiran.LastModifiedDate)
	if err != nil {
		return domain.PokokPikiran{}, err
	}

	return pokokPikiran, nil
}

func (repository *PokokPikiranRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, pokokPikiran domain.PokokPikiran) (domain.PokokPikiran, error) {
	query := `UPDATE "pokok_pikiran" SET usulan = $1, alamat = $2, uraian = $3, tahun = $4, kode_opd = $5, nama_opd = $6, status = $7, last_modified_date = NOW() WHERE id = $8 RETURNING id, usulan, alamat, uraian, tahun, kode_opd, nama_opd, status, created_date, last_modified_date`
	row := tx.QueryRowContext(ctx, query, pokokPikiran.Usulan, pokokPikiran.Alamat, pokokPikiran.Uraian, pokokPikiran.Tahun, pokokPikiran.KodeOpd, pokokPikiran.NamaOpd, pokokPikiran.Status, pokokPikiran.Id)
	err := row.Scan(&pokokPikiran.Id, &pokokPikiran.Usulan, &pokokPikiran.Alamat, &pokokPikiran.Uraian, &pokokPikiran.Tahun, &pokokPikiran.KodeOpd, &pokokPikiran.NamaOpd, &pokokPikiran.Status, &pokokPikiran.CreatedDate, &pokokPikiran.LastModifiedDate)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.PokokPikiran{}, errors.New("id tidak ditemukan")
		}
		return domain.PokokPikiran{}, err
	}

	return pokokPikiran, nil
}

func (repository *PokokPikiranRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, id int) error {
	query := `DELETE FROM "pokok_pikiran" WHERE id = $1`
	_, err := tx.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}

func (repository *PokokPikiranRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, id int) (domain.PokokPikiran, error) {
	query := `SELECT id, usulan, alamat, uraian, tahun, kode_opd, nama_opd, status, created_date, last_modified_date FROM "pokok_pikiran" WHERE id = $1`
	row := tx.QueryRowContext(ctx, query, id)

	var pokokPikiran domain.PokokPikiran
	err := row.Scan(&pokokPikiran.Id, &pokokPikiran.Usulan, &pokokPikiran.Alamat, &pokokPikiran.Uraian, &pokokPikiran.Tahun, &pokokPikiran.KodeOpd, &pokokPikiran.NamaOpd, &pokokPikiran.Status, &pokokPikiran.CreatedDate, &pokokPikiran.LastModifiedDate)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.PokokPikiran{}, errors.New("id tidak ditemukan")
		}
		return domain.PokokPikiran{}, err
	}

	return pokokPikiran, nil
}

func (repository *PokokPikiranRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) ([]domain.PokokPikiran, error) {
	query := `SELECT id, usulan, alamat, uraian, tahun, kode_opd, nama_opd, status, created_date, last_modified_date FROM "pokok_pikiran" ORDER BY id ASC`
	rows, err := tx.QueryContext(ctx, query)
	if err != nil {
		return []domain.PokokPikiran{}, err
	}
	defer rows.Close()

	var pokokPikirans []domain.PokokPikiran
	for rows.Next() {
		var pokokPikiran domain.PokokPikiran
		err := rows.Scan(&pokokPikiran.Id, &pokokPikiran.Usulan, &pokokPikiran.Alamat, &pokokPikiran.Uraian, &pokokPikiran.Tahun, &pokokPikiran.KodeOpd, &pokokPikiran.NamaOpd, &pokokPikiran.Status, &pokokPikiran.CreatedDate, &pokokPikiran.LastModifiedDate)
		if err != nil {
			return nil, err
		}
		pokokPikirans = append(pokokPikirans, pokokPikiran)
	}

	return pokokPikirans, nil
}

func (repository *PokokPikiranRepositoryImpl) FindByKodeOpd(ctx context.Context, tx *sql.Tx, kodeOpd string) ([]domain.PokokPikiran, error) {
	query := `SELECT id, usulan, alamat, uraian, tahun, kode_opd, nama_opd, status, created_date, last_modified_date FROM "pokok_pikiran" WHERE kode_opd = $1 ORDER BY id ASC`
	rows, err := tx.QueryContext(ctx, query, kodeOpd)
	if err != nil {
		return []domain.PokokPikiran{}, err
	}
	defer rows.Close()

	var pokokPikirans []domain.PokokPikiran
	for rows.Next() {
		var pokokPikiran domain.PokokPikiran
		err := rows.Scan(&pokokPikiran.Id, &pokokPikiran.Usulan, &pokokPikiran.Alamat, &pokokPikiran.Uraian, &pokokPikiran.Tahun, &pokokPikiran.KodeOpd, &pokokPikiran.NamaOpd, &pokokPikiran.Status, &pokokPikiran.CreatedDate, &pokokPikiran.LastModifiedDate)
		if err != nil {
			return nil, err
		}
		pokokPikirans = append(pokokPikirans, pokokPikiran)
	}

	return pokokPikirans, nil
}
