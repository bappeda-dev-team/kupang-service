package repository

import (
	"context"
	"database/sql"
	"kupang-service/model/domain"
)

type IndikatorPokinOpdRepository interface {
	Create(ctx context.Context, tx *sql.Tx, indikatorPokinOpd domain.IndikatorPokinOpd) (domain.IndikatorPokinOpd, error)
	Update(ctx context.Context, tx *sql.Tx, indikatorPokinOpd domain.IndikatorPokinOpd) (domain.IndikatorPokinOpd, error)
	Delete(ctx context.Context, tx *sql.Tx, id int) error
	FindById(ctx context.Context, tx *sql.Tx, id int) (domain.IndikatorPokinOpd, error)
	FindAll(ctx context.Context, tx *sql.Tx) ([]domain.IndikatorPokinOpd, error)
	FindByTujuanPokinOpdId(ctx context.Context, tx *sql.Tx, tujuanPokinOpdId int) ([]domain.IndikatorPokinOpd, error)
}
