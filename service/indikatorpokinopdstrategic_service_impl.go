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

type IndikatorPokinOpdStrategicServiceImpl struct {
	IndikatorPokinOpdStrategicRepository repository.IndikatorPokinOpdStrategicRepository
	PokinOpdStrategicRepository          repository.PokinOpdStrategicRepository
	DB                                   *sql.DB
	Validator                            *validator.Validate
}

func NewIndikatorPokinOpdStrategicServiceImpl(
	indikatorRepository repository.IndikatorPokinOpdStrategicRepository,
	pokinOpdStrategicRepository repository.PokinOpdStrategicRepository,
	db *sql.DB,
	validator *validator.Validate,
) *IndikatorPokinOpdStrategicServiceImpl {
	return &IndikatorPokinOpdStrategicServiceImpl{
		IndikatorPokinOpdStrategicRepository: indikatorRepository,
		PokinOpdStrategicRepository:          pokinOpdStrategicRepository,
		DB:                                   db,
		Validator:                            validator,
	}
}

func (service *IndikatorPokinOpdStrategicServiceImpl) Create(ctx context.Context, indikator web.IndikatorPokinOpdStrategicCreateRequest) (web.IndikatorPokinOpdStrategicResponse, error) {
	err := service.Validator.Struct(indikator)
	if err != nil {
		return web.IndikatorPokinOpdStrategicResponse{}, err
	}

	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		return web.IndikatorPokinOpdStrategicResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	if _, err := service.PokinOpdStrategicRepository.FindById(ctx, tx, indikator.PokinOpdStrategicId); err != nil {
		return web.IndikatorPokinOpdStrategicResponse{}, errors.New("pokin_opd_strategic_id tidak ditemukan")
	}

	domainIndikator := domain.IndikatorPokinOpdStrategic{
		PokinOpdStrategicId: indikator.PokinOpdStrategicId,
		NamaIndikator:       indikator.NamaIndikator,
	}

	domainIndikator, err = service.IndikatorPokinOpdStrategicRepository.Create(ctx, tx, domainIndikator)
	if err != nil {
		return web.IndikatorPokinOpdStrategicResponse{}, err
	}

	return web.IndikatorPokinOpdStrategicResponse{
		Id:                  domainIndikator.Id,
		PokinOpdStrategicId: domainIndikator.PokinOpdStrategicId,
		NamaIndikator:       domainIndikator.NamaIndikator,
	}, nil
}

func (service *IndikatorPokinOpdStrategicServiceImpl) Update(ctx context.Context, indikator web.IndikatorPokinOpdStrategicUpdateRequest) (web.IndikatorPokinOpdStrategicResponse, error) {
	if err := service.Validator.Struct(indikator); err != nil {
		return web.IndikatorPokinOpdStrategicResponse{}, err
	}

	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		return web.IndikatorPokinOpdStrategicResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	if _, err := service.PokinOpdStrategicRepository.FindById(ctx, tx, indikator.PokinOpdStrategicId); err != nil {
		return web.IndikatorPokinOpdStrategicResponse{}, errors.New("pokin_opd_strategic_id tidak ditemukan")
	}

	domainIndikator := domain.IndikatorPokinOpdStrategic{
		Id:                  indikator.Id,
		PokinOpdStrategicId: indikator.PokinOpdStrategicId,
		NamaIndikator:       indikator.NamaIndikator,
	}

	domainIndikator, err = service.IndikatorPokinOpdStrategicRepository.Update(ctx, tx, domainIndikator)
	if err != nil {
		return web.IndikatorPokinOpdStrategicResponse{}, err
	}

	return web.IndikatorPokinOpdStrategicResponse{
		Id:                  domainIndikator.Id,
		PokinOpdStrategicId: domainIndikator.PokinOpdStrategicId,
		NamaIndikator:       domainIndikator.NamaIndikator,
	}, nil
}

func (service *IndikatorPokinOpdStrategicServiceImpl) Delete(ctx context.Context, id int) error {
	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer helper.CommitOrRollback(tx)

	err = service.IndikatorPokinOpdStrategicRepository.Delete(ctx, tx, id)
	if err != nil {
		return err
	}

	return nil
}

func (service *IndikatorPokinOpdStrategicServiceImpl) FindById(ctx context.Context, id int) (web.IndikatorPokinOpdStrategicResponse, error) {
	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		return web.IndikatorPokinOpdStrategicResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	indikator, err := service.IndikatorPokinOpdStrategicRepository.FindById(ctx, tx, id)
	if err != nil {
		return web.IndikatorPokinOpdStrategicResponse{}, errors.New("id tidak ditemukan")
	}

	return web.IndikatorPokinOpdStrategicResponse{
		Id:                  indikator.Id,
		PokinOpdStrategicId: indikator.PokinOpdStrategicId,
		NamaIndikator:       indikator.NamaIndikator,
	}, nil
}

func (service *IndikatorPokinOpdStrategicServiceImpl) FindAll(ctx context.Context) ([]web.IndikatorPokinOpdStrategicResponse, error) {
	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		return []web.IndikatorPokinOpdStrategicResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	indikators, err := service.IndikatorPokinOpdStrategicRepository.FindAll(ctx, tx)
	if err != nil {
		return []web.IndikatorPokinOpdStrategicResponse{}, err
	}

	return helper.ToIndikatorPokinOpdStrategicResponses(indikators), nil
}
