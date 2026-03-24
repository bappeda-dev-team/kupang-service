package service

import (
	"context"
	"database/sql"
	"kupang-service/helper"
	"kupang-service/model/web"
	"kupang-service/repository"
)

type PohonKinerjaServiceImpl struct {
	PokinOpdRepository                     repository.PokinOpdRepository
	TujuanPokinOpdRepository               repository.TujuanPokinOpdRepository
	IndikatorPokinOpdRepository            repository.IndikatorPokinOpdRepository
	TargetPokinOpdRepository               repository.TargetPokinOpdRepository
	PokinOpdStrategicRepository            repository.PokinOpdStrategicRepository
	IndikatorPokinOpdStrategicRepository   repository.IndikatorPokinOpdStrategicRepository
	TargetPokinOpdStrategicRepository      repository.TargetPokinOpdStrategicRepository
	PokinOpdTacticalRepository             repository.PokinOpdTacticalRepository
	IndikatorPokinOpdTacticalRepository    repository.IndikatorPokinOpdTacticalRepository
	TargetPokinOpdTacticalRepository       repository.TargetPokinOpdTacticalRepository
	PokinOpdOperationalRepository          repository.PokinOpdOperationalRepository
	IndikatorPokinOpdOperationalRepository repository.IndikatorPokinOpdOperationalRepository
	TargetPokinOpdOperationalRepository    repository.TargetPokinOpdOperationalRepository
	DB                                     *sql.DB
}

func NewPohonKinerjaServiceImpl(
	pokinOpdRepository repository.PokinOpdRepository,
	tujuanPokinOpdRepository repository.TujuanPokinOpdRepository,
	indikatorPokinOpdRepository repository.IndikatorPokinOpdRepository,
	targetPokinOpdRepository repository.TargetPokinOpdRepository,
	pokinOpdStrategicRepository repository.PokinOpdStrategicRepository,
	indikatorPokinOpdStrategicRepository repository.IndikatorPokinOpdStrategicRepository,
	targetPokinOpdStrategicRepository repository.TargetPokinOpdStrategicRepository,
	pokinOpdTacticalRepository repository.PokinOpdTacticalRepository,
	indikatorPokinOpdTacticalRepository repository.IndikatorPokinOpdTacticalRepository,
	targetPokinOpdTacticalRepository repository.TargetPokinOpdTacticalRepository,
	pokinOpdOperationalRepository repository.PokinOpdOperationalRepository,
	indikatorPokinOpdOperationalRepository repository.IndikatorPokinOpdOperationalRepository,
	targetPokinOpdOperationalRepository repository.TargetPokinOpdOperationalRepository,
	db *sql.DB,
) *PohonKinerjaServiceImpl {
	return &PohonKinerjaServiceImpl{
		PokinOpdRepository:                     pokinOpdRepository,
		TujuanPokinOpdRepository:               tujuanPokinOpdRepository,
		IndikatorPokinOpdRepository:            indikatorPokinOpdRepository,
		TargetPokinOpdRepository:               targetPokinOpdRepository,
		PokinOpdStrategicRepository:            pokinOpdStrategicRepository,
		IndikatorPokinOpdStrategicRepository:   indikatorPokinOpdStrategicRepository,
		TargetPokinOpdStrategicRepository:      targetPokinOpdStrategicRepository,
		PokinOpdTacticalRepository:             pokinOpdTacticalRepository,
		IndikatorPokinOpdTacticalRepository:    indikatorPokinOpdTacticalRepository,
		TargetPokinOpdTacticalRepository:       targetPokinOpdTacticalRepository,
		PokinOpdOperationalRepository:          pokinOpdOperationalRepository,
		IndikatorPokinOpdOperationalRepository: indikatorPokinOpdOperationalRepository,
		TargetPokinOpdOperationalRepository:    targetPokinOpdOperationalRepository,
		DB:                                     db,
	}
}

func (service *PohonKinerjaServiceImpl) FindByKodeOpdAndTahun(ctx context.Context, kodeOpd string, tahun int) (web.PohonKinerjaResponse, error) {
	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		return web.PohonKinerjaResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	pokinOpd, err := service.PokinOpdRepository.FindByKodeOpdAndTahun(ctx, tx, kodeOpd, tahun)
	if err != nil {
		return web.PohonKinerjaResponse{}, err
	}

	tujuanDomains, err := service.TujuanPokinOpdRepository.FindByPokinOpdId(ctx, tx, pokinOpd.Id)
	if err != nil {
		return web.PohonKinerjaResponse{}, err
	}

	tujuanResponses := make([]web.PohonKinerjaTujuanResponse, 0, len(tujuanDomains))
	for _, tujuan := range tujuanDomains {
		indikatorDomains, err := service.IndikatorPokinOpdRepository.FindByTujuanPokinOpdId(ctx, tx, tujuan.Id)
		if err != nil {
			return web.PohonKinerjaResponse{}, err
		}

		indikatorResponses := make([]web.PohonKinerjaIndikatorResponse, 0, len(indikatorDomains))
		for _, indikator := range indikatorDomains {
			targetDomains, err := service.TargetPokinOpdRepository.FindByIndikatorPokinOpdId(ctx, tx, indikator.Id)
			if err != nil {
				return web.PohonKinerjaResponse{}, err
			}

			targetResponses := make([]web.PohonKinerjaTargetResponse, 0, len(targetDomains))
			for _, target := range targetDomains {
				targetResponses = append(targetResponses, web.PohonKinerjaTargetResponse{
					Id:     target.Id,
					Target: target.NilaiTarget,
					Satuan: target.Satuan,
				})
			}

			indikatorResponses = append(indikatorResponses, web.PohonKinerjaIndikatorResponse{
				Id:        indikator.Id,
				Indikator: indikator.NamaIndikator,
				Targets:   targetResponses,
			})
		}

		tujuanResponses = append(tujuanResponses, web.PohonKinerjaTujuanResponse{
			Id:                tujuan.Id,
			KodeOpd:           tujuan.KodeOpd,
			Tujuan:            tujuan.NamaTujuan,
			BidangUrusan:      tujuan.BidangUrusan,
			TahunAwalPeriode:  tujuan.TahunAwalPeriode,
			TahunAkhirPeriode: tujuan.TahunAkhirPeriode,
			Indikator:         indikatorResponses,
		})
	}

	strategicDomains, err := service.PokinOpdStrategicRepository.FindByKodeOpdTahunParentLevel(ctx, tx, pokinOpd.KodeOpd, pokinOpd.Tahun, 0, 4)
	if err != nil {
		return web.PohonKinerjaResponse{}, err
	}

	childResponses := make([]web.PohonKinerjaStrategicResponse, 0, len(strategicDomains))
	for _, strategic := range strategicDomains {
		indikatorResponses, err := service.buildStrategicIndikatorResponses(ctx, tx, strategic.Id)
		if err != nil {
			return web.PohonKinerjaResponse{}, err
		}

		tacticalResponses, err := service.buildTacticalChildResponses(ctx, tx, strategic.Id)
		if err != nil {
			return web.PohonKinerjaResponse{}, err
		}

		childResponses = append(childResponses, web.PohonKinerjaStrategicResponse{
			Id:           strategic.Id,
			Parent:       strategic.Parent,
			NamaPohon:    strategic.NamaPohon,
			JenisPohon:   strategic.JenisPohon,
			LevelPohon:   strategic.LevelPohon,
			KodeOpd:      strategic.KodeOpd,
			NamaOpd:      strategic.NamaOpd,
			Keterangan:   strategic.Keterangan,
			Tahun:        strategic.Tahun,
			JumlahReview: strategic.JumlahReview,
			Status:       strategic.Status,
			Pelaksana:    strategic.Pelaksana,
			UpdatedBy:    strategic.UpdatedBy,
			Indikator:    indikatorResponses,
			Childs:       tacticalResponses,
		})
	}

	return web.PohonKinerjaResponse{
		KodeOpd:   pokinOpd.KodeOpd,
		NamaOpd:   pokinOpd.NamaOpd,
		Tahun:     pokinOpd.Tahun,
		TujuanOpd: tujuanResponses,
		Childs:    childResponses,
	}, nil
}

func (service *PohonKinerjaServiceImpl) buildStrategicIndikatorResponses(ctx context.Context, tx *sql.Tx, pokinOpdStrategicId int) ([]web.PokinOpdStrategicIndikatorResponse, error) {
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

func (service *PohonKinerjaServiceImpl) buildTacticalChildResponses(ctx context.Context, tx *sql.Tx, strategicId int) ([]web.WebResponse, error) {
	tacticalDomains, err := service.PokinOpdTacticalRepository.FindByParent(ctx, tx, strategicId)
	if err != nil {
		return nil, err
	}
	if len(tacticalDomains) == 0 {
		return []web.WebResponse{}, nil
	}

	responses := make([]web.WebResponse, 0, len(tacticalDomains))
	for _, tactical := range tacticalDomains {
		indikatorResponses, err := service.buildTacticalIndikatorResponses(ctx, tx, tactical.Id)
		if err != nil {
			return nil, err
		}

		operationalResponses, err := service.buildOperationalChildResponses(ctx, tx, tactical.Id)
		if err != nil {
			return nil, err
		}

		tacticalResponse := web.PokinOpdTacticalResponse{
			Id:           tactical.Id,
			Parent:       tactical.Parent,
			NamaPohon:    tactical.NamaPohon,
			JenisPohon:   tactical.JenisPohon,
			LevelPohon:   tactical.LevelPohon,
			KodeOpd:      tactical.KodeOpd,
			NamaOpd:      tactical.NamaOpd,
			Keterangan:   tactical.Keterangan,
			Tahun:        tactical.Tahun,
			JumlahReview: tactical.JumlahReview,
			Status:       tactical.Status,
			Pelaksana:    tactical.Pelaksana,
			UpdatedBy:    tactical.UpdatedBy,
			Indikator:    indikatorResponses,
			Childs:       operationalResponses,
		}

		responses = append(responses, web.WebResponse{
			Code:   200,
			Status: "OK",
			Data:   tacticalResponse,
		})
	}

	return responses, nil
}

func (service *PohonKinerjaServiceImpl) buildTacticalIndikatorResponses(ctx context.Context, tx *sql.Tx, pokinOpdTacticalId int) ([]web.PokinOpdTacticalIndikatorResponse, error) {
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

func (service *PohonKinerjaServiceImpl) buildOperationalChildResponses(ctx context.Context, tx *sql.Tx, tacticalId int) ([]web.PokinOpdOperationalResponse, error) {
	operationalDomains, err := service.PokinOpdOperationalRepository.FindByParent(ctx, tx, tacticalId)
	if err != nil {
		return nil, err
	}
	if len(operationalDomains) == 0 {
		return []web.PokinOpdOperationalResponse{}, nil
	}

	responses := make([]web.PokinOpdOperationalResponse, 0, len(operationalDomains))
	for _, operational := range operationalDomains {
		indikatorResponses, err := service.buildOperationalIndikatorResponses(ctx, tx, operational.Id)
		if err != nil {
			return nil, err
		}

		responses = append(responses, web.PokinOpdOperationalResponse{
			Id:           operational.Id,
			Parent:       operational.Parent,
			NamaPohon:    operational.NamaPohon,
			JenisPohon:   operational.JenisPohon,
			LevelPohon:   operational.LevelPohon,
			KodeOpd:      operational.KodeOpd,
			NamaOpd:      operational.NamaOpd,
			Keterangan:   operational.Keterangan,
			Tahun:        operational.Tahun,
			JumlahReview: operational.JumlahReview,
			Status:       operational.Status,
			Pelaksana:    operational.Pelaksana,
			UpdatedBy:    operational.UpdatedBy,
			Indikator:    indikatorResponses,
		})
	}

	return responses, nil
}

func (service *PohonKinerjaServiceImpl) buildOperationalIndikatorResponses(ctx context.Context, tx *sql.Tx, pokinOpdOperationalId int) ([]web.PokinOpdOperationalIndikatorResponse, error) {
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
