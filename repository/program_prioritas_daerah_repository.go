package repository

import (
	"context"
	"database/sql"
	"kupang-service/model/domain"
)

type ProgramPrioritasDaerahRepository interface {
	Create(ctx context.Context, tx *sql.Tx, program domain.ProgramPrioritasDaerah) (domain.ProgramPrioritasDaerah, error)
	Update(ctx context.Context, tx *sql.Tx, program domain.ProgramPrioritasDaerah) (domain.ProgramPrioritasDaerah, error)
	Delete(ctx context.Context, tx *sql.Tx, id int) error
	FindById(ctx context.Context, tx *sql.Tx, id int) (domain.ProgramPrioritasDaerah, error)
	FindAll(ctx context.Context, tx *sql.Tx) ([]domain.ProgramPrioritasDaerah, error)
	FindByTahunRange(ctx context.Context, tx *sql.Tx, tahunAwal, tahunAkhir string) ([]domain.ProgramPrioritasDaerah, error)
}
