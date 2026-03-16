package controller

import (
	"kupang-service/model/web"
	"kupang-service/service"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type IndikatorPokinOpdStrategicControllerImpl struct {
	IndikatorPokinOpdStrategicService service.IndikatorPokinOpdStrategicService
}

func NewIndikatorPokinOpdStrategicControllerImpl(indikatorService service.IndikatorPokinOpdStrategicService) *IndikatorPokinOpdStrategicControllerImpl {
	return &IndikatorPokinOpdStrategicControllerImpl{
		IndikatorPokinOpdStrategicService: indikatorService,
	}
}

// @Summary Create Indikator Pokin OPD Strategic
// @Description Create new Indikator Pokin OPD Strategic
// @Tags Indikator Pokin OPD Strategic
// @Accept json
// @Produce json
// @Param data body web.IndikatorPokinOpdStrategicCreateRequest true "Indikator Pokin OPD Strategic Create Request"
// @Success 201 {object} web.WebResponse{data=web.IndikatorPokinOpdStrategicResponse} "Created"
// @Failure 400 {object} web.WebResponse "Bad Request"
// @Failure 500 {object} web.WebResponse "Internal Server Error"
// @Router /api/v1/indikator-pokin-opd-strategics [post]
func (controller *IndikatorPokinOpdStrategicControllerImpl) Create(c echo.Context) error {
	createRequest := web.IndikatorPokinOpdStrategicCreateRequest{}
	err := c.Bind(&createRequest)
	if err != nil {
		return c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD_REQUEST",
		})
	}

	response, err := controller.IndikatorPokinOpdStrategicService.Create(c.Request().Context(), createRequest)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, web.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: "INTERNAL_SERVER_ERROR",
		})
	}

	return c.JSON(http.StatusOK, web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   response,
	})
}

// @Summary Update Indikator Pokin OPD Strategic
// @Description Update existing Indikator Pokin OPD Strategic by ID
// @Tags Indikator Pokin OPD Strategic
// @Accept json
// @Produce json
// @Param id path int true "Indikator Pokin OPD Strategic ID"
// @Param data body web.IndikatorPokinOpdStrategicUpdateRequest true "Indikator Pokin OPD Strategic Update Request"
// @Success 200 {object} web.WebResponse{data=web.IndikatorPokinOpdStrategicResponse} "OK"
// @Failure 400 {object} web.WebResponse "Bad Request"
// @Failure 500 {object} web.WebResponse "Internal Server Error"
// @Router /api/v1/indikator-pokin-opd-strategics/{id} [put]
func (controller *IndikatorPokinOpdStrategicControllerImpl) Update(c echo.Context) error {
	updateRequest := web.IndikatorPokinOpdStrategicUpdateRequest{}
	err := c.Bind(&updateRequest)
	if err != nil {
		return c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD_REQUEST",
		})
	}

	updateRequest.Id, err = strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD_REQUEST",
		})
	}

	response, err := controller.IndikatorPokinOpdStrategicService.Update(c.Request().Context(), updateRequest)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, web.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: "INTERNAL_SERVER_ERROR",
		})
	}

	return c.JSON(http.StatusOK, web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   response,
	})
}

// @Summary Delete Indikator Pokin OPD Strategic
// @Description Delete existing Indikator Pokin OPD Strategic by ID
// @Tags Indikator Pokin OPD Strategic
// @Accept json
// @Produce json
// @Param id path int true "Indikator Pokin OPD Strategic ID"
// @Success 200 {object} web.WebResponse "OK"
// @Failure 400 {object} web.WebResponse "Bad Request"
// @Failure 500 {object} web.WebResponse "Internal Server Error"
// @Router /api/v1/indikator-pokin-opd-strategics/{id} [delete]
func (controller *IndikatorPokinOpdStrategicControllerImpl) Delete(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD_REQUEST",
		})
	}

	err = controller.IndikatorPokinOpdStrategicService.Delete(c.Request().Context(), id)
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

// @Summary Get Indikator Pokin OPD Strategic by ID
// @Description Get Indikator Pokin OPD Strategic detail by ID
// @Tags Indikator Pokin OPD Strategic
// @Accept json
// @Produce json
// @Param id path int true "Indikator Pokin OPD Strategic ID"
// @Success 200 {object} web.WebResponse{data=web.IndikatorPokinOpdStrategicResponse} "OK"
// @Failure 400 {object} web.WebResponse "Bad Request"
// @Failure 500 {object} web.WebResponse "Internal Server Error"
// @Router /api/v1/indikator-pokin-opd-strategics/{id} [get]
func (controller *IndikatorPokinOpdStrategicControllerImpl) FindById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD_REQUEST",
		})
	}

	response, err := controller.IndikatorPokinOpdStrategicService.FindById(c.Request().Context(), id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, web.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: "INTERNAL_SERVER_ERROR",
		})
	}

	return c.JSON(http.StatusOK, web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   response,
	})
}

// @Summary List All Indikator Pokin OPD Strategic
// @Description Get list of all Indikator Pokin OPD Strategic
// @Tags Indikator Pokin OPD Strategic
// @Accept json
// @Produce json
// @Success 200 {object} web.WebResponse{data=[]web.IndikatorPokinOpdStrategicResponse} "OK"
// @Failure 500 {object} web.WebResponse "Internal Server Error"
// @Router /api/v1/indikator-pokin-opd-strategics [get]
func (controller *IndikatorPokinOpdStrategicControllerImpl) FindAll(c echo.Context) error {
	responses, err := controller.IndikatorPokinOpdStrategicService.FindAll(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, web.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: "INTERNAL_SERVER_ERROR",
		})
	}

	return c.JSON(http.StatusOK, web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   responses,
	})
}
