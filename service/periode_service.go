package service

import (
	"context"
	"kupang-service/model/web"
)

type PeriodeService interface {
	Create(ctx context.Context, periode web.PeriodeCreateRequest) (web.PeriodeResponse, error)
	Update(ctx context.Context, periode web.PeriodeUpdateRequest) (web.PeriodeResponse, error)
	Delete(ctx context.Context, id int) error
	FindById(ctx context.Context, id int) (web.PeriodeResponse, error)
	FindAll(ctx context.Context) ([]web.PeriodeResponse, error)
}
