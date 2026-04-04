package repository

import (
	"context"
	"database/sql"
	"kupang-service/model/domain"
)

type PemdaRepository interface {
	Create(ctx context.Context, tx *sql.Tx, pemda domain.Pemda) (domain.Pemda, error)
	Update(ctx context.Context, tx *sql.Tx, pemda domain.Pemda) (domain.Pemda, error)
	Delete(ctx context.Context, tx *sql.Tx, id int) error
	FindById(ctx context.Context, tx *sql.Tx, id int) (domain.Pemda, error)
	FindAll(ctx context.Context, tx *sql.Tx) ([]domain.Pemda, error)
}
