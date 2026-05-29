package service

import (
	"context"
	"kupang-service/model/web"
)

type KegiatanService interface {
	Create(ctx context.Context, kegiatan web.KegiatanCreateRequest) (web.KegiatanResponse, error)
	Update(ctx context.Context, kegiatan web.KegiatanUpdateRequest) (web.KegiatanResponse, error)
	Delete(ctx context.Context, id int) error
	FindById(ctx context.Context, id int) (web.KegiatanResponse, error)
	FindAll(ctx context.Context) ([]web.KegiatanResponse, error)
}
