package controller

import (
	"kupang-service/model/web"
	"kupang-service/service"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type PokinOpdControllerImpl struct {
	PokinOpdService service.PokinOpdService
}

func NewPokinOpdControllerImpl(pokinOpdService service.PokinOpdService) *PokinOpdControllerImpl {
	return &PokinOpdControllerImpl{
		PokinOpdService: pokinOpdService,
	}
}

// @Summary Create Pokin Opd
// @Description Create new Pokin Opd
// @Tags Pokin Opd
// @Accept json
// @Produce json
// @Param data body web.PokinOpdCreateRequest true "Pokin Opd Create Request"
// @Success 201 {object} web.WebResponse{data=web.PokinOpdResponse} "Created"
// @Failure 400 {object} web.WebResponse "Bad Request"
// @Failure 500 {object} web.WebResponse "Internal Server Error"
// @Router /api/v1/pokin-opd [post]
func (controller *PokinOpdControllerImpl) Create(c echo.Context) error {
	pokinOpdCreateRequest := web.PokinOpdCreateRequest{}
	err := c.Bind(&pokinOpdCreateRequest)
	if err != nil {
		return c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD_REQUEST",
		})
	}

	pokinOpdResponse, err := controller.PokinOpdService.Create(c.Request().Context(), pokinOpdCreateRequest)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, web.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: "INTERNAL_SERVER_ERROR",
		})
	}

	return c.JSON(http.StatusOK, web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   pokinOpdResponse,
	})
}

// @Summary Update Pokin Opd
// @Description Update existing Pokin Opd by ID
// @Tags Pokin Opd
// @Accept json
// @Produce json
// @Param id path int true "Pokin Opd ID"
// @Param data body web.PokinOpdUpdateRequest true "Pokin Opd Update Request"
// @Success 200 {object} web.WebResponse{data=web.PokinOpdResponse} "OK"
// @Failure 400 {object} web.WebResponse "Bad Request"
// @Failure 500 {object} web.WebResponse "Internal Server Error"
// @Router /api/v1/pokin-opd/{id} [put]
func (controller *PokinOpdControllerImpl) Update(c echo.Context) error {
	pokinOpdUpdateRequest := web.PokinOpdUpdateRequest{}
	err := c.Bind(&pokinOpdUpdateRequest)
	if err != nil {
		return c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD_REQUEST",
		})
	}

	pokinOpdUpdateRequest.Id, err = strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD_REQUEST",
		})
	}

	pokinOpdResponse, err := controller.PokinOpdService.Update(c.Request().Context(), pokinOpdUpdateRequest)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, web.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: "INTERNAL_SERVER_ERROR",
		})
	}

	return c.JSON(http.StatusOK, web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   pokinOpdResponse,
	})
}

// @Summary Update Pokin Opd
// @Description Update existing Pokin Opd by ID
// @Tags Pokin Opd
// @Accept json
// @Produce json
// @Param id path int true "Pokin Opd ID"
// @Param data body web.PokinOpdUpdateRequest true "Pokin Opd Update Request"
// @Success 200 {object} web.WebResponse{data=web.PokinOpdResponse} "OK"
// @Failure 400 {object} web.WebResponse "Bad Request"
// @Failure 500 {object} web.WebResponse "Internal Server Error"
// @Router /api/v1/pokin-opd/{id} [put]
func (controller *PokinOpdControllerImpl) Delete(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD_REQUEST",
		})
	}

	err = controller.PokinOpdService.Delete(c.Request().Context(), id)
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

// @Summary Get Pokin Opd by ID
// @Description Get Pokin Opd detail by ID
// @Tags Pokin Opd
// @Accept json
// @Produce json
// @Param id path int true "Pokin Opd ID"
// @Success 200 {object} web.WebResponse{data=web.PokinOpdResponse} "OK"
// @Failure 400 {object} web.WebResponse "Bad Request"
// @Failure 500 {object} web.WebResponse "Internal Server Error"
// @Router /api/v1/pokin-opd/{id} [get]
func (controller *PokinOpdControllerImpl) FindById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD_REQUEST",
		})
	}

	pokinOpdResponse, err := controller.PokinOpdService.FindById(c.Request().Context(), id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, web.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: "INTERNAL_SERVER_ERROR",
		})
	}

	return c.JSON(http.StatusOK, web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   pokinOpdResponse,
	})
}

// @Summary List All Pokin Opd
// @Description Get list of all Pokin Opd
// @Tags Pokin Opd
// @Accept json
// @Produce json
// @Success 200 {object} web.WebResponse{data=[]web.PokinOpdResponse} "OK"
// @Failure 500 {object} web.WebResponse "Internal Server Error"
// @Router /api/v1/pokin-opd [get]
func (controller *PokinOpdControllerImpl) FindAll(c echo.Context) error {
	pokinOpdResponses, err := controller.PokinOpdService.FindAll(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, web.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: "INTERNAL_SERVER_ERROR",
		})
	}

	return c.JSON(http.StatusOK, web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   pokinOpdResponses,
	})
}
