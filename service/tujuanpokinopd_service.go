package service

import (
	"context"
	"kupang-service/model/web"
)

type TujuanPokinOpdService interface {
	Create(ctx context.Context, tujuanPokinOpd web.TujuanPokinOpdCreateRequest) (web.TujuanPokinOpdResponse, error)
	Update(ctx context.Context, tujuanPokinOpd web.TujuanPokinOpdUpdateRequest) (web.TujuanPokinOpdResponse, error)
	Delete(ctx context.Context, id int) error
	FindById(ctx context.Context, id int) (web.TujuanPokinOpdResponse, error)
	FindAll(ctx context.Context) ([]web.TujuanPokinOpdResponse, error)
}
