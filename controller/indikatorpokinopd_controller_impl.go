package controller

import (
	"kupang-service/model/web"
	"kupang-service/service"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type IndikatorPokinOpdControllerImpl struct {
	IndikatorPokinOpdService service.IndikatorPokinOpdService
}

func NewIndikatorPokinOpdControllerImpl(indikatorPokinOpdService service.IndikatorPokinOpdService) *IndikatorPokinOpdControllerImpl {
	return &IndikatorPokinOpdControllerImpl{
		IndikatorPokinOpdService: indikatorPokinOpdService,
	}
}

// @Summary Create Indikator Pokin Opd
// @Description Create new Indikator Pokin Opd
// @Tags Indikator Pokin Opd
// @Accept json
// @Produce json
// @Param data body web.IndikatorPokinOpdCreateRequest true "Indikator Pokin Opd Create Request"
// @Success 201 {object} web.WebResponse{data=web.IndikatorPokinOpdResponse} "Created"
// @Failure 400 {object} web.WebResponse "Bad Request"
// @Failure 500 {object} web.WebResponse "Internal Server Error"
// @Router /api/v1/indikator-pokin-opd [post]
func (controller *IndikatorPokinOpdControllerImpl) Create(c echo.Context) error {
	indikatorPokinOpdCreateRequest := web.IndikatorPokinOpdCreateRequest{}
	err := c.Bind(&indikatorPokinOpdCreateRequest)
	if err != nil {
		return c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD_REQUEST",
		})
	}

	indikatorPokinOpdResponse, err := controller.IndikatorPokinOpdService.Create(c.Request().Context(), indikatorPokinOpdCreateRequest)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, web.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: "INTERNAL_SERVER_ERROR",
		})
	}

	return c.JSON(http.StatusOK, web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   indikatorPokinOpdResponse,
	})
}

// @Summary Update Indikator Pokin Opd
// @Description Update existing Indikator Pokin Opd by ID
// @Tags Indikator Pokin Opd
// @Accept json
// @Produce json
// @Param id path int true "Indikator Pokin Opd ID"
// @Param data body web.IndikatorPokinOpdUpdateRequest true "Indikator Pokin Opd Update Request"
// @Success 200 {object} web.WebResponse{data=web.IndikatorPokinOpdResponse} "OK"
// @Failure 400 {object} web.WebResponse "Bad Request"
// @Failure 500 {object} web.WebResponse "Internal Server Error"
// @Router /api/v1/indikator-pokin-opd/{id} [put]
func (controller *IndikatorPokinOpdControllerImpl) Update(c echo.Context) error {
	indikatorPokinOpdUpdateRequest := web.IndikatorPokinOpdUpdateRequest{}
	err := c.Bind(&indikatorPokinOpdUpdateRequest)
	if err != nil {
		return c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD_REQUEST",
		})
	}

	indikatorPokinOpdUpdateRequest.Id, err = strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD_REQUEST",
		})
	}

	indikatorPokinOpdResponse, err := controller.IndikatorPokinOpdService.Update(c.Request().Context(), indikatorPokinOpdUpdateRequest)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, web.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: "INTERNAL_SERVER_ERROR",
		})
	}

	return c.JSON(http.StatusOK, web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   indikatorPokinOpdResponse,
	})
}

// @Summary Delete Indikator Pokin Opd
// @Description Delete existing Indikator Pokin Opd by ID
// @Tags Indikator Pokin Opd
// @Accept json
// @Produce json
// @Param id path int true "Indikator Pokin Opd ID"
// @Success 200 {object} web.WebResponse{data=web.IndikatorPokinOpdResponse} "OK"
// @Failure 400 {object} web.WebResponse "Bad Request"
// @Failure 500 {object} web.WebResponse "Internal Server Error"
// @Router /api/v1/indikator-pokin-opd/{id} [delete]
func (controller *IndikatorPokinOpdControllerImpl) Delete(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD_REQUEST",
		})
	}

	err = controller.IndikatorPokinOpdService.Delete(c.Request().Context(), id)
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

// @Summary Get Indikator Pokin Opd by ID
// @Description Get Indikator Pokin Opd detail by ID
// @Tags Indikator Pokin Opd
// @Accept json
// @Produce json
// @Param id path int true "Indikator Pokin Opd ID"
// @Success 200 {object} web.WebResponse{data=web.IndikatorPokinOpdResponse} "OK"
// @Failure 400 {object} web.WebResponse "Bad Request"
// @Failure 500 {object} web.WebResponse "Internal Server Error"
// @Router /api/v1/indikator-pokin-opd/{id} [get]
func (controller *IndikatorPokinOpdControllerImpl) FindById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD_REQUEST",
		})
	}

	indikatorPokinOpdResponse, err := controller.IndikatorPokinOpdService.FindById(c.Request().Context(), id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, web.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: "INTERNAL_SERVER_ERROR",
		})
	}

	return c.JSON(http.StatusOK, web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   indikatorPokinOpdResponse,
	})
}

// @Summary List All Indikator Pokin Opd
// @Description Get list of all Indikator Pokin Opd
// @Tags Indikator Pokin Opd
// @Accept json
// @Produce json
// @Success 200 {object} web.WebResponse{data=[]web.IndikatorPokinOpdResponse} "OK"
// @Failure 500 {object} web.WebResponse "Internal Server Error"
// @Router /api/v1/indikator-pokin-opd [get]
func (controller *IndikatorPokinOpdControllerImpl) FindAll(c echo.Context) error {
	indikatorPokinOpdResponses, err := controller.IndikatorPokinOpdService.FindAll(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, web.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: "INTERNAL_SERVER_ERROR",
		})
	}

	return c.JSON(http.StatusOK, web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   indikatorPokinOpdResponses,
	})
}
