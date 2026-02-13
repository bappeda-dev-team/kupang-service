package app

import (
	"kupang-service/controller"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func NewRouter(pokinOpdController controller.PokinOpdController, indikatorPokinOpdController controller.IndikatorPokinOpdController) *echo.Echo {
	e := echo.New()

	const apiVersion = "/api/v1"
	const pokinOpdBase = apiVersion + "/pokin-opds"
	const indikatorPokinOpdBase = apiVersion + "/indikator-pokin-opd"

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	// e.Use(middleware.CORS())

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.POST(pokinOpdBase, pokinOpdController.Create)
	e.PUT(pokinOpdBase+"/:id", pokinOpdController.Update)
	e.DELETE(pokinOpdBase+"/:id", pokinOpdController.Delete)
	e.GET(pokinOpdBase+"/:id", pokinOpdController.FindById)
	e.GET(pokinOpdBase, pokinOpdController.FindAll)

	e.POST(indikatorPokinOpdBase, indikatorPokinOpdController.Create)
	e.PUT(indikatorPokinOpdBase+"/:id", indikatorPokinOpdController.Update)
	e.DELETE(indikatorPokinOpdBase+"/:id", indikatorPokinOpdController.Delete)
	e.GET(indikatorPokinOpdBase+"/:id", indikatorPokinOpdController.FindById)
	e.GET(indikatorPokinOpdBase, indikatorPokinOpdController.FindAll)

	return e
}
