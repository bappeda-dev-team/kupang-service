package repository

import (
	"context"
	"database/sql"
	"kupang-service/model/domain"
)

type JabatanOpdRepository interface {
	Create(ctx context.Context, tx *sql.Tx, jabatanOpd domain.JabatanOpd) (domain.JabatanOpd, error)
	Update(ctx context.Context, tx *sql.Tx, jabatanOpd domain.JabatanOpd) (domain.JabatanOpd, error)
	Delete(ctx context.Context, tx *sql.Tx, id int) error
	FindById(ctx context.Context, tx *sql.Tx, id int) (domain.JabatanOpd, error)
	FindByKodeOpd(ctx context.Context, tx *sql.Tx, kodeOpd string) ([]domain.JabatanOpd, error)
}
