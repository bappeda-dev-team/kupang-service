package service

import (
	"context"
	"kupang-service/model/web"
)

type TargetPokinOpdService interface {
	Create(ctx context.Context, targetPokinOpd web.TargetPokinOpdCreateRequest) (web.TargetPokinOpdResponse, error)
	Update(ctx context.Context, targetPokinOpd web.TargetPokinOpdUpdateRequest) (web.TargetPokinOpdResponse, error)
	Delete(ctx context.Context, id int) error
	FindById(ctx context.Context, id int) (web.TargetPokinOpdResponse, error)
	FindAll(ctx context.Context) ([]web.TargetPokinOpdResponse, error)
}
