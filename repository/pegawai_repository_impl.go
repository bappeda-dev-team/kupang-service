package repository

import (
	"context"
	"database/sql"
	"errors"
	"kupang-service/model/domain"
)

type PegawaiRepositoryImpl struct {
}

func NewPegawaiRepositoryImpl() *PegawaiRepositoryImpl {
	return &PegawaiRepositoryImpl{}
}

func (repository *PegawaiRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, pegawai domain.Pegawai) (domain.Pegawai, error) {
	query := "INSERT INTO pegawai (nama, nip, jabatan_id, nama_jabatan, kode_opd, nama_opd, jenis_pegawai) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id"
	err := tx.QueryRowContext(ctx, query, pegawai.Nama, pegawai.Nip, pegawai.JabatanId, pegawai.NamaJabatan, pegawai.KodeOpd, pegawai.NamaOpd, pegawai.JenisPegawai).Scan(&pegawai.Id)
	if err != nil {
		return domain.Pegawai{}, err
	}

	return pegawai, nil
}

func (repository *PegawaiRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, pegawai domain.Pegawai) (domain.Pegawai, error) {
	query := "UPDATE pegawai SET nama = $1, nip = $2, jabatan_id = $3, nama_jabatan = $4, kode_opd = $5, nama_opd = $6, jenis_pegawai = $7, last_modified_date = NOW() WHERE id = $8"
	_, err := tx.ExecContext(ctx, query, pegawai.Nama, pegawai.Nip, pegawai.JabatanId, pegawai.NamaJabatan, pegawai.KodeOpd, pegawai.NamaOpd, pegawai.JenisPegawai, pegawai.Id)
	if err != nil {
		return domain.Pegawai{}, err
	}

	return pegawai, nil
}

func (repository *PegawaiRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, id int) error {
	query := "DELETE FROM pegawai WHERE id = $1"
	_, err := tx.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}

func (repository *PegawaiRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, id int) (domain.Pegawai, error) {
	query := "SELECT id, nama, nip, jabatan_id, nama_jabatan, kode_opd, nama_opd, jenis_pegawai FROM pegawai WHERE id = $1"
	row := tx.QueryRowContext(ctx, query, id)

	var pegawai domain.Pegawai
	err := row.Scan(&pegawai.Id, &pegawai.Nama, &pegawai.Nip, &pegawai.JabatanId, &pegawai.NamaJabatan, &pegawai.KodeOpd, &pegawai.NamaOpd, &pegawai.JenisPegawai)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.Pegawai{}, errors.New("id tidak ditemukan")
		}
		return domain.Pegawai{}, err
	}

	return pegawai, nil
}

func (repository *PegawaiRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) ([]domain.Pegawai, error) {
	query := "SELECT id, nama, nip, jabatan_id, nama_jabatan, kode_opd, nama_opd, jenis_pegawai FROM pegawai ORDER BY id ASC"
	rows, err := tx.QueryContext(ctx, query)
	if err != nil {
		return []domain.Pegawai{}, err
	}
	defer rows.Close()

	var pegawaiList []domain.Pegawai
	for rows.Next() {
		var pegawai domain.Pegawai
		err := rows.Scan(&pegawai.Id, &pegawai.Nama, &pegawai.Nip, &pegawai.JabatanId, &pegawai.NamaJabatan, &pegawai.KodeOpd, &pegawai.NamaOpd, &pegawai.JenisPegawai)
		if err != nil {
			return []domain.Pegawai{}, err
		}

		pegawaiList = append(pegawaiList, pegawai)
	}

	return pegawaiList, nil
}

func (repository *PegawaiRepositoryImpl) FindByKodeOpd(ctx context.Context, tx *sql.Tx, kodeOpd string) ([]domain.Pegawai, error) {
	query := "SELECT id, nama, nip, jabatan_id, nama_jabatan, kode_opd, nama_opd, jenis_pegawai FROM pegawai WHERE kode_opd = $1 ORDER BY id ASC"
	rows, err := tx.QueryContext(ctx, query, kodeOpd)
	if err != nil {
		return []domain.Pegawai{}, err
	}
	defer rows.Close()

	var pegawaiList []domain.Pegawai
	for rows.Next() {
		var pegawai domain.Pegawai
		err := rows.Scan(&pegawai.Id, &pegawai.Nama, &pegawai.Nip, &pegawai.JabatanId, &pegawai.NamaJabatan, &pegawai.KodeOpd, &pegawai.NamaOpd, &pegawai.JenisPegawai)
		if err != nil {
			return []domain.Pegawai{}, err
		}

		pegawaiList = append(pegawaiList, pegawai)
	}

	return pegawaiList, nil
}
