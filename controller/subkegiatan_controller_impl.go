package controller

import (
	"kupang-service/helper"
	"kupang-service/model/web"
	"kupang-service/service"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type SubkegiatanControllerImpl struct {
	SubkegiatanService service.SubkegiatanService
}

func NewSubkegiatanControllerImpl(subkegiatanService service.SubkegiatanService) *SubkegiatanControllerImpl {
	return &SubkegiatanControllerImpl{
		SubkegiatanService: subkegiatanService,
	}
}

// @Summary Create Subkegiatan
// @Description Create new Subkegiatan
// @Tags Subkegiatan
// @Accept json
// @Produce json
// @Param data body web.SubkegiatanCreateRequest true "Subkegiatan Create Request"
// @Success 201 {object} web.WebResponse{data=web.SubkegiatanResponse} "Created"
// @Failure 400 {object} web.WebResponse "Bad Request"
// @Failure 500 {object} web.WebResponse "Internal Server Error"
// @Router /subkegiatans [post]
func (controller *SubkegiatanControllerImpl) Create(c echo.Context) error {
	subkegiatanCreateRequest := web.SubkegiatanCreateRequest{}
	err := c.Bind(&subkegiatanCreateRequest)
	if err != nil {
		return c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD_REQUEST",
		})
	}

	subkegiatanResponse, err := controller.SubkegiatanService.Create(c.Request().Context(), subkegiatanCreateRequest)
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
		Data:   subkegiatanResponse,
	})
}

// @Summary Update Subkegiatan
// @Description Update existing Subkegiatan by ID
// @Tags Subkegiatan
// @Accept json
// @Produce json
// @Param id path int true "Subkegiatan ID"
// @Param data body web.SubkegiatanUpdateRequest true "Subkegiatan Update Request"
// @Success 200 {object} web.WebResponse{data=web.SubkegiatanResponse} "OK"
// @Failure 400 {object} web.WebResponse "Bad Request"
// @Failure 500 {object} web.WebResponse "Internal Server Error"
// @Router /subkegiatans/{id} [put]
func (controller *SubkegiatanControllerImpl) Update(c echo.Context) error {
	subkegiatanUpdateRequest := web.SubkegiatanUpdateRequest{}
	err := c.Bind(&subkegiatanUpdateRequest)
	if err != nil {
		return c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD_REQUEST",
		})
	}

	subkegiatanUpdateRequest.Id, err = strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD_REQUEST",
		})
	}

	subkegiatanResponse, err := controller.SubkegiatanService.Update(c.Request().Context(), subkegiatanUpdateRequest)
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
		Data:   subkegiatanResponse,
	})
}

// @Summary Delete Subkegiatan
// @Description Delete existing Subkegiatan by ID
// @Tags Subkegiatan
// @Accept json
// @Produce json
// @Param id path int true "Subkegiatan ID"
// @Success 200 {object} web.WebResponse{data=web.SubkegiatanResponse} "OK"
// @Failure 400 {object} web.WebResponse "Bad Request"
// @Failure 500 {object} web.WebResponse "Internal Server Error"
// @Router /subkegiatans/{id} [delete]
func (controller *SubkegiatanControllerImpl) Delete(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD_REQUEST",
		})
	}

	err = controller.SubkegiatanService.Delete(c.Request().Context(), id)
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

// @Summary Get Subkegiatan by ID
// @Description Get Subkegiatan detail by ID
// @Tags Subkegiatan
// @Accept json
// @Produce json
// @Param id path int true "Subkegiatan ID"
// @Success 200 {object} web.WebResponse{data=web.SubkegiatanResponse} "OK"
// @Failure 400 {object} web.WebResponse "Bad Request"
// @Failure 500 {object} web.WebResponse "Internal Server Error"
// @Router /subkegiatans/{id} [get]
func (controller *SubkegiatanControllerImpl) FindById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD_REQUEST",
		})
	}

	subkegiatanResponse, err := controller.SubkegiatanService.FindById(c.Request().Context(), id)
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
		Data:   subkegiatanResponse,
	})
}

// @Summary List All Subkegiatan
// @Description Get list of all Subkegiatan
// @Tags Subkegiatan
// @Accept json
// @Produce json
// @Success 200 {object} web.WebResponse{data=[]web.SubkegiatanResponse} "OK"
// @Failure 500 {object} web.WebResponse "Internal Server Error"
// @Router /subkegiatans [get]
func (controller *SubkegiatanControllerImpl) FindAll(c echo.Context) error {
	subkegiatanResponses, err := controller.SubkegiatanService.FindAll(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, web.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: "INTERNAL_SERVER_ERROR",
		})
	}

	return c.JSON(http.StatusOK, web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   subkegiatanResponses,
	})
}
