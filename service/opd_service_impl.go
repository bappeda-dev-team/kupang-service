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

type OpdServiceImpl struct {
	OpdRepository repository.OpdRepository
	DB            *sql.DB
	Validator     *validator.Validate
}

func NewOpdServiceImpl(opdRepository repository.OpdRepository, db *sql.DB, validator *validator.Validate) *OpdServiceImpl {
	return &OpdServiceImpl{
		OpdRepository: opdRepository,
		DB:            db,
		Validator:     validator,
	}
}

func (service *OpdServiceImpl) Create(ctx context.Context, opd web.OpdCreateRequest) (web.OpdResponse, error) {
	err := service.Validator.Struct(opd)
	if err != nil {
		return web.OpdResponse{}, err
	}

	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		return web.OpdResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	opdDomain := domain.Opd{
		KodeOpd: opd.KodeOpd,
		NamaOpd: opd.NamaOpd,
	}

	opdDomain, err = service.OpdRepository.Create(ctx, tx, opdDomain)
	if err != nil {
		return web.OpdResponse{}, err
	}

	return web.OpdResponse{
		Id:      opdDomain.Id,
		KodeOpd: opdDomain.KodeOpd,
		NamaOpd: opdDomain.NamaOpd,
	}, nil
}

func (service *OpdServiceImpl) Update(ctx context.Context, opdData web.OpdUpdateRequest) (web.OpdResponse, error) {
	err := service.Validator.Struct(opdData)
	if err != nil {
		return web.OpdResponse{}, err
	}

	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		return web.OpdResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	opdDomain := domain.Opd{
		Id:      opdData.Id,
		KodeOpd: opdData.KodeOpd,
		NamaOpd: opdData.NamaOpd,
	}

	opdDomain, err = service.OpdRepository.Update(ctx, tx, opdDomain)
	if err != nil {
		return web.OpdResponse{}, err
	}

	return web.OpdResponse{
		Id:      opdDomain.Id,
		KodeOpd: opdDomain.KodeOpd,
		NamaOpd: opdDomain.NamaOpd,
	}, nil
}

func (service *OpdServiceImpl) Delete(ctx context.Context, id int) error {
	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer helper.CommitOrRollback(tx)

	err = service.OpdRepository.Delete(ctx, tx, id)
	if err != nil {
		return err
	}

	return nil
}

func (service *OpdServiceImpl) FindById(ctx context.Context, id int) (web.OpdResponse, error) {
	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		return web.OpdResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	opdDomain, err := service.OpdRepository.FindById(ctx, tx, id)
	if err != nil {
		return web.OpdResponse{}, errors.New("id tidak ditemukan")
	}

	return web.OpdResponse{
		Id:      opdDomain.Id,
		KodeOpd: opdDomain.KodeOpd,
		NamaOpd: opdDomain.NamaOpd,
	}, nil
}

func (service *OpdServiceImpl) FindAll(ctx context.Context) ([]web.OpdResponse, error) {
	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		return []web.OpdResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	opdDomains, err := service.OpdRepository.FindAll(ctx, tx)
	if err != nil {
		return []web.OpdResponse{}, err
	}

	return helper.ToOpdResponses(opdDomains), nil
}
