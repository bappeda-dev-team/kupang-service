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

type PokokPikiranServiceImpl struct {
	PokokPikiranRepository repository.PokokPikiranRepository
	DB                     *sql.DB
	Validator              *validator.Validate
}

func NewPokokPikiranServiceImpl(pokokPikiranRepository repository.PokokPikiranRepository, db *sql.DB, validator *validator.Validate) *PokokPikiranServiceImpl {
	return &PokokPikiranServiceImpl{
		PokokPikiranRepository: pokokPikiranRepository,
		DB:                     db,
		Validator:              validator,
	}
}

func (service *PokokPikiranServiceImpl) Create(ctx context.Context, pokokPikiran web.PokokPikiranCreateRequest) (web.PokokPikiranResponse, error) {
	err := service.Validator.Struct(pokokPikiran)
	if err != nil {
		return web.PokokPikiranResponse{}, err
	}

	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		return web.PokokPikiranResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	pokokPikiranDomain := domain.PokokPikiran{
		Usulan:  pokokPikiran.Usulan,
		Alamat:  pokokPikiran.Alamat,
		Uraian:  pokokPikiran.Uraian,
		Tahun:   toNullString(pokokPikiran.Tahun),
		KodeOpd: pokokPikiran.KodeOpd,
		NamaOpd: pokokPikiran.NamaOpd,
		Status:  toNullString(pokokPikiran.Status),
	}

	pokokPikiranDomain, err = service.PokokPikiranRepository.Create(ctx, tx, pokokPikiranDomain)
	if err != nil {
		return web.PokokPikiranResponse{}, err
	}

	return helper.ToPokokPikiranResponse(pokokPikiranDomain), nil
}

func (service *PokokPikiranServiceImpl) Update(ctx context.Context, pokokPikiran web.PokokPikiranUpdateRequest) (web.PokokPikiranResponse, error) {
	err := service.Validator.Struct(pokokPikiran)
	if err != nil {
		return web.PokokPikiranResponse{}, err
	}

	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		return web.PokokPikiranResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	pokokPikiranDomain := domain.PokokPikiran{
		Id:      pokokPikiran.Id,
		Usulan:  pokokPikiran.Usulan,
		Alamat:  pokokPikiran.Alamat,
		Uraian:  pokokPikiran.Uraian,
		Tahun:   toNullString(pokokPikiran.Tahun),
		KodeOpd: pokokPikiran.KodeOpd,
		NamaOpd: pokokPikiran.NamaOpd,
		Status:  toNullString(pokokPikiran.Status),
	}

	pokokPikiranDomain, err = service.PokokPikiranRepository.Update(ctx, tx, pokokPikiranDomain)
	if err != nil {
		return web.PokokPikiranResponse{}, err
	}

	return helper.ToPokokPikiranResponse(pokokPikiranDomain), nil
}

func (service *PokokPikiranServiceImpl) Delete(ctx context.Context, id int) error {
	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer helper.CommitOrRollback(tx)

	return service.PokokPikiranRepository.Delete(ctx, tx, id)
}

func (service *PokokPikiranServiceImpl) FindById(ctx context.Context, id int) (web.PokokPikiranResponse, error) {
	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		return web.PokokPikiranResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	pokokPikiranDomain, err := service.PokokPikiranRepository.FindById(ctx, tx, id)
	if err != nil {
		return web.PokokPikiranResponse{}, err
	}

	return helper.ToPokokPikiranResponse(pokokPikiranDomain), nil
}

func (service *PokokPikiranServiceImpl) FindAll(ctx context.Context) ([]web.PokokPikiranResponse, error) {
	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		return []web.PokokPikiranResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	pokokPikiranDomains, err := service.PokokPikiranRepository.FindAll(ctx, tx)
	if err != nil {
		return []web.PokokPikiranResponse{}, err
	}

	return helper.ToPokokPikiranResponses(pokokPikiranDomains), nil
}

func (service *PokokPikiranServiceImpl) FindByKodeOpd(ctx context.Context, kodeOpd string) ([]web.PokokPikiranResponse, error) {
	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		return []web.PokokPikiranResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	pokokPikiranDomains, err := service.PokokPikiranRepository.FindByKodeOpd(ctx, tx, kodeOpd)
	if err != nil {
		return []web.PokokPikiranResponse{}, err
	}

	return helper.ToPokokPikiranResponses(pokokPikiranDomains), nil
}
