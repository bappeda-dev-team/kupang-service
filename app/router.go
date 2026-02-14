package app

import (
	"kupang-service/controller"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func NewRouter(pokinOpdController controller.PokinOpdController, indikatorPokinOpdController controller.IndikatorPokinOpdController, tujuanPokinOpdController controller.TujuanPokinOpdController, targetPokinOpdController controller.TargetPokinOpdController, pohonKinerjaController controller.PohonKinerjaController) *echo.Echo {
	e := echo.New()

	const apiVersion = "/api/v1"
	const pokinOpdBase = apiVersion + "/pokin-opds"
	const indikatorPokinOpdBase = apiVersion + "/indikator-pokin-opd"
	const tujuanPokinOpdBase = apiVersion + "/tujuan-pokin-opd"
	const targetPokinOpdBase = apiVersion + "/target-pokin-opd"
	const pohonKinerjaBase = apiVersion + "/pohon-kinerja-opd"

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

	e.POST(tujuanPokinOpdBase, tujuanPokinOpdController.Create)
	e.PUT(tujuanPokinOpdBase+"/:id", tujuanPokinOpdController.Update)
	e.DELETE(tujuanPokinOpdBase+"/:id", tujuanPokinOpdController.Delete)
	e.GET(tujuanPokinOpdBase+"/:id", tujuanPokinOpdController.FindById)
	e.GET(tujuanPokinOpdBase, tujuanPokinOpdController.FindAll)

	e.POST(targetPokinOpdBase, targetPokinOpdController.Create)
	e.PUT(targetPokinOpdBase+"/:id", targetPokinOpdController.Update)
	e.DELETE(targetPokinOpdBase+"/:id", targetPokinOpdController.Delete)
	e.GET(targetPokinOpdBase+"/:id", targetPokinOpdController.FindById)
	e.GET(targetPokinOpdBase, targetPokinOpdController.FindAll)

	e.GET(pohonKinerjaBase+"/:kode_opd/:tahun", pohonKinerjaController.FindByKodeOpdAndTahun)

	return e
}
