package service

import (
	"context"
	"database/sql"

	"kupang-service/helper"
	"kupang-service/model/domain"
	"kupang-service/model/web"
	"kupang-service/repository"

	"github.com/go-playground/validator/v10"
)

type UserServiceImpl struct {
	UserRepository repository.UserRepository
	DB             *sql.DB
	Validator      *validator.Validate
}

func NewUserServiceImpl(userRepository repository.UserRepository, db *sql.DB, validator *validator.Validate) *UserServiceImpl {
	return &UserServiceImpl{
		UserRepository: userRepository,
		DB:             db,
		Validator:      validator,
	}
}

func (service *UserServiceImpl) Create(ctx context.Context, user web.UserCreateRequest) (web.UserResponse, error) {
	err := service.Validator.Struct(user)
	if err != nil {
		return web.UserResponse{}, err
	}

	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		return web.UserResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	userDomain := domain.User{
		Nama:        user.Nama,
		Nip:         sql.NullString{String: user.Nip, Valid: true},
		Email:       user.Email,
		Status:      user.Status,
		Role:        sql.NullString{String: user.Role, Valid: true},
		KodeOpd:     sql.NullString{String: user.KodeOpd, Valid: true},
		OpdId:       sql.NullInt64{Int64: int64(user.OpdId), Valid: true},
		PegawaiId:   sql.NullInt64{Int64: int64(user.PegawaiId), Valid: true},
		RoleId:      sql.NullInt64{Int64: int64(user.RoleId), Valid: true},
		NamaOpd:     sql.NullString{String: user.NamaOpd, Valid: user.NamaOpd != ""},
		NamaPegawai: sql.NullString{String: user.NamaPegawai, Valid: user.NamaPegawai != ""},
	}

	userDomain, err = service.UserRepository.Create(ctx, tx, userDomain)
	if err != nil {
		return web.UserResponse{}, err
	}

	return helper.ToUserResponse(userDomain), nil
}

func (service *UserServiceImpl) Update(ctx context.Context, user web.UserUpdateRequest) (web.UserResponse, error) {
	err := service.Validator.Struct(user)
	if err != nil {
		return web.UserResponse{}, err
	}

	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		return web.UserResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	userDomain := domain.User{
		Id:          user.Id,
		Nama:        user.Nama,
		Nip:         sql.NullString{String: user.Nip, Valid: true},
		Email:       user.Email,
		Status:      user.Status,
		Role:        sql.NullString{String: user.Role, Valid: true},
		KodeOpd:     sql.NullString{String: user.KodeOpd, Valid: true},
		OpdId:       sql.NullInt64{Int64: int64(user.OpdId), Valid: true},
		PegawaiId:   sql.NullInt64{Int64: int64(user.PegawaiId), Valid: true},
		RoleId:      sql.NullInt64{Int64: int64(user.RoleId), Valid: true},
		NamaOpd:     sql.NullString{String: user.NamaOpd, Valid: user.NamaOpd != ""},
		NamaPegawai: sql.NullString{String: user.NamaPegawai, Valid: user.NamaPegawai != ""},
	}

	userDomain, err = service.UserRepository.Update(ctx, tx, userDomain)
	if err != nil {
		return web.UserResponse{}, err
	}

	return helper.ToUserResponse(userDomain), nil
}

func (service *UserServiceImpl) Delete(ctx context.Context, id int) error {
	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer helper.CommitOrRollback(tx)

	err = service.UserRepository.Delete(ctx, tx, id)
	if err != nil {
		return err
	}

	return nil
}

func (service *UserServiceImpl) FindById(ctx context.Context, id int) (web.UserResponse, error) {
	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		return web.UserResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	userDomain, err := service.UserRepository.FindById(ctx, tx, id)
	if err != nil {
		return web.UserResponse{}, err
	}

	return helper.ToUserResponse(userDomain), nil
}

func (service *UserServiceImpl) FindAll(ctx context.Context) ([]web.UserResponse, error) {
	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		return []web.UserResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	userDomains, err := service.UserRepository.FindAll(ctx, tx)
	if err != nil {
		return []web.UserResponse{}, err
	}

	return helper.ToUserResponses(userDomains), nil
}

func (service *UserServiceImpl) FindByKodeOpd(ctx context.Context, kodeOpd string) ([]web.UserResponse, error) {
	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		return []web.UserResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	userDomains, err := service.UserRepository.FindByKodeOpd(ctx, tx, kodeOpd)
	if err != nil {
		return []web.UserResponse{}, err
	}

	return helper.ToUserResponses(userDomains), nil
}
