package repository

import (
	"context"
	"database/sql"
	"kupang-service/model/domain"
)

type BidangUrusanRepository interface {
	Create(ctx context.Context, tx *sql.Tx, bidangUrusan domain.BidangUrusan) (domain.BidangUrusan, error)
	Update(ctx context.Context, tx *sql.Tx, bidangUrusan domain.BidangUrusan) (domain.BidangUrusan, error)
	Delete(ctx context.Context, tx *sql.Tx, id int) error
	FindById(ctx context.Context, tx *sql.Tx, id int) (domain.BidangUrusan, error)
	FindAll(ctx context.Context, tx *sql.Tx) ([]domain.BidangUrusan, error)
}
