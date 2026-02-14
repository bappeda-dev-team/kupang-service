package service

import (
	"context"
	"database/sql"
	"kupang-service/helper"
	"kupang-service/model/web"
	"kupang-service/repository"
)

type PohonKinerjaServiceImpl struct {
	PokinOpdRepository         repository.PokinOpdRepository
	TujuanPokinOpdRepository   repository.TujuanPokinOpdRepository
	IndikatorPokinOpdRepository repository.IndikatorPokinOpdRepository
	TargetPokinOpdRepository    repository.TargetPokinOpdRepository
	DB                         *sql.DB
}

func NewPohonKinerjaServiceImpl(pokinOpdRepository repository.PokinOpdRepository, tujuanPokinOpdRepository repository.TujuanPokinOpdRepository, indikatorPokinOpdRepository repository.IndikatorPokinOpdRepository, targetPokinOpdRepository repository.TargetPokinOpdRepository, db *sql.DB) *PohonKinerjaServiceImpl {
	return &PohonKinerjaServiceImpl{
		PokinOpdRepository:         pokinOpdRepository,
		TujuanPokinOpdRepository:   tujuanPokinOpdRepository,
		IndikatorPokinOpdRepository: indikatorPokinOpdRepository,
		TargetPokinOpdRepository:    targetPokinOpdRepository,
		DB:                         db,
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

	return web.PohonKinerjaResponse{
		KodeOpd:   pokinOpd.KodeOpd,
		NamaOpd:   pokinOpd.NamaOpd,
		Tahun:     pokinOpd.Tahun,
		TujuanOpd: tujuanResponses,
		Childs:    []interface{}{},
	}, nil
}
