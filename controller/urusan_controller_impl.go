package controller

import (
	"kupang-service/helper"
	"kupang-service/model/web"
	"kupang-service/service"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type UrusanControllerImpl struct {
	UrusanService service.UrusanService
}

func NewUrusanControllerImpl(urusanService service.UrusanService) *UrusanControllerImpl {
	return &UrusanControllerImpl{
		UrusanService: urusanService,
	}
}

// @Summary Create Urusan
// @Description Create new Urusan
// @Tags Urusan
// @Accept json
// @Produce json
// @Param data body web.UrusanCreateRequest true "Urusan Create Request"
// @Success 201 {object} web.WebResponse{data=web.UrusanResponse} "Created"
// @Failure 400 {object} web.WebResponse "Bad Request"
// @Failure 500 {object} web.WebResponse "Internal Server Error"
// @Router /urusans [post]
func (controller *UrusanControllerImpl) Create(c echo.Context) error {
	urusanCreateRequest := web.UrusanCreateRequest{}
	err := c.Bind(&urusanCreateRequest)
	if err != nil {
		return c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD_REQUEST",
		})
	}

	urusanResponse, err := controller.UrusanService.Create(c.Request().Context(), urusanCreateRequest)
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
		Data:   urusanResponse,
	})
}

// @Summary Update Urusan
// @Description Update existing Urusan by ID
// @Tags Urusan
// @Accept json
// @Produce json
// @Param id path int true "Urusan ID"
// @Param data body web.UrusanUpdateRequest true "Urusan Update Request"
// @Success 200 {object} web.WebResponse{data=web.UrusanResponse} "OK"
// @Failure 400 {object} web.WebResponse "Bad Request"
// @Failure 500 {object} web.WebResponse "Internal Server Error"
// @Router /urusans/{id} [put]
func (controller *UrusanControllerImpl) Update(c echo.Context) error {
	urusanUpdateRequest := web.UrusanUpdateRequest{}
	err := c.Bind(&urusanUpdateRequest)
	if err != nil {
		return c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD_REQUEST",
		})
	}

	urusanUpdateRequest.Id, err = strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD_REQUEST",
		})
	}

	urusanResponse, err := controller.UrusanService.Update(c.Request().Context(), urusanUpdateRequest)
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
		Data:   urusanResponse,
	})
}

// @Summary Delete Urusan
// @Description Delete existing Urusan by ID
// @Tags Urusan
// @Accept json
// @Produce json
// @Param id path int true "Urusan ID"
// @Success 200 {object} web.WebResponse{data=web.UrusanResponse} "OK"
// @Failure 400 {object} web.WebResponse "Bad Request"
// @Failure 500 {object} web.WebResponse "Internal Server Error"
// @Router /urusans/{id} [delete]
func (controller *UrusanControllerImpl) Delete(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD_REQUEST",
		})
	}

	err = controller.UrusanService.Delete(c.Request().Context(), id)
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

// @Summary Get Urusan by ID
// @Description Get Urusan detail by ID
// @Tags Urusan
// @Accept json
// @Produce json
// @Param id path int true "Urusan ID"
// @Success 200 {object} web.WebResponse{data=web.UrusanResponse} "OK"
// @Failure 400 {object} web.WebResponse "Bad Request"
// @Failure 500 {object} web.WebResponse "Internal Server Error"
// @Router /urusans/{id} [get]
func (controller *UrusanControllerImpl) FindById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD_REQUEST",
		})
	}

	urusanResponse, err := controller.UrusanService.FindById(c.Request().Context(), id)
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
		Data:   urusanResponse,
	})
}

// @Summary List All Urusan
// @Description Get list of all Urusan
// @Tags Urusan
// @Accept json
// @Produce json
// @Success 200 {object} web.WebResponse{data=[]web.UrusanResponse} "OK"
// @Failure 500 {object} web.WebResponse "Internal Server Error"
// @Router /urusans [get]
func (controller *UrusanControllerImpl) FindAll(c echo.Context) error {
	urusanResponses, err := controller.UrusanService.FindAll(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, web.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: "INTERNAL_SERVER_ERROR",
		})
	}

	return c.JSON(http.StatusOK, web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   urusanResponses,
	})
}
