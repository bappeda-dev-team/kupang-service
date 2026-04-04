package service

import (
	"context"
	"kupang-service/model/web"
)

type PemdaService interface {
	Create(ctx context.Context, pemda web.PemdaCreateRequest) (web.PemdaResponse, error)
	Update(ctx context.Context, pemda web.PemdaUpdateRequest) (web.PemdaResponse, error)
	Delete(ctx context.Context, id int) error
	FindById(ctx context.Context, id int) (web.PemdaResponse, error)
	FindAll(ctx context.Context) ([]web.PemdaResponse, error)
}
