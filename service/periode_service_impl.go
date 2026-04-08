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

type PeriodeServiceImpl struct {
	PeriodeRepository repository.PeriodeRepository
	DB                *sql.DB
	Validator         *validator.Validate
}

func NewPeriodeServiceImpl(periodeRepository repository.PeriodeRepository, db *sql.DB, validator *validator.Validate) *PeriodeServiceImpl {
	return &PeriodeServiceImpl{
		PeriodeRepository: periodeRepository,
		DB:                db,
		Validator:         validator,
	}
}

func (service *PeriodeServiceImpl) Create(ctx context.Context, periode web.PeriodeCreateRequest) (web.PeriodeResponse, error) {
	err := service.Validator.Struct(periode)
	if err != nil {
		return web.PeriodeResponse{}, err
	}

	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		return web.PeriodeResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	periodeDomain := domain.Periode{
		TahunAwal:    periode.TahunAwal,
		TahunAkhir:   periode.TahunAkhir,
		JenisPeriode: periode.JenisPeriode,
	}

	periodeDomain, err = service.PeriodeRepository.Create(ctx, tx, periodeDomain)
	if err != nil {
		return web.PeriodeResponse{}, err
	}

	return helper.ToPeriodeResponse(periodeDomain), nil
}

func (service *PeriodeServiceImpl) Update(ctx context.Context, periode web.PeriodeUpdateRequest) (web.PeriodeResponse, error) {
	err := service.Validator.Struct(periode)
	if err != nil {
		return web.PeriodeResponse{}, err
	}

	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		return web.PeriodeResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	periodeDomain := domain.Periode{
		Id:           periode.Id,
		TahunAwal:    periode.TahunAwal,
		TahunAkhir:   periode.TahunAkhir,
		JenisPeriode: periode.JenisPeriode,
	}

	periodeDomain, err = service.PeriodeRepository.Update(ctx, tx, periodeDomain)
	if err != nil {
		return web.PeriodeResponse{}, err
	}

	return helper.ToPeriodeResponse(periodeDomain), nil
}

func (service *PeriodeServiceImpl) Delete(ctx context.Context, id int) error {
	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer helper.CommitOrRollback(tx)

	return service.PeriodeRepository.Delete(ctx, tx, id)
}

func (service *PeriodeServiceImpl) FindById(ctx context.Context, id int) (web.PeriodeResponse, error) {
	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		return web.PeriodeResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	periodeDomain, err := service.PeriodeRepository.FindById(ctx, tx, id)
	if err != nil {
		return web.PeriodeResponse{}, errors.New("id tidak ditemukan")
	}

	return helper.ToPeriodeResponse(periodeDomain), nil
}

func (service *PeriodeServiceImpl) FindAll(ctx context.Context) ([]web.PeriodeResponse, error) {
	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		return []web.PeriodeResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	periodeDomains, err := service.PeriodeRepository.FindAll(ctx, tx)
	if err != nil {
		return []web.PeriodeResponse{}, err
	}

	return helper.ToPeriodeResponses(periodeDomains), nil
}
