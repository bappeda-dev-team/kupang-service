package controller

import (
	"kupang-service/helper"
	"kupang-service/model/web"
	"kupang-service/service"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type TargetPokinOpdOperationalNControllerImpl struct {
	TargetPokinOpdOperationalNService service.TargetPokinOpdOperationalNService
}

func NewTargetPokinOpdOperationalNControllerImpl(targetService service.TargetPokinOpdOperationalNService) *TargetPokinOpdOperationalNControllerImpl {
	return &TargetPokinOpdOperationalNControllerImpl{
		TargetPokinOpdOperationalNService: targetService,
	}
}

// @Summary Create Target Pokin OPD Operational N
// @Description Create new Target Pokin OPD Operational N
// @Tags Target Pokin OPD Operational N
// @Accept json
// @Produce json
// @Param data body web.TargetPokinOpdOperationalNCreateRequest true "Target Pokin OPD Operational N Create Request"
// @Success 201 {object} web.WebResponse{data=web.TargetPokinOpdOperationalNResponse} "Created"
// @Failure 400 {object} web.WebResponse "Bad Request"
// @Failure 500 {object} web.WebResponse "Internal Server Error"
// @Router /target-pokin-opd-operational-ns [post]
func (controller *TargetPokinOpdOperationalNControllerImpl) Create(c echo.Context) error {
	request := web.TargetPokinOpdOperationalNCreateRequest{}
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD_REQUEST",
		})
	}

	response, err := controller.TargetPokinOpdOperationalNService.Create(c.Request().Context(), request)
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

// @Summary Update Target Pokin OPD Operational N
// @Description Update existing Target Pokin OPD Operational N by ID
// @Tags Target Pokin OPD Operational N
// @Accept json
// @Produce json
// @Param id path int true "Target Pokin OPD Operational N ID"
// @Param data body web.TargetPokinOpdOperationalNUpdateRequest true "Target Pokin OPD Operational N Update Request"
// @Success 200 {object} web.WebResponse{data=web.TargetPokinOpdOperationalNResponse} "OK"
// @Failure 400 {object} web.WebResponse "Bad Request"
// @Failure 500 {object} web.WebResponse "Internal Server Error"
// @Router /target-pokin-opd-operational-ns/{id} [put]
func (controller *TargetPokinOpdOperationalNControllerImpl) Update(c echo.Context) error {
	request := web.TargetPokinOpdOperationalNUpdateRequest{}
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

	response, err := controller.TargetPokinOpdOperationalNService.Update(c.Request().Context(), request)
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

// @Summary Delete Target Pokin OPD Operational N
// @Description Delete existing Target Pokin OPD Operational N by ID
// @Tags Target Pokin OPD Operational N
// @Accept json
// @Produce json
// @Param id path int true "Target Pokin OPD Operational N ID"
// @Success 200 {object} web.WebResponse "OK"
// @Failure 400 {object} web.WebResponse "Bad Request"
// @Failure 500 {object} web.WebResponse "Internal Server Error"
// @Router /target-pokin-opd-operational-ns/{id} [delete]
func (controller *TargetPokinOpdOperationalNControllerImpl) Delete(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD_REQUEST",
		})
	}

	if err := controller.TargetPokinOpdOperationalNService.Delete(c.Request().Context(), id); err != nil {
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

// @Summary Get Target Pokin OPD Operational N by ID
// @Description Get Target Pokin OPD Operational N detail by ID
// @Tags Target Pokin OPD Operational N
// @Accept json
// @Produce json
// @Param id path int true "Target Pokin OPD Operational N ID"
// @Success 200 {object} web.WebResponse{data=web.TargetPokinOpdOperationalNResponse} "OK"
// @Failure 400 {object} web.WebResponse "Bad Request"
// @Failure 404 {object} web.WebResponse "Not Found"
// @Failure 500 {object} web.WebResponse "Internal Server Error"
// @Router /target-pokin-opd-operational-ns/{id} [get]
func (controller *TargetPokinOpdOperationalNControllerImpl) FindById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD_REQUEST",
		})
	}

	response, err := controller.TargetPokinOpdOperationalNService.FindById(c.Request().Context(), id)
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

// @Summary List All Target Pokin OPD Operational N
// @Description Get list of all Target Pokin OPD Operational N
// @Tags Target Pokin OPD Operational N
// @Accept json
// @Produce json
// @Success 200 {object} web.WebResponse{data=[]web.TargetPokinOpdOperationalNResponse} "OK"
// @Failure 500 {object} web.WebResponse "Internal Server Error"
// @Router /target-pokin-opd-operational-ns [get]
func (controller *TargetPokinOpdOperationalNControllerImpl) FindAll(c echo.Context) error {
	responses, err := controller.TargetPokinOpdOperationalNService.FindAll(c.Request().Context())
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
