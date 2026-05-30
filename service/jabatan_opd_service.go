package service

import (
	"context"
	"kupang-service/model/web"
)

type JabatanOpdService interface {
	Create(ctx context.Context, jabatanOpd web.JabatanOpdCreateRequest) (web.JabatanOpdResponse, error)
	Update(ctx context.Context, jabatanOpd web.JabatanOpdUpdateRequest) (web.JabatanOpdResponse, error)
	Delete(ctx context.Context, id int) error
	FindById(ctx context.Context, id int) (web.JabatanOpdResponse, error)
	FindByKodeOpd(ctx context.Context, kodeOpd string) ([]web.JabatanOpdResponse, error)
}
