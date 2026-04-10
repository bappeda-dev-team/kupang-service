package service

import (
	"context"
	"database/sql"
	"kupang-service/helper"
	"kupang-service/model/domain"
	"kupang-service/model/web"
	"kupang-service/repository"

	"github.com/go-playground/validator/v10"
)

type MusrenbangServiceImpl struct {
	MusrenbangRepository repository.MusrenbangRepository
	DB                   *sql.DB
	Validator            *validator.Validate
}

func NewMusrenbangServiceImpl(musrenbangRepository repository.MusrenbangRepository, db *sql.DB, validator *validator.Validate) *MusrenbangServiceImpl {
	return &MusrenbangServiceImpl{
		MusrenbangRepository: musrenbangRepository,
		DB:                   db,
		Validator:            validator,
	}
}

func (service *MusrenbangServiceImpl) Create(ctx context.Context, musrenbang web.MusrenbangCreateRequest) (web.MusrenbangResponse, error) {
	err := service.Validator.Struct(musrenbang)
	if err != nil {
		return web.MusrenbangResponse{}, err
	}

	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		return web.MusrenbangResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	musrenbangDomain := domain.Musrenbang{
		Usulan:  musrenbang.Usulan,
		Alamat:  musrenbang.Alamat,
		Uraian:  musrenbang.Uraian,
		Tahun:   toNullString(musrenbang.Tahun),
		KodeOpd: musrenbang.KodeOpd,
		NamaOpd: musrenbang.NamaOpd,
		Status:  toNullString(musrenbang.Status),
	}

	musrenbangDomain, err = service.MusrenbangRepository.Create(ctx, tx, musrenbangDomain)
	if err != nil {
		return web.MusrenbangResponse{}, err
	}

	return helper.ToMusrenbangResponse(musrenbangDomain), nil
}

func toNullString(value *string) sql.NullString {
	if value == nil {
		return sql.NullString{}
	}
	return sql.NullString{
		String: *value,
		Valid:  true,
	}
}

func (service *MusrenbangServiceImpl) Update(ctx context.Context, musrenbang web.MusrenbangUpdateRequest) (web.MusrenbangResponse, error) {
	err := service.Validator.Struct(musrenbang)
	if err != nil {
		return web.MusrenbangResponse{}, err
	}

	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		return web.MusrenbangResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	musrenbangDomain := domain.Musrenbang{
		Id:      musrenbang.Id,
		Usulan:  musrenbang.Usulan,
		Alamat:  musrenbang.Alamat,
		Uraian:  musrenbang.Uraian,
		Tahun:   toNullString(musrenbang.Tahun),
		KodeOpd: musrenbang.KodeOpd,
		NamaOpd: musrenbang.NamaOpd,
		Status:  toNullString(musrenbang.Status),
	}

	musrenbangDomain, err = service.MusrenbangRepository.Update(ctx, tx, musrenbangDomain)
	if err != nil {
		return web.MusrenbangResponse{}, err
	}

	return helper.ToMusrenbangResponse(musrenbangDomain), nil
}

func (service *MusrenbangServiceImpl) Delete(ctx context.Context, id int) error {
	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer helper.CommitOrRollback(tx)

	return service.MusrenbangRepository.Delete(ctx, tx, id)
}

func (service *MusrenbangServiceImpl) FindById(ctx context.Context, id int) (web.MusrenbangResponse, error) {
	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		return web.MusrenbangResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	musrenbangDomain, err := service.MusrenbangRepository.FindById(ctx, tx, id)
	if err != nil {
		return web.MusrenbangResponse{}, err
	}

	return helper.ToMusrenbangResponse(musrenbangDomain), nil
}

func (service *MusrenbangServiceImpl) FindAll(ctx context.Context) ([]web.MusrenbangResponse, error) {
	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		return []web.MusrenbangResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	musrenbangDomains, err := service.MusrenbangRepository.FindAll(ctx, tx)
	if err != nil {
		return []web.MusrenbangResponse{}, err
	}

	return helper.ToMusrenbangResponses(musrenbangDomains), nil
}

func (service *MusrenbangServiceImpl) FindByKodeOpd(ctx context.Context, kodeOpd string) ([]web.MusrenbangResponse, error) {
	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		return []web.MusrenbangResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	musrenbangDomains, err := service.MusrenbangRepository.FindByKodeOpd(ctx, tx, kodeOpd)
	if err != nil {
		return []web.MusrenbangResponse{}, err
	}

	return helper.ToMusrenbangResponses(musrenbangDomains), nil
}
