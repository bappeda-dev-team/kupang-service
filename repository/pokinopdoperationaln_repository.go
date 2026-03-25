package repository

import (
	"context"
	"database/sql"
	"kupang-service/model/domain"
)

type PokinOpdOperationalNRepository interface {
	Create(ctx context.Context, tx *sql.Tx, pokin domain.PokinOpdOperationalN) (domain.PokinOpdOperationalN, error)
	Update(ctx context.Context, tx *sql.Tx, pokin domain.PokinOpdOperationalN) (domain.PokinOpdOperationalN, error)
	Delete(ctx context.Context, tx *sql.Tx, id int) error
	FindById(ctx context.Context, tx *sql.Tx, id int) (domain.PokinOpdOperationalN, error)
	FindAll(ctx context.Context, tx *sql.Tx) ([]domain.PokinOpdOperationalN, error)
	FindByParent(ctx context.Context, tx *sql.Tx, parent int) ([]domain.PokinOpdOperationalN, error)
}
