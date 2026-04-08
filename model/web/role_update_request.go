package web

type RoleUpdateRequest struct {
	Id   int    `json:"id" validate:"required"`
	Role string `json:"role" validate:"required"`
}
