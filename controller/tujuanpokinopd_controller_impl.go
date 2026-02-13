package controller

import (
	"kupang-service/model/web"
	"kupang-service/service"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type TujuanPokinOpdControllerImpl struct {
	TujuanPokinOpdService service.TujuanPokinOpdService
}

func NewTujuanPokinOpdControllerImpl(tujuanPokinOpdService service.TujuanPokinOpdService) *TujuanPokinOpdControllerImpl {
	return &TujuanPokinOpdControllerImpl{
		TujuanPokinOpdService: tujuanPokinOpdService,
	}
}

// @Summary Create Tujuan Pokin Opd
// @Description Create new Tujuan Pokin Opd
// @Tags Tujuan Pokin Opd
// @Accept json
// @Produce json
// @Param data body web.TujuanPokinOpdCreateRequest true "Tujuan Pokin Opd Create Request"
// @Success 201 {object} web.WebResponse{data=web.TujuanPokinOpdResponse} "Created"
// @Failure 400 {object} web.WebResponse "Bad Request"
// @Failure 500 {object} web.WebResponse "Internal Server Error"
// @Router /api/v1/tujuan-pokin-opd [post]
func (controller *TujuanPokinOpdControllerImpl) Create(c echo.Context) error {
	tujuanPokinOpdCreateRequest := web.TujuanPokinOpdCreateRequest{}
	err := c.Bind(&tujuanPokinOpdCreateRequest)
	if err != nil {
		return c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD_REQUEST",
		})
	}

	tujuanPokinOpdResponse, err := controller.TujuanPokinOpdService.Create(c.Request().Context(), tujuanPokinOpdCreateRequest)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, web.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: "INTERNAL_SERVER_ERROR",
		})
	}

	return c.JSON(http.StatusOK, web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   tujuanPokinOpdResponse,
	})
}

// @Summary Tujuan Update Pokin Opd
// @Description Update existing Tujuan Pokin Opd by ID
// @Tags Tujuan Pokin Opd
// @Accept json
// @Produce json
// @Param id path int true "Tujuan Pokin Opd ID"
// @Param data body web.TujuanPokinOpdUpdateRequest true "Tujuan Pokin Opd Update Request"
// @Success 200 {object} web.WebResponse{data=web.TujuanPokinOpdResponse} "OK"
// @Failure 400 {object} web.WebResponse "Bad Request"
// @Failure 500 {object} web.WebResponse "Internal Server Error"
// @Router /api/v1/tujan-pokin-opd/{id} [put]
func (controller *TujuanPokinOpdControllerImpl) Update(c echo.Context) error {
	tujuanPokinOpdUpdateRequest := web.TujuanPokinOpdUpdateRequest{}
	err := c.Bind(&tujuanPokinOpdUpdateRequest)
	if err != nil {
		return c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD_REQUEST",
		})
	}

	tujuanPokinOpdUpdateRequest.Id, err = strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD_REQUEST",
		})
	}

	tujuanPokinOpdResponse, err := controller.TujuanPokinOpdService.Update(c.Request().Context(), tujuanPokinOpdUpdateRequest)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, web.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: "INTERNAL_SERVER_ERROR",
		})
	}

	return c.JSON(http.StatusOK, web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   tujuanPokinOpdResponse,
	})
}

// @Summary Delete Tujuan Pokin Opd
// @Description Delete existing Tujuan Pokin Opd by ID
// @Tags Tujuan Pokin Opd
// @Accept json
// @Produce json
// @Param id path int true "Tujuan Pokin Opd ID"
// @Success 200 {object} web.WebResponse{data=web.TujuanPokinOpdResponse} "OK"
// @Failure 400 {object} web.WebResponse "Bad Request"
// @Failure 500 {object} web.WebResponse "Internal Server Error"
// @Router /api/v1/tujuan-pokin-opd/{id} [delete]
func (controller *TujuanPokinOpdControllerImpl) Delete(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD_REQUEST",
		})
	}

	err = controller.TujuanPokinOpdService.Delete(c.Request().Context(), id)
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

// @Summary Get Tujuan Pokin Opd by ID
// @Description Get Tujuan Pokin Opd detail by ID
// @Tags Tujuan Pokin Opd
// @Accept json
// @Produce json
// @Param id path int true "Tujuan Pokin Opd ID"
// @Success 200 {object} web.WebResponse{data=web.TujuanPokinOpdResponse} "OK"
// @Failure 400 {object} web.WebResponse "Bad Request"
// @Failure 500 {object} web.WebResponse "Internal Server Error"
// @Router /api/v1/tujuan-pokin-opd/{id} [get]
func (controller *TujuanPokinOpdControllerImpl) FindById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD_REQUEST",
		})
	}

	tujuanPokinOpdResponse, err := controller.TujuanPokinOpdService.FindById(c.Request().Context(), id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, web.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: "INTERNAL_SERVER_ERROR",
		})
	}

	return c.JSON(http.StatusOK, web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   tujuanPokinOpdResponse,
	})
}

// @Summary List All Tujuan Pokin Opd
// @Description Get list of all Tujuan Pokin Opd
// @Tags Tujuan Pokin Opd
// @Accept json
// @Produce json
// @Success 200 {object} web.WebResponse{data=[]web.TujuanPokinOpdResponse} "OK"
// @Failure 500 {object} web.WebResponse "Internal Server Error"
// @Router /api/v1/tujuan-pokin-opd [get]
func (controller *TujuanPokinOpdControllerImpl) FindAll(c echo.Context) error {
	tujuanPokinOpdResponses, err := controller.TujuanPokinOpdService.FindAll(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, web.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: "INTERNAL_SERVER_ERROR",
		})
	}

	return c.JSON(http.StatusOK, web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   tujuanPokinOpdResponses,
	})
}
