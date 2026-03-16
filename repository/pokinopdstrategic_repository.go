package repository

import (
	"context"
	"database/sql"
	"kupang-service/model/domain"
)

type PokinOpdStrategicRepository interface {
	Create(ctx context.Context, tx *sql.Tx, pokin domain.PokinOpdStrategic) (domain.PokinOpdStrategic, error)
	Update(ctx context.Context, tx *sql.Tx, pokin domain.PokinOpdStrategic) (domain.PokinOpdStrategic, error)
	Delete(ctx context.Context, tx *sql.Tx, id int) error
	FindById(ctx context.Context, tx *sql.Tx, id int) (domain.PokinOpdStrategic, error)
	FindAll(ctx context.Context, tx *sql.Tx) ([]domain.PokinOpdStrategic, error)
	FindByKodeOpdAndTahun(ctx context.Context, tx *sql.Tx, kodeOpd string, tahun int) ([]domain.PokinOpdStrategic, error)
	FindByKodeOpdTahunParentLevel(ctx context.Context, tx *sql.Tx, kodeOpd string, tahun int, parent int, levelPohon int) ([]domain.PokinOpdStrategic, error)
}
