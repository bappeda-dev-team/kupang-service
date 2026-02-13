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

type TargetPokinOpdServiceImpl struct {
	TargetPokinOpdRepository repository.TargetPokinOpdRepository
	DB                       *sql.DB
	Validator                *validator.Validate
}

func NewTargetPokinOpdServiceImpl(targetPokinOpdRepository repository.TargetPokinOpdRepository, db *sql.DB, validator *validator.Validate) *TargetPokinOpdServiceImpl {
	return &TargetPokinOpdServiceImpl{
		TargetPokinOpdRepository: targetPokinOpdRepository,
		DB:                       db,
		Validator:                validator,
	}
}

func (service *TargetPokinOpdServiceImpl) Create(ctx context.Context, targetPokinOpd web.TargetPokinOpdCreateRequest) (web.TargetPokinOpdResponse, error) {
	err := service.Validator.Struct(targetPokinOpd)
	if err != nil {
		return web.TargetPokinOpdResponse{}, err
	}

	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		return web.TargetPokinOpdResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	targetPokinOpdDomain := domain.TargetPokinOpd{
		NilaiTarget: targetPokinOpd.NilaiTarget,
		Satuan:      targetPokinOpd.Satuan,
	}

	targetPokinOpdDomain, err = service.TargetPokinOpdRepository.Create(ctx, tx, targetPokinOpdDomain)
	if err != nil {
		return web.TargetPokinOpdResponse{}, err
	}

	return helper.ToTargetPokinOpdResponse(targetPokinOpdDomain), nil
}

func (service *TargetPokinOpdServiceImpl) Update(ctx context.Context, targetPokinOpd web.TargetPokinOpdUpdateRequest) (web.TargetPokinOpdResponse, error) {
	err := service.Validator.Struct(targetPokinOpd)
	if err != nil {
		return web.TargetPokinOpdResponse{}, err
	}

	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		return web.TargetPokinOpdResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	targetPokinOpdDomain := domain.TargetPokinOpd{
		Id:          targetPokinOpd.Id,
		NilaiTarget: targetPokinOpd.NamaTarget,
		Satuan:      targetPokinOpd.Satuan,
	}

	targetPokinOpdDomain, err = service.TargetPokinOpdRepository.Update(ctx, tx, targetPokinOpdDomain)
	if err != nil {
		return web.TargetPokinOpdResponse{}, err
	}

	return helper.ToTargetPokinOpdResponse(targetPokinOpdDomain), nil
}

func (service *TargetPokinOpdServiceImpl) Delete(ctx context.Context, id int) error {
	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer helper.CommitOrRollback(tx)

	err = service.TargetPokinOpdRepository.Delete(ctx, tx, id)
	if err != nil {
		return err
	}

	return nil
}

func (service *TargetPokinOpdServiceImpl) FindById(ctx context.Context, id int) (web.TargetPokinOpdResponse, error) {
	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		return web.TargetPokinOpdResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	targetPokinOpdDomain, err := service.TargetPokinOpdRepository.FindById(ctx, tx, id)
	if err != nil {
		return web.TargetPokinOpdResponse{}, errors.New("id tidak ditemukan")
	}

	return helper.ToTargetPokinOpdResponse(targetPokinOpdDomain), nil
}

func (service *TargetPokinOpdServiceImpl) FindAll(ctx context.Context) ([]web.TargetPokinOpdResponse, error) {
	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		return []web.TargetPokinOpdResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	targetPokinOpdDomains, err := service.TargetPokinOpdRepository.FindAll(ctx, tx)
	if err != nil {
		return []web.TargetPokinOpdResponse{}, err
	}

	return helper.ToTargetPokinOpdResponses(targetPokinOpdDomains), nil
}
