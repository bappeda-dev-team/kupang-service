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

type IndikatorPokinOpdOperationalServiceImpl struct {
	IndikatorPokinOpdOperationalRepository repository.IndikatorPokinOpdOperationalRepository
	DB                                     *sql.DB
	Validator                              *validator.Validate
}

func NewIndikatorPokinOpdOperationalServiceImpl(
	indikatorRepository repository.IndikatorPokinOpdOperationalRepository,
	db *sql.DB,
	validator *validator.Validate,
) *IndikatorPokinOpdOperationalServiceImpl {
	return &IndikatorPokinOpdOperationalServiceImpl{
		IndikatorPokinOpdOperationalRepository: indikatorRepository,
		DB:                                     db,
		Validator:                              validator,
	}
}

func (service *IndikatorPokinOpdOperationalServiceImpl) Create(ctx context.Context, indikator web.IndikatorPokinOpdOperationalCreateRequest) (web.IndikatorPokinOpdOperationalResponse, error) {
	if err := service.Validator.Struct(indikator); err != nil {
		return web.IndikatorPokinOpdOperationalResponse{}, err
	}

	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		return web.IndikatorPokinOpdOperationalResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	domainIndikator := domain.IndikatorPokinOpdOperational{
		PokinOpdOperationalId: indikator.PokinOpdOperationalId,
		NamaIndikator:         indikator.NamaIndikator,
	}

	domainIndikator, err = service.IndikatorPokinOpdOperationalRepository.Create(ctx, tx, domainIndikator)
	if err != nil {
		return web.IndikatorPokinOpdOperationalResponse{}, err
	}

	return web.IndikatorPokinOpdOperationalResponse{
		Id:                    domainIndikator.Id,
		PokinOpdOperationalId: domainIndikator.PokinOpdOperationalId,
		NamaIndikator:         domainIndikator.NamaIndikator,
	}, nil
}

func (service *IndikatorPokinOpdOperationalServiceImpl) Update(ctx context.Context, indikator web.IndikatorPokinOpdOperationalUpdateRequest) (web.IndikatorPokinOpdOperationalResponse, error) {
	if err := service.Validator.Struct(indikator); err != nil {
		return web.IndikatorPokinOpdOperationalResponse{}, err
	}

	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		return web.IndikatorPokinOpdOperationalResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	domainIndikator := domain.IndikatorPokinOpdOperational{
		Id:                    indikator.Id,
		PokinOpdOperationalId: indikator.PokinOpdOperationalId,
		NamaIndikator:         indikator.NamaIndikator,
	}

	domainIndikator, err = service.IndikatorPokinOpdOperationalRepository.Update(ctx, tx, domainIndikator)
	if err != nil {
		return web.IndikatorPokinOpdOperationalResponse{}, err
	}

	return web.IndikatorPokinOpdOperationalResponse{
		Id:                    domainIndikator.Id,
		PokinOpdOperationalId: domainIndikator.PokinOpdOperationalId,
		NamaIndikator:         domainIndikator.NamaIndikator,
	}, nil
}

func (service *IndikatorPokinOpdOperationalServiceImpl) Delete(ctx context.Context, id int) error {
	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer helper.CommitOrRollback(tx)

	err = service.IndikatorPokinOpdOperationalRepository.Delete(ctx, tx, id)
	if err != nil {
		return err
	}

	return nil
}

func (service *IndikatorPokinOpdOperationalServiceImpl) FindById(ctx context.Context, id int) (web.IndikatorPokinOpdOperationalResponse, error) {
	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		return web.IndikatorPokinOpdOperationalResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	indikator, err := service.IndikatorPokinOpdOperationalRepository.FindById(ctx, tx, id)
	if err != nil {
		return web.IndikatorPokinOpdOperationalResponse{}, errors.New("id tidak ditemukan")
	}

	return web.IndikatorPokinOpdOperationalResponse{
		Id:                    indikator.Id,
		PokinOpdOperationalId: indikator.PokinOpdOperationalId,
		NamaIndikator:         indikator.NamaIndikator,
	}, nil
}

func (service *IndikatorPokinOpdOperationalServiceImpl) FindAll(ctx context.Context) ([]web.IndikatorPokinOpdOperationalResponse, error) {
	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		return []web.IndikatorPokinOpdOperationalResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	indikators, err := service.IndikatorPokinOpdOperationalRepository.FindAll(ctx, tx)
	if err != nil {
		return []web.IndikatorPokinOpdOperationalResponse{}, err
	}

	return helper.ToIndikatorPokinOpdOperationalResponses(indikators), nil
}
