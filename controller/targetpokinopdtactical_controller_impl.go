package controller

import (
	"kupang-service/model/web"
	"kupang-service/service"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type TargetPokinOpdTacticalControllerImpl struct {
	TargetPokinOpdTacticalService service.TargetPokinOpdTacticalService
}

func NewTargetPokinOpdTacticalControllerImpl(targetService service.TargetPokinOpdTacticalService) *TargetPokinOpdTacticalControllerImpl {
	return &TargetPokinOpdTacticalControllerImpl{
		TargetPokinOpdTacticalService: targetService,
	}
}

// @Summary Create Target Pokin OPD Tactical
// @Description Create new Target Pokin OPD Tactical
// @Tags Target Pokin OPD Tactical
// @Accept json
// @Produce json
// @Param data body web.TargetPokinOpdTacticalCreateRequest true "Target Pokin OPD Tactical Create Request"
// @Success 201 {object} web.WebResponse{data=web.TargetPokinOpdTacticalResponse} "Created"
// @Failure 400 {object} web.WebResponse "Bad Request"
// @Failure 500 {object} web.WebResponse "Internal Server Error"
// @Router /target-pokin-opd-tacticals [post]
func (controller *TargetPokinOpdTacticalControllerImpl) Create(c echo.Context) error {
	request := web.TargetPokinOpdTacticalCreateRequest{}
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD_REQUEST",
		})
	}

	response, err := controller.TargetPokinOpdTacticalService.Create(c.Request().Context(), request)
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

// @Summary Update Target Pokin OPD Tactical
// @Description Update existing Target Pokin OPD Tactical by ID
// @Tags Target Pokin OPD Tactical
// @Accept json
// @Produce json
// @Param id path int true "Target Pokin OPD Tactical ID"
// @Param data body web.TargetPokinOpdTacticalUpdateRequest true "Target Pokin OPD Tactical Update Request"
// @Success 200 {object} web.WebResponse{data=web.TargetPokinOpdTacticalResponse} "OK"
// @Failure 400 {object} web.WebResponse "Bad Request"
// @Failure 500 {object} web.WebResponse "Internal Server Error"
// @Router /target-pokin-opd-tacticals/{id} [put]
func (controller *TargetPokinOpdTacticalControllerImpl) Update(c echo.Context) error {
	request := web.TargetPokinOpdTacticalUpdateRequest{}
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

	response, err := controller.TargetPokinOpdTacticalService.Update(c.Request().Context(), request)
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

// @Summary Delete Target Pokin OPD Tactical
// @Description Delete existing Target Pokin OPD Tactical by ID
// @Tags Target Pokin OPD Tactical
// @Accept json
// @Produce json
// @Param id path int true "Target Pokin OPD Tactical ID"
// @Success 200 {object} web.WebResponse "OK"
// @Failure 400 {object} web.WebResponse "Bad Request"
// @Failure 500 {object} web.WebResponse "Internal Server Error"
// @Router /target-pokin-opd-tacticals/{id} [delete]
func (controller *TargetPokinOpdTacticalControllerImpl) Delete(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD_REQUEST",
		})
	}

	if err := controller.TargetPokinOpdTacticalService.Delete(c.Request().Context(), id); err != nil {
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

// @Summary Get Target Pokin OPD Tactical by ID
// @Description Get Target Pokin OPD Tactical detail by ID
// @Tags Target Pokin OPD Tactical
// @Accept json
// @Produce json
// @Param id path int true "Target Pokin OPD Tactical ID"
// @Success 200 {object} web.WebResponse{data=web.TargetPokinOpdTacticalResponse} "OK"
// @Failure 400 {object} web.WebResponse "Bad Request"
// @Failure 500 {object} web.WebResponse "Internal Server Error"
// @Router /target-pokin-opd-tacticals/{id} [get]
func (controller *TargetPokinOpdTacticalControllerImpl) FindById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD_REQUEST",
		})
	}

	response, err := controller.TargetPokinOpdTacticalService.FindById(c.Request().Context(), id)
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

// @Summary List All Target Pokin OPD Tactical
// @Description Get list of all Target Pokin OPD Tactical
// @Tags Target Pokin OPD Tactical
// @Accept json
// @Produce json
// @Success 200 {object} web.WebResponse{data=[]web.TargetPokinOpdTacticalResponse} "OK"
// @Failure 500 {object} web.WebResponse "Internal Server Error"
// @Router /target-pokin-opd-tacticals [get]
func (controller *TargetPokinOpdTacticalControllerImpl) FindAll(c echo.Context) error {
	responses, err := controller.TargetPokinOpdTacticalService.FindAll(c.Request().Context())
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
