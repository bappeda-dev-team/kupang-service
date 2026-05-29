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

type UrusanServiceImpl struct {
	UrusanRepository repository.UrusanRepository
	DB               *sql.DB
	Validator        *validator.Validate
}

func NewUrusanServiceImpl(urusanRepository repository.UrusanRepository, db *sql.DB, validator *validator.Validate) *UrusanServiceImpl {
	return &UrusanServiceImpl{
		UrusanRepository: urusanRepository,
		DB:               db,
		Validator:        validator,
	}
}

func (service *UrusanServiceImpl) Create(ctx context.Context, urusan web.UrusanCreateRequest) (web.UrusanResponse, error) {
	err := service.Validator.Struct(urusan)
	if err != nil {
		return web.UrusanResponse{}, err
	}
	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		return web.UrusanResponse{}, err
	}
	defer helper.CommitOrRollback(tx)
	urusanDomain := domain.Urusan{
		KodeUrusan: urusan.KodeUrusan,
		NamaUrusan: urusan.NamaUrusan,
	}
	urusanDomain, err = service.UrusanRepository.Create(ctx, tx, urusanDomain)
	if err != nil {
		return web.UrusanResponse{}, err
	}
	return helper.ToUrusanResponse(urusanDomain), nil
}

func (service *UrusanServiceImpl) Update(ctx context.Context, urusanData web.UrusanUpdateRequest) (web.UrusanResponse, error) {
	err := service.Validator.Struct(urusanData)
	if err != nil {
		return web.UrusanResponse{}, err
	}
	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		return web.UrusanResponse{}, err
	}
	defer helper.CommitOrRollback(tx)
	urusanDomain := domain.Urusan{
		Id:         urusanData.Id,
		KodeUrusan: urusanData.KodeUrusan,
		NamaUrusan: urusanData.NamaUrusan,
	}
	urusanDomain, err = service.UrusanRepository.Update(ctx, tx, urusanDomain)
	if err != nil {
		return web.UrusanResponse{}, err
	}
	return helper.ToUrusanResponse(urusanDomain), nil
}

func (service *UrusanServiceImpl) Delete(ctx context.Context, id int) error {
	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer helper.CommitOrRollback(tx)
	err = service.UrusanRepository.Delete(ctx, tx, id)
	if err != nil {
		return err
	}
	return nil
}

func (service *UrusanServiceImpl) FindById(ctx context.Context, id int) (web.UrusanResponse, error) {
	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		return web.UrusanResponse{}, err
	}
	defer helper.CommitOrRollback(tx)
	urusanDomain, err := service.UrusanRepository.FindById(ctx, tx, id)
	if err != nil {
		return web.UrusanResponse{}, errors.New("id tidak ditemukan")
	}
	return helper.ToUrusanResponse(urusanDomain), nil
}

func (service *UrusanServiceImpl) FindAll(ctx context.Context) ([]web.UrusanResponse, error) {
	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		return []web.UrusanResponse{}, err
	}
	defer helper.CommitOrRollback(tx)
	urusanDomains, err := service.UrusanRepository.FindAll(ctx, tx)
	if err != nil {
		return []web.UrusanResponse{}, err
	}
	return helper.ToUrusanResponses(urusanDomains), nil
}
