package service

import (
	"context"
	"kupang-service/model/web"
)

type IndikatorPokinOpdService interface {
	Create(ctx context.Context, indikatorPokinOpd web.IndikatorPokinOpdCreateRequest) (web.IndikatorPokinOpdResponse, error)
	Update(ctx context.Context, indikatorPokinOpd web.IndikatorPokinOpdUpdateRequest) (web.IndikatorPokinOpdResponse, error)
	Delete(ctx context.Context, id int) error
	FindById(ctx context.Context, id int) (web.IndikatorPokinOpdResponse, error)
	FindAll(ctx context.Context) ([]web.IndikatorPokinOpdResponse, error)
}
