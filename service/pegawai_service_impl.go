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

type PegawaiServiceImpl struct {
	PegawaiRepository repository.PegawaiRepository
	DB                *sql.DB
	Validator         *validator.Validate
}

func NewPegawaiServiceImpl(pegawaiRepository repository.PegawaiRepository, db *sql.DB, validator *validator.Validate) *PegawaiServiceImpl {
	return &PegawaiServiceImpl{
		PegawaiRepository: pegawaiRepository,
		DB:                db,
		Validator:         validator,
	}
}

func (service *PegawaiServiceImpl) Create(ctx context.Context, pegawai web.PegawaiCreateRequest) (web.PegawaiResponse, error) {
	err := service.Validator.Struct(pegawai)
	if err != nil {
		return web.PegawaiResponse{}, err
	}

	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		return web.PegawaiResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	pegawaiDomain := domain.Pegawai{
		Nama:    pegawai.Nama,
		Nip:     pegawai.Nip,
		Jabatan: pegawai.Jabatan,
		KodeOpd: pegawai.KodeOpd,
		NamaOpd: pegawai.NamaOpd,
	}

	pegawaiDomain, err = service.PegawaiRepository.Create(ctx, tx, pegawaiDomain)
	if err != nil {
		return web.PegawaiResponse{}, err
	}

	return web.PegawaiResponse{
		Id:      pegawaiDomain.Id,
		Nama:    pegawaiDomain.Nama,
		Nip:     pegawaiDomain.Nip,
		Jabatan: pegawaiDomain.Jabatan,
		KodeOpd: pegawaiDomain.KodeOpd,
		NamaOpd: pegawaiDomain.NamaOpd,
	}, nil
}

func (service *PegawaiServiceImpl) Update(ctx context.Context, pegawaiData web.PegawaiUpdateRequest) (web.PegawaiResponse, error) {
	err := service.Validator.Struct(pegawaiData)
	if err != nil {
		return web.PegawaiResponse{}, err
	}

	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		return web.PegawaiResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	pegawaiDomain := domain.Pegawai{
		Id:      pegawaiData.Id,
		Nama:    pegawaiData.Nama,
		Nip:     pegawaiData.Nip,
		Jabatan: pegawaiData.Jabatan,
		KodeOpd: pegawaiData.KodeOpd,
		NamaOpd: pegawaiData.NamaOpd,
	}

	pegawaiDomain, err = service.PegawaiRepository.Update(ctx, tx, pegawaiDomain)
	if err != nil {
		return web.PegawaiResponse{}, err
	}

	return web.PegawaiResponse{
		Id:      pegawaiDomain.Id,
		Nama:    pegawaiDomain.Nama,
		Nip:     pegawaiDomain.Nip,
		Jabatan: pegawaiDomain.Jabatan,
		KodeOpd: pegawaiDomain.KodeOpd,
		NamaOpd: pegawaiDomain.NamaOpd,
	}, nil
}

func (service *PegawaiServiceImpl) Delete(ctx context.Context, id int) error {
	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer helper.CommitOrRollback(tx)

	err = service.PegawaiRepository.Delete(ctx, tx, id)
	if err != nil {
		return err
	}

	return nil
}

func (service *PegawaiServiceImpl) FindById(ctx context.Context, id int) (web.PegawaiResponse, error) {
	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		return web.PegawaiResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	pegawai, err := service.PegawaiRepository.FindById(ctx, tx, id)
	if err != nil {
		return web.PegawaiResponse{}, err
	}

	return web.PegawaiResponse{
		Id:      pegawai.Id,
		Nama:    pegawai.Nama,
		Nip:     pegawai.Nip,
		Jabatan: pegawai.Jabatan,
		KodeOpd: pegawai.KodeOpd,
		NamaOpd: pegawai.NamaOpd,
	}, nil
}

func (service *PegawaiServiceImpl) FindAll(ctx context.Context) ([]web.PegawaiResponse, error) {
	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		return []web.PegawaiResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	pegawaiList, err := service.PegawaiRepository.FindAll(ctx, tx)
	if err != nil {
		return []web.PegawaiResponse{}, err
	}

	var responses []web.PegawaiResponse
	for _, pegawai := range pegawaiList {
		responses = append(responses, web.PegawaiResponse{
			Id:      pegawai.Id,
			Nama:    pegawai.Nama,
			Nip:     pegawai.Nip,
			Jabatan: pegawai.Jabatan,
			KodeOpd: pegawai.KodeOpd,
			NamaOpd: pegawai.NamaOpd,
		})
	}

	return responses, nil
}

func (service *PegawaiServiceImpl) FindByKodeOpd(ctx context.Context, kodeOpd string) ([]web.PegawaiResponse, error) {
	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		return []web.PegawaiResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	pegawaiList, err := service.PegawaiRepository.FindByKodeOpd(ctx, tx, kodeOpd)
	if err != nil {
		return []web.PegawaiResponse{}, err
	}

	var responses []web.PegawaiResponse
	for _, pegawai := range pegawaiList {
		responses = append(responses, web.PegawaiResponse{
			Id:      pegawai.Id,
			Nama:    pegawai.Nama,
			Nip:     pegawai.Nip,
			Jabatan: pegawai.Jabatan,
			KodeOpd: pegawai.KodeOpd,
			NamaOpd: pegawai.NamaOpd,
		})
	}

	return responses, nil
}
