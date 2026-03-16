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

type TargetPokinOpdStrategicServiceImpl struct {
	Repository repository.TargetPokinOpdStrategicRepository
	DB         *sql.DB
	Validator  *validator.Validate
}

func NewTargetPokinOpdStrategicServiceImpl(repository repository.TargetPokinOpdStrategicRepository, db *sql.DB, validator *validator.Validate) *TargetPokinOpdStrategicServiceImpl {
	return &TargetPokinOpdStrategicServiceImpl{
		Repository: repository,
		DB:         db,
		Validator:  validator,
	}
}

func (service *TargetPokinOpdStrategicServiceImpl) Create(ctx context.Context, request web.TargetPokinOpdStrategicCreateRequest) (web.TargetPokinOpdStrategicResponse, error) {
	if err := service.Validator.Struct(request); err != nil {
		return web.TargetPokinOpdStrategicResponse{}, err
	}

	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		return web.TargetPokinOpdStrategicResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	domainModel := domain.TargetPokinOpdStrategic{
		IndikatorPokinOpdStrategicId: request.IndikatorPokinOpdStrategicId,
		NilaiTarget:                  request.NilaiTarget,
		Satuan:                       request.Satuan,
	}

	domainModel, err = service.Repository.Create(ctx, tx, domainModel)
	if err != nil {
		return web.TargetPokinOpdStrategicResponse{}, err
	}

	return helper.ToTargetPokinOpdStrategicResponse(domainModel), nil
}

func (service *TargetPokinOpdStrategicServiceImpl) Update(ctx context.Context, request web.TargetPokinOpdStrategicUpdateRequest) (web.TargetPokinOpdStrategicResponse, error) {
	if err := service.Validator.Struct(request); err != nil {
		return web.TargetPokinOpdStrategicResponse{}, err
	}

	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		return web.TargetPokinOpdStrategicResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	domainModel := domain.TargetPokinOpdStrategic{
		Id:                           request.Id,
		NilaiTarget:                  request.NamaTarget,
		Satuan:                       request.Satuan,
		IndikatorPokinOpdStrategicId: request.IndikatorPokinOpdStrategicId,
	}

	domainModel, err = service.Repository.Update(ctx, tx, domainModel)
	if err != nil {
		return web.TargetPokinOpdStrategicResponse{}, err
	}

	return helper.ToTargetPokinOpdStrategicResponse(domainModel), nil
}

func (service *TargetPokinOpdStrategicServiceImpl) Delete(ctx context.Context, id int) error {
	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer helper.CommitOrRollback(tx)

	return service.Repository.Delete(ctx, tx, id)
}

func (service *TargetPokinOpdStrategicServiceImpl) FindById(ctx context.Context, id int) (web.TargetPokinOpdStrategicResponse, error) {
	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		return web.TargetPokinOpdStrategicResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	domainModel, err := service.Repository.FindById(ctx, tx, id)
	if err != nil {
		return web.TargetPokinOpdStrategicResponse{}, errors.New("id tidak ditemukan")
	}

	return helper.ToTargetPokinOpdStrategicResponse(domainModel), nil
}

func (service *TargetPokinOpdStrategicServiceImpl) FindAll(ctx context.Context) ([]web.TargetPokinOpdStrategicResponse, error) {
	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		return []web.TargetPokinOpdStrategicResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	domains, err := service.Repository.FindAll(ctx, tx)
	if err != nil {
		return []web.TargetPokinOpdStrategicResponse{}, err
	}

	return helper.ToTargetPokinOpdStrategicResponses(domains), nil
}
