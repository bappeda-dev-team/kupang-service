package controller

import (
	"kupang-service/helper"
	"kupang-service/model/web"
	"kupang-service/service"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type MusrenbangControllerImpl struct {
	MusrenbangService service.MusrenbangService
}

func NewMusrenbangControllerImpl(musrenbangService service.MusrenbangService) *MusrenbangControllerImpl {
	return &MusrenbangControllerImpl{
		MusrenbangService: musrenbangService,
	}
}

// @Summary Create Musrenbang
// @Description Create new Musrenbang
// @Tags Musrenbang
// @Accept json
// @Produce json
// @Param data body web.MusrenbangCreateRequest true "Musrenbang Create Request"
// @Success 201 {object} web.WebResponse{data=web.MusrenbangResponse} "Created"
// @Failure 400 {object} web.WebResponse "Bad Request"
// @Failure 500 {object} web.WebResponse "Internal Server Error"
// @Router /musrenbangs [post]
func (controller *MusrenbangControllerImpl) Create(c echo.Context) error {
	musrenbangCreateRequest := web.MusrenbangCreateRequest{}
	err := c.Bind(&musrenbangCreateRequest)
	if err != nil {
		return c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD_REQUEST",
		})
	}

	musrenbangResponse, err := controller.MusrenbangService.Create(c.Request().Context(), musrenbangCreateRequest)
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
		Data:   musrenbangResponse,
	})
}

// @Summary Update Musrenbang
// @Description Update existing Musrenbang by ID
// @Tags Musrenbang
// @Accept json
// @Produce json
// @Param id path int true "Musrenbang ID"
// @Param data body web.MusrenbangUpdateRequest true "Musrenbang Update Request"
// @Success 200 {object} web.WebResponse{data=web.MusrenbangResponse} "OK"
// @Failure 400 {object} web.WebResponse "Bad Request"
// @Failure 500 {object} web.WebResponse "Internal Server Error"
// @Router /musrenbangs/{id} [put]
func (controller *MusrenbangControllerImpl) Update(c echo.Context) error {
	musrenbangUpdateRequest := web.MusrenbangUpdateRequest{}
	err := c.Bind(&musrenbangUpdateRequest)
	if err != nil {
		return c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD_REQUEST",
		})
	}

	musrenbangUpdateRequest.Id, err = strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD_REQUEST",
		})
	}

	musrenbangResponse, err := controller.MusrenbangService.Update(c.Request().Context(), musrenbangUpdateRequest)
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
		Data:   musrenbangResponse,
	})
}

// @Summary Delete Musrenbang
// @Description Delete existing Musrenbang by ID
// @Tags Musrenbang
// @Accept json
// @Produce json
// @Param id path int true "Musrenbang ID"
// @Success 200 {object} web.WebResponse "OK"
// @Failure 400 {object} web.WebResponse "Bad Request"
// @Failure 500 {object} web.WebResponse "Internal Server Error"
// @Router /musrenbangs/{id} [delete]
func (controller *MusrenbangControllerImpl) Delete(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD_REQUEST",
		})
	}

	err = controller.MusrenbangService.Delete(c.Request().Context(), id)
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

// @Summary Get Musrenbang by ID
// @Description Get Musrenbang detail by ID
// @Tags Musrenbang
// @Accept json
// @Produce json
// @Param id path int true "Musrenbang ID"
// @Success 200 {object} web.WebResponse{data=web.MusrenbangResponse} "OK"
// @Failure 400 {object} web.WebResponse "Bad Request"
// @Failure 500 {object} web.WebResponse "Internal Server Error"
// @Router /musrenbangs/{id} [get]
func (controller *MusrenbangControllerImpl) FindById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD_REQUEST",
		})
	}

	musrenbangResponse, err := controller.MusrenbangService.FindById(c.Request().Context(), id)
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
		Data:   musrenbangResponse,
	})
}

// @Summary List All Musrenbang
// @Description Get list of all Musrenbangs
// @Tags Musrenbang
// @Accept json
// @Produce json
// @Success 200 {object} web.WebResponse{data=[]web.MusrenbangResponse} "OK"
// @Failure 500 {object} web.WebResponse "Internal Server Error"
// @Router /musrenbangs [get]
func (controller *MusrenbangControllerImpl) FindAll(c echo.Context) error {
	musrenbangResponses, err := controller.MusrenbangService.FindAll(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, web.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: "INTERNAL_SERVER_ERROR",
		})
	}

	return c.JSON(http.StatusOK, web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   musrenbangResponses,
	})
}

// @Summary Get Musrenbang by Kode OPD
// @Description Get Musrenbang list by OPD code
// @Tags Musrenbang
// @Accept json
// @Produce json
// @Param kode_opd path string true "OPD Code"
// @Success 200 {object} web.WebResponse{data=[]web.MusrenbangResponse} "OK"
// @Failure 400 {object} web.WebResponse "Bad Request"
// @Failure 500 {object} web.WebResponse "Internal Server Error"
// @Router /musrenbangs/opd/{kode_opd} [get]
func (controller *MusrenbangControllerImpl) FindByKodeOpd(c echo.Context) error {
	kodeOpd := c.Param("kode_opd")

	musrenbangResponses, err := controller.MusrenbangService.FindByKodeOpd(c.Request().Context(), kodeOpd)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, web.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: "INTERNAL_SERVER_ERROR",
		})
	}

	return c.JSON(http.StatusOK, web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   musrenbangResponses,
	})
}
