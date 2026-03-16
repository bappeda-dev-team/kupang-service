package controller

import (
	"kupang-service/model/web"
	"kupang-service/service"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type PokinOpdStrategicControllerImpl struct {
	PokinOpdStrategicService service.PokinOpdStrategicService
}

func NewPokinOpdStrategicControllerImpl(pokinOpdStrategicService service.PokinOpdStrategicService) *PokinOpdStrategicControllerImpl {
	return &PokinOpdStrategicControllerImpl{
		PokinOpdStrategicService: pokinOpdStrategicService,
	}
}

// @Summary Create Pokin Opd Strategic
// @Description Create new Pokin Opd Strategic
// @Tags Pokin Opd Strategic
// @Accept json
// @Produce json
// @Param data body web.PokinOpdStrategicCreateRequest true "Pokin Opd Strategic Create Request"
// @Success 201 {object} web.WebResponse{data=web.PokinOpdStrategicResponse} "Created"
// @Failure 400 {object} web.WebResponse "Bad Request"
// @Failure 500 {object} web.WebResponse "Internal Server Error"
// @Router /api/v1/pokin-opd-strategics [post]
func (controller *PokinOpdStrategicControllerImpl) Create(c echo.Context) error {
	request := web.PokinOpdStrategicCreateRequest{}
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD_REQUEST",
		})
	}

	response, err := controller.PokinOpdStrategicService.Create(c.Request().Context(), request)
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

// @Summary Update Pokin Opd Strategic
// @Description Update existing Pokin Opd Strategic by ID
// @Tags Pokin Opd Strategic
// @Accept json
// @Produce json
// @Param id path int true "Pokin Opd Strategic ID"
// @Param data body web.PokinOpdStrategicUpdateRequest true "Pokin Opd Strategic Update Request"
// @Success 200 {object} web.WebResponse{data=web.PokinOpdStrategicResponse} "OK"
// @Failure 400 {object} web.WebResponse "Bad Request"
// @Failure 500 {object} web.WebResponse "Internal Server Error"
// @Router /api/v1/pokin-opd-strategics/{id} [put]
func (controller *PokinOpdStrategicControllerImpl) Update(c echo.Context) error {
	request := web.PokinOpdStrategicUpdateRequest{}
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD_REQUEST",
		})
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD_REQUEST",
		})
	}
	request.Id = id

	response, err := controller.PokinOpdStrategicService.Update(c.Request().Context(), request)
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

// @Summary Delete Pokin Opd Strategic
// @Description Delete existing Pokin Opd Strategic by ID
// @Tags Pokin Opd Strategic
// @Accept json
// @Produce json
// @Param id path int true "Pokin Opd Strategic ID"
// @Success 200 {object} web.WebResponse{data=web.PokinOpdStrategicResponse} "OK"
// @Failure 400 {object} web.WebResponse "Bad Request"
// @Failure 500 {object} web.WebResponse "Internal Server Error"
// @Router /api/v1/pokin-opd-strategics/{id} [delete]
func (controller *PokinOpdStrategicControllerImpl) Delete(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD_REQUEST",
		})
	}

	if err := controller.PokinOpdStrategicService.Delete(c.Request().Context(), id); err != nil {
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

// @Summary Get Pokin Opd Strategic by ID
// @Description Get Pokin Opd Strategic detail by ID
// @Tags Pokin Opd Strategic
// @Accept json
// @Produce json
// @Param id path int true "Pokin Opd Strategic ID"
// @Success 200 {object} web.WebResponse{data=web.PokinOpdStrategicResponse} "OK"
// @Failure 400 {object} web.WebResponse "Bad Request"
// @Failure 500 {object} web.WebResponse "Internal Server Error"
// @Router /api/v1/pokin-opd-strategics/{id} [get]
func (controller *PokinOpdStrategicControllerImpl) FindById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD_REQUEST",
		})
	}

	response, err := controller.PokinOpdStrategicService.FindById(c.Request().Context(), id)
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

// @Summary List All Pokin Opd Strategic
// @Description Get list of all Pokin Opd Strategic
// @Tags Pokin Opd Strategic
// @Accept json
// @Produce json
// @Success 200 {object} web.WebResponse{data=[]web.PokinOpdStrategicResponse} "OK"
// @Failure 500 {object} web.WebResponse "Internal Server Error"
// @Router /api/v1/pokin-opd-strategics [get]
func (controller *PokinOpdStrategicControllerImpl) FindAll(c echo.Context) error {
	responses, err := controller.PokinOpdStrategicService.FindAll(c.Request().Context())
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
