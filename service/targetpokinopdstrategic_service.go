package service

import (
	"context"
	"kupang-service/model/web"
)

type TargetPokinOpdStrategicService interface {
	Create(ctx context.Context, request web.TargetPokinOpdStrategicCreateRequest) (web.TargetPokinOpdStrategicResponse, error)
	Update(ctx context.Context, request web.TargetPokinOpdStrategicUpdateRequest) (web.TargetPokinOpdStrategicResponse, error)
	Delete(ctx context.Context, id int) error
	FindById(ctx context.Context, id int) (web.TargetPokinOpdStrategicResponse, error)
	FindAll(ctx context.Context) ([]web.TargetPokinOpdStrategicResponse, error)
}
