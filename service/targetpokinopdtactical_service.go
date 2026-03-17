package service

import (
	"context"
	"kupang-service/model/web"
)

type TargetPokinOpdTacticalService interface {
	Create(ctx context.Context, request web.TargetPokinOpdTacticalCreateRequest) (web.TargetPokinOpdTacticalResponse, error)
	Update(ctx context.Context, request web.TargetPokinOpdTacticalUpdateRequest) (web.TargetPokinOpdTacticalResponse, error)
	Delete(ctx context.Context, id int) error
	FindById(ctx context.Context, id int) (web.TargetPokinOpdTacticalResponse, error)
	FindAll(ctx context.Context) ([]web.TargetPokinOpdTacticalResponse, error)
}
