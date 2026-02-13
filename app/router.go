package app

import (
	"kupang-service/controller"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func NewRouter(pokinOpdController controller.PokinOpdController) *echo.Echo {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	// e.Use(middleware.CORS())

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.POST("/pokinopd", pokinOpdController.Create)
	e.PUT("/pokinopd/:id", pokinOpdController.Update)
	e.DELETE("/pokinopd/:id", pokinOpdController.Delete)
	e.GET("/pokinopd/:id", pokinOpdController.FindById)
	e.GET("/pokinopd", pokinOpdController.FindAll)

	return e
}
