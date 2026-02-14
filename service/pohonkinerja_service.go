package service

import (
	"context"
	"kupang-service/model/web"
)

type PohonKinerjaService interface {
	FindByKodeOpdAndTahun(ctx context.Context, kodeOpd string, tahun int) (web.PohonKinerjaResponse, error)
}
