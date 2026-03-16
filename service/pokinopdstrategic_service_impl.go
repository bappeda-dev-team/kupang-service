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

type PokinOpdStrategicServiceImpl struct {
	PokinOpdStrategicRepository          repository.PokinOpdStrategicRepository
	IndikatorPokinOpdStrategicRepository repository.IndikatorPokinOpdStrategicRepository
	TargetPokinOpdStrategicRepository    repository.TargetPokinOpdStrategicRepository
	DB                                   *sql.DB
	Validator                            *validator.Validate
}

func NewPokinOpdStrategicServiceImpl(
	pokinRepository repository.PokinOpdStrategicRepository,
	indikatorRepository repository.IndikatorPokinOpdStrategicRepository,
	targetRepository repository.TargetPokinOpdStrategicRepository,
	db *sql.DB,
	validator *validator.Validate,
) *PokinOpdStrategicServiceImpl {
	return &PokinOpdStrategicServiceImpl{
		PokinOpdStrategicRepository:          pokinRepository,
		IndikatorPokinOpdStrategicRepository: indikatorRepository,
		TargetPokinOpdStrategicRepository:    targetRepository,
		DB:                                   db,
		Validator:                            validator,
	}
}

func (service *PokinOpdStrategicServiceImpl) Create(ctx context.Context, request web.PokinOpdStrategicCreateRequest) (web.PokinOpdStrategicResponse, error) {
	if err := service.Validator.Struct(request); err != nil {
		return web.PokinOpdStrategicResponse{}, err
	}

	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		return web.PokinOpdStrategicResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	pokinDomain := domain.PokinOpdStrategic{
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

	pokinDomain, err = service.PokinOpdStrategicRepository.Create(ctx, tx, pokinDomain)
	if err != nil {
		return web.PokinOpdStrategicResponse{}, err
	}

	return web.PokinOpdStrategicResponse{
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

func (service *PokinOpdStrategicServiceImpl) Update(ctx context.Context, request web.PokinOpdStrategicUpdateRequest) (web.PokinOpdStrategicResponse, error) {
	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		return web.PokinOpdStrategicResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	pokinDomain := domain.PokinOpdStrategic{
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

	pokinDomain, err = service.PokinOpdStrategicRepository.Update(ctx, tx, pokinDomain)
	if err != nil {
		return web.PokinOpdStrategicResponse{}, err
	}

	return web.PokinOpdStrategicResponse{
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

func (service *PokinOpdStrategicServiceImpl) Delete(ctx context.Context, id int) error {
	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer helper.CommitOrRollback(tx)

	return service.PokinOpdStrategicRepository.Delete(ctx, tx, id)
}

func (service *PokinOpdStrategicServiceImpl) FindById(ctx context.Context, id int) (web.PokinOpdStrategicResponse, error) {
	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		return web.PokinOpdStrategicResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	pokinDomain, err := service.PokinOpdStrategicRepository.FindById(ctx, tx, id)
	if err != nil {
		return web.PokinOpdStrategicResponse{}, errors.New("id tidak ditemukan")
	}

	indikatorResponses, err := service.buildIndikatorResponses(ctx, tx, pokinDomain.Id)
	if err != nil {
		return web.PokinOpdStrategicResponse{}, err
	}

	return web.PokinOpdStrategicResponse{
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

func (service *PokinOpdStrategicServiceImpl) FindAll(ctx context.Context) ([]web.PokinOpdStrategicResponse, error) {
	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		return []web.PokinOpdStrategicResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	pokins, err := service.PokinOpdStrategicRepository.FindAll(ctx, tx)
	if err != nil {
		return []web.PokinOpdStrategicResponse{}, err
	}

	responses := make([]web.PokinOpdStrategicResponse, 0, len(pokins))
	for _, pokinDomain := range pokins {
		indikatorResponses, err := service.buildIndikatorResponses(ctx, tx, pokinDomain.Id)
		if err != nil {
			return []web.PokinOpdStrategicResponse{}, err
		}
		responses = append(responses, web.PokinOpdStrategicResponse{
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

func (service *PokinOpdStrategicServiceImpl) buildIndikatorResponses(ctx context.Context, tx *sql.Tx, pokinOpdStrategicId int) ([]web.PokinOpdStrategicIndikatorResponse, error) {
	indikatorDomains, err := service.IndikatorPokinOpdStrategicRepository.FindByPokinOpdStrategicId(ctx, tx, pokinOpdStrategicId)
	if err != nil {
		return nil, err
	}
	if len(indikatorDomains) == 0 {
		return nil, nil
	}

	indikatorResponses := make([]web.PokinOpdStrategicIndikatorResponse, 0, len(indikatorDomains))
	for _, indikator := range indikatorDomains {
		targetDomains, err := service.TargetPokinOpdStrategicRepository.FindByIndikatorId(ctx, tx, indikator.Id)
		if err != nil {
			return nil, err
		}

		targetResponses := make([]web.PokinOpdStrategicTargetResponse, 0, len(targetDomains))
		for _, target := range targetDomains {
			targetResponses = append(targetResponses, web.PokinOpdStrategicTargetResponse{
				IdTarget:    target.Id,
				IndikatorId: target.IndikatorPokinOpdStrategicId,
				Target:      target.NilaiTarget,
				Satuan:      target.Satuan,
			})
		}

		indikatorResponses = append(indikatorResponses, web.PokinOpdStrategicIndikatorResponse{
			IdIndikator:   indikator.Id,
			NamaIndikator: indikator.NamaIndikator,
			Targets:       targetResponses,
		})
	}

	return indikatorResponses, nil
}
