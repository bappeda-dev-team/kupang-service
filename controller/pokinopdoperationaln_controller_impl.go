package controller

import (
	"kupang-service/helper"
	"kupang-service/model/web"
	"kupang-service/service"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type PokinOpdOperationalNControllerImpl struct {
	PokinOpdOperationalNService service.PokinOpdOperationalNService
}

func NewPokinOpdOperationalNControllerImpl(pokinOpdOperationalNService service.PokinOpdOperationalNService) *PokinOpdOperationalNControllerImpl {
	return &PokinOpdOperationalNControllerImpl{
		PokinOpdOperationalNService: pokinOpdOperationalNService,
	}
}

// @Summary Create Pokin Opd Operational N
// @Description Create new Pokin Opd Operational N
// @Tags Pokin Opd Operational N
// @Accept json
// @Produce json
// @Param data body web.PokinOpdOperationalNCreateRequest true "Pokin Opd Operational N Create Request"
// @Success 201 {object} web.WebResponse{data=web.PokinOpdOperationalNResponse} "Created"
// @Failure 400 {object} web.WebResponse "Bad Request"
// @Failure 500 {object} web.WebResponse "Internal Server Error"
// @Router /pokin-opd-operational-ns [post]
func (controller *PokinOpdOperationalNControllerImpl) Create(c echo.Context) error {
	request := web.PokinOpdOperationalNCreateRequest{}
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD_REQUEST",
		})
	}

	response, err := controller.PokinOpdOperationalNService.Create(c.Request().Context(), request)
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
		Data:   response,
	})
}

// @Summary Update Pokin Opd Operational N
// @Description Update existing Pokin Opd Operational N by ID
// @Tags Pokin Opd Operational N
// @Accept json
// @Produce json
// @Param id path int true "Pokin Opd Operational N ID"
// @Param data body web.PokinOpdOperationalNUpdateRequest true "Pokin Opd Operational N Update Request"
// @Success 200 {object} web.WebResponse{data=web.PokinOpdOperationalNResponse} "OK"
// @Failure 400 {object} web.WebResponse "Bad Request"
// @Failure 500 {object} web.WebResponse "Internal Server Error"
// @Router /pokin-opd-operational-ns/{id} [put]
func (controller *PokinOpdOperationalNControllerImpl) Update(c echo.Context) error {
	request := web.PokinOpdOperationalNUpdateRequest{}
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD_REQUEST",
		})
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD_REQUEST",
		})
	}
	request.Id = id

	response, err := controller.PokinOpdOperationalNService.Update(c.Request().Context(), request)
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
		Data:   response,
	})
}

// @Summary Delete Pokin Opd Operational N
// @Description Delete existing Pokin Opd Operational N by ID
// @Tags Pokin Opd Operational N
// @Accept json
// @Produce json
// @Param id path int true "Pokin Opd Operational N ID"
// @Success 200 {object} web.WebResponse{data=web.PokinOpdOperationalNResponse} "OK"
// @Failure 400 {object} web.WebResponse "Bad Request"
// @Failure 500 {object} web.WebResponse "Internal Server Error"
// @Router /pokin-opd-operational-ns/{id} [delete]
func (controller *PokinOpdOperationalNControllerImpl) Delete(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD_REQUEST",
		})
	}

	if err := controller.PokinOpdOperationalNService.Delete(c.Request().Context(), id); err != nil {
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

// @Summary Get Pokin Opd Operational N by ID
// @Description Get Pokin Opd Operational N detail by ID
// @Tags Pokin Opd Operational N
// @Accept json
// @Produce json
// @Param id path int true "Pokin Opd Operational N ID"
// @Success 200 {object} web.WebResponse{data=web.PokinOpdOperationalNResponse} "OK"
// @Failure 400 {object} web.WebResponse "Bad Request"
// @Failure 500 {object} web.WebResponse "Internal Server Error"
// @Router /pokin-opd-operational-ns/{id} [get]
func (controller *PokinOpdOperationalNControllerImpl) FindById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD_REQUEST",
		})
	}

	response, err := controller.PokinOpdOperationalNService.FindById(c.Request().Context(), id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, web.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: "INTERNAL_SERVER_ERROR",
		})
	}

	return c.JSON(http.StatusOK, web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   response,
	})
}

// @Summary List All Pokin Opd Operational N
// @Description Get list of all Pokin Opd Operational N
// @Tags Pokin Opd Operational N
// @Accept json
// @Produce json
// @Success 200 {object} web.WebResponse{data=[]web.PokinOpdOperationalNResponse} "OK"
// @Failure 500 {object} web.WebResponse "Internal Server Error"
// @Router /pokin-opd-operational-ns [get]
func (controller *PokinOpdOperationalNControllerImpl) FindAll(c echo.Context) error {
	responses, err := controller.PokinOpdOperationalNService.FindAll(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, web.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: "INTERNAL_SERVER_ERROR",
		})
	}

	return c.JSON(http.StatusOK, web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   responses,
	})
}
