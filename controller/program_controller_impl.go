package controller

import (
	"kupang-service/helper"
	"kupang-service/model/web"
	"kupang-service/service"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ProgramControllerImpl struct {
	ProgramService service.ProgramService
}

func NewProgramControllerImpl(programService service.ProgramService) *ProgramControllerImpl {
	return &ProgramControllerImpl{
		ProgramService: programService,
	}
}

// @Summary Create Program
// @Description Create new Program
// @Tags Program
// @Accept json
// @Produce json
// @Param data body web.ProgramCreateRequest true "Program Create Request"
// @Success 201 {object} web.WebResponse{data=web.ProgramResponse} "Created"
// @Failure 400 {object} web.WebResponse "Bad Request"
// @Failure 500 {object} web.WebResponse "Internal Server Error"
// @Router /programs [post]
func (controller *ProgramControllerImpl) Create(c echo.Context) error {
	programCreateRequest := web.ProgramCreateRequest{}
	err := c.Bind(&programCreateRequest)
	if err != nil {
		return c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD_REQUEST",
		})
	}

	programResponse, err := controller.ProgramService.Create(c.Request().Context(), programCreateRequest)
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
		Data:   programResponse,
	})
}

// @Summary Update Program
// @Description Update existing Program by ID
// @Tags Program
// @Accept json
// @Produce json
// @Param id path int true "Program ID"
// @Param data body web.ProgramUpdateRequest true "Program Update Request"
// @Success 200 {object} web.WebResponse{data=web.ProgramResponse} "OK"
// @Failure 400 {object} web.WebResponse "Bad Request"
// @Failure 500 {object} web.WebResponse "Internal Server Error"
// @Router /programs/{id} [put]
func (controller *ProgramControllerImpl) Update(c echo.Context) error {
	programUpdateRequest := web.ProgramUpdateRequest{}
	err := c.Bind(&programUpdateRequest)
	if err != nil {
		return c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD_REQUEST",
		})
	}

	programUpdateRequest.Id, err = strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD_REQUEST",
		})
	}

	programResponse, err := controller.ProgramService.Update(c.Request().Context(), programUpdateRequest)
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
		Data:   programResponse,
	})
}

// @Summary Delete Program
// @Description Delete existing Program by ID
// @Tags Program
// @Accept json
// @Produce json
// @Param id path int true "Program ID"
// @Success 200 {object} web.WebResponse{data=web.ProgramResponse} "OK"
// @Failure 400 {object} web.WebResponse "Bad Request"
// @Failure 500 {object} web.WebResponse "Internal Server Error"
// @Router /programs/{id} [delete]
func (controller *ProgramControllerImpl) Delete(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD_REQUEST",
		})
	}

	err = controller.ProgramService.Delete(c.Request().Context(), id)
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

// @Summary Get Program by ID
// @Description Get Program detail by ID
// @Tags Program
// @Accept json
// @Produce json
// @Param id path int true "Program ID"
// @Success 200 {object} web.WebResponse{data=web.ProgramResponse} "OK"
// @Failure 400 {object} web.WebResponse "Bad Request"
// @Failure 500 {object} web.WebResponse "Internal Server Error"
// @Router /programs/{id} [get]
func (controller *ProgramControllerImpl) FindById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD_REQUEST",
		})
	}

	programResponse, err := controller.ProgramService.FindById(c.Request().Context(), id)
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
		Data:   programResponse,
	})
}

// @Summary List All Program
// @Description Get list of all Program
// @Tags Program
// @Accept json
// @Produce json
// @Success 200 {object} web.WebResponse{data=[]web.ProgramResponse} "OK"
// @Failure 500 {object} web.WebResponse "Internal Server Error"
// @Router /programs [get]
func (controller *ProgramControllerImpl) FindAll(c echo.Context) error {
	programResponses, err := controller.ProgramService.FindAll(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, web.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: "INTERNAL_SERVER_ERROR",
		})
	}

	return c.JSON(http.StatusOK, web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   programResponses,
	})
}
