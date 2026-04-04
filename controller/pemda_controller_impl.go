package controller

import (
	"kupang-service/helper"
	"kupang-service/model/web"
	"kupang-service/service"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type PemdaControllerImpl struct {
	PemdaService service.PemdaService
}

func NewPemdaControllerImpl(pemdaService service.PemdaService) *PemdaControllerImpl {
	return &PemdaControllerImpl{
		PemdaService: pemdaService,
	}
}

// @Summary Create Pemda
// @Description Create new Pemda
// @Tags Pemda
// @Accept json
// @Produce json
// @Param data body web.PemdaCreateRequest true "Pemda Create Request"
// @Success 201 {object} web.WebResponse{data=web.PemdaResponse} "Created"
// @Failure 400 {object} web.WebResponse "Bad Request"
// @Failure 500 {object} web.WebResponse "Internal Server Error"
// @Router /pemdas [post]
func (controller *PemdaControllerImpl) Create(c echo.Context) error {
	pemdaCreateRequest := web.PemdaCreateRequest{}
	err := c.Bind(&pemdaCreateRequest)
	if err != nil {
		return c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD_REQUEST",
		})
	}

	pemdaResponse, err := controller.PemdaService.Create(c.Request().Context(), pemdaCreateRequest)
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
		Data:   pemdaResponse,
	})
}

// @Summary Update Pemda
// @Description Update existing Pemda by ID
// @Tags Pemda
// @Accept json
// @Produce json
// @Param id path int true "Pemda ID"
// @Param data body web.PemdaUpdateRequest true "Pemda Update Request"
// @Success 200 {object} web.WebResponse{data=web.PemdaResponse} "OK"
// @Failure 400 {object} web.WebResponse "Bad Request"
// @Failure 500 {object} web.WebResponse "Internal Server Error"
// @Router /pemdas/{id} [put]
func (controller *PemdaControllerImpl) Update(c echo.Context) error {
	pemdaUpdateRequest := web.PemdaUpdateRequest{}
	err := c.Bind(&pemdaUpdateRequest)
	if err != nil {
		return c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD_REQUEST",
		})
	}

	pemdaUpdateRequest.Id, err = strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD_REQUEST",
		})
	}

	pemdaResponse, err := controller.PemdaService.Update(c.Request().Context(), pemdaUpdateRequest)
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
		Data:   pemdaResponse,
	})
}

// @Summary Delete Pemda
// @Description Delete existing Pemda by ID
// @Tags Pemda
// @Accept json
// @Produce json
// @Param id path int true "Pemda ID"
// @Success 200 {object} web.WebResponse{data=web.PemdaResponse} "OK"
// @Failure 400 {object} web.WebResponse "Bad Request"
// @Failure 500 {object} web.WebResponse "Internal Server Error"
// @Router /pemdas/{id} [delete]
func (controller *PemdaControllerImpl) Delete(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD_REQUEST",
		})
	}

	err = controller.PemdaService.Delete(c.Request().Context(), id)
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

// @Summary Get Pemda by ID
// @Description Get Pemda detail by ID
// @Tags Pemda
// @Accept json
// @Produce json
// @Param id path int true "Pemda ID"
// @Success 200 {object} web.WebResponse{data=web.PemdaResponse} "OK"
// @Failure 400 {object} web.WebResponse "Bad Request"
// @Failure 500 {object} web.WebResponse "Internal Server Error"
// @Router /pemdas/{id} [get]
func (controller *PemdaControllerImpl) FindById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD_REQUEST",
		})
	}

	pemdaResponse, err := controller.PemdaService.FindById(c.Request().Context(), id)
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
		Data:   pemdaResponse,
	})
}

// @Summary List All Pemda
// @Description Get list of all Pemda
// @Tags Pemda
// @Accept json
// @Produce json
// @Success 200 {object} web.WebResponse{data=[]web.PemdaResponse} "OK"
// @Failure 500 {object} web.WebResponse "Internal Server Error"
// @Router /pemdas [get]
func (controller *PemdaControllerImpl) FindAll(c echo.Context) error {
	pemdaResponses, err := controller.PemdaService.FindAll(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, web.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: "INTERNAL_SERVER_ERROR",
		})
	}

	return c.JSON(http.StatusOK, web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   pemdaResponses,
	})
}
