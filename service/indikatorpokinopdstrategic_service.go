package service

import (
	"context"
	"kupang-service/model/web"
)

type IndikatorPokinOpdStrategicService interface {
	Create(ctx context.Context, indikator web.IndikatorPokinOpdStrategicCreateRequest) (web.IndikatorPokinOpdStrategicResponse, error)
	Update(ctx context.Context, indikator web.IndikatorPokinOpdStrategicUpdateRequest) (web.IndikatorPokinOpdStrategicResponse, error)
	Delete(ctx context.Context, id int) error
	FindById(ctx context.Context, id int) (web.IndikatorPokinOpdStrategicResponse, error)
	FindAll(ctx context.Context) ([]web.IndikatorPokinOpdStrategicResponse, error)
}
