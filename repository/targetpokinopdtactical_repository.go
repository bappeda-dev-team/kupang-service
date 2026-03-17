package repository

import (
	"context"
	"database/sql"
	"kupang-service/model/domain"
)

type TargetPokinOpdTacticalRepository interface {
	Create(ctx context.Context, tx *sql.Tx, target domain.TargetPokinOpdTactical) (domain.TargetPokinOpdTactical, error)
	Update(ctx context.Context, tx *sql.Tx, target domain.TargetPokinOpdTactical) (domain.TargetPokinOpdTactical, error)
	Delete(ctx context.Context, tx *sql.Tx, id int) error
	FindById(ctx context.Context, tx *sql.Tx, id int) (domain.TargetPokinOpdTactical, error)
	FindAll(ctx context.Context, tx *sql.Tx) ([]domain.TargetPokinOpdTactical, error)
	FindByIndikatorId(ctx context.Context, tx *sql.Tx, indikatorId int) ([]domain.TargetPokinOpdTactical, error)
}
