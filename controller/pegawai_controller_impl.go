package controller

import (
	"kupang-service/helper"
	"kupang-service/model/web"
	"kupang-service/service"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type PegawaiControllerImpl struct {
	PegawaiService service.PegawaiService
}

func NewPegawaiControllerImpl(pegawaiService service.PegawaiService) *PegawaiControllerImpl {
	return &PegawaiControllerImpl{
		PegawaiService: pegawaiService,
	}
}

// @Summary Create Pegawai
// @Description Create new Pegawai (jabatan_id optional)
// @Tags Pegawai
// @Accept json
// @Produce json
// @Param data body web.PegawaiCreateRequest true "Pegawai Create Request"
// @Success 201 {object} web.WebResponse{data=web.PegawaiResponse} "Created"
// @Failure 400 {object} web.WebResponse "Bad Request"
// @Failure 500 {object} web.WebResponse "Internal Server Error"
// @Router /pegawais [post]
func (controller *PegawaiControllerImpl) Create(c echo.Context) error {
	pegawaiCreateRequest := web.PegawaiCreateRequest{}
	err := c.Bind(&pegawaiCreateRequest)
	if err != nil {
		return c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD_REQUEST",
		})
	}

	pegawaiResponse, err := controller.PegawaiService.Create(c.Request().Context(), pegawaiCreateRequest)
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
		Data:   pegawaiResponse,
	})
}

// @Summary Update Pegawai
// @Description Update existing Pegawai by ID
// @Tags Pegawai
// @Accept json
// @Produce json
// @Param id path int true "Pegawai ID"
// @Param data body web.PegawaiUpdateRequest true "Pegawai Update Request"
// @Success 200 {object} web.WebResponse{data=web.PegawaiResponse} "OK"
// @Failure 400 {object} web.WebResponse "Bad Request"
// @Failure 500 {object} web.WebResponse "Internal Server Error"
// @Router /pegawais/{id} [put]
func (controller *PegawaiControllerImpl) Update(c echo.Context) error {
	pegawaiUpdateRequest := web.PegawaiUpdateRequest{}
	err := c.Bind(&pegawaiUpdateRequest)
	if err != nil {
		return c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD_REQUEST",
		})
	}

	pegawaiUpdateRequest.Id, err = strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD_REQUEST",
		})
	}

	pegawaiResponse, err := controller.PegawaiService.Update(c.Request().Context(), pegawaiUpdateRequest)
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
		Data:   pegawaiResponse,
	})
}

// @Summary Tambah Jabatan ke Pegawai
// @Description Menambah atau menetapkan jabatan pada pegawai yang sudah ada (akan membuat jabatan baru jika belum tersedia)
// @Tags Pegawai
// @Accept json
// @Produce json
// @Param data body web.PegawaiAddJabatanRequest true "Tambah Jabatan"
// @Success 200 {object} web.WebResponse{data=web.PegawaiResponse} "OK"
// @Failure 400 {object} web.WebResponse "Bad Request"
// @Failure 500 {object} web.WebResponse "Internal Server Error"
// @Router /pegawais/jabatan [post]
func (controller *PegawaiControllerImpl) AddJabatan(c echo.Context) error {
	request := web.PegawaiAddJabatanRequest{}
	err := c.Bind(&request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD_REQUEST",
		})
	}

	pegawaiResponse, err := controller.PegawaiService.AddJabatan(c.Request().Context(), request)
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
		Data:   pegawaiResponse,
	})
}

// @Summary Delete Pegawai
// @Description Delete existing Pegawai by ID
// @Tags Pegawai
// @Accept json
// @Produce json
// @Param id path int true "Pegawai ID"
// @Success 200 {object} web.WebResponse "OK"
// @Failure 400 {object} web.WebResponse "Bad Request"
// @Failure 500 {object} web.WebResponse "Internal Server Error"
// @Router /pegawais/{id} [delete]
func (controller *PegawaiControllerImpl) Delete(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD_REQUEST",
		})
	}

	err = controller.PegawaiService.Delete(c.Request().Context(), id)
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

// @Summary Get Pegawai by ID
// @Description Get Pegawai details by ID
// @Tags Pegawai
// @Produce json
// @Param id path int true "Pegawai ID"
// @Success 200 {object} web.WebResponse{data=web.PegawaiResponse} "OK"
// @Failure 400 {object} web.WebResponse "Bad Request"
// @Failure 404 {object} web.WebResponse "Not Found"
// @Failure 500 {object} web.WebResponse "Internal Server Error"
// @Router /pegawais/{id} [get]
func (controller *PegawaiControllerImpl) FindById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD_REQUEST",
		})
	}

	pegawaiResponse, err := controller.PegawaiService.FindById(c.Request().Context(), id)
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
		Data:   pegawaiResponse,
	})
}

// @Summary Get All Pegawai
// @Description Get list of all Pegawai
// @Tags Pegawai
// @Produce json
// @Success 200 {object} web.WebResponse{data=[]web.PegawaiResponse} "OK"
// @Failure 500 {object} web.WebResponse "Internal Server Error"
// @Router /pegawais [get]
func (controller *PegawaiControllerImpl) FindAll(c echo.Context) error {
	pegawaiResponses, err := controller.PegawaiService.FindAll(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, web.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: "INTERNAL_SERVER_ERROR",
		})
	}

	return c.JSON(http.StatusOK, web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   pegawaiResponses,
	})
}

// @Summary Get Pegawai by Kode OPD
// @Description Get list of Pegawai by kode_opd
// @Tags Pegawai
// @Produce json
// @Param kode_opd path string true "Kode OPD"
// @Success 200 {object} web.WebResponse{data=[]web.PegawaiResponse} "OK"
// @Failure 500 {object} web.WebResponse "Internal Server Error"
// @Router /pegawais/opd/{kode_opd} [get]
func (controller *PegawaiControllerImpl) FindByKodeOpd(c echo.Context) error {
	kodeOpd := c.Param("kode_opd")

	pegawaiResponses, err := controller.PegawaiService.FindByKodeOpd(c.Request().Context(), kodeOpd)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, web.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: "INTERNAL_SERVER_ERROR",
		})
	}

	return c.JSON(http.StatusOK, web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   pegawaiResponses,
	})
}
