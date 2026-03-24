package controller

import (
	"kupang-service/helper"
	"kupang-service/model/web"
	"kupang-service/service"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type PokinOpdOperationalControllerImpl struct {
	PokinOpdOperationalService service.PokinOpdOperationalService
}

func NewPokinOpdOperationalControllerImpl(pokinOpdOperationalService service.PokinOpdOperationalService) *PokinOpdOperationalControllerImpl {
	return &PokinOpdOperationalControllerImpl{
		PokinOpdOperationalService: pokinOpdOperationalService,
	}
}

// @Summary Create Pokin Opd Operational
// @Description Create new Pokin Opd Operational
// @Tags Pokin Opd Operational
// @Accept json
// @Produce json
// @Param data body web.PokinOpdOperationalCreateRequest true "Pokin Opd Operational Create Request"
// @Success 201 {object} web.WebResponse{data=web.PokinOpdOperationalResponse} "Created"
// @Failure 400 {object} web.WebResponse "Bad Request"
// @Failure 500 {object} web.WebResponse "Internal Server Error"
// @Router /pokin-opd-operationals [post]
func (controller *PokinOpdOperationalControllerImpl) Create(c echo.Context) error {
	request := web.PokinOpdOperationalCreateRequest{}
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD_REQUEST",
		})
	}

	response, err := controller.PokinOpdOperationalService.Create(c.Request().Context(), request)
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

// @Summary Update Pokin Opd Operational
// @Description Update existing Pokin Opd Operational by ID
// @Tags Pokin Opd Operational
// @Accept json
// @Produce json
// @Param id path int true "Pokin Opd Operational ID"
// @Param data body web.PokinOpdOperationalUpdateRequest true "Pokin Opd Operational Update Request"
// @Success 200 {object} web.WebResponse{data=web.PokinOpdOperationalResponse} "OK"
// @Failure 400 {object} web.WebResponse "Bad Request"
// @Failure 500 {object} web.WebResponse "Internal Server Error"
// @Router /pokin-opd-operationals/{id} [put]
func (controller *PokinOpdOperationalControllerImpl) Update(c echo.Context) error {
	request := web.PokinOpdOperationalUpdateRequest{}
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

	response, err := controller.PokinOpdOperationalService.Update(c.Request().Context(), request)
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

// @Summary Delete Pokin Opd Operational
// @Description Delete existing Pokin Opd Operational by ID
// @Tags Pokin Opd Operational
// @Accept json
// @Produce json
// @Param id path int true "Pokin Opd Operational ID"
// @Success 200 {object} web.WebResponse{data=web.PokinOpdOperationalResponse} "OK"
// @Failure 400 {object} web.WebResponse "Bad Request"
// @Failure 500 {object} web.WebResponse "Internal Server Error"
// @Router /pokin-opd-operationals/{id} [delete]
func (controller *PokinOpdOperationalControllerImpl) Delete(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD_REQUEST",
		})
	}

	if err := controller.PokinOpdOperationalService.Delete(c.Request().Context(), id); err != nil {
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

// @Summary Get Pokin Opd Operational by ID
// @Description Get Pokin Opd Operational detail by ID
// @Tags Pokin Opd Operational
// @Accept json
// @Produce json
// @Param id path int true "Pokin Opd Operational ID"
// @Success 200 {object} web.WebResponse{data=web.PokinOpdOperationalResponse} "OK"
// @Failure 400 {object} web.WebResponse "Bad Request"
// @Failure 500 {object} web.WebResponse "Internal Server Error"
// @Router /pokin-opd-operationals/{id} [get]
func (controller *PokinOpdOperationalControllerImpl) FindById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD_REQUEST",
		})
	}

	response, err := controller.PokinOpdOperationalService.FindById(c.Request().Context(), id)
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

// @Summary List All Pokin Opd Operational
// @Description Get list of all Pokin Opd Operational
// @Tags Pokin Opd Operational
// @Accept json
// @Produce json
// @Success 200 {object} web.WebResponse{data=[]web.PokinOpdOperationalResponse} "OK"
// @Failure 500 {object} web.WebResponse "Internal Server Error"
// @Router /pokin-opd-operationals [get]
func (controller *PokinOpdOperationalControllerImpl) FindAll(c echo.Context) error {
	responses, err := controller.PokinOpdOperationalService.FindAll(c.Request().Context())
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
