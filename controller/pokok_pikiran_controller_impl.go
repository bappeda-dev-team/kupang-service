package controller

import (
	"kupang-service/helper"
	"kupang-service/model/web"
	"kupang-service/service"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type PokokPikiranControllerImpl struct {
	PokokPikiranService service.PokokPikiranService
}

func NewPokokPikiranControllerImpl(pokokPikiranService service.PokokPikiranService) *PokokPikiranControllerImpl {
	return &PokokPikiranControllerImpl{
		PokokPikiranService: pokokPikiranService,
	}
}

// @Summary Create Pokok Pikiran
// @Description Create new Pokok Pikiran
// @Tags Pokok Pikiran
// @Accept json
// @Produce json
// @Param data body web.PokokPikiranCreateRequest true "Pokok Pikiran Create Request"
// @Success 201 {object} web.WebResponse{data=web.PokokPikiranResponse} "Created"
// @Failure 400 {object} web.WebResponse "Bad Request"
// @Failure 500 {object} web.WebResponse "Internal Server Error"
// @Router /pokok-pikirans [post]
func (controller *PokokPikiranControllerImpl) Create(c echo.Context) error {
	pokokPikiranCreateRequest := web.PokokPikiranCreateRequest{}
	err := c.Bind(&pokokPikiranCreateRequest)
	if err != nil {
		return c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD_REQUEST",
		})
	}

	pokokPikiranResponse, err := controller.PokokPikiranService.Create(c.Request().Context(), pokokPikiranCreateRequest)
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
		Data:   pokokPikiranResponse,
	})
}

// @Summary Update Pokok Pikiran
// @Description Update existing Pokok Pikiran by ID
// @Tags Pokok Pikiran
// @Accept json
// @Produce json
// @Param id path int true "Pokok Pikiran ID"
// @Param data body web.PokokPikiranUpdateRequest true "Pokok Pikiran Update Request"
// @Success 200 {object} web.WebResponse{data=web.PokokPikiranResponse} "OK"
// @Failure 400 {object} web.WebResponse "Bad Request"
// @Failure 500 {object} web.WebResponse "Internal Server Error"
// @Router /pokok-pikirans/{id} [put]
func (controller *PokokPikiranControllerImpl) Update(c echo.Context) error {
	pokokPikiranUpdateRequest := web.PokokPikiranUpdateRequest{}
	err := c.Bind(&pokokPikiranUpdateRequest)
	if err != nil {
		return c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD_REQUEST",
		})
	}

	pokokPikiranUpdateRequest.Id, err = strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD_REQUEST",
		})
	}

	pokokPikiranResponse, err := controller.PokokPikiranService.Update(c.Request().Context(), pokokPikiranUpdateRequest)
	if err != nil {
		if helper.IsValidationError(err) {
			return c.JSON(http.StatusBadRequest, web.WebResponse{
				Code:   http.StatusBadRequest,
				Status: "BAD_REQUEST",
				Data:   err.Error(),
			})
		}
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
		Data:   pokokPikiranResponse,
	})
}

// @Summary Delete Pokok Pikiran
// @Description Delete existing Pokok Pikiran by ID
// @Tags Pokok Pikiran
// @Accept json
// @Produce json
// @Param id path int true "Pokok Pikiran ID"
// @Success 200 {object} web.WebResponse "OK"
// @Failure 400 {object} web.WebResponse "Bad Request"
// @Failure 500 {object} web.WebResponse "Internal Server Error"
// @Router /pokok-pikirans/{id} [delete]
func (controller *PokokPikiranControllerImpl) Delete(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD_REQUEST",
		})
	}

	err = controller.PokokPikiranService.Delete(c.Request().Context(), id)
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

// @Summary Get Pokok Pikiran by ID
// @Description Get Pokok Pikiran detail by ID
// @Tags Pokok Pikiran
// @Accept json
// @Produce json
// @Param id path int true "Pokok Pikiran ID"
// @Success 200 {object} web.WebResponse{data=web.PokokPikiranResponse} "OK"
// @Failure 400 {object} web.WebResponse "Bad Request"
// @Failure 500 {object} web.WebResponse "Internal Server Error"
// @Router /pokok-pikirans/{id} [get]
func (controller *PokokPikiranControllerImpl) FindById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD_REQUEST",
		})
	}

	pokokPikiranResponse, err := controller.PokokPikiranService.FindById(c.Request().Context(), id)
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
		Data:   pokokPikiranResponse,
	})
}

// @Summary List All Pokok Pikiran
// @Description Get list of all Pokok Pikiran
// @Tags Pokok Pikiran
// @Accept json
// @Produce json
// @Success 200 {object} web.WebResponse{data=[]web.PokokPikiranResponse} "OK"
// @Failure 500 {object} web.WebResponse "Internal Server Error"
// @Router /pokok-pikirans [get]
func (controller *PokokPikiranControllerImpl) FindAll(c echo.Context) error {
	pokokPikiranResponses, err := controller.PokokPikiranService.FindAll(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, web.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: "INTERNAL_SERVER_ERROR",
		})
	}

	return c.JSON(http.StatusOK, web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   pokokPikiranResponses,
	})
}

// @Summary Get Pokok Pikiran by Kode OPD
// @Description Get Pokok Pikiran list by OPD code
// @Tags Pokok Pikiran
// @Accept json
// @Produce json
// @Param kode_opd path string true "OPD Code"
// @Success 200 {object} web.WebResponse{data=[]web.PokokPikiranResponse} "OK"
// @Failure 400 {object} web.WebResponse "Bad Request"
// @Failure 500 {object} web.WebResponse "Internal Server Error"
// @Router /pokok-pikirans/opd/{kode_opd} [get]
func (controller *PokokPikiranControllerImpl) FindByKodeOpd(c echo.Context) error {
	kodeOpd := c.Param("kode_opd")

	pokokPikiranResponses, err := controller.PokokPikiranService.FindByKodeOpd(c.Request().Context(), kodeOpd)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, web.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: "INTERNAL_SERVER_ERROR",
		})
	}

	return c.JSON(http.StatusOK, web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   pokokPikiranResponses,
	})
}
