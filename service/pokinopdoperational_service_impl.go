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

type PokinOpdOperationalServiceImpl struct {
	PokinOpdOperationalRepository          repository.PokinOpdOperationalRepository
	IndikatorPokinOpdOperationalRepository repository.IndikatorPokinOpdOperationalRepository
	TargetPokinOpdOperationalRepository    repository.TargetPokinOpdOperationalRepository
	DB                                     *sql.DB
	Validator                              *validator.Validate
}

func NewPokinOpdOperationalServiceImpl(
	pokinRepository repository.PokinOpdOperationalRepository,
	indikatorRepository repository.IndikatorPokinOpdOperationalRepository,
	targetRepository repository.TargetPokinOpdOperationalRepository,
	db *sql.DB,
	validator *validator.Validate,
) *PokinOpdOperationalServiceImpl {
	return &PokinOpdOperationalServiceImpl{
		PokinOpdOperationalRepository:          pokinRepository,
		IndikatorPokinOpdOperationalRepository: indikatorRepository,
		TargetPokinOpdOperationalRepository:    targetRepository,
		DB:                                     db,
		Validator:                              validator,
	}
}

func (service *PokinOpdOperationalServiceImpl) Create(ctx context.Context, request web.PokinOpdOperationalCreateRequest) (web.PokinOpdOperationalResponse, error) {
	if err := service.Validator.Struct(request); err != nil {
		return web.PokinOpdOperationalResponse{}, err
	}

	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		return web.PokinOpdOperationalResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	pokinDomain := domain.PokinOpdOperational{
		Parent:       request.Parent,
		NamaPohon:    request.NamaPohon,
		JenisPohon:   request.JenisPohon,
		LevelPohon:   request.LevelPohon,
		KodeOpd:      request.KodeOpd,
		NamaOpd:      request.NamaOpd,
		Keterangan:   request.Keterangan,
		Tahun:        request.Tahun,
		JumlahReview: request.JumlahReview,
		Status:       request.Status,
		Pelaksana:    request.Pelaksana,
		UpdatedBy:    request.UpdatedBy,
	}

	pokinDomain, err = service.PokinOpdOperationalRepository.Create(ctx, tx, pokinDomain)
	if err != nil {
		return web.PokinOpdOperationalResponse{}, err
	}

	return web.PokinOpdOperationalResponse{
		Id:           pokinDomain.Id,
		Parent:       pokinDomain.Parent,
		NamaPohon:    pokinDomain.NamaPohon,
		JenisPohon:   pokinDomain.JenisPohon,
		LevelPohon:   pokinDomain.LevelPohon,
		KodeOpd:      pokinDomain.KodeOpd,
		NamaOpd:      pokinDomain.NamaOpd,
		Keterangan:   pokinDomain.Keterangan,
		Tahun:        pokinDomain.Tahun,
		JumlahReview: pokinDomain.JumlahReview,
		Status:       pokinDomain.Status,
		Pelaksana:    pokinDomain.Pelaksana,
		UpdatedBy:    pokinDomain.UpdatedBy,
		Indikator:    nil,
	}, nil
}

func (service *PokinOpdOperationalServiceImpl) Update(ctx context.Context, request web.PokinOpdOperationalUpdateRequest) (web.PokinOpdOperationalResponse, error) {
	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		return web.PokinOpdOperationalResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	pokinDomain := domain.PokinOpdOperational{
		Id:           request.Id,
		Parent:       request.Parent,
		NamaPohon:    request.NamaPohon,
		JenisPohon:   request.JenisPohon,
		LevelPohon:   request.LevelPohon,
		KodeOpd:      request.KodeOpd,
		NamaOpd:      request.NamaOpd,
		Keterangan:   request.Keterangan,
		Tahun:        request.Tahun,
		JumlahReview: request.JumlahReview,
		Status:       request.Status,
		Pelaksana:    request.Pelaksana,
		UpdatedBy:    request.UpdatedBy,
	}

	pokinDomain, err = service.PokinOpdOperationalRepository.Update(ctx, tx, pokinDomain)
	if err != nil {
		return web.PokinOpdOperationalResponse{}, err
	}

	return web.PokinOpdOperationalResponse{
		Id:           pokinDomain.Id,
		Parent:       pokinDomain.Parent,
		NamaPohon:    pokinDomain.NamaPohon,
		JenisPohon:   pokinDomain.JenisPohon,
		LevelPohon:   pokinDomain.LevelPohon,
		KodeOpd:      pokinDomain.KodeOpd,
		NamaOpd:      pokinDomain.NamaOpd,
		Keterangan:   pokinDomain.Keterangan,
		Tahun:        pokinDomain.Tahun,
		JumlahReview: pokinDomain.JumlahReview,
		Status:       pokinDomain.Status,
		Pelaksana:    pokinDomain.Pelaksana,
		UpdatedBy:    pokinDomain.UpdatedBy,
		Indikator:    nil,
	}, nil
}

func (service *PokinOpdOperationalServiceImpl) Delete(ctx context.Context, id int) error {
	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer helper.CommitOrRollback(tx)

	return service.PokinOpdOperationalRepository.Delete(ctx, tx, id)
}

func (service *PokinOpdOperationalServiceImpl) FindById(ctx context.Context, id int) (web.PokinOpdOperationalResponse, error) {
	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		return web.PokinOpdOperationalResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	pokinDomain, err := service.PokinOpdOperationalRepository.FindById(ctx, tx, id)
	if err != nil {
		return web.PokinOpdOperationalResponse{}, errors.New("id tidak ditemukan")
	}

	indikatorResponses, err := service.buildIndikatorResponses(ctx, tx, pokinDomain.Id)
	if err != nil {
		return web.PokinOpdOperationalResponse{}, err
	}

	return web.PokinOpdOperationalResponse{
		Id:           pokinDomain.Id,
		Parent:       pokinDomain.Parent,
		NamaPohon:    pokinDomain.NamaPohon,
		JenisPohon:   pokinDomain.JenisPohon,
		LevelPohon:   pokinDomain.LevelPohon,
		KodeOpd:      pokinDomain.KodeOpd,
		NamaOpd:      pokinDomain.NamaOpd,
		Keterangan:   pokinDomain.Keterangan,
		Tahun:        pokinDomain.Tahun,
		JumlahReview: pokinDomain.JumlahReview,
		Status:       pokinDomain.Status,
		Pelaksana:    pokinDomain.Pelaksana,
		UpdatedBy:    pokinDomain.UpdatedBy,
		Indikator:    indikatorResponses,
	}, nil
}

func (service *PokinOpdOperationalServiceImpl) FindAll(ctx context.Context) ([]web.PokinOpdOperationalResponse, error) {
	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		return []web.PokinOpdOperationalResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	pokins, err := service.PokinOpdOperationalRepository.FindAll(ctx, tx)
	if err != nil {
		return []web.PokinOpdOperationalResponse{}, err
	}

	responses := make([]web.PokinOpdOperationalResponse, 0, len(pokins))
	for _, pokinDomain := range pokins {
		indikatorResponses, err := service.buildIndikatorResponses(ctx, tx, pokinDomain.Id)
		if err != nil {
			return []web.PokinOpdOperationalResponse{}, err
		}
		responses = append(responses, web.PokinOpdOperationalResponse{
			Id:           pokinDomain.Id,
			Parent:       pokinDomain.Parent,
			NamaPohon:    pokinDomain.NamaPohon,
			JenisPohon:   pokinDomain.JenisPohon,
			LevelPohon:   pokinDomain.LevelPohon,
			KodeOpd:      pokinDomain.KodeOpd,
			NamaOpd:      pokinDomain.NamaOpd,
			Keterangan:   pokinDomain.Keterangan,
			Tahun:        pokinDomain.Tahun,
			JumlahReview: pokinDomain.JumlahReview,
			Status:       pokinDomain.Status,
			Pelaksana:    pokinDomain.Pelaksana,
			UpdatedBy:    pokinDomain.UpdatedBy,
			Indikator:    indikatorResponses,
		})
	}

	return responses, nil
}

func (service *PokinOpdOperationalServiceImpl) buildIndikatorResponses(ctx context.Context, tx *sql.Tx, pokinOpdOperationalId int) ([]web.PokinOpdOperationalIndikatorResponse, error) {
	indikatorDomains, err := service.IndikatorPokinOpdOperationalRepository.FindByPokinOpdOperationalId(ctx, tx, pokinOpdOperationalId)
	if err != nil {
		return nil, err
	}
	if len(indikatorDomains) == 0 {
		return nil, nil
	}

	indikatorResponses := make([]web.PokinOpdOperationalIndikatorResponse, 0, len(indikatorDomains))
	for _, indikator := range indikatorDomains {
		targetDomains, err := service.TargetPokinOpdOperationalRepository.FindByIndikatorId(ctx, tx, indikator.Id)
		if err != nil {
			return nil, err
		}

		targetResponses := make([]web.PokinOpdOperationalTargetResponse, 0, len(targetDomains))
		for _, target := range targetDomains {
			targetResponses = append(targetResponses, web.PokinOpdOperationalTargetResponse{
				IdTarget:    target.Id,
				IndikatorId: target.IndikatorPokinOpdOperationalId,
				Target:      target.NilaiTarget,
				Satuan:      target.Satuan,
			})
		}

		indikatorResponses = append(indikatorResponses, web.PokinOpdOperationalIndikatorResponse{
			IdIndikator:   indikator.Id,
			NamaIndikator: indikator.NamaIndikator,
			Targets:       targetResponses,
		})
	}

	return indikatorResponses, nil
}
