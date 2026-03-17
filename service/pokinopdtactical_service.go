package service

import (
	"context"
	"kupang-service/model/web"
)

type PokinOpdTacticalService interface {
	Create(ctx context.Context, request web.PokinOpdTacticalCreateRequest) (web.PokinOpdTacticalResponse, error)
	Update(ctx context.Context, request web.PokinOpdTacticalUpdateRequest) (web.PokinOpdTacticalResponse, error)
	Delete(ctx context.Context, id int) error
	FindById(ctx context.Context, id int) (web.PokinOpdTacticalResponse, error)
	FindAll(ctx context.Context) ([]web.PokinOpdTacticalResponse, error)
}
