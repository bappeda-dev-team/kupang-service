package controller

import (
	"kupang-service/model/web"
	"kupang-service/service"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type PohonKinerjaControllerImpl struct {
	PohonKinerjaService service.PohonKinerjaService
}

func NewPohonKinerjaControllerImpl(pohonKinerjaService service.PohonKinerjaService) *PohonKinerjaControllerImpl {
	return &PohonKinerjaControllerImpl{
		PohonKinerjaService: pohonKinerjaService,
	}
}

// @Summary Get Pohon Kinerja Opd
// @Description Get pohon kinerja opd by kode opd and tahun
// @Tags Pohon Kinerja Opd
// @Accept json
// @Produce json
// @Param kode_opd path string true "Kode OPD"
// @Param tahun path int true "Tahun"
// @Success 200 {object} web.WebResponse{data=web.PohonKinerjaResponse} "OK"
// @Failure 400 {object} web.WebResponse "Bad Request"
// @Failure 500 {object} web.WebResponse "Internal Server Error"
// @Router /api/v1/pohon-kinerja-opd/{kode_opd}/{tahun} [get]
func (controller *PohonKinerjaControllerImpl) FindByKodeOpdAndTahun(c echo.Context) error {
	kodeOpd := c.Param("kode_opd")
	if kodeOpd == "" {
		return c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD_REQUEST",
		})
	}

	tahun, err := strconv.Atoi(c.Param("tahun"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD_REQUEST",
		})
	}

	response, err := controller.PohonKinerjaService.FindByKodeOpdAndTahun(c.Request().Context(), kodeOpd, tahun)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, web.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: "INTERNAL_SERVER_ERROR",
		})
	}

	return c.JSON(http.StatusOK, web.WebResponse{
		Code:   http.StatusOK,
		Status: "Success Get All Pohon Kinerja",
		Data:   response,
	})
}
