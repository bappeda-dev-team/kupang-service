package repository

import (
	"context"
	"database/sql"

	"kupang-service/model/domain"
)

type UserRepository interface {
	Create(ctx context.Context, tx *sql.Tx, user domain.User) (domain.User, error)
	Update(ctx context.Context, tx *sql.Tx, user domain.User) (domain.User, error)
	Delete(ctx context.Context, tx *sql.Tx, id int) error
	FindById(ctx context.Context, tx *sql.Tx, id int) (domain.User, error)
	FindAll(ctx context.Context, tx *sql.Tx) ([]domain.User, error)
	FindByKodeOpd(ctx context.Context, tx *sql.Tx, kodeOpd string) ([]domain.User, error)
}
