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

type LembagaServiceImpl struct {
	LembagaRepository repository.LembagaRepository
	DB                *sql.DB
	Validator         *validator.Validate
}

func NewLembagaServiceImpl(lembagaRepository repository.LembagaRepository, db *sql.DB, validator *validator.Validate) *LembagaServiceImpl {
	return &LembagaServiceImpl{
		LembagaRepository: lembagaRepository,
		DB:                db,
		Validator:         validator,
	}
}

func (service *LembagaServiceImpl) Create(ctx context.Context, lembaga web.LembagaCreateRequest) (web.LembagaResponse, error) {
	err := service.Validator.Struct(lembaga)
	if err != nil {
		return web.LembagaResponse{}, err
	}

	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		return web.LembagaResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	lembagaDomain := domain.Lembaga{
		KodeLembaga: lembaga.KodeLembaga,
		NamaLembaga: lembaga.NamaLembaga,
	}

	lembagaDomain, err = service.LembagaRepository.Create(ctx, tx, lembagaDomain)
	if err != nil {
		return web.LembagaResponse{}, err
	}

	return web.LembagaResponse{
		Id:          lembagaDomain.Id,
		KodeLembaga: lembagaDomain.KodeLembaga,
		NamaLembaga: lembagaDomain.NamaLembaga,
	}, nil
}

func (service *LembagaServiceImpl) Update(ctx context.Context, lembagaData web.LembagaUpdateRequest) (web.LembagaResponse, error) {
	err := service.Validator.Struct(lembagaData)
	if err != nil {
		return web.LembagaResponse{}, err
	}

	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		return web.LembagaResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	lembagaDomain := domain.Lembaga{
		Id:          lembagaData.Id,
		KodeLembaga: lembagaData.KodeLembaga,
		NamaLembaga: lembagaData.NamaLembaga,
	}

	lembagaDomain, err = service.LembagaRepository.Update(ctx, tx, lembagaDomain)
	if err != nil {
		return web.LembagaResponse{}, err
	}

	return web.LembagaResponse{
		Id:          lembagaDomain.Id,
		KodeLembaga: lembagaDomain.KodeLembaga,
		NamaLembaga: lembagaDomain.NamaLembaga,
	}, nil
}

func (service *LembagaServiceImpl) Delete(ctx context.Context, id int) error {
	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer helper.CommitOrRollback(tx)

	return service.LembagaRepository.Delete(ctx, tx, id)
}

func (service *LembagaServiceImpl) FindById(ctx context.Context, id int) (web.LembagaResponse, error) {
	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		return web.LembagaResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	lembagaDomain, err := service.LembagaRepository.FindById(ctx, tx, id)
	if err != nil {
		return web.LembagaResponse{}, errors.New("id tidak ditemukan")
	}

	return web.LembagaResponse{
		Id:          lembagaDomain.Id,
		KodeLembaga: lembagaDomain.KodeLembaga,
		NamaLembaga: lembagaDomain.NamaLembaga,
	}, nil
}

func (service *LembagaServiceImpl) FindAll(ctx context.Context) ([]web.LembagaResponse, error) {
	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		return []web.LembagaResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	lembagaDomains, err := service.LembagaRepository.FindAll(ctx, tx)
	if err != nil {
		return []web.LembagaResponse{}, err
	}

	return helper.ToLembagaResponses(lembagaDomains), nil
}
