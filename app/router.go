package app

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
	"kupang-service/controller"
)

func NewRouter(lembagaController controller.LembagaController, periodeController controller.PeriodeController, opdController controller.OpdController, pegawaiController controller.PegawaiController, pokinOpdController controller.PokinOpdController, pokinOpdStrategicController controller.PokinOpdStrategicController, pokinOpdTacticalController controller.PokinOpdTacticalController, pokinOpdOperationalController controller.PokinOpdOperationalController, pokinOpdOperationalNController controller.PokinOpdOperationalNController, indikatorPokinOpdController controller.IndikatorPokinOpdController, indikatorPokinOpdStrategicController controller.IndikatorPokinOpdStrategicController, indikatorPokinOpdTacticalController controller.IndikatorPokinOpdTacticalController, indikatorPokinOpdOperationalController controller.IndikatorPokinOpdOperationalController, indikatorPokinOpdOperationalNController controller.IndikatorPokinOpdOperationalNController, tujuanPokinOpdController controller.TujuanPokinOpdController, targetPokinOpdController controller.TargetPokinOpdController, targetPokinOpdStrategicController controller.TargetPokinOpdStrategicController, targetPokinOpdTacticalController controller.TargetPokinOpdTacticalController, targetPokinOpdOperationalController controller.TargetPokinOpdOperationalController, targetPokinOpdOperationalNController controller.TargetPokinOpdOperationalNController, pohonKinerjaController controller.PohonKinerjaController) *echo.Echo {
	e := echo.New()

	const lembagaBase = "/lembagas"
	const periodeBase = "/periodes"
	const opdBase = "/opds"
	const pegawaiBase = "/pegawais"
	const pokinOpdBase = "/pokin-opds"
	const pokinOpdStrategicBase = "/pokin-opd-strategics"
	const pokinOpdTacticalBase = "/pokin-opd-tacticals"
	const pokinOpdOperationalBase = "/pokin-opd-operationals"
	const pokinOpdOperationalNBase = "/pokin-opd-operational-ns"
	const indikatorPokinOpdBase = "/indikator-pokin-opds"
	const indikatorPokinOpdStrategicBase = "/indikator-pokin-opd-strategics"
	const indikatorPokinOpdTacticalBase = "/indikator-pokin-opd-tacticals"
	const indikatorPokinOpdOperationalBase = "/indikator-pokin-opd-operationals"
	const indikatorPokinOpdOperationalNBase = "/indikator-pokin-opd-operational-ns"
	const tujuanPokinOpdBase = "/tujuan-pokin-opds"
	const targetPokinOpdBase = "/target-pokin-opds"
	const targetPokinOpdStrategicBase = "/target-pokin-opd-strategics"
	const targetPokinOpdTacticalBase = "/target-pokin-opd-tacticals"
	const targetPokinOpdOperationalBase = "/target-pokin-opd-operationals"
	const targetPokinOpdOperationalNBase = "/target-pokin-opd-operational-ns"
	const pohonKinerjaBase = "/pohon-kinerja-opds"

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000"},
		AllowMethods: []string{
			echo.GET,
			echo.POST,
			echo.PUT,
			echo.DELETE,
			echo.OPTIONS,
		},
		AllowHeaders: []string{"Content-Type", "Authorization"},
	}))

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.POST(lembagaBase, lembagaController.Create)
	e.PUT(lembagaBase+"/:id", lembagaController.Update)
	e.DELETE(lembagaBase+"/:id", lembagaController.Delete)
	e.GET(lembagaBase+"/:id", lembagaController.FindById)
	e.GET(lembagaBase, lembagaController.FindAll)

	e.POST(periodeBase, periodeController.Create)
	e.PUT(periodeBase+"/:id", periodeController.Update)
	e.DELETE(periodeBase+"/:id", periodeController.Delete)
	e.GET(periodeBase+"/:id", periodeController.FindById)
	e.GET(periodeBase, periodeController.FindAll)

	e.POST(opdBase, opdController.Create)
	e.PUT(opdBase+"/:id", opdController.Update)
	e.DELETE(opdBase+"/:id", opdController.Delete)
	e.GET(opdBase+"/:id", opdController.FindById)
	e.GET(opdBase, opdController.FindAll)

	e.POST(pegawaiBase, pegawaiController.Create)
	e.PUT(pegawaiBase+"/:id", pegawaiController.Update)
	e.POST(pegawaiBase+"/jabatan", pegawaiController.AddJabatan)
	e.PUT(pegawaiBase+"/jabatan/:id", pegawaiController.UpdateJabatan)
	e.GET(pegawaiBase+"/jabatan", pegawaiController.FindAllJabatan)
	e.DELETE(pegawaiBase+"/:id", pegawaiController.Delete)
	e.GET(pegawaiBase+"/search", pegawaiController.Search)
	e.GET(pegawaiBase+"/opd/:kode_opd", pegawaiController.FindByKodeOpd)
	e.GET(pegawaiBase+"/:id", pegawaiController.FindById)
	e.GET(pegawaiBase, pegawaiController.FindAll)

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

	e.POST(pokinOpdTacticalBase, pokinOpdTacticalController.Create)
	e.PUT(pokinOpdTacticalBase+"/:id", pokinOpdTacticalController.Update)
	e.DELETE(pokinOpdTacticalBase+"/:id", pokinOpdTacticalController.Delete)
	e.GET(pokinOpdTacticalBase+"/:id", pokinOpdTacticalController.FindById)
	e.GET(pokinOpdTacticalBase, pokinOpdTacticalController.FindAll)

	e.POST(pokinOpdOperationalBase, pokinOpdOperationalController.Create)
	e.PUT(pokinOpdOperationalBase+"/:id", pokinOpdOperationalController.Update)
	e.DELETE(pokinOpdOperationalBase+"/:id", pokinOpdOperationalController.Delete)
	e.GET(pokinOpdOperationalBase+"/:id", pokinOpdOperationalController.FindById)
	e.GET(pokinOpdOperationalBase, pokinOpdOperationalController.FindAll)

	e.POST(pokinOpdOperationalNBase, pokinOpdOperationalNController.Create)
	e.PUT(pokinOpdOperationalNBase+"/:id", pokinOpdOperationalNController.Update)
	e.DELETE(pokinOpdOperationalNBase+"/:id", pokinOpdOperationalNController.Delete)
	e.GET(pokinOpdOperationalNBase+"/:id", pokinOpdOperationalNController.FindById)
	e.GET(pokinOpdOperationalNBase, pokinOpdOperationalNController.FindAll)

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

	e.POST(indikatorPokinOpdTacticalBase, indikatorPokinOpdTacticalController.Create)
	e.PUT(indikatorPokinOpdTacticalBase+"/:id", indikatorPokinOpdTacticalController.Update)
	e.DELETE(indikatorPokinOpdTacticalBase+"/:id", indikatorPokinOpdTacticalController.Delete)
	e.GET(indikatorPokinOpdTacticalBase+"/:id", indikatorPokinOpdTacticalController.FindById)
	e.GET(indikatorPokinOpdTacticalBase, indikatorPokinOpdTacticalController.FindAll)

	e.POST(indikatorPokinOpdOperationalBase, indikatorPokinOpdOperationalController.Create)
	e.PUT(indikatorPokinOpdOperationalBase+"/:id", indikatorPokinOpdOperationalController.Update)
	e.DELETE(indikatorPokinOpdOperationalBase+"/:id", indikatorPokinOpdOperationalController.Delete)
	e.GET(indikatorPokinOpdOperationalBase+"/:id", indikatorPokinOpdOperationalController.FindById)
	e.GET(indikatorPokinOpdOperationalBase, indikatorPokinOpdOperationalController.FindAll)

	e.POST(indikatorPokinOpdOperationalNBase, indikatorPokinOpdOperationalNController.Create)
	e.PUT(indikatorPokinOpdOperationalNBase+"/:id", indikatorPokinOpdOperationalNController.Update)
	e.DELETE(indikatorPokinOpdOperationalNBase+"/:id", indikatorPokinOpdOperationalNController.Delete)
	e.GET(indikatorPokinOpdOperationalNBase+"/:id", indikatorPokinOpdOperationalNController.FindById)
	e.GET(indikatorPokinOpdOperationalNBase, indikatorPokinOpdOperationalNController.FindAll)

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

	e.POST(targetPokinOpdTacticalBase, targetPokinOpdTacticalController.Create)
	e.PUT(targetPokinOpdTacticalBase+"/:id", targetPokinOpdTacticalController.Update)
	e.DELETE(targetPokinOpdTacticalBase+"/:id", targetPokinOpdTacticalController.Delete)
	e.GET(targetPokinOpdTacticalBase+"/:id", targetPokinOpdTacticalController.FindById)
	e.GET(targetPokinOpdTacticalBase, targetPokinOpdTacticalController.FindAll)

	e.POST(targetPokinOpdOperationalBase, targetPokinOpdOperationalController.Create)
	e.PUT(targetPokinOpdOperationalBase+"/:id", targetPokinOpdOperationalController.Update)
	e.DELETE(targetPokinOpdOperationalBase+"/:id", targetPokinOpdOperationalController.Delete)
	e.GET(targetPokinOpdOperationalBase+"/:id", targetPokinOpdOperationalController.FindById)
	e.GET(targetPokinOpdOperationalBase, targetPokinOpdOperationalController.FindAll)

	e.POST(targetPokinOpdOperationalNBase, targetPokinOpdOperationalNController.Create)
	e.PUT(targetPokinOpdOperationalNBase+"/:id", targetPokinOpdOperationalNController.Update)
	e.DELETE(targetPokinOpdOperationalNBase+"/:id", targetPokinOpdOperationalNController.Delete)
	e.GET(targetPokinOpdOperationalNBase+"/:id", targetPokinOpdOperationalNController.FindById)
	e.GET(targetPokinOpdOperationalNBase, targetPokinOpdOperationalNController.FindAll)

	e.GET(pohonKinerjaBase+"/:kode_opd/:tahun", pohonKinerjaController.FindByKodeOpdAndTahun)

	return e
}
