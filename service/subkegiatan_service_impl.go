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

type SubkegiatanServiceImpl struct {
	SubkegiatanRepository repository.SubkegiatanRepository
	DB                    *sql.DB
	Validator             *validator.Validate
}

func NewSubkegiatanServiceImpl(subkegiatanRepository repository.SubkegiatanRepository, db *sql.DB, validator *validator.Validate) *SubkegiatanServiceImpl {
	return &SubkegiatanServiceImpl{
		SubkegiatanRepository: subkegiatanRepository,
		DB:                    db,
		Validator:             validator,
	}
}

func (service *SubkegiatanServiceImpl) Create(ctx context.Context, subkegiatan web.SubkegiatanCreateRequest) (web.SubkegiatanResponse, error) {
	err := service.Validator.Struct(subkegiatan)
	if err != nil {
		return web.SubkegiatanResponse{}, err
	}

	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		return web.SubkegiatanResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	subkegiatanDomain := domain.Subkegiatan{
		KodeSubkegiatan: subkegiatan.KodeSubkegiatan,
		NamaSubkegiatan: subkegiatan.NamaSubkegiatan,
		Tahun:           subkegiatan.Tahun,
	}

	subkegiatanDomain, err = service.SubkegiatanRepository.Create(ctx, tx, subkegiatanDomain)
	if err != nil {
		return web.SubkegiatanResponse{}, err
	}

	return helper.ToSubkegiatanResponse(subkegiatanDomain), nil
}

func (service *SubkegiatanServiceImpl) Update(ctx context.Context, subkegiatanData web.SubkegiatanUpdateRequest) (web.SubkegiatanResponse, error) {
	err := service.Validator.Struct(subkegiatanData)
	if err != nil {
		return web.SubkegiatanResponse{}, err
	}

	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		return web.SubkegiatanResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	subkegiatanDomain := domain.Subkegiatan{
		Id:              subkegiatanData.Id,
		KodeSubkegiatan: subkegiatanData.KodeSubkegiatan,
		NamaSubkegiatan: subkegiatanData.NamaSubkegiatan,
		Tahun:           subkegiatanData.Tahun,
	}

	subkegiatanDomain, err = service.SubkegiatanRepository.Update(ctx, tx, subkegiatanDomain)
	if err != nil {
		return web.SubkegiatanResponse{}, err
	}

	return helper.ToSubkegiatanResponse(subkegiatanDomain), nil
}

func (service *SubkegiatanServiceImpl) Delete(ctx context.Context, id int) error {
	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer helper.CommitOrRollback(tx)

	err = service.SubkegiatanRepository.Delete(ctx, tx, id)
	if err != nil {
		return err
	}

	return nil
}

func (service *SubkegiatanServiceImpl) FindById(ctx context.Context, id int) (web.SubkegiatanResponse, error) {
	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		return web.SubkegiatanResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	subkegiatanDomain, err := service.SubkegiatanRepository.FindById(ctx, tx, id)
	if err != nil {
		return web.SubkegiatanResponse{}, errors.New("id tidak ditemukan")
	}

	return helper.ToSubkegiatanResponse(subkegiatanDomain), nil
}

func (service *SubkegiatanServiceImpl) FindAll(ctx context.Context) ([]web.SubkegiatanResponse, error) {
	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		return []web.SubkegiatanResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	subkegiatanDomains, err := service.SubkegiatanRepository.FindAll(ctx, tx)
	if err != nil {
		return []web.SubkegiatanResponse{}, err
	}

	return helper.ToSubkegiatanResponses(subkegiatanDomains), nil
}
