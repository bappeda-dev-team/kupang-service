package controller

import (
	"kupang-service/helper"
	"kupang-service/model/web"
	"kupang-service/service"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type KegiatanControllerImpl struct {
	KegiatanService service.KegiatanService
}

func NewKegiatanControllerImpl(kegiatanService service.KegiatanService) *KegiatanControllerImpl {
	return &KegiatanControllerImpl{
		KegiatanService: kegiatanService,
	}
}

// @Summary Create Kegiatan
// @Description Create new Kegiatan
// @Tags Kegiatan
// @Accept json
// @Produce json
// @Param data body web.KegiatanCreateRequest true "Kegiatan Create Request"
// @Success 201 {object} web.WebResponse{data=web.KegiatanResponse} "Created"
// @Failure 400 {object} web.WebResponse "Bad Request"
// @Failure 500 {object} web.WebResponse "Internal Server Error"
// @Router /kegiatans [post]
func (controller *KegiatanControllerImpl) Create(c echo.Context) error {
	kegiatanCreateRequest := web.KegiatanCreateRequest{}
	err := c.Bind(&kegiatanCreateRequest)
	if err != nil {
		return c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD_REQUEST",
		})
	}

	kegiatanResponse, err := controller.KegiatanService.Create(c.Request().Context(), kegiatanCreateRequest)
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
		Data:   kegiatanResponse,
	})
}

// @Summary Update Kegiatan
// @Description Update existing Kegiatan by ID
// @Tags Kegiatan
// @Accept json
// @Produce json
// @Param id path int true "Kegiatan ID"
// @Param data body web.KegiatanUpdateRequest true "Kegiatan Update Request"
// @Success 200 {object} web.WebResponse{data=web.KegiatanResponse} "OK"
// @Failure 400 {object} web.WebResponse "Bad Request"
// @Failure 500 {object} web.WebResponse "Internal Server Error"
// @Router /kegiatans/{id} [put]
func (controller *KegiatanControllerImpl) Update(c echo.Context) error {
	kegiatanUpdateRequest := web.KegiatanUpdateRequest{}
	err := c.Bind(&kegiatanUpdateRequest)
	if err != nil {
		return c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD_REQUEST",
		})
	}

	kegiatanUpdateRequest.Id, err = strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD_REQUEST",
		})
	}

	kegiatanResponse, err := controller.KegiatanService.Update(c.Request().Context(), kegiatanUpdateRequest)
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
		Data:   kegiatanResponse,
	})
}

// @Summary Delete Kegiatan
// @Description Delete existing Kegiatan by ID
// @Tags Kegiatan
// @Accept json
// @Produce json
// @Param id path int true "Kegiatan ID"
// @Success 200 {object} web.WebResponse{data=web.KegiatanResponse} "OK"
// @Failure 400 {object} web.WebResponse "Bad Request"
// @Failure 500 {object} web.WebResponse "Internal Server Error"
// @Router /kegiatans/{id} [delete]
func (controller *KegiatanControllerImpl) Delete(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD_REQUEST",
		})
	}

	err = controller.KegiatanService.Delete(c.Request().Context(), id)
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

// @Summary Get Kegiatan by ID
// @Description Get Kegiatan detail by ID
// @Tags Kegiatan
// @Accept json
// @Produce json
// @Param id path int true "Kegiatan ID"
// @Success 200 {object} web.WebResponse{data=web.KegiatanResponse} "OK"
// @Failure 400 {object} web.WebResponse "Bad Request"
// @Failure 500 {object} web.WebResponse "Internal Server Error"
// @Router /kegiatans/{id} [get]
func (controller *KegiatanControllerImpl) FindById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD_REQUEST",
		})
	}

	kegiatanResponse, err := controller.KegiatanService.FindById(c.Request().Context(), id)
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
		Data:   kegiatanResponse,
	})
}

// @Summary List All Kegiatan
// @Description Get list of all Kegiatan
// @Tags Kegiatan
// @Accept json
// @Produce json
// @Success 200 {object} web.WebResponse{data=[]web.KegiatanResponse} "OK"
// @Failure 500 {object} web.WebResponse "Internal Server Error"
// @Router /kegiatans [get]
func (controller *KegiatanControllerImpl) FindAll(c echo.Context) error {
	kegiatanResponses, err := controller.KegiatanService.FindAll(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, web.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: "INTERNAL_SERVER_ERROR",
		})
	}

	return c.JSON(http.StatusOK, web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   kegiatanResponses,
	})
}
