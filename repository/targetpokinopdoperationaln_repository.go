package repository

import (
	"context"
	"database/sql"
	"kupang-service/model/domain"
)

type TargetPokinOpdOperationalNRepository interface {
	Create(ctx context.Context, tx *sql.Tx, target domain.TargetPokinOpdOperationalN) (domain.TargetPokinOpdOperationalN, error)
	Update(ctx context.Context, tx *sql.Tx, target domain.TargetPokinOpdOperationalN) (domain.TargetPokinOpdOperationalN, error)
	Delete(ctx context.Context, tx *sql.Tx, id int) error
	FindById(ctx context.Context, tx *sql.Tx, id int) (domain.TargetPokinOpdOperationalN, error)
	FindAll(ctx context.Context, tx *sql.Tx) ([]domain.TargetPokinOpdOperationalN, error)
	FindByIndikatorId(ctx context.Context, tx *sql.Tx, indikatorId int) ([]domain.TargetPokinOpdOperationalN, error)
}
