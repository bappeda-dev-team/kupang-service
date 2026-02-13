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

type PokinOpdServiceImpl struct {
	PokinOpdRepository repository.PokinOpdRepository
	DB                 *sql.DB
	Validator          *validator.Validate
}

func NewPokinOpdServiceImpl(pokinOpdRepository repository.PokinOpdRepository, db *sql.DB, validator *validator.Validate) *PokinOpdServiceImpl {
	return &PokinOpdServiceImpl{
		PokinOpdRepository: pokinOpdRepository,
		DB:                 db,
		Validator:          validator,
	}
}

func (service *PokinOpdServiceImpl) Create(ctx context.Context, pokinOpd web.PokinOpdCreateRequest) (web.PokinOpdResponse, error) {
	err := service.Validator.Struct(pokinOpd)
	if err != nil {
		return web.PokinOpdResponse{}, err
	}

	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		return web.PokinOpdResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	pokinOpdDomain := domain.PokinOpd{
		KodeOpd: pokinOpd.KodeOpd,
		NamaOpd: pokinOpd.NamaOpd,
		Tahun:   pokinOpd.Tahun,
	}

	pokinOpdDomain, err = service.PokinOpdRepository.Create(ctx, tx, pokinOpdDomain)
	if err != nil {
		return web.PokinOpdResponse{}, err
	}

	return web.PokinOpdResponse{
		KodeOpd: pokinOpdDomain.KodeOpd,
		NamaOpd: pokinOpd.NamaOpd,
		Tahun:   pokinOpd.Tahun,
	}, nil
}

func (service *PokinOpdServiceImpl) Update(ctx context.Context, pokinOpdData web.PokinOpdUpdateRequest) (web.PokinOpdResponse, error) {
	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		return web.PokinOpdResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	pokinOpdDomain := domain.PokinOpd{
		Id:      pokinOpdData.Id,
		KodeOpd: pokinOpdData.KodeOpd,
		NamaOpd: pokinOpdData.NamaOpd,
		Tahun:   pokinOpdData.Tahun,
	}

	pokinOpdDomain, err = service.PokinOpdRepository.Update(ctx, tx, pokinOpdDomain)
	if err != nil {
		return web.PokinOpdResponse{}, err
	}

	return web.PokinOpdResponse{
		Id:      pokinOpdDomain.Id,
		KodeOpd: pokinOpdDomain.KodeOpd,
		NamaOpd: pokinOpdDomain.NamaOpd,
		Tahun:   pokinOpdDomain.Tahun,
	}, nil
}

func (service *PokinOpdServiceImpl) Delete(ctx context.Context, id int) error {
	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer helper.CommitOrRollback(tx)

	err = service.PokinOpdRepository.Delete(ctx, tx, id)
	if err != nil {
		return err
	}

	return nil
}

func (service *PokinOpdServiceImpl) FindById(ctx context.Context, id int) (web.PokinOpdResponse, error) {
	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		return web.PokinOpdResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	pokinOpdDomain, err := service.PokinOpdRepository.FindById(ctx, tx, id)
	if err != nil {
		return web.PokinOpdResponse{}, errors.New("id tidak ditemukan")
	}

	return web.PokinOpdResponse{
		Id:      pokinOpdDomain.Id,
		KodeOpd: pokinOpdDomain.KodeOpd,
		NamaOpd: pokinOpdDomain.NamaOpd,
		Tahun:   pokinOpdDomain.Tahun,
	}, nil
}

func (service *PokinOpdServiceImpl) FindAll(ctx context.Context) ([]web.PokinOpdResponse, error) {
	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		return []web.PokinOpdResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	PokinOpdDomains, err := service.PokinOpdRepository.FindAll(ctx, tx)
	if err != nil {
		return []web.PokinOpdResponse{}, err
	}

	return helper.ToPokinOpdResponses(PokinOpdDomains), nil
}
