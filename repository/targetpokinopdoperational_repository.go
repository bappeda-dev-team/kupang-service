package repository

import (
	"context"
	"database/sql"
	"kupang-service/model/domain"
)

type TargetPokinOpdOperationalRepository interface {
	Create(ctx context.Context, tx *sql.Tx, target domain.TargetPokinOpdOperational) (domain.TargetPokinOpdOperational, error)
	Update(ctx context.Context, tx *sql.Tx, target domain.TargetPokinOpdOperational) (domain.TargetPokinOpdOperational, error)
	Delete(ctx context.Context, tx *sql.Tx, id int) error
	FindById(ctx context.Context, tx *sql.Tx, id int) (domain.TargetPokinOpdOperational, error)
	FindAll(ctx context.Context, tx *sql.Tx) ([]domain.TargetPokinOpdOperational, error)
	FindByIndikatorId(ctx context.Context, tx *sql.Tx, indikatorId int) ([]domain.TargetPokinOpdOperational, error)
}
