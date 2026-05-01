package controller

import (
	"kupang-service/helper"
	"kupang-service/model/web"
	"kupang-service/service"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type BidangUrusanControllerImpl struct {
	BidangUrusanService service.BidangUrusanService
}

func NewBidangUrusanControllerImpl(bidangUrusanService service.BidangUrusanService) *BidangUrusanControllerImpl {
	return &BidangUrusanControllerImpl{
		BidangUrusanService: bidangUrusanService,
	}
}

// @Summary Create Bidang Urusan
// @Description Create new Bidang Urusan
// @Tags Bidang Urusan
// @Accept json
// @Produce json
// @Param data body web.BidangUrusanCreateRequest true "Bidang Urusan Create Request"
// @Success 201 {object} web.WebResponse{data=web.BidangUrusanResponse} "Created"
// @Failure 400 {object} web.WebResponse "Bad Request"
// @Failure 500 {object} web.WebResponse "Internal Server Error"
// @Router /bidang-urusans [post]
func (controller *BidangUrusanControllerImpl) Create(c echo.Context) error {
	bidangUrusanCreateRequest := web.BidangUrusanCreateRequest{}
	err := c.Bind(&bidangUrusanCreateRequest)
	if err != nil {
		return c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD_REQUEST",
		})
	}

	bidangUrusanResponse, err := controller.BidangUrusanService.Create(c.Request().Context(), bidangUrusanCreateRequest)
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
		Data:   bidangUrusanResponse,
	})
}

// @Summary Update Bidang Urusan
// @Description Update existing Bidang Urusan by ID
// @Tags Bidang Urusan
// @Accept json
// @Produce json
// @Param id path int true "Bidang Urusan ID"
// @Param data body web.BidangUrusanUpdateRequest true "Bidang Urusan Update Request"
// @Success 200 {object} web.WebResponse{data=web.BidangUrusanResponse} "OK"
// @Failure 400 {object} web.WebResponse "Bad Request"
// @Failure 500 {object} web.WebResponse "Internal Server Error"
// @Router /bidang-urusans/{id} [put]
func (controller *BidangUrusanControllerImpl) Update(c echo.Context) error {
	bidangUrusanUpdateRequest := web.BidangUrusanUpdateRequest{}
	err := c.Bind(&bidangUrusanUpdateRequest)
	if err != nil {
		return c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD_REQUEST",
		})
	}

	bidangUrusanUpdateRequest.Id, err = strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD_REQUEST",
		})
	}

	bidangUrusanResponse, err := controller.BidangUrusanService.Update(c.Request().Context(), bidangUrusanUpdateRequest)
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
		Data:   bidangUrusanResponse,
	})
}

// @Summary Delete Bidang Urusan
// @Description Delete existing Bidang Urusan by ID
// @Tags Bidang Urusan
// @Accept json
// @Produce json
// @Param id path int true "Bidang Urusan ID"
// @Success 200 {object} web.WebResponse{data=web.BidangUrusanResponse} "OK"
// @Failure 400 {object} web.WebResponse "Bad Request"
// @Failure 500 {object} web.WebResponse "Internal Server Error"
// @Router /bidang-urusans/{id} [delete]
func (controller *BidangUrusanControllerImpl) Delete(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD_REQUEST",
		})
	}

	err = controller.BidangUrusanService.Delete(c.Request().Context(), id)
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

// @Summary Get Bidang Urusan by ID
// @Description Get Bidang Urusan detail by ID
// @Tags Bidang Urusan
// @Accept json
// @Produce json
// @Param id path int true "Bidang Urusan ID"
// @Success 200 {object} web.WebResponse{data=web.BidangUrusanResponse} "OK"
// @Failure 400 {object} web.WebResponse "Bad Request"
// @Failure 500 {object} web.WebResponse "Internal Server Error"
// @Router /bidang-urusans/{id} [get]
func (controller *BidangUrusanControllerImpl) FindById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD_REQUEST",
		})
	}

	bidangUrusanResponse, err := controller.BidangUrusanService.FindById(c.Request().Context(), id)
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
		Data:   bidangUrusanResponse,
	})
}

// @Summary List All Bidang Urusan
// @Description Get list of all Bidang Urusan
// @Tags Bidang Urusan
// @Accept json
// @Produce json
// @Success 200 {object} web.WebResponse{data=[]web.BidangUrusanResponse} "OK"
// @Failure 500 {object} web.WebResponse "Internal Server Error"
// @Router /bidang-urusans [get]
func (controller *BidangUrusanControllerImpl) FindAll(c echo.Context) error {
	bidangUrusanResponses, err := controller.BidangUrusanService.FindAll(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, web.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: "INTERNAL_SERVER_ERROR",
		})
	}

	return c.JSON(http.StatusOK, web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   bidangUrusanResponses,
	})
}
