package repository

import (
	"context"
	"database/sql"
	"kupang-service/model/domain"
)

type ProgramRepository interface {
	Create(ctx context.Context, tx *sql.Tx, program domain.Program) (domain.Program, error)
	Update(ctx context.Context, tx *sql.Tx, program domain.Program) (domain.Program, error)
	Delete(ctx context.Context, tx *sql.Tx, id int) error
	FindById(ctx context.Context, tx *sql.Tx, id int) (domain.Program, error)
	FindAll(ctx context.Context, tx *sql.Tx) ([]domain.Program, error)
}
