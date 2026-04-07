package service

import (
	"context"
	"kupang-service/model/web"
)

type PegawaiService interface {
	Create(ctx context.Context, pegawai web.PegawaiCreateRequest) (web.PegawaiResponse, error)
	Update(ctx context.Context, pegawai web.PegawaiUpdateRequest) (web.PegawaiResponse, error)
	AddJabatan(ctx context.Context, request web.PegawaiAddJabatanRequest) (web.PegawaiResponse, error)
	UpdateJabatan(ctx context.Context, request web.PegawaiUpdateJabatanRequest) (web.PegawaiResponse, error)
	Delete(ctx context.Context, id int) error
	FindById(ctx context.Context, id int) (web.PegawaiResponse, error)
	FindAll(ctx context.Context) ([]web.PegawaiResponse, error)
	FindByKodeOpd(ctx context.Context, kodeOpd string) ([]web.PegawaiResponse, error)
}
