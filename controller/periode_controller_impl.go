package controller

import (
	"kupang-service/helper"
	"kupang-service/model/web"
	"kupang-service/service"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type PeriodeControllerImpl struct {
	PeriodeService service.PeriodeService
}

func NewPeriodeControllerImpl(periodeService service.PeriodeService) *PeriodeControllerImpl {
	return &PeriodeControllerImpl{
		PeriodeService: periodeService,
	}
}

// @Summary Create Periode
// @Description Create new Periode
// @Tags Periode
// @Accept json
// @Produce json
// @Param data body web.PeriodeCreateRequest true "Periode Create Request"
// @Success 201 {object} web.WebResponse{data=web.PeriodeResponse} "Created"
// @Failure 400 {object} web.WebResponse "Bad Request"
// @Failure 500 {object} web.WebResponse "Internal Server Error"
// @Router /periodes [post]
func (controller *PeriodeControllerImpl) Create(c echo.Context) error {
	periodeCreateRequest := web.PeriodeCreateRequest{}
	err := c.Bind(&periodeCreateRequest)
	if err != nil {
		return c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD_REQUEST",
		})
	}

	periodeResponse, err := controller.PeriodeService.Create(c.Request().Context(), periodeCreateRequest)
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
		Data:   periodeResponse,
	})
}

// @Summary Update Periode
// @Description Update existing Periode by ID
// @Tags Periode
// @Accept json
// @Produce json
// @Param id path int true "Periode ID"
// @Param data body web.PeriodeUpdateRequest true "Periode Update Request"
// @Success 200 {object} web.WebResponse{data=web.PeriodeResponse} "OK"
// @Failure 400 {object} web.WebResponse "Bad Request"
// @Failure 500 {object} web.WebResponse "Internal Server Error"
// @Router /periodes/{id} [put]
func (controller *PeriodeControllerImpl) Update(c echo.Context) error {
	periodeUpdateRequest := web.PeriodeUpdateRequest{}
	err := c.Bind(&periodeUpdateRequest)
	if err != nil {
		return c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD_REQUEST",
		})
	}

	periodeUpdateRequest.Id, err = strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD_REQUEST",
		})
	}

	periodeResponse, err := controller.PeriodeService.Update(c.Request().Context(), periodeUpdateRequest)
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
		Data:   periodeResponse,
	})
}

// @Summary Delete Periode
// @Description Delete existing Periode by ID
// @Tags Periode
// @Accept json
// @Produce json
// @Param id path int true "Periode ID"
// @Success 200 {object} web.WebResponse "OK"
// @Failure 400 {object} web.WebResponse "Bad Request"
// @Failure 500 {object} web.WebResponse "Internal Server Error"
// @Router /periodes/{id} [delete]
func (controller *PeriodeControllerImpl) Delete(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD_REQUEST",
		})
	}

	err = controller.PeriodeService.Delete(c.Request().Context(), id)
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

// @Summary Get Periode by ID
// @Description Get Periode detail by ID
// @Tags Periode
// @Accept json
// @Produce json
// @Param id path int true "Periode ID"
// @Success 200 {object} web.WebResponse{data=web.PeriodeResponse} "OK"
// @Failure 400 {object} web.WebResponse "Bad Request"
// @Failure 500 {object} web.WebResponse "Internal Server Error"
// @Router /periodes/{id} [get]
func (controller *PeriodeControllerImpl) FindById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD_REQUEST",
		})
	}

	periodeResponse, err := controller.PeriodeService.FindById(c.Request().Context(), id)
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
		Data:   periodeResponse,
	})
}

// @Summary List All Periode
// @Description Get list of all Periode
// @Tags Periode
// @Accept json
// @Produce json
// @Success 200 {object} web.WebResponse{data=[]web.PeriodeResponse} "OK"
// @Failure 500 {object} web.WebResponse "Internal Server Error"
// @Router /periodes [get]
func (controller *PeriodeControllerImpl) FindAll(c echo.Context) error {
	periodeResponses, err := controller.PeriodeService.FindAll(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, web.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: "INTERNAL_SERVER_ERROR",
		})
	}

	return c.JSON(http.StatusOK, web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   periodeResponses,
	})
}
