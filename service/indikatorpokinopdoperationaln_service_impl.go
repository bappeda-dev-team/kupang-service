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

type IndikatorPokinOpdOperationalNServiceImpl struct {
	Repository repository.IndikatorPokinOpdOperationalNRepository
	DB         *sql.DB
	Validator  *validator.Validate
}

func NewIndikatorPokinOpdOperationalNServiceImpl(
	repository repository.IndikatorPokinOpdOperationalNRepository,
	db *sql.DB,
	validator *validator.Validate,
) *IndikatorPokinOpdOperationalNServiceImpl {
	return &IndikatorPokinOpdOperationalNServiceImpl{
		Repository: repository,
		DB:         db,
		Validator:  validator,
	}
}

func (service *IndikatorPokinOpdOperationalNServiceImpl) Create(ctx context.Context, indikator web.IndikatorPokinOpdOperationalNCreateRequest) (web.IndikatorPokinOpdOperationalNResponse, error) {
	if err := service.Validator.Struct(indikator); err != nil {
		return web.IndikatorPokinOpdOperationalNResponse{}, err
	}

	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		return web.IndikatorPokinOpdOperationalNResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	domainIndikator := domain.IndikatorPokinOpdOperationalN{
		PokinOpdOperationalNId: indikator.PokinOpdOperationalNId,
		NamaIndikator:          indikator.NamaIndikator,
	}

	domainIndikator, err = service.Repository.Create(ctx, tx, domainIndikator)
	if err != nil {
		return web.IndikatorPokinOpdOperationalNResponse{}, err
	}

	return helper.ToIndikatorPokinOpdOperationalNResponse(domainIndikator), nil
}

func (service *IndikatorPokinOpdOperationalNServiceImpl) Update(ctx context.Context, indikator web.IndikatorPokinOpdOperationalNUpdateRequest) (web.IndikatorPokinOpdOperationalNResponse, error) {
	if err := service.Validator.Struct(indikator); err != nil {
		return web.IndikatorPokinOpdOperationalNResponse{}, err
	}

	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		return web.IndikatorPokinOpdOperationalNResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	domainIndikator := domain.IndikatorPokinOpdOperationalN{
		Id:                     indikator.Id,
		PokinOpdOperationalNId: indikator.PokinOpdOperationalNId,
		NamaIndikator:          indikator.NamaIndikator,
	}

	domainIndikator, err = service.Repository.Update(ctx, tx, domainIndikator)
	if err != nil {
		return web.IndikatorPokinOpdOperationalNResponse{}, err
	}

	return helper.ToIndikatorPokinOpdOperationalNResponse(domainIndikator), nil
}

func (service *IndikatorPokinOpdOperationalNServiceImpl) Delete(ctx context.Context, id int) error {
	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer helper.CommitOrRollback(tx)

	return service.Repository.Delete(ctx, tx, id)
}

func (service *IndikatorPokinOpdOperationalNServiceImpl) FindById(ctx context.Context, id int) (web.IndikatorPokinOpdOperationalNResponse, error) {
	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		return web.IndikatorPokinOpdOperationalNResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	indikator, err := service.Repository.FindById(ctx, tx, id)
	if err != nil {
		return web.IndikatorPokinOpdOperationalNResponse{}, errors.New("id tidak ditemukan")
	}

	return helper.ToIndikatorPokinOpdOperationalNResponse(indikator), nil
}

func (service *IndikatorPokinOpdOperationalNServiceImpl) FindAll(ctx context.Context) ([]web.IndikatorPokinOpdOperationalNResponse, error) {
	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		return []web.IndikatorPokinOpdOperationalNResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	indikators, err := service.Repository.FindAll(ctx, tx)
	if err != nil {
		return []web.IndikatorPokinOpdOperationalNResponse{}, err
	}

	return helper.ToIndikatorPokinOpdOperationalNResponses(indikators), nil
}
