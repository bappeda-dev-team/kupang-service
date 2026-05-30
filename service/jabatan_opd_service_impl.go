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

type JabatanOpdServiceImpl struct {
	JabatanOpdRepository repository.JabatanOpdRepository
	DB                   *sql.DB
	Validator            *validator.Validate
}

func NewJabatanOpdServiceImpl(jabatanOpdRepository repository.JabatanOpdRepository, db *sql.DB, validator *validator.Validate) *JabatanOpdServiceImpl {
	return &JabatanOpdServiceImpl{
		JabatanOpdRepository: jabatanOpdRepository,
		DB:                   db,
		Validator:            validator,
	}
}

func (service *JabatanOpdServiceImpl) Create(ctx context.Context, jabatanOpd web.JabatanOpdCreateRequest) (web.JabatanOpdResponse, error) {
	err := service.Validator.Struct(jabatanOpd)
	if err != nil {
		return web.JabatanOpdResponse{}, err
	}

	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		return web.JabatanOpdResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	jabatanOpdDomain := domain.JabatanOpd{
		KodeJabatan: jabatanOpd.KodeJabatan,
		NamaJabatan: jabatanOpd.NamaJabatan,
		KodeOpd:     jabatanOpd.KodeOpd,
		Tahun:       jabatanOpd.Tahun,
	}

	jabatanOpdDomain, err = service.JabatanOpdRepository.Create(ctx, tx, jabatanOpdDomain)
	if err != nil {
		return web.JabatanOpdResponse{}, err
	}

	return helper.ToJabatanOpdResponse(jabatanOpdDomain), nil
}

func (service *JabatanOpdServiceImpl) Update(ctx context.Context, jabatanOpdData web.JabatanOpdUpdateRequest) (web.JabatanOpdResponse, error) {
	err := service.Validator.Struct(jabatanOpdData)
	if err != nil {
		return web.JabatanOpdResponse{}, err
	}

	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		return web.JabatanOpdResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	jabatanOpdDomain := domain.JabatanOpd{
		Id:          jabatanOpdData.Id,
		KodeJabatan: jabatanOpdData.KodeJabatan,
		NamaJabatan: jabatanOpdData.NamaJabatan,
		KodeOpd:     jabatanOpdData.KodeOpd,
		Tahun:       jabatanOpdData.Tahun,
	}

	jabatanOpdDomain, err = service.JabatanOpdRepository.Update(ctx, tx, jabatanOpdDomain)
	if err != nil {
		return web.JabatanOpdResponse{}, err
	}

	return helper.ToJabatanOpdResponse(jabatanOpdDomain), nil
}

func (service *JabatanOpdServiceImpl) Delete(ctx context.Context, id int) error {
	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer helper.CommitOrRollback(tx)

	err = service.JabatanOpdRepository.Delete(ctx, tx, id)
	if err != nil {
		return err
	}

	return nil
}

func (service *JabatanOpdServiceImpl) FindById(ctx context.Context, id int) (web.JabatanOpdResponse, error) {
	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		return web.JabatanOpdResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	jabatanOpdDomain, err := service.JabatanOpdRepository.FindById(ctx, tx, id)
	if err != nil {
		return web.JabatanOpdResponse{}, errors.New("id tidak ditemukan")
	}

	return helper.ToJabatanOpdResponse(jabatanOpdDomain), nil
}

func (service *JabatanOpdServiceImpl) FindAll(ctx context.Context) ([]web.JabatanOpdResponse, error) {
	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		return []web.JabatanOpdResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	jabatanOpdDomains, err := service.JabatanOpdRepository.FindAll(ctx, tx)
	if err != nil {
		return []web.JabatanOpdResponse{}, err
	}

	return helper.ToJabatanOpdResponses(jabatanOpdDomains), nil
}
