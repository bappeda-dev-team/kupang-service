package repository

import (
	"context"
	"database/sql"
	"errors"
	"kupang-service/model/domain"
)

type ProgramRepositoryImpl struct {
}

func NewProgramRepositoryImpl() *ProgramRepositoryImpl {
	return &ProgramRepositoryImpl{}
}

func (repository *ProgramRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, program domain.Program) (domain.Program, error) {
	query := "INSERT INTO program (kode_urusan, nama_urusan, tahun) VALUES ($1, $2, $3) RETURNING id"
	err := tx.QueryRowContext(ctx, query, program.KodeProgram, program.NamaProgram, program.Tahun).Scan(&program.Id)
	if err != nil {
		return domain.Program{}, err
	}

	return program, nil
}

func (repository *ProgramRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, program domain.Program) (domain.Program, error) {
	query := "UPDATE program SET kode_urusan = $1, nama_urusan = $2, tahun = $3 WHERE id = $4 RETURNING id"
	err := tx.QueryRowContext(ctx, query, program.KodeProgram, program.NamaProgram, program.Tahun, program.Id).Scan(&program.Id)
	if err != nil {
		return domain.Program{}, err
	}

	return program, nil
}

func (repository *ProgramRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, id int) (domain.Program, error) {
	query := "SELECT id, kode_urusan, nama_urusan, tahun FROM program WHERE id = $1"
	row := tx.QueryRowContext(ctx, query, id)

	var program domain.Program
	err := row.Scan(&program.Id, &program.KodeProgram, &program.NamaProgram, &program.Tahun)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.Program{}, errors.New("id tidak ditemukan")
		}
		return domain.Program{}, err
	}

	return program, nil
}

func (repository *ProgramRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) ([]domain.Program, error) {
	query := "SELECT id, kode_urusan, nama_urusan, tahun FROM program ORDER BY id ASC"
	rows, err := tx.QueryContext(ctx, query)
	if err != nil {
		return []domain.Program{}, err
	}
	defer rows.Close()

	var programList []domain.Program
	for rows.Next() {
		var program domain.Program
		err := rows.Scan(&program.Id, &program.KodeProgram, &program.NamaProgram, &program.Tahun)
		if err != nil {
			return []domain.Program{}, err
		}

		programList = append(programList, program)
	}

	return programList, nil
}

func (repository *ProgramRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, id int) error {
	query := "DELETE FROM program WHERE id = $1"
	_, err := tx.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}
