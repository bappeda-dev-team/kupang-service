package helper

import (
	"kupang-service/model/domain"
	"kupang-service/model/web"
)

func ToPokinOpdResponse(pokinOpd domain.PokinOpd) web.PokinOpdResponse {
	return web.PokinOpdResponse{
		Id:      pokinOpd.Id,
		KodeOpd: pokinOpd.KodeOpd,
		NamaOpd: pokinOpd.NamaOpd,
		Tahun:   pokinOpd.Tahun,
	}
}

func ToPokinOpdResponses(pokinOpds []domain.PokinOpd) []web.PokinOpdResponse {
	var responses []web.PokinOpdResponse
	for _, pokinOpd := range pokinOpds {
		responses = append(responses, ToPokinOpdResponse(pokinOpd))
	}
	return responses
}

func ToOpdResponse(opd domain.Opd) web.OpdResponse {
	return web.OpdResponse{
		Id:                           opd.Id,
		KodeOpd:                      opd.KodeOpd,
		NamaOpd:                      opd.NamaOpd,
		KodeLembaga:                  opd.KodeLembaga,
		NamaKepalaPerangkatDaerah:    opd.NamaKepalaPerangkatDaerah,
		NipKepalaPerangkatDaerah:     opd.NipKepalaPerangkatDaerah,
		PangkatKepalaPerangkatDaerah: opd.PangkatKepalaPerangkatDaerah,
	}
}

func ToOpdResponses(opds []domain.Opd) []web.OpdResponse {
	var responses []web.OpdResponse
	for _, opd := range opds {
		responses = append(responses, ToOpdResponse(opd))
	}
	return responses
}

func ToLembagaResponse(lembaga domain.Lembaga) web.LembagaResponse {
	return web.LembagaResponse{
		Id:                   lembaga.Id,
		KodeLembaga:          lembaga.KodeLembaga,
		NamaLembaga:          lembaga.NamaLembaga,
		JabatanKepalaLembaga: lembaga.JabatanKepalaLembaga,
		NamaKepalaLembaga:    lembaga.NamaKepalaLembaga,
		NipKepalaLembaga:     lembaga.NipKepalaLembaga,
	}
}

func ToLembagaResponses(lembagas []domain.Lembaga) []web.LembagaResponse {
	var responses []web.LembagaResponse
	for _, lembaga := range lembagas {
		responses = append(responses, ToLembagaResponse(lembaga))
	}
	return responses
}

func ToPeriodeResponse(periode domain.Periode) web.PeriodeResponse {
	return web.PeriodeResponse{
		Id:           periode.Id,
		TahunAwal:    periode.TahunAwal,
		TahunAkhir:   periode.TahunAkhir,
		JenisPeriode: periode.JenisPeriode,
	}
}

func ToPeriodeResponses(periodes []domain.Periode) []web.PeriodeResponse {
	var responses []web.PeriodeResponse
	for _, periode := range periodes {
		responses = append(responses, ToPeriodeResponse(periode))
	}
	return responses
}

func ToRoleResponse(role domain.Role) web.RoleResponse {
	return web.RoleResponse{
		Id:   role.Id,
		Role: role.Role,
	}
}

func ToRoleResponses(roles []domain.Role) []web.RoleResponse {
	var responses []web.RoleResponse
	for _, role := range roles {
		responses = append(responses, ToRoleResponse(role))
	}
	return responses
}

func ToMusrenbangResponse(musrenbang domain.Musrenbang) web.MusrenbangResponse {
	return web.MusrenbangResponse{
		Id:      musrenbang.Id,
		Usulan:  musrenbang.Usulan,
		Alamat:  musrenbang.Alamat,
		Uraian:  musrenbang.Uraian,
		Tahun:   musrenbang.Tahun,
		KodeOpd: musrenbang.KodeOpd,
		NamaOpd: musrenbang.NamaOpd,
		Status:  musrenbang.Status,
	}
}

func ToMusrenbangResponses(musrenbangs []domain.Musrenbang) []web.MusrenbangResponse {
	var responses []web.MusrenbangResponse
	for _, musrenbang := range musrenbangs {
		responses = append(responses, ToMusrenbangResponse(musrenbang))
	}
	return responses
}

func ToUserResponse(user domain.User) web.UserResponse {
	var kodeOpd *string
	if user.KodeOpd.Valid {
		val := user.KodeOpd.String
		kodeOpd = &val
	}

	var nip *string
	if user.Nip.Valid {
		val := user.Nip.String
		nip = &val
	}

	var role *string
	if user.Role.Valid {
		val := user.Role.String
		role = &val
	}

	var opdId *int
	if user.OpdId.Valid {
		val := int(user.OpdId.Int64)
		opdId = &val
	}

	var pegawaiId *int
	if user.PegawaiId.Valid {
		val := int(user.PegawaiId.Int64)
		pegawaiId = &val
	}

	var roleId *int
	if user.RoleId.Valid {
		val := int(user.RoleId.Int64)
		roleId = &val
	}

	var namaOpd *string
	if user.NamaOpd.Valid {
		val := user.NamaOpd.String
		namaOpd = &val
	}

	var namaPegawai *string
	if user.NamaPegawai.Valid {
		val := user.NamaPegawai.String
		namaPegawai = &val
	}

	return web.UserResponse{
		Id:          user.Id,
		Nama:        user.Nama,
		Email:       user.Email,
		Status:      user.Status,
		Nip:         nip,
		KodeOpd:     kodeOpd,
		Role:        role,
		OpdId:       opdId,
		PegawaiId:   pegawaiId,
		RoleId:      roleId,
		NamaOpd:     namaOpd,
		NamaPegawai: namaPegawai,
	}
}

func ToUserResponses(users []domain.User) []web.UserResponse {
	var responses []web.UserResponse
	for _, user := range users {
		responses = append(responses, ToUserResponse(user))
	}
	return responses
}

func ToIndikatorPokinOpdResponse(indikatorPokinOpd domain.IndikatorPokinOpd) web.IndikatorPokinOpdResponse {
	return web.IndikatorPokinOpdResponse{
		Id:            indikatorPokinOpd.Id,
		NamaIndikator: indikatorPokinOpd.NamaIndikator,
	}
}

func ToIndikatorPokinOpdResponses(indikatorPokinOpds []domain.IndikatorPokinOpd) []web.IndikatorPokinOpdResponse {
	var responses []web.IndikatorPokinOpdResponse
	for _, indikatorPokinOpd := range indikatorPokinOpds {
		responses = append(responses, ToIndikatorPokinOpdResponse(indikatorPokinOpd))
	}
	return responses
}

func ToIndikatorPokinOpdStrategicResponse(indikator domain.IndikatorPokinOpdStrategic) web.IndikatorPokinOpdStrategicResponse {
	return web.IndikatorPokinOpdStrategicResponse{
		Id:                  indikator.Id,
		PokinOpdStrategicId: indikator.PokinOpdStrategicId,
		NamaIndikator:       indikator.NamaIndikator,
	}
}

func ToIndikatorPokinOpdStrategicResponses(indikators []domain.IndikatorPokinOpdStrategic) []web.IndikatorPokinOpdStrategicResponse {
	var responses []web.IndikatorPokinOpdStrategicResponse
	for _, indikator := range indikators {
		responses = append(responses, ToIndikatorPokinOpdStrategicResponse(indikator))
	}
	return responses
}

func ToIndikatorPokinOpdTacticalResponse(indikator domain.IndikatorPokinOpdTactical) web.IndikatorPokinOpdTacticalResponse {
	return web.IndikatorPokinOpdTacticalResponse{
		Id:                 indikator.Id,
		PokinOpdTacticalId: indikator.PokinOpdTacticalId,
		NamaIndikator:      indikator.NamaIndikator,
	}
}

func ToIndikatorPokinOpdTacticalResponses(indikators []domain.IndikatorPokinOpdTactical) []web.IndikatorPokinOpdTacticalResponse {
	var responses []web.IndikatorPokinOpdTacticalResponse
	for _, indikator := range indikators {
		responses = append(responses, ToIndikatorPokinOpdTacticalResponse(indikator))
	}
	return responses
}

func ToIndikatorPokinOpdOperationalResponse(indikator domain.IndikatorPokinOpdOperational) web.IndikatorPokinOpdOperationalResponse {
	return web.IndikatorPokinOpdOperationalResponse{
		Id:                    indikator.Id,
		PokinOpdOperationalId: indikator.PokinOpdOperationalId,
		NamaIndikator:         indikator.NamaIndikator,
	}
}

func ToIndikatorPokinOpdOperationalResponses(indikators []domain.IndikatorPokinOpdOperational) []web.IndikatorPokinOpdOperationalResponse {
	var responses []web.IndikatorPokinOpdOperationalResponse
	for _, indikator := range indikators {
		responses = append(responses, ToIndikatorPokinOpdOperationalResponse(indikator))
	}
	return responses
}

func ToIndikatorPokinOpdOperationalNResponse(indikator domain.IndikatorPokinOpdOperationalN) web.IndikatorPokinOpdOperationalNResponse {
	return web.IndikatorPokinOpdOperationalNResponse{
		Id:                     indikator.Id,
		PokinOpdOperationalNId: indikator.PokinOpdOperationalNId,
		NamaIndikator:          indikator.NamaIndikator,
	}
}

func ToIndikatorPokinOpdOperationalNResponses(indikators []domain.IndikatorPokinOpdOperationalN) []web.IndikatorPokinOpdOperationalNResponse {
	var responses []web.IndikatorPokinOpdOperationalNResponse
	for _, indikator := range indikators {
		responses = append(responses, ToIndikatorPokinOpdOperationalNResponse(indikator))
	}
	return responses
}

func ToTujuanPokinOpdResponse(tujuanPokinOpd domain.TujuanPokinOpd) web.TujuanPokinOpdResponse {
	return web.TujuanPokinOpdResponse{
		Id:                tujuanPokinOpd.Id,
		PokinOpdId:        tujuanPokinOpd.PokinOpdId,
		KodeOpd:           tujuanPokinOpd.KodeOpd,
		NamaTujuan:        tujuanPokinOpd.NamaTujuan,
		BidangUrusan:      tujuanPokinOpd.BidangUrusan,
		TahunAwalPeriode:  tujuanPokinOpd.TahunAwalPeriode,
		TahunAkhirPeriode: tujuanPokinOpd.TahunAkhirPeriode,
	}
}

func ToTujuanPokinOpdResponses(tujuanPokinOpds []domain.TujuanPokinOpd) []web.TujuanPokinOpdResponse {
	var responses []web.TujuanPokinOpdResponse
	for _, tujuanPokinOpd := range tujuanPokinOpds {
		responses = append(responses, ToTujuanPokinOpdResponse(tujuanPokinOpd))
	}
	return responses
}

func ToTargetPokinOpdResponse(targetPokinOpd domain.TargetPokinOpd) web.TargetPokinOpdResponse {
	return web.TargetPokinOpdResponse{
		Id:         targetPokinOpd.Id,
		NamaTarget: targetPokinOpd.NilaiTarget,
		Satuan:     targetPokinOpd.Satuan,
	}
}

func ToTargetPokinOpdResponses(targetPokinOpds []domain.TargetPokinOpd) []web.TargetPokinOpdResponse {
	var responses []web.TargetPokinOpdResponse
	for _, targetPokinOpd := range targetPokinOpds {
		responses = append(responses, ToTargetPokinOpdResponse(targetPokinOpd))
	}
	return responses
}

func ToTargetPokinOpdStrategicResponse(target domain.TargetPokinOpdStrategic) web.TargetPokinOpdStrategicResponse {
	return web.TargetPokinOpdStrategicResponse{
		Id:                  target.Id,
		IdPokinOpdStrategic: target.IndikatorPokinOpdStrategicId,
		NamaTarget:          target.NilaiTarget,
		Satuan:              target.Satuan,
	}
}

func ToTargetPokinOpdStrategicResponses(targets []domain.TargetPokinOpdStrategic) []web.TargetPokinOpdStrategicResponse {
	var responses []web.TargetPokinOpdStrategicResponse
	for _, target := range targets {
		responses = append(responses, ToTargetPokinOpdStrategicResponse(target))
	}
	return responses
}

func ToTargetPokinOpdTacticalResponse(target domain.TargetPokinOpdTactical) web.TargetPokinOpdTacticalResponse {
	return web.TargetPokinOpdTacticalResponse{
		Id:                          target.Id,
		IndikatorPokinOpdTacticalId: target.IndikatorPokinOpdTacticalId,
		NilaiTarget:                 target.NilaiTarget,
		Satuan:                      target.Satuan,
	}
}

func ToTargetPokinOpdTacticalResponses(targets []domain.TargetPokinOpdTactical) []web.TargetPokinOpdTacticalResponse {
	var responses []web.TargetPokinOpdTacticalResponse
	for _, target := range targets {
		responses = append(responses, ToTargetPokinOpdTacticalResponse(target))
	}
	return responses
}

func ToTargetPokinOpdOperationalResponse(target domain.TargetPokinOpdOperational) web.TargetPokinOpdOperationalResponse {
	return web.TargetPokinOpdOperationalResponse{
		Id:                             target.Id,
		IndikatorPokinOpdOperationalId: target.IndikatorPokinOpdOperationalId,
		NilaiTarget:                    target.NilaiTarget,
		Satuan:                         target.Satuan,
	}
}

func ToTargetPokinOpdOperationalResponses(targets []domain.TargetPokinOpdOperational) []web.TargetPokinOpdOperationalResponse {
	var responses []web.TargetPokinOpdOperationalResponse
	for _, target := range targets {
		responses = append(responses, ToTargetPokinOpdOperationalResponse(target))
	}
	return responses
}

func ToTargetPokinOpdOperationalNResponse(target domain.TargetPokinOpdOperationalN) web.TargetPokinOpdOperationalNResponse {
	return web.TargetPokinOpdOperationalNResponse{
		Id:                              target.Id,
		IndikatorPokinOpdOperationalNId: target.IndikatorPokinOpdOperationalNId,
		NilaiTarget:                     target.NilaiTarget,
		Satuan:                          target.Satuan,
	}
}

func ToTargetPokinOpdOperationalNResponses(targets []domain.TargetPokinOpdOperationalN) []web.TargetPokinOpdOperationalNResponse {
	var responses []web.TargetPokinOpdOperationalNResponse
	for _, target := range targets {
		responses = append(responses, ToTargetPokinOpdOperationalNResponse(target))
	}
	return responses
}
