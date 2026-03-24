package repository

import (
	"context"
	"database/sql"
	"kupang-service/model/domain"
)

type IndikatorPokinOpdOperationalRepository interface {
	Create(ctx context.Context, tx *sql.Tx, indikator domain.IndikatorPokinOpdOperational) (domain.IndikatorPokinOpdOperational, error)
	Update(ctx context.Context, tx *sql.Tx, indikator domain.IndikatorPokinOpdOperational) (domain.IndikatorPokinOpdOperational, error)
	Delete(ctx context.Context, tx *sql.Tx, id int) error
	FindById(ctx context.Context, tx *sql.Tx, id int) (domain.IndikatorPokinOpdOperational, error)
	FindAll(ctx context.Context, tx *sql.Tx) ([]domain.IndikatorPokinOpdOperational, error)
	FindByPokinOpdOperationalId(ctx context.Context, tx *sql.Tx, pokinOpdOperationalId int) ([]domain.IndikatorPokinOpdOperational, error)
}
