package repository

import (
	"context"
	"database/sql"
	"kupang-service/model/domain"
)

type IndikatorPokinOpdOperationalNRepository interface {
	Create(ctx context.Context, tx *sql.Tx, indikator domain.IndikatorPokinOpdOperationalN) (domain.IndikatorPokinOpdOperationalN, error)
	Update(ctx context.Context, tx *sql.Tx, indikator domain.IndikatorPokinOpdOperationalN) (domain.IndikatorPokinOpdOperationalN, error)
	Delete(ctx context.Context, tx *sql.Tx, id int) error
	FindById(ctx context.Context, tx *sql.Tx, id int) (domain.IndikatorPokinOpdOperationalN, error)
	FindAll(ctx context.Context, tx *sql.Tx) ([]domain.IndikatorPokinOpdOperationalN, error)
	FindByPokinOpdOperationalNId(ctx context.Context, tx *sql.Tx, pokinOpdOperationalNId int) ([]domain.IndikatorPokinOpdOperationalN, error)
}
