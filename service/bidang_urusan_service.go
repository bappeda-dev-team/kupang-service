package service

import (
	"context"
	"kupang-service/model/web"
)

type BidangUrusanService interface {
	Create(ctx context.Context, bidangUrusan web.BidangUrusanCreateRequest) (web.BidangUrusanResponse, error)
	Update(ctx context.Context, bidangUrusan web.BidangUrusanUpdateRequest) (web.BidangUrusanResponse, error)
	Delete(ctx context.Context, id int) error
	FindById(ctx context.Context, id int) (web.BidangUrusanResponse, error)
	FindAll(ctx context.Context) ([]web.BidangUrusanResponse, error)
}
