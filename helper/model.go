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

func ToTujuanPokinOpdResponse(tujuanPokinOpd domain.TujuanPokinOpd) web.TujuanPokinOpdResponse {
	return web.TujuanPokinOpdResponse{
		Id:                tujuanPokinOpd.Id,
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
