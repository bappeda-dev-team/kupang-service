package service

import (
	"context"
	"kupang-service/model/web"
)

type PokokPikiranService interface {
	Create(ctx context.Context, pokokPikiran web.PokokPikiranCreateRequest) (web.PokokPikiranResponse, error)
	Update(ctx context.Context, pokokPikiran web.PokokPikiranUpdateRequest) (web.PokokPikiranResponse, error)
	Delete(ctx context.Context, id int) error
	FindById(ctx context.Context, id int) (web.PokokPikiranResponse, error)
	FindAll(ctx context.Context) ([]web.PokokPikiranResponse, error)
	FindByKodeOpd(ctx context.Context, kodeOpd string) ([]web.PokokPikiranResponse, error)
}
