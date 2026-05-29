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

type ProgramServiceImpl struct {
	ProgramRepository repository.ProgramRepository
	DB                *sql.DB
	Validator         *validator.Validate
}

func NewProgramServiceImpl(programRepository repository.ProgramRepository, db *sql.DB, validator *validator.Validate) *ProgramServiceImpl {
	return &ProgramServiceImpl{
		ProgramRepository: programRepository,
		DB:                db,
		Validator:         validator,
	}
}

func (service *ProgramServiceImpl) Create(ctx context.Context, program web.ProgramCreateRequest) (web.ProgramResponse, error) {
	err := service.Validator.Struct(program)
	if err != nil {
		return web.ProgramResponse{}, err
	}

	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		return web.ProgramResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	programDomain := domain.Program{
		KodeProgram: program.KodeProgram,
		NamaProgram: program.NamaProgram,
		Tahun:       program.Tahun,
		KodeOpd:     program.KodeOpd,
	}

	programDomain, err = service.ProgramRepository.Create(ctx, tx, programDomain)
	if err != nil {
		return web.ProgramResponse{}, err
	}

	return helper.ToProgramResponse(programDomain), nil
}

func (service *ProgramServiceImpl) Update(ctx context.Context, programData web.ProgramUpdateRequest) (web.ProgramResponse, error) {
	err := service.Validator.Struct(programData)
	if err != nil {
		return web.ProgramResponse{}, err
	}

	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		return web.ProgramResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	programDomain := domain.Program{
		Id:          programData.Id,
		KodeProgram: programData.KodeProgram,
		NamaProgram: programData.NamaProgram,
		Tahun:       programData.Tahun,
		KodeOpd:     programData.KodeOpd,
	}

	programDomain, err = service.ProgramRepository.Update(ctx, tx, programDomain)
	if err != nil {
		return web.ProgramResponse{}, err
	}

	return helper.ToProgramResponse(programDomain), nil
}

func (service *ProgramServiceImpl) Delete(ctx context.Context, id int) error {
	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer helper.CommitOrRollback(tx)

	err = service.ProgramRepository.Delete(ctx, tx, id)
	if err != nil {
		return err
	}

	return nil
}

func (service *ProgramServiceImpl) FindById(ctx context.Context, id int) (web.ProgramResponse, error) {
	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		return web.ProgramResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	programDomain, err := service.ProgramRepository.FindById(ctx, tx, id)
	if err != nil {
		return web.ProgramResponse{}, errors.New("id tidak ditemukan")
	}

	return helper.ToProgramResponse(programDomain), nil
}

func (service *ProgramServiceImpl) FindAll(ctx context.Context) ([]web.ProgramResponse, error) {
	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		return []web.ProgramResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	programDomains, err := service.ProgramRepository.FindAll(ctx, tx)
	if err != nil {
		return []web.ProgramResponse{}, err
	}

	return helper.ToProgramResponses(programDomains), nil
}
