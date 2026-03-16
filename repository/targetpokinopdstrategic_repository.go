package repository

import (
	"context"
	"database/sql"
	"kupang-service/model/domain"
)

type TargetPokinOpdStrategicRepository interface {
	Create(ctx context.Context, tx *sql.Tx, target domain.TargetPokinOpdStrategic) (domain.TargetPokinOpdStrategic, error)
	Update(ctx context.Context, tx *sql.Tx, target domain.TargetPokinOpdStrategic) (domain.TargetPokinOpdStrategic, error)
	Delete(ctx context.Context, tx *sql.Tx, id int) error
	FindById(ctx context.Context, tx *sql.Tx, id int) (domain.TargetPokinOpdStrategic, error)
	FindAll(ctx context.Context, tx *sql.Tx) ([]domain.TargetPokinOpdStrategic, error)
	FindByIndikatorId(ctx context.Context, tx *sql.Tx, indikatorId int) ([]domain.TargetPokinOpdStrategic, error)
}
