package controller

import (
	"kupang-service/helper"
	"kupang-service/model/web"
	"kupang-service/service"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type JabatanOpdControllerImpl struct {
	JabatanOpdService service.JabatanOpdService
}

func NewJabatanOpdControllerImpl(jabatanOpdService service.JabatanOpdService) *JabatanOpdControllerImpl {
	return &JabatanOpdControllerImpl{
		JabatanOpdService: jabatanOpdService,
	}
}

// @Summary Create JabatanOpd
// @Description Create new JabatanOpd
// @Tags JabatanOpd
// @Accept json
// @Produce json
// @Param data body web.JabatanOpdCreateRequest true "JabatanOpd Create Request"
// @Success 201 {object} web.WebResponse{data=web.JabatanOpdResponse} "Created"
// @Failure 400 {object} web.WebResponse "Bad Request"
// @Failure 500 {object} web.WebResponse "Internal Server Error"
// @Router /jabatan-opds [post]
func (controller *JabatanOpdControllerImpl) Create(c echo.Context) error {
	jabatanOpdCreateRequest := web.JabatanOpdCreateRequest{}
	err := c.Bind(&jabatanOpdCreateRequest)
	if err != nil {
		return c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD_REQUEST",
		})
	}

	jabatanOpdResponse, err := controller.JabatanOpdService.Create(c.Request().Context(), jabatanOpdCreateRequest)
	if err != nil {
		if helper.IsValidationError(err) {
			return c.JSON(http.StatusBadRequest, web.WebResponse{
				Code:   http.StatusBadRequest,
				Status: "BAD_REQUEST",
				Data:   err.Error(),
			})
		}
		return c.JSON(http.StatusInternalServerError, web.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: "INTERNAL_SERVER_ERROR",
		})
	}

	return c.JSON(http.StatusCreated, web.WebResponse{
		Code:   http.StatusCreated,
		Status: "CREATED",
		Data:   jabatanOpdResponse,
	})
}

// @Summary Update JabatanOpd
// @Description Update existing JabatanOpd by ID
// @Tags JabatanOpd
// @Accept json
// @Produce json
// @Param id path int true "JabatanOpd ID"
// @Param data body web.JabatanOpdUpdateRequest true "JabatanOpd Update Request"
// @Success 200 {object} web.WebResponse{data=web.JabatanOpdResponse} "OK"
// @Failure 400 {object} web.WebResponse "Bad Request"
// @Failure 500 {object} web.WebResponse "Internal Server Error"
// @Router /jabatan-opds/{id} [put]
func (controller *JabatanOpdControllerImpl) Update(c echo.Context) error {
	jabatanOpdUpdateRequest := web.JabatanOpdUpdateRequest{}
	err := c.Bind(&jabatanOpdUpdateRequest)
	if err != nil {
		return c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD_REQUEST",
		})
	}

	jabatanOpdUpdateRequest.Id, err = strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD_REQUEST",
		})
	}

	jabatanOpdResponse, err := controller.JabatanOpdService.Update(c.Request().Context(), jabatanOpdUpdateRequest)
	if err != nil {
		if helper.IsValidationError(err) {
			return c.JSON(http.StatusBadRequest, web.WebResponse{
				Code:   http.StatusBadRequest,
				Status: "BAD_REQUEST",
				Data:   err.Error(),
			})
		}
		return c.JSON(http.StatusInternalServerError, web.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: "INTERNAL_SERVER_ERROR",
		})
	}

	return c.JSON(http.StatusOK, web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   jabatanOpdResponse,
	})
}

// @Summary Delete JabatanOpd
// @Description Delete existing JabatanOpd by ID
// @Tags JabatanOpd
// @Accept json
// @Produce json
// @Param id path int true "JabatanOpd ID"
// @Success 200 {object} web.WebResponse{data=web.JabatanOpdResponse} "OK"
// @Failure 400 {object} web.WebResponse "Bad Request"
// @Failure 500 {object} web.WebResponse "Internal Server Error"
// @Router /jabatan-opds/{id} [delete]
func (controller *JabatanOpdControllerImpl) Delete(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD_REQUEST",
		})
	}

	err = controller.JabatanOpdService.Delete(c.Request().Context(), id)
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

// @Summary Get JabatanOpd by ID
// @Description Get JabatanOpd detail by ID
// @Tags JabatanOpd
// @Accept json
// @Produce json
// @Param id path int true "JabatanOpd ID"
// @Success 200 {object} web.WebResponse{data=web.JabatanOpdResponse} "OK"
// @Failure 400 {object} web.WebResponse "Bad Request"
// @Failure 500 {object} web.WebResponse "Internal Server Error"
// @Router /jabatan-opds/{id} [get]
func (controller *JabatanOpdControllerImpl) FindById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD_REQUEST",
		})
	}

	jabatanOpdResponse, err := controller.JabatanOpdService.FindById(c.Request().Context(), id)
	if err != nil {
		if err.Error() == "id tidak ditemukan" {
			return c.JSON(http.StatusBadRequest, web.WebResponse{
				Code:   http.StatusBadRequest,
				Status: "BAD_REQUEST",
				Data:   err.Error(),
			})
		}
		return c.JSON(http.StatusInternalServerError, web.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: "INTERNAL_SERVER_ERROR",
		})
	}

	return c.JSON(http.StatusOK, web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   jabatanOpdResponse,
	})
}

// @Summary List All JabatanOpd
// @Description Get list of all JabatanOpd
// @Tags JabatanOpd
// @Accept json
// @Produce json
// @Success 200 {object} web.WebResponse{data=[]web.JabatanOpdResponse} "OK"
// @Failure 500 {object} web.WebResponse "Internal Server Error"
// @Router /jabatan-opds [get]
func (controller *JabatanOpdControllerImpl) FindAll(c echo.Context) error {
	jabatanOpdResponses, err := controller.JabatanOpdService.FindAll(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, web.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: "INTERNAL_SERVER_ERROR",
		})
	}

	return c.JSON(http.StatusOK, web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   jabatanOpdResponses,
	})
}
