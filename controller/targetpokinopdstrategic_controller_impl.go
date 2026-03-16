package controller

import (
	"kupang-service/model/web"
	"kupang-service/service"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type TargetPokinOpdStrategicControllerImpl struct {
	TargetPokinOpdStrategicService service.TargetPokinOpdStrategicService
}

func NewTargetPokinOpdStrategicControllerImpl(targetService service.TargetPokinOpdStrategicService) *TargetPokinOpdStrategicControllerImpl {
	return &TargetPokinOpdStrategicControllerImpl{
		TargetPokinOpdStrategicService: targetService,
	}
}

// @Summary Create Target Pokin OPD Strategic
// @Description Create new Target Pokin OPD Strategic
// @Tags Target Pokin OPD Strategic
// @Accept json
// @Produce json
// @Param data body web.TargetPokinOpdStrategicCreateRequest true "Target Pokin OPD Strategic Create Request"
// @Success 201 {object} web.WebResponse{data=web.TargetPokinOpdStrategicResponse} "Created"
// @Failure 400 {object} web.WebResponse "Bad Request"
// @Failure 500 {object} web.WebResponse "Internal Server Error"
// @Router /api/v1/target-pokin-opd-strategics [post]
func (controller *TargetPokinOpdStrategicControllerImpl) Create(c echo.Context) error {
	request := web.TargetPokinOpdStrategicCreateRequest{}
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD_REQUEST",
		})
	}

	response, err := controller.TargetPokinOpdStrategicService.Create(c.Request().Context(), request)
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

// @Summary Update Target Pokin OPD Strategic
// @Description Update existing Target Pokin OPD Strategic by ID
// @Tags Target Pokin OPD Strategic
// @Accept json
// @Produce json
// @Param id path int true "Target Pokin OPD Strategic ID"
// @Param data body web.TargetPokinOpdStrategicUpdateRequest true "Target Pokin OPD Strategic Update Request"
// @Success 200 {object} web.WebResponse{data=web.TargetPokinOpdStrategicResponse} "OK"
// @Failure 400 {object} web.WebResponse "Bad Request"
// @Failure 500 {object} web.WebResponse "Internal Server Error"
// @Router /api/v1/target-pokin-opd-strategics/{id} [put]
func (controller *TargetPokinOpdStrategicControllerImpl) Update(c echo.Context) error {
	request := web.TargetPokinOpdStrategicUpdateRequest{}
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

	response, err := controller.TargetPokinOpdStrategicService.Update(c.Request().Context(), request)
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

// @Summary Delete Target Pokin OPD Strategic
// @Description Delete existing Target Pokin OPD Strategic by ID
// @Tags Target Pokin OPD Strategic
// @Accept json
// @Produce json
// @Param id path int true "Target Pokin OPD Strategic ID"
// @Success 200 {object} web.WebResponse "OK"
// @Failure 400 {object} web.WebResponse "Bad Request"
// @Failure 500 {object} web.WebResponse "Internal Server Error"
// @Router /api/v1/target-pokin-opd-strategics/{id} [delete]
func (controller *TargetPokinOpdStrategicControllerImpl) Delete(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD_REQUEST",
		})
	}

	if err := controller.TargetPokinOpdStrategicService.Delete(c.Request().Context(), id); err != nil {
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

// @Summary Get Target Pokin OPD Strategic by ID
// @Description Get Target Pokin OPD Strategic detail by ID
// @Tags Target Pokin OPD Strategic
// @Accept json
// @Produce json
// @Param id path int true "Target Pokin OPD Strategic ID"
// @Success 200 {object} web.WebResponse{data=web.TargetPokinOpdStrategicResponse} "OK"
// @Failure 400 {object} web.WebResponse "Bad Request"
// @Failure 500 {object} web.WebResponse "Internal Server Error"
// @Router /api/v1/target-pokin-opd-strategics/{id} [get]
func (controller *TargetPokinOpdStrategicControllerImpl) FindById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD_REQUEST",
		})
	}

	response, err := controller.TargetPokinOpdStrategicService.FindById(c.Request().Context(), id)
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

// @Summary List All Target Pokin OPD Strategic
// @Description Get list of all Target Pokin OPD Strategic
// @Tags Target Pokin OPD Strategic
// @Accept json
// @Produce json
// @Success 200 {object} web.WebResponse{data=[]web.TargetPokinOpdStrategicResponse} "OK"
// @Failure 500 {object} web.WebResponse "Internal Server Error"
// @Router /api/v1/target-pokin-opd-strategics [get]
func (controller *TargetPokinOpdStrategicControllerImpl) FindAll(c echo.Context) error {
	responses, err := controller.TargetPokinOpdStrategicService.FindAll(c.Request().Context())
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
