package controller

import (
	"kupang-service/helper"
	"kupang-service/model/web"
	"kupang-service/service"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type IndikatorPokinOpdTacticalControllerImpl struct {
	IndikatorPokinOpdTacticalService service.IndikatorPokinOpdTacticalService
}

func NewIndikatorPokinOpdTacticalControllerImpl(indikatorService service.IndikatorPokinOpdTacticalService) *IndikatorPokinOpdTacticalControllerImpl {
	return &IndikatorPokinOpdTacticalControllerImpl{
		IndikatorPokinOpdTacticalService: indikatorService,
	}
}

// @Summary Create Indikator Pokin OPD Tactical
// @Description Create new Indikator Pokin OPD Tactical
// @Tags Indikator Pokin OPD Tactical
// @Accept json
// @Produce json
// @Param data body web.IndikatorPokinOpdTacticalCreateRequest true "Indikator Pokin OPD Tactical Create Request"
// @Success 201 {object} web.WebResponse{data=web.IndikatorPokinOpdTacticalResponse} "Created"
// @Failure 400 {object} web.WebResponse "Bad Request"
// @Failure 500 {object} web.WebResponse "Internal Server Error"
// @Router /indikator-pokin-opd-tacticals [post]
func (controller *IndikatorPokinOpdTacticalControllerImpl) Create(c echo.Context) error {
	createRequest := web.IndikatorPokinOpdTacticalCreateRequest{}
	err := c.Bind(&createRequest)
	if err != nil {
		return c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD_REQUEST",
		})
	}

	response, err := controller.IndikatorPokinOpdTacticalService.Create(c.Request().Context(), createRequest)
	if err != nil {
		if err.Error() == "pokin_opd_tactical_id tidak ditemukan" {
			return c.JSON(http.StatusBadRequest, web.WebResponse{
				Code:   http.StatusBadRequest,
				Status: "BAD_REQUEST",
				Data:   err.Error(),
			})
		}
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

// @Summary Update Indikator Pokin OPD Tactical
// @Description Update existing Indikator Pokin OPD Tactical by ID
// @Tags Indikator Pokin OPD Tactical
// @Accept json
// @Produce json
// @Param id path int true "Indikator Pokin OPD Tactical ID"
// @Param data body web.IndikatorPokinOpdTacticalUpdateRequest true "Indikator Pokin OPD Tactical Update Request"
// @Success 200 {object} web.WebResponse{data=web.IndikatorPokinOpdTacticalResponse} "OK"
// @Failure 400 {object} web.WebResponse "Bad Request"
// @Failure 500 {object} web.WebResponse "Internal Server Error"
// @Router /indikator-pokin-opd-tacticals/{id} [put]
func (controller *IndikatorPokinOpdTacticalControllerImpl) Update(c echo.Context) error {
	updateRequest := web.IndikatorPokinOpdTacticalUpdateRequest{}
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

	response, err := controller.IndikatorPokinOpdTacticalService.Update(c.Request().Context(), updateRequest)
	if err != nil {
		if err.Error() == "pokin_opd_tactical_id tidak ditemukan" {
			return c.JSON(http.StatusBadRequest, web.WebResponse{
				Code:   http.StatusBadRequest,
				Status: "BAD_REQUEST",
				Data:   err.Error(),
			})
		}
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

// @Summary Delete Indikator Pokin OPD Tactical
// @Description Delete existing Indikator Pokin OPD Tactical by ID
// @Tags Indikator Pokin OPD Tactical
// @Accept json
// @Produce json
// @Param id path int true "Indikator Pokin OPD Tactical ID"
// @Success 200 {object} web.WebResponse "OK"
// @Failure 400 {object} web.WebResponse "Bad Request"
// @Failure 500 {object} web.WebResponse "Internal Server Error"
// @Router /indikator-pokin-opd-tacticals/{id} [delete]
func (controller *IndikatorPokinOpdTacticalControllerImpl) Delete(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD_REQUEST",
		})
	}

	err = controller.IndikatorPokinOpdTacticalService.Delete(c.Request().Context(), id)
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

// @Summary Get Indikator Pokin OPD Tactical by ID
// @Description Get Indikator Pokin OPD Tactical detail by ID
// @Tags Indikator Pokin OPD Tactical
// @Accept json
// @Produce json
// @Param id path int true "Indikator Pokin OPD Tactical ID"
// @Success 200 {object} web.WebResponse{data=web.IndikatorPokinOpdTacticalResponse} "OK"
// @Failure 400 {object} web.WebResponse "Bad Request"
// @Failure 500 {object} web.WebResponse "Internal Server Error"
// @Router /indikator-pokin-opd-tacticals/{id} [get]
func (controller *IndikatorPokinOpdTacticalControllerImpl) FindById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD_REQUEST",
		})
	}

	response, err := controller.IndikatorPokinOpdTacticalService.FindById(c.Request().Context(), id)
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

// @Summary List All Indikator Pokin OPD Tactical
// @Description Get list of all Indikator Pokin OPD Tactical
// @Tags Indikator Pokin OPD Tactical
// @Accept json
// @Produce json
// @Success 200 {object} web.WebResponse{data=[]web.IndikatorPokinOpdTacticalResponse} "OK"
// @Failure 500 {object} web.WebResponse "Internal Server Error"
// @Router /indikator-pokin-opd-tacticals [get]
func (controller *IndikatorPokinOpdTacticalControllerImpl) FindAll(c echo.Context) error {
	responses, err := controller.IndikatorPokinOpdTacticalService.FindAll(c.Request().Context())
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
