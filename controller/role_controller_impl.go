package controller

import (
	"kupang-service/helper"
	"kupang-service/model/web"
	"kupang-service/service"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type RoleControllerImpl struct {
	RoleService service.RoleService
}

func NewRoleControllerImpl(roleService service.RoleService) *RoleControllerImpl {
	return &RoleControllerImpl{
		RoleService: roleService,
	}
}

// @Summary Create Role
// @Description Create new Role
// @Tags Role
// @Accept json
// @Produce json
// @Param data body web.RoleCreateRequest true "Role Create Request"
// @Success 201 {object} web.WebResponse{data=web.RoleResponse} "Created"
// @Failure 400 {object} web.WebResponse "Bad Request"
// @Failure 500 {object} web.WebResponse "Internal Server Error"
// @Router /roles [post]
func (controller *RoleControllerImpl) Create(c echo.Context) error {
	roleCreateRequest := web.RoleCreateRequest{}
	err := c.Bind(&roleCreateRequest)
	if err != nil {
		return c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD_REQUEST",
		})
	}

	roleResponse, err := controller.RoleService.Create(c.Request().Context(), roleCreateRequest)
	if err != nil {
		if helper.IsValidationError(err) {
			return c.JSON(http.StatusBadRequest, web.WebResponse{
				Code:   http.StatusBadRequest,
				Status: "BAD_REQUEST",
				Data:   err.Error(),
			})
		}
		return c.JSON(http.StatusInternalServerError, web.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: "INTERNAL_SERVER_ERROR",
		})
	}

	return c.JSON(http.StatusCreated, web.WebResponse{
		Code:   http.StatusCreated,
		Status: "CREATED",
		Data:   roleResponse,
	})
}

// @Summary Update Role
// @Description Update existing Role by ID
// @Tags Role
// @Accept json
// @Produce json
// @Param id path int true "Role ID"
// @Param data body web.RoleUpdateRequest true "Role Update Request"
// @Success 200 {object} web.WebResponse{data=web.RoleResponse} "OK"
// @Failure 400 {object} web.WebResponse "Bad Request"
// @Failure 500 {object} web.WebResponse "Internal Server Error"
// @Router /roles/{id} [put]
func (controller *RoleControllerImpl) Update(c echo.Context) error {
	roleUpdateRequest := web.RoleUpdateRequest{}
	err := c.Bind(&roleUpdateRequest)
	if err != nil {
		return c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD_REQUEST",
		})
	}

	roleUpdateRequest.Id, err = strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD_REQUEST",
		})
	}

	roleResponse, err := controller.RoleService.Update(c.Request().Context(), roleUpdateRequest)
	if err != nil {
		if helper.IsValidationError(err) {
			return c.JSON(http.StatusBadRequest, web.WebResponse{
				Code:   http.StatusBadRequest,
				Status: "BAD_REQUEST",
				Data:   err.Error(),
			})
		}
		return c.JSON(http.StatusInternalServerError, web.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: "INTERNAL_SERVER_ERROR",
		})
	}

	return c.JSON(http.StatusOK, web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   roleResponse,
	})
}

// @Summary Delete Role
// @Description Delete existing Role by ID
// @Tags Role
// @Accept json
// @Produce json
// @Param id path int true "Role ID"
// @Success 200 {object} web.WebResponse "OK"
// @Failure 400 {object} web.WebResponse "Bad Request"
// @Failure 500 {object} web.WebResponse "Internal Server Error"
// @Router /roles/{id} [delete]
func (controller *RoleControllerImpl) Delete(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD_REQUEST",
		})
	}

	err = controller.RoleService.Delete(c.Request().Context(), id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, web.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: "INTERNAL_SERVER_ERROR",
		})
	}

	return c.JSON(http.StatusOK, web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
	})
}

// @Summary Get Role by ID
// @Description Get Role detail by ID
// @Tags Role
// @Accept json
// @Produce json
// @Param id path int true "Role ID"
// @Success 200 {object} web.WebResponse{data=web.RoleResponse} "OK"
// @Failure 400 {object} web.WebResponse "Bad Request"
// @Failure 500 {object} web.WebResponse "Internal Server Error"
// @Router /roles/{id} [get]
func (controller *RoleControllerImpl) FindById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD_REQUEST",
		})
	}

	roleResponse, err := controller.RoleService.FindById(c.Request().Context(), id)
	if err != nil {
		if err.Error() == "id tidak ditemukan" {
			return c.JSON(http.StatusBadRequest, web.WebResponse{
				Code:   http.StatusBadRequest,
				Status: "BAD_REQUEST",
				Data:   err.Error(),
			})
		}
		return c.JSON(http.StatusInternalServerError, web.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: "INTERNAL_SERVER_ERROR",
		})
	}

	return c.JSON(http.StatusOK, web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   roleResponse,
	})
}

// @Summary List All Role
// @Description Get list of all Role
// @Tags Role
// @Accept json
// @Produce json
// @Success 200 {object} web.WebResponse{data=[]web.RoleResponse} "OK"
// @Failure 500 {object} web.WebResponse "Internal Server Error"
// @Router /roles [get]
func (controller *RoleControllerImpl) FindAll(c echo.Context) error {
	roleResponses, err := controller.RoleService.FindAll(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, web.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: "INTERNAL_SERVER_ERROR",
		})
	}

	return c.JSON(http.StatusOK, web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   roleResponses,
	})
}
