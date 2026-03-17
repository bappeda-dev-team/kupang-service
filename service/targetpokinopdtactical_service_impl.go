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

type TargetPokinOpdTacticalServiceImpl struct {
	Repository repository.TargetPokinOpdTacticalRepository
	DB         *sql.DB
	Validator  *validator.Validate
}

func NewTargetPokinOpdTacticalServiceImpl(repository repository.TargetPokinOpdTacticalRepository, db *sql.DB, validator *validator.Validate) *TargetPokinOpdTacticalServiceImpl {
	return &TargetPokinOpdTacticalServiceImpl{
		Repository: repository,
		DB:         db,
		Validator:  validator,
	}
}

func (service *TargetPokinOpdTacticalServiceImpl) Create(ctx context.Context, request web.TargetPokinOpdTacticalCreateRequest) (web.TargetPokinOpdTacticalResponse, error) {
	if err := service.Validator.Struct(request); err != nil {
		return web.TargetPokinOpdTacticalResponse{}, err
	}

	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		return web.TargetPokinOpdTacticalResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	domainModel := domain.TargetPokinOpdTactical{
		IndikatorPokinOpdTacticalId: request.IndikatorPokinOpdTacticalId,
		NilaiTarget:                 request.NilaiTarget,
		Satuan:                      request.Satuan,
	}

	domainModel, err = service.Repository.Create(ctx, tx, domainModel)
	if err != nil {
		return web.TargetPokinOpdTacticalResponse{}, err
	}

	return helper.ToTargetPokinOpdTacticalResponse(domainModel), nil
}

func (service *TargetPokinOpdTacticalServiceImpl) Update(ctx context.Context, request web.TargetPokinOpdTacticalUpdateRequest) (web.TargetPokinOpdTacticalResponse, error) {
	if err := service.Validator.Struct(request); err != nil {
		return web.TargetPokinOpdTacticalResponse{}, err
	}

	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		return web.TargetPokinOpdTacticalResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	domainModel := domain.TargetPokinOpdTactical{
		Id:                          request.Id,
		IndikatorPokinOpdTacticalId: request.IndikatorPokinOpdTacticalId,
		NilaiTarget:                 request.NilaiTarget,
		Satuan:                      request.Satuan,
	}

	domainModel, err = service.Repository.Update(ctx, tx, domainModel)
	if err != nil {
		return web.TargetPokinOpdTacticalResponse{}, err
	}

	return helper.ToTargetPokinOpdTacticalResponse(domainModel), nil
}

func (service *TargetPokinOpdTacticalServiceImpl) Delete(ctx context.Context, id int) error {
	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer helper.CommitOrRollback(tx)

	return service.Repository.Delete(ctx, tx, id)
}

func (service *TargetPokinOpdTacticalServiceImpl) FindById(ctx context.Context, id int) (web.TargetPokinOpdTacticalResponse, error) {
	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		return web.TargetPokinOpdTacticalResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	domainModel, err := service.Repository.FindById(ctx, tx, id)
	if err != nil {
		return web.TargetPokinOpdTacticalResponse{}, errors.New("id tidak ditemukan")
	}

	return helper.ToTargetPokinOpdTacticalResponse(domainModel), nil
}

func (service *TargetPokinOpdTacticalServiceImpl) FindAll(ctx context.Context) ([]web.TargetPokinOpdTacticalResponse, error) {
	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		return []web.TargetPokinOpdTacticalResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	domains, err := service.Repository.FindAll(ctx, tx)
	if err != nil {
		return []web.TargetPokinOpdTacticalResponse{}, err
	}

	return helper.ToTargetPokinOpdTacticalResponses(domains), nil
}
