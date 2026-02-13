package service

import (
	"context"
	"kupang-service/model/web"
)

type PokinOpdService interface {
	Create(ctx context.Context, pokinOpd web.PokinOpdCreateRequest) (web.PokinOpdResponse, error)
	Update(ctx context.Context, pokinOpd web.PokinOpdUpdateRequest) (web.PokinOpdResponse, error)
	Delete(ctx context.Context, id int) error
	FindById(ctx context.Context, id int) (web.PokinOpdResponse, error)
	FindAll(ctx context.Context) ([]web.PokinOpdResponse, error)
}
