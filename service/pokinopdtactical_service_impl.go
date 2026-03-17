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

type PokinOpdTacticalServiceImpl struct {
	PokinOpdTacticalRepository          repository.PokinOpdTacticalRepository
	IndikatorPokinOpdTacticalRepository repository.IndikatorPokinOpdTacticalRepository
	TargetPokinOpdTacticalRepository    repository.TargetPokinOpdTacticalRepository
	DB                                  *sql.DB
	Validator                           *validator.Validate
}

func NewPokinOpdTacticalServiceImpl(
	pokinRepository repository.PokinOpdTacticalRepository,
	indikatorRepository repository.IndikatorPokinOpdTacticalRepository,
	targetRepository repository.TargetPokinOpdTacticalRepository,
	db *sql.DB,
	validator *validator.Validate,
) *PokinOpdTacticalServiceImpl {
	return &PokinOpdTacticalServiceImpl{
		PokinOpdTacticalRepository:          pokinRepository,
		IndikatorPokinOpdTacticalRepository: indikatorRepository,
		TargetPokinOpdTacticalRepository:    targetRepository,
		DB:                                  db,
		Validator:                           validator,
	}
}

func (service *PokinOpdTacticalServiceImpl) Create(ctx context.Context, request web.PokinOpdTacticalCreateRequest) (web.PokinOpdTacticalResponse, error) {
	if err := service.Validator.Struct(request); err != nil {
		return web.PokinOpdTacticalResponse{}, err
	}

	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		return web.PokinOpdTacticalResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	pokinDomain := domain.PokinOpdTactical{
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

	pokinDomain, err = service.PokinOpdTacticalRepository.Create(ctx, tx, pokinDomain)
	if err != nil {
		return web.PokinOpdTacticalResponse{}, err
	}

	return web.PokinOpdTacticalResponse{
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

func (service *PokinOpdTacticalServiceImpl) Update(ctx context.Context, request web.PokinOpdTacticalUpdateRequest) (web.PokinOpdTacticalResponse, error) {
	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		return web.PokinOpdTacticalResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	pokinDomain := domain.PokinOpdTactical{
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

	pokinDomain, err = service.PokinOpdTacticalRepository.Update(ctx, tx, pokinDomain)
	if err != nil {
		return web.PokinOpdTacticalResponse{}, err
	}

	return web.PokinOpdTacticalResponse{
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

func (service *PokinOpdTacticalServiceImpl) Delete(ctx context.Context, id int) error {
	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer helper.CommitOrRollback(tx)

	return service.PokinOpdTacticalRepository.Delete(ctx, tx, id)
}

func (service *PokinOpdTacticalServiceImpl) FindById(ctx context.Context, id int) (web.PokinOpdTacticalResponse, error) {
	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		return web.PokinOpdTacticalResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	pokinDomain, err := service.PokinOpdTacticalRepository.FindById(ctx, tx, id)
	if err != nil {
		return web.PokinOpdTacticalResponse{}, errors.New("id tidak ditemukan")
	}

	indikatorResponses, err := service.buildIndikatorResponses(ctx, tx, pokinDomain.Id)
	if err != nil {
		return web.PokinOpdTacticalResponse{}, err
	}

	return web.PokinOpdTacticalResponse{
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

func (service *PokinOpdTacticalServiceImpl) FindAll(ctx context.Context) ([]web.PokinOpdTacticalResponse, error) {
	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		return []web.PokinOpdTacticalResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	pokins, err := service.PokinOpdTacticalRepository.FindAll(ctx, tx)
	if err != nil {
		return []web.PokinOpdTacticalResponse{}, err
	}

	responses := make([]web.PokinOpdTacticalResponse, 0, len(pokins))
	for _, pokinDomain := range pokins {
		indikatorResponses, err := service.buildIndikatorResponses(ctx, tx, pokinDomain.Id)
		if err != nil {
			return []web.PokinOpdTacticalResponse{}, err
		}
		responses = append(responses, web.PokinOpdTacticalResponse{
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

func (service *PokinOpdTacticalServiceImpl) buildIndikatorResponses(ctx context.Context, tx *sql.Tx, pokinOpdTacticalId int) ([]web.PokinOpdTacticalIndikatorResponse, error) {
	indikatorDomains, err := service.IndikatorPokinOpdTacticalRepository.FindByPokinOpdTacticalId(ctx, tx, pokinOpdTacticalId)
	if err != nil {
		return nil, err
	}
	if len(indikatorDomains) == 0 {
		return nil, nil
	}

	indikatorResponses := make([]web.PokinOpdTacticalIndikatorResponse, 0, len(indikatorDomains))
	for _, indikator := range indikatorDomains {
		targetDomains, err := service.TargetPokinOpdTacticalRepository.FindByIndikatorId(ctx, tx, indikator.Id)
		if err != nil {
			return nil, err
		}

		targetResponses := make([]web.PokinOpdTacticalTargetResponse, 0, len(targetDomains))
		for _, target := range targetDomains {
			targetResponses = append(targetResponses, web.PokinOpdTacticalTargetResponse{
				IdTarget:    target.Id,
				IndikatorId: target.IndikatorPokinOpdTacticalId,
				Target:      target.NilaiTarget,
				Satuan:      target.Satuan,
			})
		}

		indikatorResponses = append(indikatorResponses, web.PokinOpdTacticalIndikatorResponse{
			IdIndikator:   indikator.Id,
			NamaIndikator: indikator.NamaIndikator,
			Targets:       targetResponses,
		})
	}

	return indikatorResponses, nil
}
