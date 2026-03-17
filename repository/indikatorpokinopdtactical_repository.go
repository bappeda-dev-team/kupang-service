package repository

import (
	"context"
	"database/sql"
	"kupang-service/model/domain"
)

type IndikatorPokinOpdTacticalRepository interface {
	Create(ctx context.Context, tx *sql.Tx, indikator domain.IndikatorPokinOpdTactical) (domain.IndikatorPokinOpdTactical, error)
	Update(ctx context.Context, tx *sql.Tx, indikator domain.IndikatorPokinOpdTactical) (domain.IndikatorPokinOpdTactical, error)
	Delete(ctx context.Context, tx *sql.Tx, id int) error
	FindById(ctx context.Context, tx *sql.Tx, id int) (domain.IndikatorPokinOpdTactical, error)
	FindAll(ctx context.Context, tx *sql.Tx) ([]domain.IndikatorPokinOpdTactical, error)
	FindByPokinOpdTacticalId(ctx context.Context, tx *sql.Tx, pokinOpdTacticalId int) ([]domain.IndikatorPokinOpdTactical, error)
}
