package service

import (
	"context"
	"kupang-service/model/web"
)

type PokinOpdStrategicService interface {
	Create(ctx context.Context, request web.PokinOpdStrategicCreateRequest) (web.PokinOpdStrategicResponse, error)
	Update(ctx context.Context, request web.PokinOpdStrategicUpdateRequest) (web.PokinOpdStrategicResponse, error)
	Delete(ctx context.Context, id int) error
	FindById(ctx context.Context, id int) (web.PokinOpdStrategicResponse, error)
	FindAll(ctx context.Context) ([]web.PokinOpdStrategicResponse, error)
}
