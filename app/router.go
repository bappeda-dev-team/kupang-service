package app

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
	"kupang-service/controller"
)

func NewRouter(pokinOpdController controller.PokinOpdController, pokinOpdStrategicController controller.PokinOpdStrategicController, indikatorPokinOpdController controller.IndikatorPokinOpdController, indikatorPokinOpdStrategicController controller.IndikatorPokinOpdStrategicController, tujuanPokinOpdController controller.TujuanPokinOpdController, targetPokinOpdController controller.TargetPokinOpdController, targetPokinOpdStrategicController controller.TargetPokinOpdStrategicController, pohonKinerjaController controller.PohonKinerjaController) *echo.Echo {
	e := echo.New()

	const pokinOpdBase = "/pokin-opds"
	const pokinOpdStrategicBase = "/pokin-opd-strategics"
	const indikatorPokinOpdBase = "/indikator-pokin-opds"
	const indikatorPokinOpdStrategicBase = "/indikator-pokin-opd-strategics"
	const tujuanPokinOpdBase = "/tujuan-pokin-opds"
	const targetPokinOpdBase = "/target-pokin-opds"
	const targetPokinOpdStrategicBase = "/target-pokin-opd-strategics"
	const pohonKinerjaBase = "/pohon-kinerja-opds"

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	// e.Use(middleware.CORS())

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.POST(pokinOpdBase, pokinOpdController.Create)
	e.PUT(pokinOpdBase+"/:id", pokinOpdController.Update)
	e.DELETE(pokinOpdBase+"/:id", pokinOpdController.Delete)
	e.GET(pokinOpdBase+"/:id", pokinOpdController.FindById)
	e.GET(pokinOpdBase, pokinOpdController.FindAll)

	e.POST(pokinOpdStrategicBase, pokinOpdStrategicController.Create)
	e.PUT(pokinOpdStrategicBase+"/:id", pokinOpdStrategicController.Update)
	e.DELETE(pokinOpdStrategicBase+"/:id", pokinOpdStrategicController.Delete)
	e.GET(pokinOpdStrategicBase+"/:id", pokinOpdStrategicController.FindById)
	e.GET(pokinOpdStrategicBase, pokinOpdStrategicController.FindAll)

	e.POST(indikatorPokinOpdBase, indikatorPokinOpdController.Create)
	e.PUT(indikatorPokinOpdBase+"/:id", indikatorPokinOpdController.Update)
	e.DELETE(indikatorPokinOpdBase+"/:id", indikatorPokinOpdController.Delete)
	e.GET(indikatorPokinOpdBase+"/:id", indikatorPokinOpdController.FindById)
	e.GET(indikatorPokinOpdBase, indikatorPokinOpdController.FindAll)

	e.POST(indikatorPokinOpdStrategicBase, indikatorPokinOpdStrategicController.Create)
	e.PUT(indikatorPokinOpdStrategicBase+"/:id", indikatorPokinOpdStrategicController.Update)
	e.DELETE(indikatorPokinOpdStrategicBase+"/:id", indikatorPokinOpdStrategicController.Delete)
	e.GET(indikatorPokinOpdStrategicBase+"/:id", indikatorPokinOpdStrategicController.FindById)
	e.GET(indikatorPokinOpdStrategicBase, indikatorPokinOpdStrategicController.FindAll)

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

	e.POST(targetPokinOpdStrategicBase, targetPokinOpdStrategicController.Create)
	e.PUT(targetPokinOpdStrategicBase+"/:id", targetPokinOpdStrategicController.Update)
	e.DELETE(targetPokinOpdStrategicBase+"/:id", targetPokinOpdStrategicController.Delete)
	e.GET(targetPokinOpdStrategicBase+"/:id", targetPokinOpdStrategicController.FindById)
	e.GET(targetPokinOpdStrategicBase, targetPokinOpdStrategicController.FindAll)

	e.GET(pohonKinerjaBase+"/:kode_opd/:tahun", pohonKinerjaController.FindByKodeOpdAndTahun)

	return e
}
