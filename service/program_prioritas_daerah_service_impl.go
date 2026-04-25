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

type ProgramPrioritasDaerahServiceImpl struct {
	ProgramPrioritasDaerahRepository repository.ProgramPrioritasDaerahRepository
	DB                               *sql.DB
	Validator                        *validator.Validate
}

func NewProgramPrioritasDaerahServiceImpl(programRepository repository.ProgramPrioritasDaerahRepository, db *sql.DB, validator *validator.Validate) *ProgramPrioritasDaerahServiceImpl {
	return &ProgramPrioritasDaerahServiceImpl{
		ProgramPrioritasDaerahRepository: programRepository,
		DB:                               db,
		Validator:                        validator,
	}
}

func (service *ProgramPrioritasDaerahServiceImpl) Create(ctx context.Context, program web.ProgramPrioritasDaerahCreateRequest) (web.ProgramPrioritasDaerahResponse, error) {
	err := service.Validator.Struct(program)
	if err != nil {
		return web.ProgramPrioritasDaerahResponse{}, err
	}

	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		return web.ProgramPrioritasDaerahResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	programDomain := domain.ProgramPrioritasDaerah{
		KodeProgramPrioritasDaerah: program.KodeProgramPrioritasDaerah,
		NamaProgramPrioritasDaerah: program.NamaProgramPrioritasDaerah,
		RencanaImplementasi:        program.RencanaImplementasi,
		Keterangan:                 program.Keterangan,
		TahunAwal:                  program.TahunAwal,
		TahunAkhir:                 program.TahunAkhir,
		IsActive:                   program.IsActive,
	}

	programDomain, err = service.ProgramPrioritasDaerahRepository.Create(ctx, tx, programDomain)
	if err != nil {
		return web.ProgramPrioritasDaerahResponse{}, err
	}

	return helper.ToProgramPrioritasDaerahResponse(programDomain), nil
}

func (service *ProgramPrioritasDaerahServiceImpl) Update(ctx context.Context, program web.ProgramPrioritasDaerahUpdateRequest) (web.ProgramPrioritasDaerahResponse, error) {
	err := service.Validator.Struct(program)
	if err != nil {
		return web.ProgramPrioritasDaerahResponse{}, err
	}

	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		return web.ProgramPrioritasDaerahResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	programDomain := domain.ProgramPrioritasDaerah{
		Id:                         program.Id,
		KodeProgramPrioritasDaerah: program.KodeProgramPrioritasDaerah,
		NamaProgramPrioritasDaerah: program.NamaProgramPrioritasDaerah,
		RencanaImplementasi:        program.RencanaImplementasi,
		Keterangan:                 program.Keterangan,
		TahunAwal:                  program.TahunAwal,
		TahunAkhir:                 program.TahunAkhir,
		IsActive:                   program.IsActive,
	}

	programDomain, err = service.ProgramPrioritasDaerahRepository.Update(ctx, tx, programDomain)
	if err != nil {
		return web.ProgramPrioritasDaerahResponse{}, err
	}

	return helper.ToProgramPrioritasDaerahResponse(programDomain), nil
}

func (service *ProgramPrioritasDaerahServiceImpl) Delete(ctx context.Context, id int) error {
	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer helper.CommitOrRollback(tx)

	return service.ProgramPrioritasDaerahRepository.Delete(ctx, tx, id)
}

func (service *ProgramPrioritasDaerahServiceImpl) FindById(ctx context.Context, id int) (web.ProgramPrioritasDaerahResponse, error) {
	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		return web.ProgramPrioritasDaerahResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	programDomain, err := service.ProgramPrioritasDaerahRepository.FindById(ctx, tx, id)
	if err != nil {
		return web.ProgramPrioritasDaerahResponse{}, err
	}

	return helper.ToProgramPrioritasDaerahResponse(programDomain), nil
}

func (service *ProgramPrioritasDaerahServiceImpl) FindAll(ctx context.Context) ([]web.ProgramPrioritasDaerahResponse, error) {
	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		return []web.ProgramPrioritasDaerahResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	programs, err := service.ProgramPrioritasDaerahRepository.FindAll(ctx, tx)
	if err != nil {
		return []web.ProgramPrioritasDaerahResponse{}, err
	}

	return helper.ToProgramPrioritasDaerahResponses(programs), nil
}

func (service *ProgramPrioritasDaerahServiceImpl) FindByTahunRange(ctx context.Context, tahunAwal, tahunAkhir string) ([]web.ProgramPrioritasDaerahResponse, error) {
	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		return []web.ProgramPrioritasDaerahResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	programs, err := service.ProgramPrioritasDaerahRepository.FindByTahunRange(ctx, tx, tahunAwal, tahunAkhir)
	if err != nil {
		return []web.ProgramPrioritasDaerahResponse{}, err
	}

	return helper.ToProgramPrioritasDaerahResponses(programs), nil
}
