package repository

import (
	"context"
	"database/sql"
	"kupang-service/model/domain"
)

type IndikatorPokinOpdStrategicRepository interface {
	Create(ctx context.Context, tx *sql.Tx, indikator domain.IndikatorPokinOpdStrategic) (domain.IndikatorPokinOpdStrategic, error)
	Update(ctx context.Context, tx *sql.Tx, indikator domain.IndikatorPokinOpdStrategic) (domain.IndikatorPokinOpdStrategic, error)
	Delete(ctx context.Context, tx *sql.Tx, id int) error
	FindById(ctx context.Context, tx *sql.Tx, id int) (domain.IndikatorPokinOpdStrategic, error)
	FindAll(ctx context.Context, tx *sql.Tx) ([]domain.IndikatorPokinOpdStrategic, error)
	FindByPokinOpdStrategicId(ctx context.Context, tx *sql.Tx, pokinOpdStrategicId int) ([]domain.IndikatorPokinOpdStrategic, error)
}
