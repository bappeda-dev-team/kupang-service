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
		Nama: pegawai.Nama,
		Nip:  pegawai.Nip,
		JabatanId: sql.NullInt64{
			Int64: 0,
			Valid: false,
		},
		NamaJabatan: sql.NullString{
			String: "",
			Valid:  false,
		},
		KodeOpd:      pegawai.KodeOpd,
		NamaOpd:      pegawai.NamaOpd,
		JenisPegawai: ptrToNullString(pegawai.JenisPegawai),
	}

	pegawaiDomain, err = service.PegawaiRepository.Create(ctx, tx, pegawaiDomain)
	if err != nil {
		return web.PegawaiResponse{}, err
	}

	return web.PegawaiResponse{
		Id:           pegawaiDomain.Id,
		Nama:         pegawaiDomain.Nama,
		Nip:          pegawaiDomain.Nip,
		JabatanId:    nil,
		NamaJabatan:  nullStringToPtr(pegawaiDomain.NamaJabatan),
		KodeOpd:      pegawaiDomain.KodeOpd,
		NamaOpd:      pegawaiDomain.NamaOpd,
		JenisPegawai: nullStringToPtr(pegawaiDomain.JenisPegawai),
	}, nil
}

func (service *PegawaiServiceImpl) AddJabatan(ctx context.Context, request web.PegawaiAddJabatanRequest) (web.PegawaiResponse, error) {
	err := service.Validator.Struct(request)
	if err != nil {
		return web.PegawaiResponse{}, err
	}

	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		return web.PegawaiResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	pegawai, err := service.PegawaiRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		return web.PegawaiResponse{}, err
	}

	jabatanId, err := service.ensureJabatan(ctx, tx, request.NamaJabatan)
	if err != nil {
		return web.PegawaiResponse{}, err
	}

	pegawai.JabatanId = sql.NullInt64{Int64: jabatanId, Valid: true}
	pegawai.NamaJabatan = sql.NullString{String: request.NamaJabatan, Valid: true}

	pegawai, err = service.PegawaiRepository.Update(ctx, tx, pegawai)
	if err != nil {
		return web.PegawaiResponse{}, err
	}

	return web.PegawaiResponse{
		Id:           pegawai.Id,
		Nama:         pegawai.Nama,
		Nip:          pegawai.Nip,
		JabatanId:    nullIntToPtr(pegawai.JabatanId),
		NamaJabatan:  nullStringToPtr(pegawai.NamaJabatan),
		KodeOpd:      pegawai.KodeOpd,
		NamaOpd:      pegawai.NamaOpd,
		JenisPegawai: nullStringToPtr(pegawai.JenisPegawai),
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

	namaJabatan, err := service.fetchNamaJabatan(ctx, tx, int64(pegawaiData.JabatanId))
	if err != nil {
		return web.PegawaiResponse{}, err
	}

	pegawaiDomain := domain.Pegawai{
		Id:   pegawaiData.Id,
		Nama: pegawaiData.Nama,
		Nip:  pegawaiData.Nip,
		JabatanId: sql.NullInt64{
			Int64: int64(pegawaiData.JabatanId),
			Valid: true,
		},
		NamaJabatan: sql.NullString{
			String: namaJabatan,
			Valid:  true,
		},
		KodeOpd:      pegawaiData.KodeOpd,
		NamaOpd:      pegawaiData.NamaOpd,
		JenisPegawai: ptrToNullString(pegawaiData.JenisPegawai),
	}

	pegawaiDomain, err = service.PegawaiRepository.Update(ctx, tx, pegawaiDomain)
	if err != nil {
		return web.PegawaiResponse{}, err
	}

	return web.PegawaiResponse{
		Id:           pegawaiDomain.Id,
		Nama:         pegawaiDomain.Nama,
		Nip:          pegawaiDomain.Nip,
		JabatanId:    nullIntToPtr(pegawaiDomain.JabatanId),
		NamaJabatan:  nullStringToPtr(pegawaiDomain.NamaJabatan),
		KodeOpd:      pegawaiDomain.KodeOpd,
		NamaOpd:      pegawaiDomain.NamaOpd,
		JenisPegawai: nullStringToPtr(pegawaiDomain.JenisPegawai),
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
		Id:           pegawai.Id,
		Nama:         pegawai.Nama,
		Nip:          pegawai.Nip,
		JabatanId:    nullIntToPtr(pegawai.JabatanId),
		NamaJabatan:  nullStringToPtr(pegawai.NamaJabatan),
		KodeOpd:      pegawai.KodeOpd,
		NamaOpd:      pegawai.NamaOpd,
		JenisPegawai: nullStringToPtr(pegawai.JenisPegawai),
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
			Id:           pegawai.Id,
			Nama:         pegawai.Nama,
			Nip:          pegawai.Nip,
			JabatanId:    nullIntToPtr(pegawai.JabatanId),
			NamaJabatan:  nullStringToPtr(pegawai.NamaJabatan),
			KodeOpd:      pegawai.KodeOpd,
			NamaOpd:      pegawai.NamaOpd,
			JenisPegawai: nullStringToPtr(pegawai.JenisPegawai),
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
			Id:           pegawai.Id,
			Nama:         pegawai.Nama,
			Nip:          pegawai.Nip,
			JabatanId:    nullIntToPtr(pegawai.JabatanId),
			NamaJabatan:  nullStringToPtr(pegawai.NamaJabatan),
			KodeOpd:      pegawai.KodeOpd,
			NamaOpd:      pegawai.NamaOpd,
			JenisPegawai: nullStringToPtr(pegawai.JenisPegawai),
		})
	}

	return responses, nil
}

func nullIntToPtr(value sql.NullInt64) *int {
	if value.Valid {
		v := int(value.Int64)
		return &v
	}

	return nil
}

func nullStringToPtr(value sql.NullString) *string {
	if value.Valid {
		v := value.String
		return &v
	}

	return nil
}

func ptrToNullString(value *string) sql.NullString {
	if value == nil {
		return sql.NullString{String: "", Valid: false}
	}

	return sql.NullString{String: *value, Valid: true}
}

func (service *PegawaiServiceImpl) ensureJabatan(ctx context.Context, tx *sql.Tx, namaJabatan string) (int64, error) {
	var id int64
	query := "INSERT INTO jabatan (nama_jabatan) VALUES ($1) ON CONFLICT (nama_jabatan) DO UPDATE SET nama_jabatan = EXCLUDED.nama_jabatan RETURNING id"
	err := tx.QueryRowContext(ctx, query, namaJabatan).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (service *PegawaiServiceImpl) fetchNamaJabatan(ctx context.Context, tx *sql.Tx, jabatanId int64) (string, error) {
	var nama string
	err := tx.QueryRowContext(ctx, "SELECT nama_jabatan FROM jabatan WHERE id = $1", jabatanId).Scan(&nama)
	if err != nil {
		return "", err
	}

	return nama, nil
}
