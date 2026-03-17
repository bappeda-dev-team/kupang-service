package repository

import (
	"context"
	"database/sql"
	"errors"
	"kupang-service/model/domain"
)

type PokinOpdTacticalRepositoryImpl struct {
}

func NewPokinOpdTacticalRepositoryImpl() *PokinOpdTacticalRepositoryImpl {
	return &PokinOpdTacticalRepositoryImpl{}
}

func (repository *PokinOpdTacticalRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, pokin domain.PokinOpdTactical) (domain.PokinOpdTactical, error) {
	query := "INSERT INTO pokin_opd_tactical (parent, nama_pohon, jenis_pohon, level_pohon, kode_opd, nama_opd, keterangan, tahun, jumlah_review, status, pelaksana, updated_by) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12) RETURNING id"
	err := tx.QueryRowContext(
		ctx,
		query,
		pokin.Parent,
		pokin.NamaPohon,
		pokin.JenisPohon,
		pokin.LevelPohon,
		pokin.KodeOpd,
		pokin.NamaOpd,
		pokin.Keterangan,
		pokin.Tahun,
		pokin.JumlahReview,
		pokin.Status,
		pokin.Pelaksana,
		pokin.UpdatedBy,
	).Scan(&pokin.Id)
	if err != nil {
		return domain.PokinOpdTactical{}, err
	}

	return pokin, nil
}

func (repository *PokinOpdTacticalRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, pokin domain.PokinOpdTactical) (domain.PokinOpdTactical, error) {
	query := "UPDATE pokin_opd_tactical SET parent = $1, nama_pohon = $2, jenis_pohon = $3, level_pohon = $4, kode_opd = $5, nama_opd = $6, keterangan = $7, tahun = $8, jumlah_review = $9, status = $10, pelaksana = $11, updated_by = $12, last_modified_date = NOW() WHERE id = $13"
	_, err := tx.ExecContext(
		ctx,
		query,
		pokin.Parent,
		pokin.NamaPohon,
		pokin.JenisPohon,
		pokin.LevelPohon,
		pokin.KodeOpd,
		pokin.NamaOpd,
		pokin.Keterangan,
		pokin.Tahun,
		pokin.JumlahReview,
		pokin.Status,
		pokin.Pelaksana,
		pokin.UpdatedBy,
		pokin.Id,
	)
	if err != nil {
		return domain.PokinOpdTactical{}, err
	}

	return pokin, nil
}

func (repository *PokinOpdTacticalRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, id int) error {
	query := "DELETE FROM pokin_opd_tactical WHERE id = $1"
	_, err := tx.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}
	return nil
}

func (repository *PokinOpdTacticalRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, id int) (domain.PokinOpdTactical, error) {
	query := "SELECT id, parent, nama_pohon, jenis_pohon, level_pohon, kode_opd, nama_opd, keterangan, tahun, jumlah_review, status, pelaksana, updated_by FROM pokin_opd_tactical WHERE id = $1"
	row := tx.QueryRowContext(ctx, query, id)

	var pokin domain.PokinOpdTactical
	err := row.Scan(
		&pokin.Id,
		&pokin.Parent,
		&pokin.NamaPohon,
		&pokin.JenisPohon,
		&pokin.LevelPohon,
		&pokin.KodeOpd,
		&pokin.NamaOpd,
		&pokin.Keterangan,
		&pokin.Tahun,
		&pokin.JumlahReview,
		&pokin.Status,
		&pokin.Pelaksana,
		&pokin.UpdatedBy,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.PokinOpdTactical{}, errors.New("id tidak ditemukan")
		}
		return domain.PokinOpdTactical{}, err
	}

	return pokin, nil
}

func (repository *PokinOpdTacticalRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) ([]domain.PokinOpdTactical, error) {
	query := "SELECT id, parent, nama_pohon, jenis_pohon, level_pohon, kode_opd, nama_opd, keterangan, tahun, jumlah_review, status, pelaksana, updated_by FROM pokin_opd_tactical ORDER BY id ASC"
	rows, err := tx.QueryContext(ctx, query)
	if err != nil {
		return []domain.PokinOpdTactical{}, err
	}
	defer rows.Close()

	var pokins []domain.PokinOpdTactical
	for rows.Next() {
		var pokin domain.PokinOpdTactical
		err := rows.Scan(
			&pokin.Id,
			&pokin.Parent,
			&pokin.NamaPohon,
			&pokin.JenisPohon,
			&pokin.LevelPohon,
			&pokin.KodeOpd,
			&pokin.NamaOpd,
			&pokin.Keterangan,
			&pokin.Tahun,
			&pokin.JumlahReview,
			&pokin.Status,
			&pokin.Pelaksana,
			&pokin.UpdatedBy,
		)
		if err != nil {
			return []domain.PokinOpdTactical{}, err
		}
		pokins = append(pokins, pokin)
	}

	return pokins, nil
}

func (repository *PokinOpdTacticalRepositoryImpl) FindByParent(ctx context.Context, tx *sql.Tx, parent int) ([]domain.PokinOpdTactical, error) {
	query := "SELECT id, parent, nama_pohon, jenis_pohon, level_pohon, kode_opd, nama_opd, keterangan, tahun, jumlah_review, status, pelaksana, updated_by FROM pokin_opd_tactical WHERE parent = $1 ORDER BY id ASC"
	rows, err := tx.QueryContext(ctx, query, parent)
	if err != nil {
		return []domain.PokinOpdTactical{}, err
	}
	defer rows.Close()

	var pokins []domain.PokinOpdTactical
	for rows.Next() {
		var pokin domain.PokinOpdTactical
		err := rows.Scan(
			&pokin.Id,
			&pokin.Parent,
			&pokin.NamaPohon,
			&pokin.JenisPohon,
			&pokin.LevelPohon,
			&pokin.KodeOpd,
			&pokin.NamaOpd,
			&pokin.Keterangan,
			&pokin.Tahun,
			&pokin.JumlahReview,
			&pokin.Status,
			&pokin.Pelaksana,
			&pokin.UpdatedBy,
		)
		if err != nil {
			return []domain.PokinOpdTactical{}, err
		}
		pokins = append(pokins, pokin)
	}

	return pokins, nil
}
