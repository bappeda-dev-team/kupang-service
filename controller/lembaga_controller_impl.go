package controller

import (
	"kupang-service/helper"
	"kupang-service/model/web"
	"kupang-service/service"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type LembagaControllerImpl struct {
	LembagaService service.LembagaService
}

func NewLembagaControllerImpl(lembagaService service.LembagaService) *LembagaControllerImpl {
	return &LembagaControllerImpl{
		LembagaService: lembagaService,
	}
}

// @Summary Create Lembaga
// @Description Create new Lembaga
// @Tags Lembaga
// @Accept json
// @Produce json
// @Param data body web.LembagaCreateRequest true "Lembaga Create Request"
// @Success 201 {object} web.WebResponse{data=web.LembagaResponse} "Created"
// @Failure 400 {object} web.WebResponse "Bad Request"
// @Failure 500 {object} web.WebResponse "Internal Server Error"
// @Router /lembagas [post]
func (controller *LembagaControllerImpl) Create(c echo.Context) error {
	lembagaCreateRequest := web.LembagaCreateRequest{}
	err := c.Bind(&lembagaCreateRequest)
	if err != nil {
		return c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD_REQUEST",
		})
	}

	lembagaResponse, err := controller.LembagaService.Create(c.Request().Context(), lembagaCreateRequest)
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
		Data:   lembagaResponse,
	})
}

// @Summary Update Lembaga
// @Description Update existing Lembaga by ID
// @Tags Lembaga
// @Accept json
// @Produce json
// @Param id path int true "Lembaga ID"
// @Param data body web.LembagaUpdateRequest true "Lembaga Update Request"
// @Success 200 {object} web.WebResponse{data=web.LembagaResponse} "OK"
// @Failure 400 {object} web.WebResponse "Bad Request"
// @Failure 500 {object} web.WebResponse "Internal Server Error"
// @Router /lembagas/{id} [put]
func (controller *LembagaControllerImpl) Update(c echo.Context) error {
	lembagaUpdateRequest := web.LembagaUpdateRequest{}
	err := c.Bind(&lembagaUpdateRequest)
	if err != nil {
		return c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD_REQUEST",
		})
	}

	lembagaUpdateRequest.Id, err = strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD_REQUEST",
		})
	}

	lembagaResponse, err := controller.LembagaService.Update(c.Request().Context(), lembagaUpdateRequest)
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
		Data:   lembagaResponse,
	})
}

// @Summary Delete Lembaga
// @Description Delete existing Lembaga by ID
// @Tags Lembaga
// @Accept json
// @Produce json
// @Param id path int true "Lembaga ID"
// @Success 200 {object} web.WebResponse "OK"
// @Failure 400 {object} web.WebResponse "Bad Request"
// @Failure 500 {object} web.WebResponse "Internal Server Error"
// @Router /lembagas/{id} [delete]
func (controller *LembagaControllerImpl) Delete(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD_REQUEST",
		})
	}

	err = controller.LembagaService.Delete(c.Request().Context(), id)
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

// @Summary Get Lembaga by ID
// @Description Get Lembaga detail by ID
// @Tags Lembaga
// @Accept json
// @Produce json
// @Param id path int true "Lembaga ID"
// @Success 200 {object} web.WebResponse{data=web.LembagaResponse} "OK"
// @Failure 400 {object} web.WebResponse "Bad Request"
// @Failure 500 {object} web.WebResponse "Internal Server Error"
// @Router /lembagas/{id} [get]
func (controller *LembagaControllerImpl) FindById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD_REQUEST",
		})
	}

	lembagaResponse, err := controller.LembagaService.FindById(c.Request().Context(), id)
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
		Data:   lembagaResponse,
	})
}

// @Summary List All Lembaga
// @Description Get list of all Lembaga
// @Tags Lembaga
// @Accept json
// @Produce json
// @Success 200 {object} web.WebResponse{data=[]web.LembagaResponse} "OK"
// @Failure 500 {object} web.WebResponse "Internal Server Error"
// @Router /lembagas [get]
func (controller *LembagaControllerImpl) FindAll(c echo.Context) error {
	lembagaResponses, err := controller.LembagaService.FindAll(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, web.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: "INTERNAL_SERVER_ERROR",
		})
	}

	return c.JSON(http.StatusOK, web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   lembagaResponses,
	})
}
