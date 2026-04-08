package service

import (
	"context"
	"kupang-service/model/web"
)

type RoleService interface {
	Create(ctx context.Context, role web.RoleCreateRequest) (web.RoleResponse, error)
	Update(ctx context.Context, role web.RoleUpdateRequest) (web.RoleResponse, error)
	Delete(ctx context.Context, id int) error
	FindById(ctx context.Context, id int) (web.RoleResponse, error)
	FindAll(ctx context.Context) ([]web.RoleResponse, error)
}
