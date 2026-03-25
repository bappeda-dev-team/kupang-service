package service

import (
	"context"
	"kupang-service/model/web"
)

type IndikatorPokinOpdOperationalNService interface {
	Create(ctx context.Context, indikator web.IndikatorPokinOpdOperationalNCreateRequest) (web.IndikatorPokinOpdOperationalNResponse, error)
	Update(ctx context.Context, indikator web.IndikatorPokinOpdOperationalNUpdateRequest) (web.IndikatorPokinOpdOperationalNResponse, error)
	Delete(ctx context.Context, id int) error
	FindById(ctx context.Context, id int) (web.IndikatorPokinOpdOperationalNResponse, error)
	FindAll(ctx context.Context) ([]web.IndikatorPokinOpdOperationalNResponse, error)
}