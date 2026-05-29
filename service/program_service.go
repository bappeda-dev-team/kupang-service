package service

import (
	"context"
	"kupang-service/model/web"
)

type ProgramService interface {
	Create(ctx context.Context, program web.ProgramCreateRequest) (web.ProgramResponse, error)
	Update(ctx context.Context, program web.ProgramUpdateRequest) (web.ProgramResponse, error)
	Delete(ctx context.Context, id int) error
	FindById(ctx context.Context, id int) (web.ProgramResponse, error)
	FindAll(ctx context.Context) ([]web.ProgramResponse, error)
}
