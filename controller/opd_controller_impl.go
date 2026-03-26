package controller

import (
	"kupang-service/helper"
	"kupang-service/model/web"
	"kupang-service/service"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type OpdControllerImpl struct {
	OpdService service.OpdService
}

func NewOpdControllerImpl(opdService service.OpdService) *OpdControllerImpl {
	return &OpdControllerImpl{
		OpdService: opdService,
	}
}

// @Summary Create OPD
// @Description Create new OPD
// @Tags OPD
// @Accept json
// @Produce json
// @Param data body web.OpdCreateRequest true "OPD Create Request"
// @Success 201 {object} web.WebResponse{data=web.OpdResponse} "Created"
// @Failure 400 {object} web.WebResponse "Bad Request"
// @Failure 500 {object} web.WebResponse "Internal Server Error"
// @Router /opds [post]
func (controller *OpdControllerImpl) Create(c echo.Context) error {
	opdCreateRequest := web.OpdCreateRequest{}
	err := c.Bind(&opdCreateRequest)
	if err != nil {
		return c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD_REQUEST",
		})
	}

	opdResponse, err := controller.OpdService.Create(c.Request().Context(), opdCreateRequest)
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
		Data:   opdResponse,
	})
}

// @Summary Update OPD
// @Description Update existing OPD by ID
// @Tags OPD
// @Accept json
// @Produce json
// @Param id path int true "OPD ID"
// @Param data body web.OpdUpdateRequest true "OPD Update Request"
// @Success 200 {object} web.WebResponse{data=web.OpdResponse} "OK"
// @Failure 400 {object} web.WebResponse "Bad Request"
// @Failure 500 {object} web.WebResponse "Internal Server Error"
// @Router /opds/{id} [put]
func (controller *OpdControllerImpl) Update(c echo.Context) error {
	opdUpdateRequest := web.OpdUpdateRequest{}
	err := c.Bind(&opdUpdateRequest)
	if err != nil {
		return c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD_REQUEST",
		})
	}

	opdUpdateRequest.Id, err = strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD_REQUEST",
		})
	}

	opdResponse, err := controller.OpdService.Update(c.Request().Context(), opdUpdateRequest)
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
		Data:   opdResponse,
	})
}

// @Summary Delete OPD
// @Description Delete existing OPD by ID
// @Tags OPD
// @Accept json
// @Produce json
// @Param id path int true "OPD ID"
// @Success 200 {object} web.WebResponse{data=web.OpdResponse} "OK"
// @Failure 400 {object} web.WebResponse "Bad Request"
// @Failure 500 {object} web.WebResponse "Internal Server Error"
// @Router /opds/{id} [delete]
func (controller *OpdControllerImpl) Delete(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD_REQUEST",
		})
	}

	err = controller.OpdService.Delete(c.Request().Context(), id)
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

// @Summary Get OPD by ID
// @Description Get OPD detail by ID
// @Tags OPD
// @Accept json
// @Produce json
// @Param id path int true "OPD ID"
// @Success 200 {object} web.WebResponse{data=web.OpdResponse} "OK"
// @Failure 400 {object} web.WebResponse "Bad Request"
// @Failure 500 {object} web.WebResponse "Internal Server Error"
// @Router /opds/{id} [get]
func (controller *OpdControllerImpl) FindById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD_REQUEST",
		})
	}

	opdResponse, err := controller.OpdService.FindById(c.Request().Context(), id)
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
		Data:   opdResponse,
	})
}

// @Summary List All OPD
// @Description Get list of all OPD
// @Tags OPD
// @Accept json
// @Produce json
// @Success 200 {object} web.WebResponse{data=[]web.OpdResponse} "OK"
// @Failure 500 {object} web.WebResponse "Internal Server Error"
// @Router /opds [get]
func (controller *OpdControllerImpl) FindAll(c echo.Context) error {
	opdResponses, err := controller.OpdService.FindAll(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, web.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: "INTERNAL_SERVER_ERROR",
		})
	}

	return c.JSON(http.StatusOK, web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   opdResponses,
	})
}
