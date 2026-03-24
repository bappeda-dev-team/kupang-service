package service

import (
	"context"
	"kupang-service/model/web"
)

type PokinOpdOperationalService interface {
	Create(ctx context.Context, request web.PokinOpdOperationalCreateRequest) (web.PokinOpdOperationalResponse, error)
	Update(ctx context.Context, request web.PokinOpdOperationalUpdateRequest) (web.PokinOpdOperationalResponse, error)
	Delete(ctx context.Context, id int) error
	FindById(ctx context.Context, id int) (web.PokinOpdOperationalResponse, error)
	FindAll(ctx context.Context) ([]web.PokinOpdOperationalResponse, error)
}
