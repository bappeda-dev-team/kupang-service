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

type IndikatorPokinOpdTacticalServiceImpl struct {
	IndikatorPokinOpdTacticalRepository repository.IndikatorPokinOpdTacticalRepository
	DB                                  *sql.DB
	Validator                           *validator.Validate
}

func NewIndikatorPokinOpdTacticalServiceImpl(indikatorRepository repository.IndikatorPokinOpdTacticalRepository, db *sql.DB, validator *validator.Validate) *IndikatorPokinOpdTacticalServiceImpl {
	return &IndikatorPokinOpdTacticalServiceImpl{
		IndikatorPokinOpdTacticalRepository: indikatorRepository,
		DB:                                  db,
		Validator:                           validator,
	}
}

func (service *IndikatorPokinOpdTacticalServiceImpl) Create(ctx context.Context, indikator web.IndikatorPokinOpdTacticalCreateRequest) (web.IndikatorPokinOpdTacticalResponse, error) {
	err := service.Validator.Struct(indikator)
	if err != nil {
		return web.IndikatorPokinOpdTacticalResponse{}, err
	}

	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		return web.IndikatorPokinOpdTacticalResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	domainIndikator := domain.IndikatorPokinOpdTactical{
		PokinOpdTacticalId: indikator.PokinOpdTacticalId,
		NamaIndikator:      indikator.NamaIndikator,
	}

	domainIndikator, err = service.IndikatorPokinOpdTacticalRepository.Create(ctx, tx, domainIndikator)
	if err != nil {
		return web.IndikatorPokinOpdTacticalResponse{}, err
	}

	return web.IndikatorPokinOpdTacticalResponse{
		Id:                 domainIndikator.Id,
		PokinOpdTacticalId: domainIndikator.PokinOpdTacticalId,
		NamaIndikator:      domainIndikator.NamaIndikator,
	}, nil
}

func (service *IndikatorPokinOpdTacticalServiceImpl) Update(ctx context.Context, indikator web.IndikatorPokinOpdTacticalUpdateRequest) (web.IndikatorPokinOpdTacticalResponse, error) {
	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		return web.IndikatorPokinOpdTacticalResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	domainIndikator := domain.IndikatorPokinOpdTactical{
		Id:                 indikator.Id,
		PokinOpdTacticalId: indikator.PokinOpdTacticalId,
		NamaIndikator:      indikator.NamaIndikator,
	}

	domainIndikator, err = service.IndikatorPokinOpdTacticalRepository.Update(ctx, tx, domainIndikator)
	if err != nil {
		return web.IndikatorPokinOpdTacticalResponse{}, err
	}

	return web.IndikatorPokinOpdTacticalResponse{
		Id:                 domainIndikator.Id,
		PokinOpdTacticalId: domainIndikator.PokinOpdTacticalId,
		NamaIndikator:      domainIndikator.NamaIndikator,
	}, nil
}

func (service *IndikatorPokinOpdTacticalServiceImpl) Delete(ctx context.Context, id int) error {
	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer helper.CommitOrRollback(tx)

	err = service.IndikatorPokinOpdTacticalRepository.Delete(ctx, tx, id)
	if err != nil {
		return err
	}

	return nil
}

func (service *IndikatorPokinOpdTacticalServiceImpl) FindById(ctx context.Context, id int) (web.IndikatorPokinOpdTacticalResponse, error) {
	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		return web.IndikatorPokinOpdTacticalResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	indikator, err := service.IndikatorPokinOpdTacticalRepository.FindById(ctx, tx, id)
	if err != nil {
		return web.IndikatorPokinOpdTacticalResponse{}, errors.New("id tidak ditemukan")
	}

	return web.IndikatorPokinOpdTacticalResponse{
		Id:                 indikator.Id,
		PokinOpdTacticalId: indikator.PokinOpdTacticalId,
		NamaIndikator:      indikator.NamaIndikator,
	}, nil
}

func (service *IndikatorPokinOpdTacticalServiceImpl) FindAll(ctx context.Context) ([]web.IndikatorPokinOpdTacticalResponse, error) {
	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		return []web.IndikatorPokinOpdTacticalResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	indikators, err := service.IndikatorPokinOpdTacticalRepository.FindAll(ctx, tx)
	if err != nil {
		return []web.IndikatorPokinOpdTacticalResponse{}, err
	}

	return helper.ToIndikatorPokinOpdTacticalResponses(indikators), nil
}
