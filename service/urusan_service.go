package service

import (
	"context"
	"kupang-service/model/web"
)

type UrusanService interface {
	Create(ctx context.Context, urusan web.UrusanCreateRequest) (web.UrusanResponse, error)
	Update(ctx context.Context, urusan web.UrusanUpdateRequest) (web.UrusanResponse, error)
	Delete(ctx context.Context, id int) error
	FindById(ctx context.Context, id int) (web.UrusanResponse, error)
	FindAll(ctx context.Context) ([]web.UrusanResponse, error)
}
