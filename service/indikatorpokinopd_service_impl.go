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

type IndikatorPokinOpdServiceImpl struct {
	IndikatorPokinOpdRepository repository.IndikatorPokinOpdRepository
	DB                          *sql.DB
	Validator                   *validator.Validate
}

func NewIndikatorPokinOpdServiceImpl(indikatorPokinOpdRepository repository.IndikatorPokinOpdRepository, db *sql.DB, validator *validator.Validate) *IndikatorPokinOpdServiceImpl {
	return &IndikatorPokinOpdServiceImpl{
		IndikatorPokinOpdRepository: indikatorPokinOpdRepository,
		DB:                          db,
		Validator:                   validator,
	}
}

func (service *IndikatorPokinOpdServiceImpl) Create(ctx context.Context, indikatorPokinOpd web.IndikatorPokinOpdCreateRequest) (web.IndikatorPokinOpdResponse, error) {
	err := service.Validator.Struct(indikatorPokinOpd)
	if err != nil {
		return web.IndikatorPokinOpdResponse{}, err
	}

	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		return web.IndikatorPokinOpdResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	indikatorPokinOpdDomain := domain.IndikatorPokinOpd{
		NamaIndikator: indikatorPokinOpd.NamaIndikator,
	}

	indikatorPokinOpdDomain, err = service.IndikatorPokinOpdRepository.Create(ctx, tx, indikatorPokinOpdDomain)
	if err != nil {
		return web.IndikatorPokinOpdResponse{}, err
	}

	return web.IndikatorPokinOpdResponse{
		NamaIndikator: indikatorPokinOpdDomain.NamaIndikator,
	}, nil
}

func (service *IndikatorPokinOpdServiceImpl) Update(ctx context.Context, indikatorPokinOpdData web.IndikatorPokinOpdUpdateRequest) (web.IndikatorPokinOpdResponse, error) {
	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		return web.IndikatorPokinOpdResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	indikatorPokinOpdDomain := domain.IndikatorPokinOpd{
		Id:            indikatorPokinOpdData.Id,
		NamaIndikator: indikatorPokinOpdData.NamaIndikator,
	}

	indikatorPokinOpdDomain, err = service.IndikatorPokinOpdRepository.Update(ctx, tx, indikatorPokinOpdDomain)
	if err != nil {
		return web.IndikatorPokinOpdResponse{}, err
	}

	return web.IndikatorPokinOpdResponse{
		Id:            indikatorPokinOpdDomain.Id,
		NamaIndikator: indikatorPokinOpdDomain.NamaIndikator,
	}, nil
}

func (service *IndikatorPokinOpdServiceImpl) Delete(ctx context.Context, id int) error {
	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer helper.CommitOrRollback(tx)

	err = service.IndikatorPokinOpdRepository.Delete(ctx, tx, id)
	if err != nil {
		return err
	}

	return nil
}

func (service *IndikatorPokinOpdServiceImpl) FindById(ctx context.Context, id int) (web.IndikatorPokinOpdResponse, error) {
	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		return web.IndikatorPokinOpdResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	indikatorPokinOpdDomain, err := service.IndikatorPokinOpdRepository.FindById(ctx, tx, id)
	if err != nil {
		return web.IndikatorPokinOpdResponse{}, errors.New("id tidak ditemukan")
	}

	return web.IndikatorPokinOpdResponse{
		Id:            indikatorPokinOpdDomain.Id,
		NamaIndikator: indikatorPokinOpdDomain.NamaIndikator,
	}, nil
}

func (service *IndikatorPokinOpdServiceImpl) FindAll(ctx context.Context) ([]web.IndikatorPokinOpdResponse, error) {
	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		return []web.IndikatorPokinOpdResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	IndikatorPokinOpdDomains, err := service.IndikatorPokinOpdRepository.FindAll(ctx, tx)
	if err != nil {
		return []web.IndikatorPokinOpdResponse{}, err
	}

	return helper.ToIndikatorPokinOpdResponses(IndikatorPokinOpdDomains), nil
}
