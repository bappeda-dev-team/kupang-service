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

type TujuanPokinOpdServiceImpl struct {
	TujuanPokinOpdRepository repository.TujuanPokinOpdRepository
	PokinOpdRepository       repository.PokinOpdRepository
	DB                       *sql.DB
	Validator                *validator.Validate
}

func NewTujuanPokinOpdServiceImpl(tujuanPokinOpdRepository repository.TujuanPokinOpdRepository, pokinOpdRepository repository.PokinOpdRepository, db *sql.DB, validator *validator.Validate) *TujuanPokinOpdServiceImpl {
	return &TujuanPokinOpdServiceImpl{
		TujuanPokinOpdRepository: tujuanPokinOpdRepository,
		PokinOpdRepository:       pokinOpdRepository,
		DB:                       db,
		Validator:                validator,
	}
}

func (service *TujuanPokinOpdServiceImpl) Create(ctx context.Context, tujuanPokinOpd web.TujuanPokinOpdCreateRequest) (web.TujuanPokinOpdResponse, error) {
	err := service.Validator.Struct(tujuanPokinOpd)
	if err != nil {
		return web.TujuanPokinOpdResponse{}, err
	}

	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		return web.TujuanPokinOpdResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	if err := service.validatePokinOpdId(ctx, tx, tujuanPokinOpd.PokinOpdId); err != nil {
		return web.TujuanPokinOpdResponse{}, err
	}

	tujuanPokinOpdDomain := domain.TujuanPokinOpd{
		PokinOpdId:         tujuanPokinOpd.PokinOpdId,
		KodeOpd:           tujuanPokinOpd.KodeOpd,
		NamaTujuan:        tujuanPokinOpd.NamaTujuan,
		BidangUrusan:      tujuanPokinOpd.BidangUrusan,
		TahunAwalPeriode:  tujuanPokinOpd.TahunAwalPeriode,
		TahunAkhirPeriode: tujuanPokinOpd.TahunAkhirPeriode,
	}

	tujuanPokinOpdDomain, err = service.TujuanPokinOpdRepository.Create(ctx, tx, tujuanPokinOpdDomain)
	if err != nil {
		return web.TujuanPokinOpdResponse{}, err
	}

	return web.TujuanPokinOpdResponse{
		Id:                tujuanPokinOpdDomain.Id,
		PokinOpdId:        tujuanPokinOpdDomain.PokinOpdId,
		KodeOpd:           tujuanPokinOpdDomain.KodeOpd,
		NamaTujuan:        tujuanPokinOpdDomain.NamaTujuan,
		BidangUrusan:      tujuanPokinOpdDomain.BidangUrusan,
		TahunAwalPeriode:  tujuanPokinOpdDomain.TahunAwalPeriode,
		TahunAkhirPeriode: tujuanPokinOpdDomain.TahunAkhirPeriode,
	}, nil
}

func (service *TujuanPokinOpdServiceImpl) Update(ctx context.Context, tujuanPokinOpdData web.TujuanPokinOpdUpdateRequest) (web.TujuanPokinOpdResponse, error) {
	err := service.Validator.Struct(tujuanPokinOpdData)
	if err != nil {
		return web.TujuanPokinOpdResponse{}, err
	}

	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		return web.TujuanPokinOpdResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	if err := service.validatePokinOpdId(ctx, tx, tujuanPokinOpdData.PokinOpdId); err != nil {
		return web.TujuanPokinOpdResponse{}, err
	}

	tujuanPokinOpdDomain := domain.TujuanPokinOpd{
		Id:                tujuanPokinOpdData.Id,
		PokinOpdId:        tujuanPokinOpdData.PokinOpdId,
		KodeOpd:           tujuanPokinOpdData.KodeOpd,
		NamaTujuan:        tujuanPokinOpdData.NamaTujuan,
		BidangUrusan:      tujuanPokinOpdData.BidangUrusan,
		TahunAwalPeriode:  tujuanPokinOpdData.TahunAwalPeriode,
		TahunAkhirPeriode: tujuanPokinOpdData.TahunAkhirPeriode,
	}

	tujuanPokinOpdDomain, err = service.TujuanPokinOpdRepository.Update(ctx, tx, tujuanPokinOpdDomain)
	if err != nil {
		return web.TujuanPokinOpdResponse{}, err
	}

	return web.TujuanPokinOpdResponse{
		Id:                tujuanPokinOpdDomain.Id,
		PokinOpdId:        tujuanPokinOpdDomain.PokinOpdId,
		KodeOpd:           tujuanPokinOpdDomain.KodeOpd,
		NamaTujuan:        tujuanPokinOpdDomain.NamaTujuan,
		BidangUrusan:      tujuanPokinOpdDomain.BidangUrusan,
		TahunAwalPeriode:  tujuanPokinOpdDomain.TahunAwalPeriode,
		TahunAkhirPeriode: tujuanPokinOpdDomain.TahunAkhirPeriode,
	}, nil
}

func (service *TujuanPokinOpdServiceImpl) Delete(ctx context.Context, id int) error {
	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer helper.CommitOrRollback(tx)

	_, err = service.TujuanPokinOpdRepository.FindById(ctx, tx, id)
	if err != nil {
		return errors.New("id tidak ditemukan")
	}

	err = service.TujuanPokinOpdRepository.Delete(ctx, tx, id)
	if err != nil {
		return err
	}

	return nil
}

func (service *TujuanPokinOpdServiceImpl) FindById(ctx context.Context, id int) (web.TujuanPokinOpdResponse, error) {
	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		return web.TujuanPokinOpdResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	tujuanPokinOpdDomain, err := service.TujuanPokinOpdRepository.FindById(ctx, tx, id)
	if err != nil {
		return web.TujuanPokinOpdResponse{}, errors.New("id tidak ditemukan")
	}

	return web.TujuanPokinOpdResponse{
		Id:                tujuanPokinOpdDomain.Id,
		PokinOpdId:        tujuanPokinOpdDomain.PokinOpdId,
		KodeOpd:           tujuanPokinOpdDomain.KodeOpd,
		NamaTujuan:        tujuanPokinOpdDomain.NamaTujuan,
		BidangUrusan:      tujuanPokinOpdDomain.BidangUrusan,
		TahunAwalPeriode:  tujuanPokinOpdDomain.TahunAwalPeriode,
		TahunAkhirPeriode: tujuanPokinOpdDomain.TahunAkhirPeriode,
	}, nil
}

func (service *TujuanPokinOpdServiceImpl) FindAll(ctx context.Context) ([]web.TujuanPokinOpdResponse, error) {
	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		return []web.TujuanPokinOpdResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	tujuanPokinOpdDomains, err := service.TujuanPokinOpdRepository.FindAll(ctx, tx)
	if err != nil {
		return []web.TujuanPokinOpdResponse{}, err
	}

	return helper.ToTujuanPokinOpdResponses(tujuanPokinOpdDomains), nil
}

type pokinOpdIdValidation struct {
	PokinOpdId int `validate:"gt=0"`
}

func (service *TujuanPokinOpdServiceImpl) validatePokinOpdId(ctx context.Context, tx *sql.Tx, id int) error {
	_, err := service.PokinOpdRepository.FindById(ctx, tx, id)
	if err != nil {
		if err.Error() == "id tidak ditemukan" {
			return service.Validator.Struct(pokinOpdIdValidation{})
		}
		return err
	}

	return nil
}
