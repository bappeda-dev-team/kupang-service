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

type PemdaServiceImpl struct {
	PemdaRepository repository.PemdaRepository
	DB              *sql.DB
	Validator       *validator.Validate
}

func NewPemdaServiceImpl(pemdaRepository repository.PemdaRepository, db *sql.DB, validator *validator.Validate) *PemdaServiceImpl {
	return &PemdaServiceImpl{
		PemdaRepository: pemdaRepository,
		DB:              db,
		Validator:       validator,
	}
}

func (service *PemdaServiceImpl) Create(ctx context.Context, pemda web.PemdaCreateRequest) (web.PemdaResponse, error) {
	err := service.Validator.Struct(pemda)
	if err != nil {
		return web.PemdaResponse{}, err
	}

	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		return web.PemdaResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	pemdaDomain := domain.Pemda{
		KodePemda: pemda.KodePemda,
		NamaPemda: pemda.NamaPemda,
	}

	pemdaDomain, err = service.PemdaRepository.Create(ctx, tx, pemdaDomain)
	if err != nil {
		return web.PemdaResponse{}, err
	}

	return web.PemdaResponse{
		Id:        pemdaDomain.Id,
		KodePemda: pemdaDomain.KodePemda,
		NamaPemda: pemdaDomain.NamaPemda,
	}, nil
}

func (service *PemdaServiceImpl) Update(ctx context.Context, pemdaData web.PemdaUpdateRequest) (web.PemdaResponse, error) {
	err := service.Validator.Struct(pemdaData)
	if err != nil {
		return web.PemdaResponse{}, err
	}

	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		return web.PemdaResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	pemdaDomain := domain.Pemda{
		Id:        pemdaData.Id,
		KodePemda: pemdaData.KodePemda,
		NamaPemda: pemdaData.NamaPemda,
	}

	pemdaDomain, err = service.PemdaRepository.Update(ctx, tx, pemdaDomain)
	if err != nil {
		return web.PemdaResponse{}, err
	}

	return web.PemdaResponse{
		Id:        pemdaDomain.Id,
		KodePemda: pemdaDomain.KodePemda,
		NamaPemda: pemdaDomain.NamaPemda,
	}, nil
}

func (service *PemdaServiceImpl) Delete(ctx context.Context, id int) error {
	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer helper.CommitOrRollback(tx)

	err = service.PemdaRepository.Delete(ctx, tx, id)
	if err != nil {
		return err
	}

	return nil
}

func (service *PemdaServiceImpl) FindById(ctx context.Context, id int) (web.PemdaResponse, error) {
	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		return web.PemdaResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	pemdaDomain, err := service.PemdaRepository.FindById(ctx, tx, id)
	if err != nil {
		return web.PemdaResponse{}, errors.New("id tidak ditemukan")
	}

	return web.PemdaResponse{
		Id:        pemdaDomain.Id,
		KodePemda: pemdaDomain.KodePemda,
		NamaPemda: pemdaDomain.NamaPemda,
	}, nil
}

func (service *PemdaServiceImpl) FindAll(ctx context.Context) ([]web.PemdaResponse, error) {
	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		return []web.PemdaResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	pemdaDomains, err := service.PemdaRepository.FindAll(ctx, tx)
	if err != nil {
		return []web.PemdaResponse{}, err
	}

	return helper.ToPemdaResponses(pemdaDomains), nil
}
