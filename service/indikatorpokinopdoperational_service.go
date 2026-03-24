package service

import (
	"context"
	"kupang-service/model/web"
)

type IndikatorPokinOpdOperationalService interface {
	Create(ctx context.Context, indikator web.IndikatorPokinOpdOperationalCreateRequest) (web.IndikatorPokinOpdOperationalResponse, error)
	Update(ctx context.Context, indikator web.IndikatorPokinOpdOperationalUpdateRequest) (web.IndikatorPokinOpdOperationalResponse, error)
	Delete(ctx context.Context, id int) error
	FindById(ctx context.Context, id int) (web.IndikatorPokinOpdOperationalResponse, error)
	FindAll(ctx context.Context) ([]web.IndikatorPokinOpdOperationalResponse, error)
}
