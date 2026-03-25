package service

import (
	"context"
	"kupang-service/model/web"
)

type PokinOpdOperationalNService interface {
	Create(ctx context.Context, request web.PokinOpdOperationalNCreateRequest) (web.PokinOpdOperationalNResponse, error)
	Update(ctx context.Context, request web.PokinOpdOperationalNUpdateRequest) (web.PokinOpdOperationalNResponse, error)
	Delete(ctx context.Context, id int) error
	FindById(ctx context.Context, id int) (web.PokinOpdOperationalNResponse, error)
	FindAll(ctx context.Context) ([]web.PokinOpdOperationalNResponse, error)
}
