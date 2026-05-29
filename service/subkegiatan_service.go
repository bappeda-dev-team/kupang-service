package service

import (
	"context"
	"kupang-service/model/web"
)

type SubkegiatanService interface {
	Create(ctx context.Context, subkegiatan web.SubkegiatanCreateRequest) (web.SubkegiatanResponse, error)
	Update(ctx context.Context, subkegiatan web.SubkegiatanUpdateRequest) (web.SubkegiatanResponse, error)
	Delete(ctx context.Context, id int) error
	FindById(ctx context.Context, id int) (web.SubkegiatanResponse, error)
	FindAll(ctx context.Context) ([]web.SubkegiatanResponse, error)
}
