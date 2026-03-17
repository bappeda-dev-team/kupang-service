package service

import (
	"context"
	"kupang-service/model/web"
)

type IndikatorPokinOpdTacticalService interface {
	Create(ctx context.Context, indikator web.IndikatorPokinOpdTacticalCreateRequest) (web.IndikatorPokinOpdTacticalResponse, error)
	Update(ctx context.Context, indikator web.IndikatorPokinOpdTacticalUpdateRequest) (web.IndikatorPokinOpdTacticalResponse, error)
	Delete(ctx context.Context, id int) error
	FindById(ctx context.Context, id int) (web.IndikatorPokinOpdTacticalResponse, error)
	FindAll(ctx context.Context) ([]web.IndikatorPokinOpdTacticalResponse, error)
}
