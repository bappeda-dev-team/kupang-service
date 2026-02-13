package controller

import (
	"kupang-service/model/web"
	"kupang-service/service"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type TargetPokinOpdControllerImpl struct {
	TargetPokinOpdService service.TargetPokinOpdService
}

func NewTargetPokinOpdControllerImpl(targetPokinOpdService service.TargetPokinOpdService) *TargetPokinOpdControllerImpl {
	return &TargetPokinOpdControllerImpl{
		TargetPokinOpdService: targetPokinOpdService,
	}
}

// @Summary Create Target Pokin Opd
// @Description Create new Target Pokin Opd
// @Tags Target Pokin Opd
// @Accept json
// @Produce json
// @Param data body web.TargetPokinOpdCreateRequest true "Target Pokin Opd Create Request"
// @Success 201 {object} web.WebResponse{data=web.TargetPokinOpdResponse} "Created"
// @Failure 400 {object} web.WebResponse "Bad Request"
// @Failure 500 {object} web.WebResponse "Internal Server Error"
// @Router /api/v1/target-pokin-opd [post]
func (controller *TargetPokinOpdControllerImpl) Create(c echo.Context) error {
	targetPokinOpdCreateRequest := web.TargetPokinOpdCreateRequest{}
	err := c.Bind(&targetPokinOpdCreateRequest)
	if err != nil {
		return c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD_REQUEST",
		})
	}

	targetPokinOpdResponse, err := controller.TargetPokinOpdService.Create(c.Request().Context(), targetPokinOpdCreateRequest)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, web.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: "INTERNAL_SERVER_ERROR",
		})
	}

	return c.JSON(http.StatusOK, web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   targetPokinOpdResponse,
	})
}

// @Summary Update Target Pokin Opd
// @Description Update existing Target Pokin Opd by ID
// @Tags Target Pokin Opd
// @Accept json
// @Produce json
// @Param id path int true "Target Pokin Opd ID"
// @Param data body web.TargetPokinOpdUpdateRequest true "Target Pokin Opd Update Request"
// @Success 200 {object} web.WebResponse{data=web.TargetPokinOpdResponse} "OK"
// @Failure 400 {object} web.WebResponse "Bad Request"
// @Failure 500 {object} web.WebResponse "Internal Server Error"
// @Router /api/v1/target-pokin-opd/{id} [put]
func (controller *TargetPokinOpdControllerImpl) Update(c echo.Context) error {
	targetPokinOpdUpdateRequest := web.TargetPokinOpdUpdateRequest{}
	err := c.Bind(&targetPokinOpdUpdateRequest)
	if err != nil {
		return c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD_REQUEST",
		})
	}

	targetPokinOpdUpdateRequest.Id, err = strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD_REQUEST",
		})
	}

	targetPokinOpdResponse, err := controller.TargetPokinOpdService.Update(c.Request().Context(), targetPokinOpdUpdateRequest)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, web.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: "INTERNAL_SERVER_ERROR",
		})
	}

	return c.JSON(http.StatusOK, web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   targetPokinOpdResponse,
	})
}

// @Summary Delete Target Pokin Opd
// @Description Delete existing Target Pokin Opd by ID
// @Tags Target Pokin Opd
// @Accept json
// @Produce json
// @Param id path int true "Target Pokin Opd ID"
// @Success 200 {object} web.WebResponse "OK"
// @Failure 400 {object} web.WebResponse "Bad Request"
// @Failure 500 {object} web.WebResponse "Internal Server Error"
// @Router /api/v1/target-pokin-opd/{id} [delete]
func (controller *TargetPokinOpdControllerImpl) Delete(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD_REQUEST",
		})
	}

	err = controller.TargetPokinOpdService.Delete(c.Request().Context(), id)
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

// @Summary Get Target Pokin Opd by ID
// @Description Get Target Pokin Opd detail by ID
// @Tags Target Pokin Opd
// @Accept json
// @Produce json
// @Param id path int true "Target Pokin Opd ID"
// @Success 200 {object} web.WebResponse{data=web.TargetPokinOpdResponse} "OK"
// @Failure 400 {object} web.WebResponse "Bad Request"
// @Failure 500 {object} web.WebResponse "Internal Server Error"
// @Router /api/v1/target-pokin-opd/{id} [get]
func (controller *TargetPokinOpdControllerImpl) FindById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD_REQUEST",
		})
	}

	targetPokinOpdResponse, err := controller.TargetPokinOpdService.FindById(c.Request().Context(), id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, web.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: "INTERNAL_SERVER_ERROR",
		})
	}

	return c.JSON(http.StatusOK, web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   targetPokinOpdResponse,
	})
}

// @Summary List All Target Pokin Opd
// @Description Get list of all Target Pokin Opd
// @Tags Target Pokin Opd
// @Accept json
// @Produce json
// @Success 200 {object} web.WebResponse{data=[]web.TargetPokinOpdResponse} "OK"
// @Failure 500 {object} web.WebResponse "Internal Server Error"
// @Router /api/v1/target-pokin-opd [get]
func (controller *TargetPokinOpdControllerImpl) FindAll(c echo.Context) error {
	targetPokinOpdResponses, err := controller.TargetPokinOpdService.FindAll(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, web.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: "INTERNAL_SERVER_ERROR",
		})
	}

	return c.JSON(http.StatusOK, web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   targetPokinOpdResponses,
	})
}
