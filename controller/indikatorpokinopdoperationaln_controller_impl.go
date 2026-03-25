package controller

import (
	"kupang-service/helper"
	"kupang-service/model/web"
	"kupang-service/service"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type IndikatorPokinOpdOperationalNControllerImpl struct {
	IndikatorPokinOpdOperationalNService service.IndikatorPokinOpdOperationalNService
}

func NewIndikatorPokinOpdOperationalNControllerImpl(indikatorService service.IndikatorPokinOpdOperationalNService) *IndikatorPokinOpdOperationalNControllerImpl {
	return &IndikatorPokinOpdOperationalNControllerImpl{
		IndikatorPokinOpdOperationalNService: indikatorService,
	}
}

// @Summary Create Indikator Pokin OPD Operational N
// @Description Create new Indikator Pokin OPD Operational N
// @Tags Indikator Pokin OPD Operational N
// @Accept json
// @Produce json
// @Param data body web.IndikatorPokinOpdOperationalNCreateRequest true "Indikator Pokin OPD Operational N Create Request"
// @Success 201 {object} web.WebResponse{data=web.IndikatorPokinOpdOperationalNResponse} "Created"
// @Failure 400 {object} web.WebResponse "Bad Request"
// @Failure 500 {object} web.WebResponse "Internal Server Error"
// @Router /indikator-pokin-opd-operational-ns [post]
func (controller *IndikatorPokinOpdOperationalNControllerImpl) Create(c echo.Context) error {
	createRequest := web.IndikatorPokinOpdOperationalNCreateRequest{}
	if err := c.Bind(&createRequest); err != nil {
		return c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD_REQUEST",
		})
	}

	response, err := controller.IndikatorPokinOpdOperationalNService.Create(c.Request().Context(), createRequest)
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

// @Summary Update Indikator Pokin OPD Operational N
// @Description Update existing Indikator Pokin OPD Operational N by ID
// @Tags Indikator Pokin OPD Operational N
// @Accept json
// @Produce json
// @Param id path int true "Indikator Pokin OPD Operational N ID"
// @Param data body web.IndikatorPokinOpdOperationalNUpdateRequest true "Indikator Pokin OPD Operational N Update Request"
// @Success 200 {object} web.WebResponse{data=web.IndikatorPokinOpdOperationalNResponse} "OK"
// @Failure 400 {object} web.WebResponse "Bad Request"
// @Failure 500 {object} web.WebResponse "Internal Server Error"
// @Router /indikator-pokin-opd-operational-ns/{id} [put]
func (controller *IndikatorPokinOpdOperationalNControllerImpl) Update(c echo.Context) error {
	updateRequest := web.IndikatorPokinOpdOperationalNUpdateRequest{}
	if err := c.Bind(&updateRequest); err != nil {
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
	updateRequest.Id = id

	response, err := controller.IndikatorPokinOpdOperationalNService.Update(c.Request().Context(), updateRequest)
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

// @Summary Delete Indikator Pokin OPD Operational N
// @Description Delete existing Indikator Pokin OPD Operational N by ID
// @Tags Indikator Pokin OPD Operational N
// @Accept json
// @Produce json
// @Param id path int true "Indikator Pokin OPD Operational N ID"
// @Success 200 {object} web.WebResponse "OK"
// @Failure 400 {object} web.WebResponse "Bad Request"
// @Failure 500 {object} web.WebResponse "Internal Server Error"
// @Router /indikator-pokin-opd-operational-ns/{id} [delete]
func (controller *IndikatorPokinOpdOperationalNControllerImpl) Delete(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD_REQUEST",
		})
	}

	if err := controller.IndikatorPokinOpdOperationalNService.Delete(c.Request().Context(), id); err != nil {
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

// @Summary Get Indikator Pokin OPD Operational N by ID
// @Description Get Indikator Pokin OPD Operational N detail by ID
// @Tags Indikator Pokin OPD Operational N
// @Accept json
// @Produce json
// @Param id path int true "Indikator Pokin OPD Operational N ID"
// @Success 200 {object} web.WebResponse{data=web.IndikatorPokinOpdOperationalNResponse} "OK"
// @Failure 400 {object} web.WebResponse "Bad Request"
// @Failure 500 {object} web.WebResponse "Internal Server Error"
// @Router /indikator-pokin-opd-operational-ns/{id} [get]
func (controller *IndikatorPokinOpdOperationalNControllerImpl) FindById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD_REQUEST",
		})
	}

	response, err := controller.IndikatorPokinOpdOperationalNService.FindById(c.Request().Context(), id)
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

// @Summary List All Indikator Pokin OPD Operational N
// @Description Get list of all Indikator Pokin OPD Operational N
// @Tags Indikator Pokin OPD Operational N
// @Accept json
// @Produce json
// @Success 200 {object} web.WebResponse{data=[]web.IndikatorPokinOpdOperationalNResponse} "OK"
// @Failure 500 {object} web.WebResponse "Internal Server Error"
// @Router /indikator-pokin-opd-operational-ns [get]
func (controller *IndikatorPokinOpdOperationalNControllerImpl) FindAll(c echo.Context) error {
	responses, err := controller.IndikatorPokinOpdOperationalNService.FindAll(c.Request().Context())
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