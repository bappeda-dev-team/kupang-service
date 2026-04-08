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

type RoleServiceImpl struct {
	RoleRepository repository.RoleRepository
	DB             *sql.DB
	Validator      *validator.Validate
}

func NewRoleServiceImpl(roleRepository repository.RoleRepository, db *sql.DB, validator *validator.Validate) *RoleServiceImpl {
	return &RoleServiceImpl{
		RoleRepository: roleRepository,
		DB:             db,
		Validator:      validator,
	}
}

func (service *RoleServiceImpl) Create(ctx context.Context, role web.RoleCreateRequest) (web.RoleResponse, error) {
	err := service.Validator.Struct(role)
	if err != nil {
		return web.RoleResponse{}, err
	}

	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		return web.RoleResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	roleDomain := domain.Role{
		Role: role.Role,
	}

	roleDomain, err = service.RoleRepository.Create(ctx, tx, roleDomain)
	if err != nil {
		return web.RoleResponse{}, err
	}

	return helper.ToRoleResponse(roleDomain), nil
}

func (service *RoleServiceImpl) Update(ctx context.Context, role web.RoleUpdateRequest) (web.RoleResponse, error) {
	err := service.Validator.Struct(role)
	if err != nil {
		return web.RoleResponse{}, err
	}

	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		return web.RoleResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	roleDomain := domain.Role{
		Id:   role.Id,
		Role: role.Role,
	}

	roleDomain, err = service.RoleRepository.Update(ctx, tx, roleDomain)
	if err != nil {
		return web.RoleResponse{}, err
	}

	return helper.ToRoleResponse(roleDomain), nil
}

func (service *RoleServiceImpl) Delete(ctx context.Context, id int) error {
	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer helper.CommitOrRollback(tx)

	return service.RoleRepository.Delete(ctx, tx, id)
}

func (service *RoleServiceImpl) FindById(ctx context.Context, id int) (web.RoleResponse, error) {
	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		return web.RoleResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	roleDomain, err := service.RoleRepository.FindById(ctx, tx, id)
	if err != nil {
		return web.RoleResponse{}, errors.New("id tidak ditemukan")
	}

	return helper.ToRoleResponse(roleDomain), nil
}

func (service *RoleServiceImpl) FindAll(ctx context.Context) ([]web.RoleResponse, error) {
	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		return []web.RoleResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	roleDomains, err := service.RoleRepository.FindAll(ctx, tx)
	if err != nil {
		return []web.RoleResponse{}, err
	}

	return helper.ToRoleResponses(roleDomains), nil
}
