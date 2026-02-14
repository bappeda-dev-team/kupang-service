package repository

import (
	"context"
	"database/sql"
	"kupang-service/model/domain"
)

type TargetPokinOpdRepository interface {
	Create(ctx context.Context, tx *sql.Tx, targetPokinOpd domain.TargetPokinOpd) (domain.TargetPokinOpd, error)
	Update(ctx context.Context, tx *sql.Tx, targetPokinOpd domain.TargetPokinOpd) (domain.TargetPokinOpd, error)
	Delete(ctx context.Context, tx *sql.Tx, id int) error
	FindById(ctx context.Context, tx *sql.Tx, id int) (domain.TargetPokinOpd, error)
	FindAll(ctx context.Context, tx *sql.Tx) ([]domain.TargetPokinOpd, error)
	FindByIndikatorPokinOpdId(ctx context.Context, tx *sql.Tx, indikatorPokinOpdId int) ([]domain.TargetPokinOpd, error)
}
