package service

import (
	"context"
	"kupang-service/model/web"
)

type OpdService interface {
	Create(ctx context.Context, opd web.OpdCreateRequest) (web.OpdResponse, error)
	Update(ctx context.Context, opd web.OpdUpdateRequest) (web.OpdResponse, error)
	Delete(ctx context.Context, id int) error
	FindById(ctx context.Context, id int) (web.OpdResponse, error)
	FindAll(ctx context.Context) ([]web.OpdResponse, error)
}
