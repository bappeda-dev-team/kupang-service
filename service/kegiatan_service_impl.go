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

type KegiatanServiceImpl struct {
	KegiatanRepository repository.KegiatanRepository
	DB                 *sql.DB
	Validator          *validator.Validate
}

func NewKegiatanServiceImpl(kegiatanRepository repository.KegiatanRepository, db *sql.DB, validator *validator.Validate) *KegiatanServiceImpl {
	return &KegiatanServiceImpl{
		KegiatanRepository: kegiatanRepository,
		DB:                 db,
		Validator:          validator,
	}
}

func (service *KegiatanServiceImpl) Create(ctx context.Context, kegiatan web.KegiatanCreateRequest) (web.KegiatanResponse, error) {
	err := service.Validator.Struct(kegiatan)
	if err != nil {
		return web.KegiatanResponse{}, err
	}

	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		return web.KegiatanResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	kegiatanDomain := domain.Kegiatan{
		KodeKegiatan: kegiatan.KodeKegiatan,
		NamaKegiatan: kegiatan.NamaKegiatan,
		Tahun:        kegiatan.Tahun,
		KodeOpd:      kegiatan.KodeOpd,
	}

	kegiatanDomain, err = service.KegiatanRepository.Create(ctx, tx, kegiatanDomain)
	if err != nil {
		return web.KegiatanResponse{}, err
	}

	return helper.ToKegiatanResponse(kegiatanDomain), nil
}

func (service *KegiatanServiceImpl) Update(ctx context.Context, kegiatanData web.KegiatanUpdateRequest) (web.KegiatanResponse, error) {
	err := service.Validator.Struct(kegiatanData)
	if err != nil {
		return web.KegiatanResponse{}, err
	}

	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		return web.KegiatanResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	kegiatanDomain := domain.Kegiatan{
		Id:           kegiatanData.Id,
		KodeKegiatan: kegiatanData.KodeKegiatan,
		NamaKegiatan: kegiatanData.NamaKegiatan,
		Tahun:        kegiatanData.Tahun,
		KodeOpd:      kegiatanData.KodeOpd,
	}

	kegiatanDomain, err = service.KegiatanRepository.Update(ctx, tx, kegiatanDomain)
	if err != nil {
		return web.KegiatanResponse{}, err
	}

	return helper.ToKegiatanResponse(kegiatanDomain), nil
}

func (service *KegiatanServiceImpl) Delete(ctx context.Context, id int) error {
	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer helper.CommitOrRollback(tx)

	err = service.KegiatanRepository.Delete(ctx, tx, id)
	if err != nil {
		return err
	}

	return nil
}

func (service *KegiatanServiceImpl) FindById(ctx context.Context, id int) (web.KegiatanResponse, error) {
	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		return web.KegiatanResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	kegiatanDomain, err := service.KegiatanRepository.FindById(ctx, tx, id)
	if err != nil {
		return web.KegiatanResponse{}, errors.New("id tidak ditemukan")
	}

	return helper.ToKegiatanResponse(kegiatanDomain), nil
}

func (service *KegiatanServiceImpl) FindAll(ctx context.Context) ([]web.KegiatanResponse, error) {
	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		return []web.KegiatanResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	kegiatanDomains, err := service.KegiatanRepository.FindAll(ctx, tx)
	if err != nil {
		return []web.KegiatanResponse{}, err
	}

	return helper.ToKegiatanResponses(kegiatanDomains), nil
}
