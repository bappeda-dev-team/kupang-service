package controller

import (
	"kupang-service/helper"
	"kupang-service/model/web"
	"kupang-service/service"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type TargetPokinOpdOperationalControllerImpl struct {
	TargetPokinOpdOperationalService service.TargetPokinOpdOperationalService
}

func NewTargetPokinOpdOperationalControllerImpl(targetService service.TargetPokinOpdOperationalService) *TargetPokinOpdOperationalControllerImpl {
	return &TargetPokinOpdOperationalControllerImpl{
		TargetPokinOpdOperationalService: targetService,
	}
}

// @Summary Create Target Pokin OPD Operational
// @Description Create new Target Pokin OPD Operational
// @Tags Target Pokin OPD Operational
// @Accept json
// @Produce json
// @Param data body web.TargetPokinOpdOperationalCreateRequest true "Target Pokin OPD Operational Create Request"
// @Success 201 {object} web.WebResponse{data=web.TargetPokinOpdOperationalResponse} "Created"
// @Failure 400 {object} web.WebResponse "Bad Request"
// @Failure 500 {object} web.WebResponse "Internal Server Error"
// @Router /target-pokin-opd-operationals [post]
func (controller *TargetPokinOpdOperationalControllerImpl) Create(c echo.Context) error {
	request := web.TargetPokinOpdOperationalCreateRequest{}
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD_REQUEST",
		})
	}

	response, err := controller.TargetPokinOpdOperationalService.Create(c.Request().Context(), request)
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

// @Summary Update Target Pokin OPD Operational
// @Description Update existing Target Pokin OPD Operational by ID
// @Tags Target Pokin OPD Operational
// @Accept json
// @Produce json
// @Param id path int true "Target Pokin OPD Operational ID"
// @Param data body web.TargetPokinOpdOperationalUpdateRequest true "Target Pokin OPD Operational Update Request"
// @Success 200 {object} web.WebResponse{data=web.TargetPokinOpdOperationalResponse} "OK"
// @Failure 400 {object} web.WebResponse "Bad Request"
// @Failure 500 {object} web.WebResponse "Internal Server Error"
// @Router /target-pokin-opd-operationals/{id} [put]
func (controller *TargetPokinOpdOperationalControllerImpl) Update(c echo.Context) error {
	request := web.TargetPokinOpdOperationalUpdateRequest{}
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

	response, err := controller.TargetPokinOpdOperationalService.Update(c.Request().Context(), request)
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

// @Summary Delete Target Pokin OPD Operational
// @Description Delete existing Target Pokin OPD Operational by ID
// @Tags Target Pokin OPD Operational
// @Accept json
// @Produce json
// @Param id path int true "Target Pokin OPD Operational ID"
// @Success 200 {object} web.WebResponse "OK"
// @Failure 400 {object} web.WebResponse "Bad Request"
// @Failure 500 {object} web.WebResponse "Internal Server Error"
// @Router /target-pokin-opd-operationals/{id} [delete]
func (controller *TargetPokinOpdOperationalControllerImpl) Delete(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD_REQUEST",
		})
	}

	if err := controller.TargetPokinOpdOperationalService.Delete(c.Request().Context(), id); err != nil {
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

// @Summary Get Target Pokin OPD Operational by ID
// @Description Get Target Pokin OPD Operational detail by ID
// @Tags Target Pokin OPD Operational
// @Accept json
// @Produce json
// @Param id path int true "Target Pokin OPD Operational ID"
// @Success 200 {object} web.WebResponse{data=web.TargetPokinOpdOperationalResponse} "OK"
// @Failure 400 {object} web.WebResponse "Bad Request"
// @Failure 404 {object} web.WebResponse "Not Found"
// @Failure 500 {object} web.WebResponse "Internal Server Error"
// @Router /target-pokin-opd-operationals/{id} [get]
func (controller *TargetPokinOpdOperationalControllerImpl) FindById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD_REQUEST",
		})
	}

	response, err := controller.TargetPokinOpdOperationalService.FindById(c.Request().Context(), id)
	if err != nil {
		if err.Error() == "id tidak ditemukan" {
			return c.JSON(http.StatusNotFound, web.WebResponse{
				Code:   http.StatusNotFound,
				Status: "NOT_FOUND",
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

// @Summary List All Target Pokin OPD Operational
// @Description Get list of all Target Pokin OPD Operational
// @Tags Target Pokin OPD Operational
// @Accept json
// @Produce json
// @Success 200 {object} web.WebResponse{data=[]web.TargetPokinOpdOperationalResponse} "OK"
// @Failure 500 {object} web.WebResponse "Internal Server Error"
// @Router /target-pokin-opd-operationals [get]
func (controller *TargetPokinOpdOperationalControllerImpl) FindAll(c echo.Context) error {
	responses, err := controller.TargetPokinOpdOperationalService.FindAll(c.Request().Context())
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
