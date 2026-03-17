package controller

import (
	"kupang-service/model/web"
	"kupang-service/service"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type PokinOpdTacticalControllerImpl struct {
	PokinOpdTacticalService service.PokinOpdTacticalService
}

func NewPokinOpdTacticalControllerImpl(pokinOpdTacticalService service.PokinOpdTacticalService) *PokinOpdTacticalControllerImpl {
	return &PokinOpdTacticalControllerImpl{
		PokinOpdTacticalService: pokinOpdTacticalService,
	}
}

// @Summary Create Pokin Opd Tactical
// @Description Create new Pokin Opd Tactical
// @Tags Pokin Opd Tactical
// @Accept json
// @Produce json
// @Param data body web.PokinOpdTacticalCreateRequest true "Pokin Opd Tactical Create Request"
// @Success 201 {object} web.WebResponse{data=web.PokinOpdTacticalResponse} "Created"
// @Failure 400 {object} web.WebResponse "Bad Request"
// @Failure 500 {object} web.WebResponse "Internal Server Error"
// @Router /pokin-opd-tacticals [post]
func (controller *PokinOpdTacticalControllerImpl) Create(c echo.Context) error {
	request := web.PokinOpdTacticalCreateRequest{}
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD_REQUEST",
		})
	}

	response, err := controller.PokinOpdTacticalService.Create(c.Request().Context(), request)
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

// @Summary Update Pokin Opd Tactical
// @Description Update existing Pokin Opd Tactical by ID
// @Tags Pokin Opd Tactical
// @Accept json
// @Produce json
// @Param id path int true "Pokin Opd Tactical ID"
// @Param data body web.PokinOpdTacticalUpdateRequest true "Pokin Opd Tactical Update Request"
// @Success 200 {object} web.WebResponse{data=web.PokinOpdTacticalResponse} "OK"
// @Failure 400 {object} web.WebResponse "Bad Request"
// @Failure 500 {object} web.WebResponse "Internal Server Error"
// @Router /pokin-opd-tacticals/{id} [put]
func (controller *PokinOpdTacticalControllerImpl) Update(c echo.Context) error {
	request := web.PokinOpdTacticalUpdateRequest{}
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

	response, err := controller.PokinOpdTacticalService.Update(c.Request().Context(), request)
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

// @Summary Delete Pokin Opd Tactical
// @Description Delete existing Pokin Opd Tactical by ID
// @Tags Pokin Opd Tactical
// @Accept json
// @Produce json
// @Param id path int true "Pokin Opd Tactical ID"
// @Success 200 {object} web.WebResponse{data=web.PokinOpdTacticalResponse} "OK"
// @Failure 400 {object} web.WebResponse "Bad Request"
// @Failure 500 {object} web.WebResponse "Internal Server Error"
// @Router /pokin-opd-tacticals/{id} [delete]
func (controller *PokinOpdTacticalControllerImpl) Delete(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD_REQUEST",
		})
	}

	if err := controller.PokinOpdTacticalService.Delete(c.Request().Context(), id); err != nil {
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

// @Summary Get Pokin Opd Tactical by ID
// @Description Get Pokin Opd Tactical detail by ID
// @Tags Pokin Opd Tactical
// @Accept json
// @Produce json
// @Param id path int true "Pokin Opd Tactical ID"
// @Success 200 {object} web.WebResponse{data=web.PokinOpdTacticalResponse} "OK"
// @Failure 400 {object} web.WebResponse "Bad Request"
// @Failure 500 {object} web.WebResponse "Internal Server Error"
// @Router /pokin-opd-tacticals/{id} [get]
func (controller *PokinOpdTacticalControllerImpl) FindById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD_REQUEST",
		})
	}

	response, err := controller.PokinOpdTacticalService.FindById(c.Request().Context(), id)
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

// @Summary List All Pokin Opd Tactical
// @Description Get list of all Pokin Opd Tactical
// @Tags Pokin Opd Tactical
// @Accept json
// @Produce json
// @Success 200 {object} web.WebResponse{data=[]web.PokinOpdTacticalResponse} "OK"
// @Failure 500 {object} web.WebResponse "Internal Server Error"
// @Router /pokin-opd-tacticals [get]
func (controller *PokinOpdTacticalControllerImpl) FindAll(c echo.Context) error {
	responses, err := controller.PokinOpdTacticalService.FindAll(c.Request().Context())
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
