package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"kupang-service/model/domain"
	"strings"
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

func (repository *PegawaiRepositoryImpl) UpdateJabatan(ctx context.Context, tx *sql.Tx, jabatan domain.Jabatan) (domain.Jabatan, error) {
	query := "UPDATE jabatan SET nama_jabatan = $1, tahun = COALESCE($2, tahun), last_modified_date = NOW() WHERE id = $3"
	result, err := tx.ExecContext(ctx, query, jabatan.NamaJabatan, jabatan.Tahun, jabatan.Id)
	if err != nil {
		return domain.Jabatan{}, err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return domain.Jabatan{}, err
	}

	if rows == 0 {
		return domain.Jabatan{}, errors.New("id tidak ditemukan")
	}

	return jabatan, nil
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
	query := "SELECT p.id, p.nama, p.nip, p.jabatan_id, p.nama_jabatan, p.kode_opd, p.nama_opd, p.jenis_pegawai, j.tahun FROM pegawai p LEFT JOIN jabatan j ON p.jabatan_id = j.id WHERE p.id = $1"
	row := tx.QueryRowContext(ctx, query, id)

	var pegawai domain.Pegawai
	err := row.Scan(&pegawai.Id, &pegawai.Nama, &pegawai.Nip, &pegawai.JabatanId, &pegawai.NamaJabatan, &pegawai.KodeOpd, &pegawai.NamaOpd, &pegawai.JenisPegawai, &pegawai.TahunJabatan)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.Pegawai{}, errors.New("id tidak ditemukan")
		}
		return domain.Pegawai{}, err
	}

	return pegawai, nil
}

func (repository *PegawaiRepositoryImpl) FindJabatanById(ctx context.Context, tx *sql.Tx, id int) (domain.Jabatan, error) {
	query := "SELECT id, nama_jabatan, tahun FROM jabatan WHERE id = $1"
	row := tx.QueryRowContext(ctx, query, id)

	var jabatan domain.Jabatan
	err := row.Scan(&jabatan.Id, &jabatan.NamaJabatan, &jabatan.Tahun)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.Jabatan{}, errors.New("id tidak ditemukan")
		}
		return domain.Jabatan{}, err
	}

	return jabatan, nil
}

func (repository *PegawaiRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) ([]domain.Pegawai, error) {
	query := "SELECT p.id, p.nama, p.nip, p.jabatan_id, p.nama_jabatan, p.kode_opd, p.nama_opd, p.jenis_pegawai, j.tahun FROM pegawai p LEFT JOIN jabatan j ON p.jabatan_id = j.id ORDER BY p.id ASC"
	rows, err := tx.QueryContext(ctx, query)
	if err != nil {
		return []domain.Pegawai{}, err
	}
	defer rows.Close()

	var pegawaiList []domain.Pegawai
	for rows.Next() {
		var pegawai domain.Pegawai
		err := rows.Scan(&pegawai.Id, &pegawai.Nama, &pegawai.Nip, &pegawai.JabatanId, &pegawai.NamaJabatan, &pegawai.KodeOpd, &pegawai.NamaOpd, &pegawai.JenisPegawai, &pegawai.TahunJabatan)
		if err != nil {
			return []domain.Pegawai{}, err
		}

		pegawaiList = append(pegawaiList, pegawai)
	}

	return pegawaiList, nil
}

func (repository *PegawaiRepositoryImpl) FindByKodeOpd(ctx context.Context, tx *sql.Tx, kodeOpd string) ([]domain.Pegawai, error) {
	query := "SELECT p.id, p.nama, p.nip, p.jabatan_id, p.nama_jabatan, p.kode_opd, p.nama_opd, p.jenis_pegawai, j.tahun FROM pegawai p LEFT JOIN jabatan j ON p.jabatan_id = j.id WHERE p.kode_opd = $1 ORDER BY p.id ASC"
	rows, err := tx.QueryContext(ctx, query, kodeOpd)
	if err != nil {
		return []domain.Pegawai{}, err
	}
	defer rows.Close()

	var pegawaiList []domain.Pegawai
	for rows.Next() {
		var pegawai domain.Pegawai
		err := rows.Scan(&pegawai.Id, &pegawai.Nama, &pegawai.Nip, &pegawai.JabatanId, &pegawai.NamaJabatan, &pegawai.KodeOpd, &pegawai.NamaOpd, &pegawai.JenisPegawai, &pegawai.TahunJabatan)
		if err != nil {
			return []domain.Pegawai{}, err
		}

		pegawaiList = append(pegawaiList, pegawai)
	}

	return pegawaiList, nil
}

func (repository *PegawaiRepositoryImpl) SearchByNamaOrNip(ctx context.Context, tx *sql.Tx, nama, nip *string) ([]domain.Pegawai, error) {
	query := "SELECT p.id, p.nama, p.nip, p.jabatan_id, p.nama_jabatan, p.kode_opd, p.nama_opd, p.jenis_pegawai, j.tahun FROM pegawai p LEFT JOIN jabatan j ON p.jabatan_id = j.id"
	var conditions []string
	var args []interface{}
	paramIndex := 1

	if nama != nil && *nama != "" {
		conditions = append(conditions, fmt.Sprintf("LOWER(p.nama) LIKE $%d", paramIndex))
		args = append(args, fmt.Sprintf("%%%s%%", strings.ToLower(*nama)))
		paramIndex++
	}

	if nip != nil && *nip != "" {
		conditions = append(conditions, fmt.Sprintf("LOWER(p.nip) LIKE $%d", paramIndex))
		args = append(args, fmt.Sprintf("%%%s%%", strings.ToLower(*nip)))
		paramIndex++
	}

	if len(conditions) > 0 {
		query += " WHERE " + strings.Join(conditions, " AND ")
	}

	query += " ORDER BY p.id ASC"

	rows, err := tx.QueryContext(ctx, query, args...)
	if err != nil {
		return []domain.Pegawai{}, err
	}
	defer rows.Close()

	var pegawaiList []domain.Pegawai
	for rows.Next() {
		var pegawai domain.Pegawai
		err := rows.Scan(&pegawai.Id, &pegawai.Nama, &pegawai.Nip, &pegawai.JabatanId, &pegawai.NamaJabatan, &pegawai.KodeOpd, &pegawai.NamaOpd, &pegawai.JenisPegawai, &pegawai.TahunJabatan)
		if err != nil {
			return []domain.Pegawai{}, err
		}

		pegawaiList = append(pegawaiList, pegawai)
	}

	return pegawaiList, nil
}

func (repository *PegawaiRepositoryImpl) UpdatePegawaiNamaJabatanByJabatanId(ctx context.Context, tx *sql.Tx, jabatanId int, namaJabatan string) error {
	query := "UPDATE pegawai SET nama_jabatan = $1 WHERE jabatan_id = $2"
	_, err := tx.ExecContext(ctx, query, namaJabatan, jabatanId)
	if err != nil {
		return err
	}

	return nil
}

func (repository *PegawaiRepositoryImpl) FindAllJabatan(ctx context.Context, tx *sql.Tx) ([]domain.Jabatan, error) {
	query := "SELECT id, nama_jabatan, tahun FROM jabatan ORDER BY id ASC"
	rows, err := tx.QueryContext(ctx, query)
	if err != nil {
		return []domain.Jabatan{}, err
	}
	defer rows.Close()

	var jabatanList []domain.Jabatan
	for rows.Next() {
		var jabatan domain.Jabatan
		err := rows.Scan(&jabatan.Id, &jabatan.NamaJabatan, &jabatan.Tahun)
		if err != nil {
			return []domain.Jabatan{}, err
		}

		jabatanList = append(jabatanList, jabatan)
	}

	return jabatanList, nil
}
