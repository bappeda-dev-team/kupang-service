package repository

import (
	"context"
	"database/sql"
	"errors"
	"kupang-service/model/domain"
)

type PokinOpdStrategicRepositoryImpl struct {
}

func NewPokinOpdStrategicRepositoryImpl() *PokinOpdStrategicRepositoryImpl {
	return &PokinOpdStrategicRepositoryImpl{}
}

func (repository *PokinOpdStrategicRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, pokin domain.PokinOpdStrategic) (domain.PokinOpdStrategic, error) {
	query := "INSERT INTO pokin_opd_strategic (parent, nama_pohon, jenis_pohon, level_pohon, kode_opd, nama_opd, keterangan, tahun, jumlah_review, status, pelaksana, updated_by) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12) RETURNING id"
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
		return domain.PokinOpdStrategic{}, err
	}

	return pokin, nil
}

func (repository *PokinOpdStrategicRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, pokin domain.PokinOpdStrategic) (domain.PokinOpdStrategic, error) {
	query := "UPDATE pokin_opd_strategic SET parent = $1, nama_pohon = $2, jenis_pohon = $3, level_pohon = $4, kode_opd = $5, nama_opd = $6, keterangan = $7, tahun = $8, jumlah_review = $9, status = $10, pelaksana = $11, updated_by = $12, last_modified_date = NOW() WHERE id = $13"
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
		return domain.PokinOpdStrategic{}, err
	}

	return pokin, nil
}

func (repository *PokinOpdStrategicRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, id int) error {
	query := "DELETE FROM pokin_opd_strategic WHERE id = $1"
	_, err := tx.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}
	return nil
}

func (repository *PokinOpdStrategicRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, id int) (domain.PokinOpdStrategic, error) {
	query := "SELECT id, parent, nama_pohon, jenis_pohon, level_pohon, kode_opd, nama_opd, keterangan, tahun, jumlah_review, status, pelaksana, updated_by FROM pokin_opd_strategic WHERE id = $1"
	row := tx.QueryRowContext(ctx, query, id)

	var pokin domain.PokinOpdStrategic
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
			return domain.PokinOpdStrategic{}, errors.New("id tidak ditemukan")
		}
		return domain.PokinOpdStrategic{}, err
	}

	return pokin, nil
}

func (repository *PokinOpdStrategicRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) ([]domain.PokinOpdStrategic, error) {
	query := "SELECT id, parent, nama_pohon, jenis_pohon, level_pohon, kode_opd, nama_opd, keterangan, tahun, jumlah_review, status, pelaksana, updated_by FROM pokin_opd_strategic ORDER BY id ASC"
	rows, err := tx.QueryContext(ctx, query)
	if err != nil {
		return []domain.PokinOpdStrategic{}, err
	}
	defer rows.Close()

	var pokins []domain.PokinOpdStrategic
	for rows.Next() {
		var pokin domain.PokinOpdStrategic
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
			return []domain.PokinOpdStrategic{}, err
		}
		pokins = append(pokins, pokin)
	}

	return pokins, nil
}

func (repository *PokinOpdStrategicRepositoryImpl) FindByKodeOpdAndTahun(ctx context.Context, tx *sql.Tx, kodeOpd string, tahun int) ([]domain.PokinOpdStrategic, error) {
	query := "SELECT id, parent, nama_pohon, jenis_pohon, level_pohon, kode_opd, nama_opd, keterangan, tahun, jumlah_review, status, pelaksana, updated_by FROM pokin_opd_strategic WHERE kode_opd = $1 AND tahun = $2 ORDER BY id ASC"
	rows, err := tx.QueryContext(ctx, query, kodeOpd, tahun)
	if err != nil {
		return []domain.PokinOpdStrategic{}, err
	}
	defer rows.Close()

	var pokins []domain.PokinOpdStrategic
	for rows.Next() {
		var pokin domain.PokinOpdStrategic
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
			return []domain.PokinOpdStrategic{}, err
		}
		pokins = append(pokins, pokin)
	}

	return pokins, nil
}

func (repository *PokinOpdStrategicRepositoryImpl) FindByKodeOpdTahunParentLevel(ctx context.Context, tx *sql.Tx, kodeOpd string, tahun int, parent int, levelPohon int) ([]domain.PokinOpdStrategic, error) {
	query := "SELECT id, parent, nama_pohon, jenis_pohon, level_pohon, kode_opd, nama_opd, keterangan, tahun, jumlah_review, status, pelaksana, updated_by FROM pokin_opd_strategic WHERE kode_opd = $1 AND tahun = $2 AND parent = $3 AND level_pohon = $4 ORDER BY id ASC"
	rows, err := tx.QueryContext(ctx, query, kodeOpd, tahun, parent, levelPohon)
	if err != nil {
		return []domain.PokinOpdStrategic{}, err
	}
	defer rows.Close()

	var pokins []domain.PokinOpdStrategic
	for rows.Next() {
		var pokin domain.PokinOpdStrategic
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
			return []domain.PokinOpdStrategic{}, err
		}
		pokins = append(pokins, pokin)
	}

	return pokins, nil
}
