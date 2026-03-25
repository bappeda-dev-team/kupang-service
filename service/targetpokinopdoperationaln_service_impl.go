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

type TargetPokinOpdOperationalNServiceImpl struct {
	Repository repository.TargetPokinOpdOperationalNRepository
	DB         *sql.DB
	Validator  *validator.Validate
}

func NewTargetPokinOpdOperationalNServiceImpl(repository repository.TargetPokinOpdOperationalNRepository, db *sql.DB, validator *validator.Validate) *TargetPokinOpdOperationalNServiceImpl {
	return &TargetPokinOpdOperationalNServiceImpl{
		Repository: repository,
		DB:         db,
		Validator:  validator,
	}
}

func (service *TargetPokinOpdOperationalNServiceImpl) Create(ctx context.Context, request web.TargetPokinOpdOperationalNCreateRequest) (web.TargetPokinOpdOperationalNResponse, error) {
	if err := service.Validator.Struct(request); err != nil {
		return web.TargetPokinOpdOperationalNResponse{}, err
	}

	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		return web.TargetPokinOpdOperationalNResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	domainModel := domain.TargetPokinOpdOperationalN{
		IndikatorPokinOpdOperationalNId: request.IndikatorPokinOpdOperationalNId,
		NilaiTarget:                     request.NilaiTarget,
		Satuan:                          request.Satuan,
	}

	domainModel, err = service.Repository.Create(ctx, tx, domainModel)
	if err != nil {
		return web.TargetPokinOpdOperationalNResponse{}, err
	}

	return helper.ToTargetPokinOpdOperationalNResponse(domainModel), nil
}

func (service *TargetPokinOpdOperationalNServiceImpl) Update(ctx context.Context, request web.TargetPokinOpdOperationalNUpdateRequest) (web.TargetPokinOpdOperationalNResponse, error) {
	if err := service.Validator.Struct(request); err != nil {
		return web.TargetPokinOpdOperationalNResponse{}, err
	}

	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		return web.TargetPokinOpdOperationalNResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	domainModel := domain.TargetPokinOpdOperationalN{
		Id:                              request.Id,
		IndikatorPokinOpdOperationalNId: request.IndikatorPokinOpdOperationalNId,
		NilaiTarget:                     request.NilaiTarget,
		Satuan:                          request.Satuan,
	}

	domainModel, err = service.Repository.Update(ctx, tx, domainModel)
	if err != nil {
		return web.TargetPokinOpdOperationalNResponse{}, err
	}

	return helper.ToTargetPokinOpdOperationalNResponse(domainModel), nil
}

func (service *TargetPokinOpdOperationalNServiceImpl) Delete(ctx context.Context, id int) error {
	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer helper.CommitOrRollback(tx)

	return service.Repository.Delete(ctx, tx, id)
}

func (service *TargetPokinOpdOperationalNServiceImpl) FindById(ctx context.Context, id int) (web.TargetPokinOpdOperationalNResponse, error) {
	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		return web.TargetPokinOpdOperationalNResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	domainModel, err := service.Repository.FindById(ctx, tx, id)
	if err != nil {
		return web.TargetPokinOpdOperationalNResponse{}, errors.New("id tidak ditemukan")
	}

	return helper.ToTargetPokinOpdOperationalNResponse(domainModel), nil
}

func (service *TargetPokinOpdOperationalNServiceImpl) FindAll(ctx context.Context) ([]web.TargetPokinOpdOperationalNResponse, error) {
	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		return []web.TargetPokinOpdOperationalNResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	domains, err := service.Repository.FindAll(ctx, tx)
	if err != nil {
		return []web.TargetPokinOpdOperationalNResponse{}, err
	}

	return helper.ToTargetPokinOpdOperationalNResponses(domains), nil
}
