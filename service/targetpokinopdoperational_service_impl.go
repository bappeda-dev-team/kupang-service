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

type TargetPokinOpdOperationalServiceImpl struct {
	Repository repository.TargetPokinOpdOperationalRepository
	DB         *sql.DB
	Validator  *validator.Validate
}

func NewTargetPokinOpdOperationalServiceImpl(repository repository.TargetPokinOpdOperationalRepository, db *sql.DB, validator *validator.Validate) *TargetPokinOpdOperationalServiceImpl {
	return &TargetPokinOpdOperationalServiceImpl{
		Repository: repository,
		DB:         db,
		Validator:  validator,
	}
}

func (service *TargetPokinOpdOperationalServiceImpl) Create(ctx context.Context, request web.TargetPokinOpdOperationalCreateRequest) (web.TargetPokinOpdOperationalResponse, error) {
	if err := service.Validator.Struct(request); err != nil {
		return web.TargetPokinOpdOperationalResponse{}, err
	}

	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		return web.TargetPokinOpdOperationalResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	domainModel := domain.TargetPokinOpdOperational{
		IndikatorPokinOpdOperationalId: request.IndikatorPokinOpdOperationalId,
		NilaiTarget:                    request.NilaiTarget,
		Satuan:                         request.Satuan,
	}

	domainModel, err = service.Repository.Create(ctx, tx, domainModel)
	if err != nil {
		return web.TargetPokinOpdOperationalResponse{}, err
	}

	return helper.ToTargetPokinOpdOperationalResponse(domainModel), nil
}

func (service *TargetPokinOpdOperationalServiceImpl) Update(ctx context.Context, request web.TargetPokinOpdOperationalUpdateRequest) (web.TargetPokinOpdOperationalResponse, error) {
	if err := service.Validator.Struct(request); err != nil {
		return web.TargetPokinOpdOperationalResponse{}, err
	}

	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		return web.TargetPokinOpdOperationalResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	domainModel := domain.TargetPokinOpdOperational{
		Id:                             request.Id,
		IndikatorPokinOpdOperationalId: request.IndikatorPokinOpdOperationalId,
		NilaiTarget:                    request.NilaiTarget,
		Satuan:                         request.Satuan,
	}

	domainModel, err = service.Repository.Update(ctx, tx, domainModel)
	if err != nil {
		return web.TargetPokinOpdOperationalResponse{}, err
	}

	return helper.ToTargetPokinOpdOperationalResponse(domainModel), nil
}

func (service *TargetPokinOpdOperationalServiceImpl) Delete(ctx context.Context, id int) error {
	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer helper.CommitOrRollback(tx)

	return service.Repository.Delete(ctx, tx, id)
}

func (service *TargetPokinOpdOperationalServiceImpl) FindById(ctx context.Context, id int) (web.TargetPokinOpdOperationalResponse, error) {
	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		return web.TargetPokinOpdOperationalResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	domainModel, err := service.Repository.FindById(ctx, tx, id)
	if err != nil {
		return web.TargetPokinOpdOperationalResponse{}, errors.New("id tidak ditemukan")
	}

	return helper.ToTargetPokinOpdOperationalResponse(domainModel), nil
}

func (service *TargetPokinOpdOperationalServiceImpl) FindAll(ctx context.Context) ([]web.TargetPokinOpdOperationalResponse, error) {
	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		return []web.TargetPokinOpdOperationalResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	domains, err := service.Repository.FindAll(ctx, tx)
	if err != nil {
		return []web.TargetPokinOpdOperationalResponse{}, err
	}

	return helper.ToTargetPokinOpdOperationalResponses(domains), nil
}
