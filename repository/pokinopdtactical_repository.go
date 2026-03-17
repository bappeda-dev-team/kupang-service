package repository

import (
	"context"
	"database/sql"
	"kupang-service/model/domain"
)

type PokinOpdTacticalRepository interface {
	Create(ctx context.Context, tx *sql.Tx, pokin domain.PokinOpdTactical) (domain.PokinOpdTactical, error)
	Update(ctx context.Context, tx *sql.Tx, pokin domain.PokinOpdTactical) (domain.PokinOpdTactical, error)
	Delete(ctx context.Context, tx *sql.Tx, id int) error
	FindById(ctx context.Context, tx *sql.Tx, id int) (domain.PokinOpdTactical, error)
	FindAll(ctx context.Context, tx *sql.Tx) ([]domain.PokinOpdTactical, error)
}
