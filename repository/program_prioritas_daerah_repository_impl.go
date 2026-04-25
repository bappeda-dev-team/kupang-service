package repository

import (
	"context"
	"database/sql"
	"errors"
	"kupang-service/model/domain"
)

type ProgramPrioritasDaerahRepositoryImpl struct{}

func NewProgramPrioritasDaerahRepositoryImpl() *ProgramPrioritasDaerahRepositoryImpl {
	return &ProgramPrioritasDaerahRepositoryImpl{}
}

func (repository *ProgramPrioritasDaerahRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, program domain.ProgramPrioritasDaerah) (domain.ProgramPrioritasDaerah, error) {
	query := `INSERT INTO "program_prioritas_daerah" (kode_program_prioritas_daerah, nama_program_prioritas_daerah, rencana_implementasi, keterangan, tahun_awal, tahun_akhir, is_active) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id, kode_program_prioritas_daerah, nama_program_prioritas_daerah, rencana_implementasi, keterangan, tahun_awal, tahun_akhir, is_active, created_date, last_modified_date`
	row := tx.QueryRowContext(ctx, query, program.KodeProgramPrioritasDaerah, program.NamaProgramPrioritasDaerah, program.RencanaImplementasi, program.Keterangan, program.TahunAwal, program.TahunAkhir, program.IsActive)
	err := row.Scan(&program.Id, &program.KodeProgramPrioritasDaerah, &program.NamaProgramPrioritasDaerah, &program.RencanaImplementasi, &program.Keterangan, &program.TahunAwal, &program.TahunAkhir, &program.IsActive, &program.CreatedDate, &program.LastModifiedDate)
	if err != nil {
		return domain.ProgramPrioritasDaerah{}, err
	}

	return program, nil
}

func (repository *ProgramPrioritasDaerahRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, program domain.ProgramPrioritasDaerah) (domain.ProgramPrioritasDaerah, error) {
	query := `UPDATE "program_prioritas_daerah" SET kode_program_prioritas_daerah = $1, nama_program_prioritas_daerah = $2, rencana_implementasi = $3, keterangan = $4, tahun_awal = $5, tahun_akhir = $6, is_active = $7, last_modified_date = NOW() WHERE id = $8 RETURNING id, kode_program_prioritas_daerah, nama_program_prioritas_daerah, rencana_implementasi, keterangan, tahun_awal, tahun_akhir, is_active, created_date, last_modified_date`
	row := tx.QueryRowContext(ctx, query, program.KodeProgramPrioritasDaerah, program.NamaProgramPrioritasDaerah, program.RencanaImplementasi, program.Keterangan, program.TahunAwal, program.TahunAkhir, program.IsActive, program.Id)
	err := row.Scan(&program.Id, &program.KodeProgramPrioritasDaerah, &program.NamaProgramPrioritasDaerah, &program.RencanaImplementasi, &program.Keterangan, &program.TahunAwal, &program.TahunAkhir, &program.IsActive, &program.CreatedDate, &program.LastModifiedDate)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.ProgramPrioritasDaerah{}, errors.New("id tidak ditemukan")
		}
		return domain.ProgramPrioritasDaerah{}, err
	}

	return program, nil
}

func (repository *ProgramPrioritasDaerahRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, id int) error {
	query := `DELETE FROM "program_prioritas_daerah" WHERE id = $1`
	_, err := tx.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}

func (repository *ProgramPrioritasDaerahRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, id int) (domain.ProgramPrioritasDaerah, error) {
	query := `SELECT id, kode_program_prioritas_daerah, nama_program_prioritas_daerah, rencana_implementasi, keterangan, tahun_awal, tahun_akhir, is_active, created_date, last_modified_date FROM "program_prioritas_daerah" WHERE id = $1`
	row := tx.QueryRowContext(ctx, query, id)

	var program domain.ProgramPrioritasDaerah
	err := row.Scan(&program.Id, &program.KodeProgramPrioritasDaerah, &program.NamaProgramPrioritasDaerah, &program.RencanaImplementasi, &program.Keterangan, &program.TahunAwal, &program.TahunAkhir, &program.IsActive, &program.CreatedDate, &program.LastModifiedDate)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.ProgramPrioritasDaerah{}, errors.New("id tidak ditemukan")
		}
		return domain.ProgramPrioritasDaerah{}, err
	}

	return program, nil
}

func (repository *ProgramPrioritasDaerahRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) ([]domain.ProgramPrioritasDaerah, error) {
	query := `SELECT id, kode_program_prioritas_daerah, nama_program_prioritas_daerah, rencana_implementasi, keterangan, tahun_awal, tahun_akhir, is_active, created_date, last_modified_date FROM "program_prioritas_daerah" ORDER BY id ASC`
	rows, err := tx.QueryContext(ctx, query)
	if err != nil {
		return []domain.ProgramPrioritasDaerah{}, err
	}
	defer rows.Close()

	var programs []domain.ProgramPrioritasDaerah
	for rows.Next() {
		var program domain.ProgramPrioritasDaerah
		err := rows.Scan(&program.Id, &program.KodeProgramPrioritasDaerah, &program.NamaProgramPrioritasDaerah, &program.RencanaImplementasi, &program.Keterangan, &program.TahunAwal, &program.TahunAkhir, &program.IsActive, &program.CreatedDate, &program.LastModifiedDate)
		if err != nil {
			return nil, err
		}
		programs = append(programs, program)
	}

	return programs, nil
}

func (repository *ProgramPrioritasDaerahRepositoryImpl) FindByTahunRange(ctx context.Context, tx *sql.Tx, tahunAwal, tahunAkhir string) ([]domain.ProgramPrioritasDaerah, error) {
	query := `SELECT id, kode_program_prioritas_daerah, nama_program_prioritas_daerah, rencana_implementasi, keterangan, tahun_awal, tahun_akhir, is_active, created_date, last_modified_date FROM "program_prioritas_daerah" WHERE tahun_awal = $1 AND tahun_akhir = $2 ORDER BY id ASC`
	rows, err := tx.QueryContext(ctx, query, tahunAwal, tahunAkhir)
	if err != nil {
		return []domain.ProgramPrioritasDaerah{}, err
	}
	defer rows.Close()

	var programs []domain.ProgramPrioritasDaerah
	for rows.Next() {
		var program domain.ProgramPrioritasDaerah
		err := rows.Scan(&program.Id, &program.KodeProgramPrioritasDaerah, &program.NamaProgramPrioritasDaerah, &program.RencanaImplementasi, &program.Keterangan, &program.TahunAwal, &program.TahunAkhir, &program.IsActive, &program.CreatedDate, &program.LastModifiedDate)
		if err != nil {
			return nil, err
		}
		programs = append(programs, program)
	}

	return programs, nil
}
