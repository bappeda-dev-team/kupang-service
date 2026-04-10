package service

import (
	"context"
	"kupang-service/model/web"
)

type MusrenbangService interface {
	Create(ctx context.Context, musrenbang web.MusrenbangCreateRequest) (web.MusrenbangResponse, error)
	Update(ctx context.Context, musrenbang web.MusrenbangUpdateRequest) (web.MusrenbangResponse, error)
	Delete(ctx context.Context, id int) error
	FindById(ctx context.Context, id int) (web.MusrenbangResponse, error)
	FindAll(ctx context.Context) ([]web.MusrenbangResponse, error)
	FindByKodeOpd(ctx context.Context, kodeOpd string) ([]web.MusrenbangResponse, error)
}
