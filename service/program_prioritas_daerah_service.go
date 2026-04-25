package service

import (
	"context"
	"kupang-service/model/web"
)

type ProgramPrioritasDaerahService interface {
	Create(ctx context.Context, program web.ProgramPrioritasDaerahCreateRequest) (web.ProgramPrioritasDaerahResponse, error)
	Update(ctx context.Context, program web.ProgramPrioritasDaerahUpdateRequest) (web.ProgramPrioritasDaerahResponse, error)
	Delete(ctx context.Context, id int) error
	FindById(ctx context.Context, id int) (web.ProgramPrioritasDaerahResponse, error)
	FindAll(ctx context.Context) ([]web.ProgramPrioritasDaerahResponse, error)
	FindByTahunRange(ctx context.Context, tahunAwal, tahunAkhir string) ([]web.ProgramPrioritasDaerahResponse, error)
}
