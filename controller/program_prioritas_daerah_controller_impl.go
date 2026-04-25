package controller

import (
	"kupang-service/helper"
	"kupang-service/model/web"
	"kupang-service/service"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ProgramPrioritasDaerahControllerImpl struct {
	ProgramPrioritasDaerahService service.ProgramPrioritasDaerahService
}

func NewProgramPrioritasDaerahControllerImpl(programService service.ProgramPrioritasDaerahService) *ProgramPrioritasDaerahControllerImpl {
	return &ProgramPrioritasDaerahControllerImpl{
		ProgramPrioritasDaerahService: programService,
	}
}

// @Summary Create Program Prioritas Daerah
// @Description Create new Program Prioritas Daerah
// @Tags Program Prioritas Daerah
// @Accept json
// @Produce json
// @Param data body web.ProgramPrioritasDaerahCreateRequest true "Program Prioritas Daerah Create Request"
// @Success 201 {object} web.WebResponse{data=web.ProgramPrioritasDaerahResponse} "Created"
// @Failure 400 {object} web.WebResponse "Bad Request"
// @Failure 500 {object} web.WebResponse "Internal Server Error"
// @Router /program-prioritas-daerahs [post]
func (controller *ProgramPrioritasDaerahControllerImpl) Create(c echo.Context) error {
	request := web.ProgramPrioritasDaerahCreateRequest{}
	err := c.Bind(&request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD_REQUEST",
		})
	}

	response, err := controller.ProgramPrioritasDaerahService.Create(c.Request().Context(), request)
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
		Data:   response,
	})
}

// @Summary Update Program Prioritas Daerah
// @Description Update existing Program Prioritas Daerah by ID
// @Tags Program Prioritas Daerah
// @Accept json
// @Produce json
// @Param id path int true "Program Prioritas Daerah ID"
// @Param data body web.ProgramPrioritasDaerahUpdateRequest true "Program Prioritas Daerah Update Request"
// @Success 200 {object} web.WebResponse{data=web.ProgramPrioritasDaerahResponse} "OK"
// @Failure 400 {object} web.WebResponse "Bad Request"
// @Failure 500 {object} web.WebResponse "Internal Server Error"
// @Router /program-prioritas-daerahs/{id} [put]
func (controller *ProgramPrioritasDaerahControllerImpl) Update(c echo.Context) error {
	request := web.ProgramPrioritasDaerahUpdateRequest{}
	err := c.Bind(&request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD_REQUEST",
		})
	}

	request.Id, err = strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD_REQUEST",
		})
	}

	response, err := controller.ProgramPrioritasDaerahService.Update(c.Request().Context(), request)
	if err != nil {
		if helper.IsValidationError(err) {
			return c.JSON(http.StatusBadRequest, web.WebResponse{
				Code:   http.StatusBadRequest,
				Status: "BAD_REQUEST",
				Data:   err.Error(),
			})
		}
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
		Data:   response,
	})
}

// @Summary Delete Program Prioritas Daerah
// @Description Delete existing Program Prioritas Daerah by ID
// @Tags Program Prioritas Daerah
// @Accept json
// @Produce json
// @Param id path int true "Program Prioritas Daerah ID"
// @Success 200 {object} web.WebResponse "OK"
// @Failure 400 {object} web.WebResponse "Bad Request"
// @Failure 500 {object} web.WebResponse "Internal Server Error"
// @Router /program-prioritas-daerahs/{id} [delete]
func (controller *ProgramPrioritasDaerahControllerImpl) Delete(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD_REQUEST",
		})
	}

	err = controller.ProgramPrioritasDaerahService.Delete(c.Request().Context(), id)
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

// @Summary Get Program Prioritas Daerah by ID
// @Description Get Program Prioritas Daerah detail by ID
// @Tags Program Prioritas Daerah
// @Accept json
// @Produce json
// @Param id path int true "Program Prioritas Daerah ID"
// @Success 200 {object} web.WebResponse{data=web.ProgramPrioritasDaerahResponse} "OK"
// @Failure 400 {object} web.WebResponse "Bad Request"
// @Failure 500 {object} web.WebResponse "Internal Server Error"
// @Router /program-prioritas-daerahs/{id} [get]
func (controller *ProgramPrioritasDaerahControllerImpl) FindById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD_REQUEST",
		})
	}

	response, err := controller.ProgramPrioritasDaerahService.FindById(c.Request().Context(), id)
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
		Data:   response,
	})
}

// @Summary List All Program Prioritas Daerah
// @Description Get list of all Program Prioritas Daerahs
// @Tags Program Prioritas Daerah
// @Accept json
// @Produce json
// @Success 200 {object} web.WebResponse{data=[]web.ProgramPrioritasDaerahResponse} "OK"
// @Failure 500 {object} web.WebResponse "Internal Server Error"
// @Router /program-prioritas-daerahs [get]
func (controller *ProgramPrioritasDaerahControllerImpl) FindAll(c echo.Context) error {
	responses, err := controller.ProgramPrioritasDaerahService.FindAll(c.Request().Context())
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

// @Summary Get Program Prioritas Daerah by Tahun Range
// @Description Get Program Prioritas Daerah list filtered by tahun_awal and tahun_akhir
// @Tags Program Prioritas Daerah
// @Accept json
// @Produce json
// @Param tahun_awal path string true "Tahun Awal"
// @Param tahun_akhir path string true "Tahun Akhir"
// @Success 200 {object} web.WebResponse{data=[]web.ProgramPrioritasDaerahResponse} "OK"
// @Failure 400 {object} web.WebResponse "Bad Request"
// @Failure 500 {object} web.WebResponse "Internal Server Error"
// @Router /program-prioritas-daerahs/tahun/{tahun_awal}/{tahun_akhir} [get]
func (controller *ProgramPrioritasDaerahControllerImpl) FindByTahunRange(c echo.Context) error {
	tahunAwal := c.Param("tahun_awal")
	tahunAkhir := c.Param("tahun_akhir")

	responses, err := controller.ProgramPrioritasDaerahService.FindByTahunRange(c.Request().Context(), tahunAwal, tahunAkhir)
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
