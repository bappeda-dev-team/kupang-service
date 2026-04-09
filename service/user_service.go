package service

import (
	"context"

	"kupang-service/model/web"
)

type UserService interface {
	Create(ctx context.Context, user web.UserCreateRequest) (web.UserResponse, error)
	Update(ctx context.Context, user web.UserUpdateRequest) (web.UserResponse, error)
	Delete(ctx context.Context, id int) error
	FindById(ctx context.Context, id int) (web.UserResponse, error)
	FindAll(ctx context.Context) ([]web.UserResponse, error)
	FindByKodeOpd(ctx context.Context, kodeOpd string) ([]web.UserResponse, error)
	Search(ctx context.Context, nama, nip *string) ([]web.UserResponse, error)
}
