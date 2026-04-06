package service

import (
	"context"
	"kupang-service/model/web"
)

type LembagaService interface {
	Create(ctx context.Context, lembaga web.LembagaCreateRequest) (web.LembagaResponse, error)
	Update(ctx context.Context, lembaga web.LembagaUpdateRequest) (web.LembagaResponse, error)
	Delete(ctx context.Context, id int) error
	FindById(ctx context.Context, id int) (web.LembagaResponse, error)
	FindAll(ctx context.Context) ([]web.LembagaResponse, error)
}
