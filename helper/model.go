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
