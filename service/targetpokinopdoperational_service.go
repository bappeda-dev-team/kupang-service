package service

import (
	"context"
	"kupang-service/model/web"
)

type TargetPokinOpdOperationalService interface {
	Create(ctx context.Context, request web.TargetPokinOpdOperationalCreateRequest) (web.TargetPokinOpdOperationalResponse, error)
	Update(ctx context.Context, request web.TargetPokinOpdOperationalUpdateRequest) (web.TargetPokinOpdOperationalResponse, error)
	Delete(ctx context.Context, id int) error
	FindById(ctx context.Context, id int) (web.TargetPokinOpdOperationalResponse, error)
	FindAll(ctx context.Context) ([]web.TargetPokinOpdOperationalResponse, error)
}
