package controller

import (
	"kupang-service/helper"
	"kupang-service/model/web"
	"kupang-service/service"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type IndikatorPokinOpdOperationalControllerImpl struct {
	IndikatorPokinOpdOperationalService service.IndikatorPokinOpdOperationalService
}

func NewIndikatorPokinOpdOperationalControllerImpl(indikatorService service.IndikatorPokinOpdOperationalService) *IndikatorPokinOpdOperationalControllerImpl {
	return &IndikatorPokinOpdOperationalControllerImpl{
		IndikatorPokinOpdOperationalService: indikatorService,
	}
}

// @Summary Create Indikator Pokin OPD Operational
// @Description Create new Indikator Pokin OPD Operational
// @Tags Indikator Pokin OPD Operational
// @Accept json
// @Produce json
// @Param data body web.IndikatorPokinOpdOperationalCreateRequest true "Indikator Pokin OPD Operational Create Request"
// @Success 201 {object} web.WebResponse{data=web.IndikatorPokinOpdOperationalResponse} "Created"
// @Failure 400 {object} web.WebResponse "Bad Request"
// @Failure 500 {object} web.WebResponse "Internal Server Error"
// @Router /indikator-pokin-opd-operationals [post]
func (controller *IndikatorPokinOpdOperationalControllerImpl) Create(c echo.Context) error {
	createRequest := web.IndikatorPokinOpdOperationalCreateRequest{}
	err := c.Bind(&createRequest)
	if err != nil {
		return c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD_REQUEST",
		})
	}

	response, err := controller.IndikatorPokinOpdOperationalService.Create(c.Request().Context(), createRequest)
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

// @Summary Update Indikator Pokin OPD Operational
// @Description Update existing Indikator Pokin OPD Operational by ID
// @Tags Indikator Pokin OPD Operational
// @Accept json
// @Produce json
// @Param id path int true "Indikator Pokin OPD Operational ID"
// @Param data body web.IndikatorPokinOpdOperationalUpdateRequest true "Indikator Pokin OPD Operational Update Request"
// @Success 200 {object} web.WebResponse{data=web.IndikatorPokinOpdOperationalResponse} "OK"
// @Failure 400 {object} web.WebResponse "Bad Request"
// @Failure 500 {object} web.WebResponse "Internal Server Error"
// @Router /indikator-pokin-opd-operationals/{id} [put]
func (controller *IndikatorPokinOpdOperationalControllerImpl) Update(c echo.Context) error {
	updateRequest := web.IndikatorPokinOpdOperationalUpdateRequest{}
	err := c.Bind(&updateRequest)
	if err != nil {
		return c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD_REQUEST",
		})
	}

	updateRequest.Id, err = strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD_REQUEST",
		})
	}

	response, err := controller.IndikatorPokinOpdOperationalService.Update(c.Request().Context(), updateRequest)
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

// @Summary Delete Indikator Pokin OPD Operational
// @Description Delete existing Indikator Pokin OPD Operational by ID
// @Tags Indikator Pokin OPD Operational
// @Accept json
// @Produce json
// @Param id path int true "Indikator Pokin OPD Operational ID"
// @Success 200 {object} web.WebResponse "OK"
// @Failure 400 {object} web.WebResponse "Bad Request"
// @Failure 500 {object} web.WebResponse "Internal Server Error"
// @Router /indikator-pokin-opd-operationals/{id} [delete]
func (controller *IndikatorPokinOpdOperationalControllerImpl) Delete(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD_REQUEST",
		})
	}

	err = controller.IndikatorPokinOpdOperationalService.Delete(c.Request().Context(), id)
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

// @Summary Get Indikator Pokin OPD Operational by ID
// @Description Get Indikator Pokin OPD Operational detail by ID
// @Tags Indikator Pokin OPD Operational
// @Accept json
// @Produce json
// @Param id path int true "Indikator Pokin OPD Operational ID"
// @Success 200 {object} web.WebResponse{data=web.IndikatorPokinOpdOperationalResponse} "OK"
// @Failure 400 {object} web.WebResponse "Bad Request"
// @Failure 500 {object} web.WebResponse "Internal Server Error"
// @Router /indikator-pokin-opd-operationals/{id} [get]
func (controller *IndikatorPokinOpdOperationalControllerImpl) FindById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD_REQUEST",
		})
	}

	response, err := controller.IndikatorPokinOpdOperationalService.FindById(c.Request().Context(), id)
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

// @Summary List All Indikator Pokin OPD Operational
// @Description Get list of all Indikator Pokin OPD Operational
// @Tags Indikator Pokin OPD Operational
// @Accept json
// @Produce json
// @Success 200 {object} web.WebResponse{data=[]web.IndikatorPokinOpdOperationalResponse} "OK"
// @Failure 500 {object} web.WebResponse "Internal Server Error"
// @Router /indikator-pokin-opd-operationals [get]
func (controller *IndikatorPokinOpdOperationalControllerImpl) FindAll(c echo.Context) error {
	responses, err := controller.IndikatorPokinOpdOperationalService.FindAll(c.Request().Context())
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
