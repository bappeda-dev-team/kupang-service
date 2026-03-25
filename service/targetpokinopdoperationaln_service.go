package service

import (
	"context"
	"kupang-service/model/web"
)

type TargetPokinOpdOperationalNService interface {
	Create(ctx context.Context, request web.TargetPokinOpdOperationalNCreateRequest) (web.TargetPokinOpdOperationalNResponse, error)
	Update(ctx context.Context, request web.TargetPokinOpdOperationalNUpdateRequest) (web.TargetPokinOpdOperationalNResponse, error)
	Delete(ctx context.Context, id int) error
	FindById(ctx context.Context, id int) (web.TargetPokinOpdOperationalNResponse, error)
	FindAll(ctx context.Context) ([]web.TargetPokinOpdOperationalNResponse, error)
}
