package service

import (
	"context"
	"database/sql"
	"errors"
	"kupang-service/helper"
	"kupang-service/model/domain"
	"kupang-service/model/web"
	"kupang-service/repository"

	"github.com/go-playground/validator/v10"
)

type BidangUrusanServiceImpl struct {
	BidangUrusanRepository repository.BidangUrusanRepository
	DB                     *sql.DB
	Validator              *validator.Validate
}

func NewBidangUrusanServiceImpl(bidangUrusanRepository repository.BidangUrusanRepository, db *sql.DB, validator *validator.Validate) *BidangUrusanServiceImpl {
	return &BidangUrusanServiceImpl{
		BidangUrusanRepository: bidangUrusanRepository,
		DB:                     db,
		Validator:              validator,
	}
}

func (service *BidangUrusanServiceImpl) Create(ctx context.Context, bidangUrusan web.BidangUrusanCreateRequest) (web.BidangUrusanResponse, error) {
	err := service.Validator.Struct(bidangUrusan)
	if err != nil {
		return web.BidangUrusanResponse{}, err
	}

	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		return web.BidangUrusanResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	bidangUrusanDomain := domain.BidangUrusan{
		KodeUrusan: bidangUrusan.KodeUrusan,
		NamaUrusan: bidangUrusan.NamaUrusan,
	}

	bidangUrusanDomain, err = service.BidangUrusanRepository.Create(ctx, tx, bidangUrusanDomain)
	if err != nil {
		return web.BidangUrusanResponse{}, err
	}

	return web.BidangUrusanResponse{
		Id:         bidangUrusanDomain.Id,
		KodeUrusan: bidangUrusanDomain.KodeUrusan,
		NamaUrusan: bidangUrusanDomain.NamaUrusan,
	}, nil
}

func (service *BidangUrusanServiceImpl) Update(ctx context.Context, bidangUrusanData web.BidangUrusanUpdateRequest) (web.BidangUrusanResponse, error) {
	err := service.Validator.Struct(bidangUrusanData)
	if err != nil {
		return web.BidangUrusanResponse{}, err
	}

	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		return web.BidangUrusanResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	bidangUrusanDomain := domain.BidangUrusan{
		Id:         bidangUrusanData.Id,
		KodeUrusan: bidangUrusanData.KodeUrusan,
		NamaUrusan: bidangUrusanData.NamaUrusan,
	}

	bidangUrusanDomain, err = service.BidangUrusanRepository.Update(ctx, tx, bidangUrusanDomain)
	if err != nil {
		return web.BidangUrusanResponse{}, err
	}

	return web.BidangUrusanResponse{
		Id:         bidangUrusanDomain.Id,
		KodeUrusan: bidangUrusanDomain.KodeUrusan,
		NamaUrusan: bidangUrusanDomain.NamaUrusan,
	}, nil
}

func (service *BidangUrusanServiceImpl) Delete(ctx context.Context, id int) error {
	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer helper.CommitOrRollback(tx)

	err = service.BidangUrusanRepository.Delete(ctx, tx, id)
	if err != nil {
		return err
	}

	return nil
}

func (service *BidangUrusanServiceImpl) FindById(ctx context.Context, id int) (web.BidangUrusanResponse, error) {
	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		return web.BidangUrusanResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	bidangUrusanDomain, err := service.BidangUrusanRepository.FindById(ctx, tx, id)
	if err != nil {
		return web.BidangUrusanResponse{}, errors.New("id tidak ditemukan")
	}

	return web.BidangUrusanResponse{
		Id:         bidangUrusanDomain.Id,
		KodeUrusan: bidangUrusanDomain.KodeUrusan,
		NamaUrusan: bidangUrusanDomain.NamaUrusan,
	}, nil
}

func (service *BidangUrusanServiceImpl) FindAll(ctx context.Context) ([]web.BidangUrusanResponse, error) {
	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		return []web.BidangUrusanResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	bidangUrusanDomains, err := service.BidangUrusanRepository.FindAll(ctx, tx)
	if err != nil {
		return []web.BidangUrusanResponse{}, err
	}

	return helper.ToBidangUrusanResponses(bidangUrusanDomains), nil
}
