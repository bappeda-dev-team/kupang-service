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

type PokinOpdOperationalNServiceImpl struct {
	PokinOpdOperationalNRepository          repository.PokinOpdOperationalNRepository
	IndikatorPokinOpdOperationalNRepository repository.IndikatorPokinOpdOperationalNRepository
	TargetPokinOpdOperationalNRepository    repository.TargetPokinOpdOperationalNRepository
	DB                                      *sql.DB
	Validator                               *validator.Validate
}

func NewPokinOpdOperationalNServiceImpl(
	pokinRepository repository.PokinOpdOperationalNRepository,
	indikatorRepository repository.IndikatorPokinOpdOperationalNRepository,
	targetRepository repository.TargetPokinOpdOperationalNRepository,
	db *sql.DB,
	validator *validator.Validate,
) *PokinOpdOperationalNServiceImpl {
	return &PokinOpdOperationalNServiceImpl{
		PokinOpdOperationalNRepository:          pokinRepository,
		IndikatorPokinOpdOperationalNRepository: indikatorRepository,
		TargetPokinOpdOperationalNRepository:    targetRepository,
		DB:                                      db,
		Validator:                               validator,
	}
}

func (service *PokinOpdOperationalNServiceImpl) Create(ctx context.Context, request web.PokinOpdOperationalNCreateRequest) (web.PokinOpdOperationalNResponse, error) {
	if err := service.Validator.Struct(request); err != nil {
		return web.PokinOpdOperationalNResponse{}, err
	}

	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		return web.PokinOpdOperationalNResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	pokinDomain := domain.PokinOpdOperationalN{
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

	pokinDomain, err = service.PokinOpdOperationalNRepository.Create(ctx, tx, pokinDomain)
	if err != nil {
		return web.PokinOpdOperationalNResponse{}, err
	}

	return web.PokinOpdOperationalNResponse{
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

func (service *PokinOpdOperationalNServiceImpl) Update(ctx context.Context, request web.PokinOpdOperationalNUpdateRequest) (web.PokinOpdOperationalNResponse, error) {
	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		return web.PokinOpdOperationalNResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	pokinDomain := domain.PokinOpdOperationalN{
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

	pokinDomain, err = service.PokinOpdOperationalNRepository.Update(ctx, tx, pokinDomain)
	if err != nil {
		return web.PokinOpdOperationalNResponse{}, err
	}

	return web.PokinOpdOperationalNResponse{
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

func (service *PokinOpdOperationalNServiceImpl) Delete(ctx context.Context, id int) error {
	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer helper.CommitOrRollback(tx)

	return service.PokinOpdOperationalNRepository.Delete(ctx, tx, id)
}

func (service *PokinOpdOperationalNServiceImpl) FindById(ctx context.Context, id int) (web.PokinOpdOperationalNResponse, error) {
	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		return web.PokinOpdOperationalNResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	pokinDomain, err := service.PokinOpdOperationalNRepository.FindById(ctx, tx, id)
	if err != nil {
		return web.PokinOpdOperationalNResponse{}, errors.New("id tidak ditemukan")
	}

	indikatorResponses, err := service.buildIndikatorResponses(ctx, tx, pokinDomain.Id)
	if err != nil {
		return web.PokinOpdOperationalNResponse{}, err
	}

	return web.PokinOpdOperationalNResponse{
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

func (service *PokinOpdOperationalNServiceImpl) FindAll(ctx context.Context) ([]web.PokinOpdOperationalNResponse, error) {
	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		return []web.PokinOpdOperationalNResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	pokins, err := service.PokinOpdOperationalNRepository.FindAll(ctx, tx)
	if err != nil {
		return []web.PokinOpdOperationalNResponse{}, err
	}

	responses := make([]web.PokinOpdOperationalNResponse, 0, len(pokins))
	for _, pokinDomain := range pokins {
		indikatorResponses, err := service.buildIndikatorResponses(ctx, tx, pokinDomain.Id)
		if err != nil {
			return []web.PokinOpdOperationalNResponse{}, err
		}
		responses = append(responses, web.PokinOpdOperationalNResponse{
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

func (service *PokinOpdOperationalNServiceImpl) buildIndikatorResponses(ctx context.Context, tx *sql.Tx, pokinOpdOperationalNId int) ([]web.PokinOpdOperationalNIndikatorResponse, error) {
	indikatorDomains, err := service.IndikatorPokinOpdOperationalNRepository.FindByPokinOpdOperationalNId(ctx, tx, pokinOpdOperationalNId)
	if err != nil {
		return nil, err
	}
	if len(indikatorDomains) == 0 {
		return nil, nil
	}

	indikatorResponses := make([]web.PokinOpdOperationalNIndikatorResponse, 0, len(indikatorDomains))
	for _, indikator := range indikatorDomains {
		targetDomains, err := service.TargetPokinOpdOperationalNRepository.FindByIndikatorId(ctx, tx, indikator.Id)
		if err != nil {
			return nil, err
		}

		targetResponses := make([]web.PokinOpdOperationalNTargetResponse, 0, len(targetDomains))
		for _, target := range targetDomains {
			targetResponses = append(targetResponses, web.PokinOpdOperationalNTargetResponse{
				IdTarget:    target.Id,
				IndikatorId: target.IndikatorPokinOpdOperationalNId,
				Target:      target.NilaiTarget,
				Satuan:      target.Satuan,
			})
		}

		indikatorResponses = append(indikatorResponses, web.PokinOpdOperationalNIndikatorResponse{
			IdIndikator:   indikator.Id,
			NamaIndikator: indikator.NamaIndikator,
			Targets:       targetResponses,
		})
	}

	return indikatorResponses, nil
}
