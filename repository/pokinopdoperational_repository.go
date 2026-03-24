package repository

import (
	"context"
	"database/sql"
	"kupang-service/model/domain"
)

type PokinOpdOperationalRepository interface {
	Create(ctx context.Context, tx *sql.Tx, pokin domain.PokinOpdOperational) (domain.PokinOpdOperational, error)
	Update(ctx context.Context, tx *sql.Tx, pokin domain.PokinOpdOperational) (domain.PokinOpdOperational, error)
	Delete(ctx context.Context, tx *sql.Tx, id int) error
	FindById(ctx context.Context, tx *sql.Tx, id int) (domain.PokinOpdOperational, error)
	FindAll(ctx context.Context, tx *sql.Tx) ([]domain.PokinOpdOperational, error)
	FindByParent(ctx context.Context, tx *sql.Tx, parent int) ([]domain.PokinOpdOperational, error)
}
