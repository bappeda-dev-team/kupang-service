package repository

import (
	"context"
	"database/sql"
	"errors"
	"kupang-service/model/domain"
)

type TujuanPokinOpdRepositoryImpl struct {
}

func NewTujuanPokinOpdRepositoryImpl() *TujuanPokinOpdRepositoryImpl {
	return &TujuanPokinOpdRepositoryImpl{}
}

func (repository *TujuanPokinOpdRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, tujuanPokinOpd domain.TujuanPokinOpd) (domain.TujuanPokinOpd, error) {
	query := "INSERT INTO tujuan_pokin_opd (kode_opd, nama_tujuan, bidang_urusan, tahun_awal_periode, tahun_akhir_periode) VALUES ($1, $2, $3, $4, $5) RETURNING id"
	err := tx.QueryRowContext(ctx, query, tujuanPokinOpd.KodeOpd, tujuanPokinOpd.NamaTujuan, tujuanPokinOpd.BidangUrusan, tujuanPokinOpd.TahunAwalPeriode, tujuanPokinOpd.TahunAkhirPeriode).Scan(&tujuanPokinOpd.Id)
	if err != nil {
		return domain.TujuanPokinOpd{}, err
	}

	return tujuanPokinOpd, nil
}

func (repository *TujuanPokinOpdRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, tujuanPokinOpd domain.TujuanPokinOpd) (domain.TujuanPokinOpd, error) {
	query := "UPDATE tujuan_pokin_opd SET kode_opd = $1, nama_tujuan = $2, bidang_urusan = $3, tahun_awal_periode = $4, tahun_akhir_periode = $5, last_modified_date = NOW() WHERE id = $6"
	_, err := tx.ExecContext(ctx, query, tujuanPokinOpd.KodeOpd, tujuanPokinOpd.NamaTujuan, tujuanPokinOpd.BidangUrusan, tujuanPokinOpd.TahunAwalPeriode, tujuanPokinOpd.TahunAkhirPeriode, tujuanPokinOpd.Id)
	if err != nil {
		return domain.TujuanPokinOpd{}, err
	}

	return tujuanPokinOpd, nil
}

func (repository *TujuanPokinOpdRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, id int) (domain.TujuanPokinOpd, error) {
	query := "SELECT id, kode_opd, nama_tujuan, bidang_urusan, tahun_awal_periode, tahun_akhir_periode FROM tujuan_pokin_opd WHERE id = $1"
	row := tx.QueryRowContext(ctx, query, id)

	var tujuanPokinOpd domain.TujuanPokinOpd
	err := row.Scan(&tujuanPokinOpd.Id, &tujuanPokinOpd.KodeOpd, &tujuanPokinOpd.NamaTujuan, &tujuanPokinOpd.BidangUrusan, &tujuanPokinOpd.TahunAwalPeriode, &tujuanPokinOpd.TahunAkhirPeriode)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.TujuanPokinOpd{}, errors.New("id tidak ditemukan")
		}
		return domain.TujuanPokinOpd{}, err
	}

	return tujuanPokinOpd, nil
}

func (repository *TujuanPokinOpdRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) ([]domain.TujuanPokinOpd, error) {
	query := "SELECT id, kode_opd, nama_tujuan, bidang_urusan, tahun_awal_periode, tahun_akhir_periode FROM tujuan_pokin_opd ORDER BY id ASC"
	rows, err := tx.QueryContext(ctx, query)
	if err != nil {
		return []domain.TujuanPokinOpd{}, err
	}
	defer rows.Close()

	var tujuanPokinOpdList []domain.TujuanPokinOpd
	for rows.Next() {
		var tujuanPokinOpd domain.TujuanPokinOpd
		err := rows.Scan(&tujuanPokinOpd.Id, &tujuanPokinOpd.KodeOpd, &tujuanPokinOpd.NamaTujuan, &tujuanPokinOpd.BidangUrusan, &tujuanPokinOpd.TahunAwalPeriode, &tujuanPokinOpd.TahunAkhirPeriode)
		if err != nil {
			return []domain.TujuanPokinOpd{}, err
		}

		tujuanPokinOpdList = append(tujuanPokinOpdList, tujuanPokinOpd)
	}

	return tujuanPokinOpdList, nil
}

func (repository *TujuanPokinOpdRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, id int) error {
	query := "DELETE FROM tujuan_pokin_opd WHERE id = $1"
	_, err := tx.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}
