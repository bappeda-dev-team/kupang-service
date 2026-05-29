//go:build wireinject
// +build wireinject

package main

import (
	"kupang-service/app"
	"kupang-service/controller"
	"kupang-service/repository"
	"kupang-service/service"

	"github.com/go-playground/validator/v10"
	"github.com/google/wire"
	"github.com/labstack/echo/v4"
)

var pokinOpdSet = wire.NewSet(
	repository.NewPokinOpdRepositoryImpl,
	wire.Bind(new(repository.PokinOpdRepository), new(*repository.PokinOpdRepositoryImpl)),
	service.NewPokinOpdServiceImpl,
	wire.Bind(new(service.PokinOpdService), new(*service.PokinOpdServiceImpl)),
	controller.NewPokinOpdControllerImpl,
	wire.Bind(new(controller.PokinOpdController), new(*controller.PokinOpdControllerImpl)),
)

var opdSet = wire.NewSet(
	repository.NewOpdRepositoryImpl,
	wire.Bind(new(repository.OpdRepository), new(*repository.OpdRepositoryImpl)),
	service.NewOpdServiceImpl,
	wire.Bind(new(service.OpdService), new(*service.OpdServiceImpl)),
	controller.NewOpdControllerImpl,
	wire.Bind(new(controller.OpdController), new(*controller.OpdControllerImpl)),
)

var pegawaiSet = wire.NewSet(
	repository.NewPegawaiRepositoryImpl,
	wire.Bind(new(repository.PegawaiRepository), new(*repository.PegawaiRepositoryImpl)),
	service.NewPegawaiServiceImpl,
	wire.Bind(new(service.PegawaiService), new(*service.PegawaiServiceImpl)),
	controller.NewPegawaiControllerImpl,
	wire.Bind(new(controller.PegawaiController), new(*controller.PegawaiControllerImpl)),
)

var lembagaSet = wire.NewSet(
	repository.NewLembagaRepositoryImpl,
	wire.Bind(new(repository.LembagaRepository), new(*repository.LembagaRepositoryImpl)),
	service.NewLembagaServiceImpl,
	wire.Bind(new(service.LembagaService), new(*service.LembagaServiceImpl)),
	controller.NewLembagaControllerImpl,
	wire.Bind(new(controller.LembagaController), new(*controller.LembagaControllerImpl)),
)

var periodeSet = wire.NewSet(
	repository.NewPeriodeRepositoryImpl,
	wire.Bind(new(repository.PeriodeRepository), new(*repository.PeriodeRepositoryImpl)),
	service.NewPeriodeServiceImpl,
	wire.Bind(new(service.PeriodeService), new(*service.PeriodeServiceImpl)),
	controller.NewPeriodeControllerImpl,
	wire.Bind(new(controller.PeriodeController), new(*controller.PeriodeControllerImpl)),
)

var roleSet = wire.NewSet(
	repository.NewRoleRepositoryImpl,
	wire.Bind(new(repository.RoleRepository), new(*repository.RoleRepositoryImpl)),
	service.NewRoleServiceImpl,
	wire.Bind(new(service.RoleService), new(*service.RoleServiceImpl)),
	controller.NewRoleControllerImpl,
	wire.Bind(new(controller.RoleController), new(*controller.RoleControllerImpl)),
)

var musrenbangSet = wire.NewSet(
	repository.NewMusrenbangRepositoryImpl,
	wire.Bind(new(repository.MusrenbangRepository), new(*repository.MusrenbangRepositoryImpl)),
	service.NewMusrenbangServiceImpl,
	wire.Bind(new(service.MusrenbangService), new(*service.MusrenbangServiceImpl)),
	controller.NewMusrenbangControllerImpl,
	wire.Bind(new(controller.MusrenbangController), new(*controller.MusrenbangControllerImpl)),
)

var programPrioritasDaerahSet = wire.NewSet(
	repository.NewProgramPrioritasDaerahRepositoryImpl,
	wire.Bind(new(repository.ProgramPrioritasDaerahRepository), new(*repository.ProgramPrioritasDaerahRepositoryImpl)),
	service.NewProgramPrioritasDaerahServiceImpl,
	wire.Bind(new(service.ProgramPrioritasDaerahService), new(*service.ProgramPrioritasDaerahServiceImpl)),
	controller.NewProgramPrioritasDaerahControllerImpl,
	wire.Bind(new(controller.ProgramPrioritasDaerahController), new(*controller.ProgramPrioritasDaerahControllerImpl)),
)

var pokokPikiranSet = wire.NewSet(
	repository.NewPokokPikiranRepositoryImpl,
	wire.Bind(new(repository.PokokPikiranRepository), new(*repository.PokokPikiranRepositoryImpl)),
	service.NewPokokPikiranServiceImpl,
	wire.Bind(new(service.PokokPikiranService), new(*service.PokokPikiranServiceImpl)),
	controller.NewPokokPikiranControllerImpl,
	wire.Bind(new(controller.PokokPikiranController), new(*controller.PokokPikiranControllerImpl)),
)

var userSet = wire.NewSet(
	repository.NewUserRepositoryImpl,
	wire.Bind(new(repository.UserRepository), new(*repository.UserRepositoryImpl)),
	service.NewUserServiceImpl,
	wire.Bind(new(service.UserService), new(*service.UserServiceImpl)),
	controller.NewUserControllerImpl,
	wire.Bind(new(controller.UserController), new(*controller.UserControllerImpl)),
)

var pokinOpdStrategicSet = wire.NewSet(
	repository.NewPokinOpdStrategicRepositoryImpl,
	wire.Bind(new(repository.PokinOpdStrategicRepository), new(*repository.PokinOpdStrategicRepositoryImpl)),
	service.NewPokinOpdStrategicServiceImpl,
	wire.Bind(new(service.PokinOpdStrategicService), new(*service.PokinOpdStrategicServiceImpl)),
	controller.NewPokinOpdStrategicControllerImpl,
	wire.Bind(new(controller.PokinOpdStrategicController), new(*controller.PokinOpdStrategicControllerImpl)),
)

var pokinOpdTacticalSet = wire.NewSet(
	repository.NewPokinOpdTacticalRepositoryImpl,
	wire.Bind(new(repository.PokinOpdTacticalRepository), new(*repository.PokinOpdTacticalRepositoryImpl)),
	service.NewPokinOpdTacticalServiceImpl,
	wire.Bind(new(service.PokinOpdTacticalService), new(*service.PokinOpdTacticalServiceImpl)),
	controller.NewPokinOpdTacticalControllerImpl,
	wire.Bind(new(controller.PokinOpdTacticalController), new(*controller.PokinOpdTacticalControllerImpl)),
)

var pokinOpdOperationalSet = wire.NewSet(
	repository.NewPokinOpdOperationalRepositoryImpl,
	wire.Bind(new(repository.PokinOpdOperationalRepository), new(*repository.PokinOpdOperationalRepositoryImpl)),
	service.NewPokinOpdOperationalServiceImpl,
	wire.Bind(new(service.PokinOpdOperationalService), new(*service.PokinOpdOperationalServiceImpl)),
	controller.NewPokinOpdOperationalControllerImpl,
	wire.Bind(new(controller.PokinOpdOperationalController), new(*controller.PokinOpdOperationalControllerImpl)),
)

var pokinOpdOperationalNSet = wire.NewSet(
	repository.NewPokinOpdOperationalNRepositoryImpl,
	wire.Bind(new(repository.PokinOpdOperationalNRepository), new(*repository.PokinOpdOperationalNRepositoryImpl)),
	service.NewPokinOpdOperationalNServiceImpl,
	wire.Bind(new(service.PokinOpdOperationalNService), new(*service.PokinOpdOperationalNServiceImpl)),
	controller.NewPokinOpdOperationalNControllerImpl,
	wire.Bind(new(controller.PokinOpdOperationalNController), new(*controller.PokinOpdOperationalNControllerImpl)),
)

var indikatorPokinOpdSet = wire.NewSet(
	repository.NewIndikatorPokinOpdRepositoryImpl,
	wire.Bind(new(repository.IndikatorPokinOpdRepository), new(*repository.IndikatorPokinOpdRepositoryImpl)),
	service.NewIndikatorPokinOpdServiceImpl,
	wire.Bind(new(service.IndikatorPokinOpdService), new(*service.IndikatorPokinOpdServiceImpl)),
	controller.NewIndikatorPokinOpdControllerImpl,
	wire.Bind(new(controller.IndikatorPokinOpdController), new(*controller.IndikatorPokinOpdControllerImpl)),
)

var indikatorPokinOpdStrategicSet = wire.NewSet(
	repository.NewIndikatorPokinOpdStrategicRepositoryImpl,
	wire.Bind(new(repository.IndikatorPokinOpdStrategicRepository), new(*repository.IndikatorPokinOpdStrategicRepositoryImpl)),
	service.NewIndikatorPokinOpdStrategicServiceImpl,
	wire.Bind(new(service.IndikatorPokinOpdStrategicService), new(*service.IndikatorPokinOpdStrategicServiceImpl)),
	controller.NewIndikatorPokinOpdStrategicControllerImpl,
	wire.Bind(new(controller.IndikatorPokinOpdStrategicController), new(*controller.IndikatorPokinOpdStrategicControllerImpl)),
)

var indikatorPokinOpdTacticalSet = wire.NewSet(
	repository.NewIndikatorPokinOpdTacticalRepositoryImpl,
	wire.Bind(new(repository.IndikatorPokinOpdTacticalRepository), new(*repository.IndikatorPokinOpdTacticalRepositoryImpl)),
	service.NewIndikatorPokinOpdTacticalServiceImpl,
	wire.Bind(new(service.IndikatorPokinOpdTacticalService), new(*service.IndikatorPokinOpdTacticalServiceImpl)),
	controller.NewIndikatorPokinOpdTacticalControllerImpl,
	wire.Bind(new(controller.IndikatorPokinOpdTacticalController), new(*controller.IndikatorPokinOpdTacticalControllerImpl)),
)

var indikatorPokinOpdOperationalSet = wire.NewSet(
	repository.NewIndikatorPokinOpdOperationalRepositoryImpl,
	wire.Bind(new(repository.IndikatorPokinOpdOperationalRepository), new(*repository.IndikatorPokinOpdOperationalRepositoryImpl)),
	service.NewIndikatorPokinOpdOperationalServiceImpl,
	wire.Bind(new(service.IndikatorPokinOpdOperationalService), new(*service.IndikatorPokinOpdOperationalServiceImpl)),
	controller.NewIndikatorPokinOpdOperationalControllerImpl,
	wire.Bind(new(controller.IndikatorPokinOpdOperationalController), new(*controller.IndikatorPokinOpdOperationalControllerImpl)),
)

var indikatorPokinOpdOperationalNSet = wire.NewSet(
	repository.NewIndikatorPokinOpdOperationalNRepositoryImpl,
	wire.Bind(new(repository.IndikatorPokinOpdOperationalNRepository), new(*repository.IndikatorPokinOpdOperationalNRepositoryImpl)),
	service.NewIndikatorPokinOpdOperationalNServiceImpl,
	wire.Bind(new(service.IndikatorPokinOpdOperationalNService), new(*service.IndikatorPokinOpdOperationalNServiceImpl)),
	controller.NewIndikatorPokinOpdOperationalNControllerImpl,
	wire.Bind(new(controller.IndikatorPokinOpdOperationalNController), new(*controller.IndikatorPokinOpdOperationalNControllerImpl)),
)

var tujuanPokinOpdSet = wire.NewSet(
	repository.NewTujuanPokinOpdRepositoryImpl,
	wire.Bind(new(repository.TujuanPokinOpdRepository), new(*repository.TujuanPokinOpdRepositoryImpl)),
	service.NewTujuanPokinOpdServiceImpl,
	wire.Bind(new(service.TujuanPokinOpdService), new(*service.TujuanPokinOpdServiceImpl)),
	controller.NewTujuanPokinOpdControllerImpl,
	wire.Bind(new(controller.TujuanPokinOpdController), new(*controller.TujuanPokinOpdControllerImpl)),
)

var targetPokinOpdSet = wire.NewSet(
	repository.NewTargetPokinOpdRepositoryImpl,
	wire.Bind(new(repository.TargetPokinOpdRepository), new(*repository.TargetPokinOpdRepositoryImpl)),
	service.NewTargetPokinOpdServiceImpl,
	wire.Bind(new(service.TargetPokinOpdService), new(*service.TargetPokinOpdServiceImpl)),
	controller.NewTargetPokinOpdControllerImpl,
	wire.Bind(new(controller.TargetPokinOpdController), new(*controller.TargetPokinOpdControllerImpl)),
)

var targetPokinOpdStrategicSet = wire.NewSet(
	repository.NewTargetPokinOpdStrategicRepositoryImpl,
	wire.Bind(new(repository.TargetPokinOpdStrategicRepository), new(*repository.TargetPokinOpdStrategicRepositoryImpl)),
	service.NewTargetPokinOpdStrategicServiceImpl,
	wire.Bind(new(service.TargetPokinOpdStrategicService), new(*service.TargetPokinOpdStrategicServiceImpl)),
	controller.NewTargetPokinOpdStrategicControllerImpl,
	wire.Bind(new(controller.TargetPokinOpdStrategicController), new(*controller.TargetPokinOpdStrategicControllerImpl)),
)

var targetPokinOpdTacticalSet = wire.NewSet(
	repository.NewTargetPokinOpdTacticalRepositoryImpl,
	wire.Bind(new(repository.TargetPokinOpdTacticalRepository), new(*repository.TargetPokinOpdTacticalRepositoryImpl)),
	service.NewTargetPokinOpdTacticalServiceImpl,
	wire.Bind(new(service.TargetPokinOpdTacticalService), new(*service.TargetPokinOpdTacticalServiceImpl)),
	controller.NewTargetPokinOpdTacticalControllerImpl,
	wire.Bind(new(controller.TargetPokinOpdTacticalController), new(*controller.TargetPokinOpdTacticalControllerImpl)),
)

var targetPokinOpdOperationalSet = wire.NewSet(
	repository.NewTargetPokinOpdOperationalRepositoryImpl,
	wire.Bind(new(repository.TargetPokinOpdOperationalRepository), new(*repository.TargetPokinOpdOperationalRepositoryImpl)),
	service.NewTargetPokinOpdOperationalServiceImpl,
	wire.Bind(new(service.TargetPokinOpdOperationalService), new(*service.TargetPokinOpdOperationalServiceImpl)),
	controller.NewTargetPokinOpdOperationalControllerImpl,
	wire.Bind(new(controller.TargetPokinOpdOperationalController), new(*controller.TargetPokinOpdOperationalControllerImpl)),
)

var targetPokinOpdOperationalNSet = wire.NewSet(
	repository.NewTargetPokinOpdOperationalNRepositoryImpl,
	wire.Bind(new(repository.TargetPokinOpdOperationalNRepository), new(*repository.TargetPokinOpdOperationalNRepositoryImpl)),
	service.NewTargetPokinOpdOperationalNServiceImpl,
	wire.Bind(new(service.TargetPokinOpdOperationalNService), new(*service.TargetPokinOpdOperationalNServiceImpl)),
	controller.NewTargetPokinOpdOperationalNControllerImpl,
	wire.Bind(new(controller.TargetPokinOpdOperationalNController), new(*controller.TargetPokinOpdOperationalNControllerImpl)),
)

var bidangUrusanSet = wire.NewSet(
	repository.NewBidangUrusanRepositoryImpl,
	wire.Bind(new(repository.BidangUrusanRepository), new(*repository.BidangUrusanRepositoryImpl)),
	service.NewBidangUrusanServiceImpl,
	wire.Bind(new(service.BidangUrusanService), new(*service.BidangUrusanServiceImpl)),
	controller.NewBidangUrusanControllerImpl,
	wire.Bind(new(controller.BidangUrusanController), new(*controller.BidangUrusanControllerImpl)),
)

var urusanSet = wire.NewSet(
	repository.NewUrusanRepositoryImpl,
	wire.Bind(new(repository.UrusanRepository), new(*repository.UrusanRepositoryImpl)),
	service.NewUrusanServiceImpl,
	wire.Bind(new(service.UrusanService), new(*service.UrusanServiceImpl)),
	controller.NewUrusanControllerImpl,
	wire.Bind(new(controller.UrusanController), new(*controller.UrusanControllerImpl)),
)

var kegiatanSet = wire.NewSet(
	repository.NewKegiatanRepositoryImpl,
	wire.Bind(new(repository.KegiatanRepository), new(*repository.KegiatanRepositoryImpl)),
	service.NewKegiatanServiceImpl,
	wire.Bind(new(service.KegiatanService), new(*service.KegiatanServiceImpl)),
	controller.NewKegiatanControllerImpl,
	wire.Bind(new(controller.KegiatanController), new(*controller.KegiatanControllerImpl)),
)

var programSet = wire.NewSet(
	repository.NewProgramRepositoryImpl,
	wire.Bind(new(repository.ProgramRepository), new(*repository.ProgramRepositoryImpl)),
	service.NewProgramServiceImpl,
	wire.Bind(new(service.ProgramService), new(*service.ProgramServiceImpl)),
	controller.NewProgramControllerImpl,
	wire.Bind(new(controller.ProgramController), new(*controller.ProgramControllerImpl)),
)

var pohonKinerjaSet = wire.NewSet(
	service.NewPohonKinerjaServiceImpl,
	wire.Bind(new(service.PohonKinerjaService), new(*service.PohonKinerjaServiceImpl)),
	controller.NewPohonKinerjaControllerImpl,
	wire.Bind(new(controller.PohonKinerjaController), new(*controller.PohonKinerjaControllerImpl)),
)

func InitializedServer() *echo.Echo {
	wire.Build(
		app.GetConnection,
		wire.Value([]validator.Option{}),
		validator.New,
		periodeSet,
		roleSet,
		musrenbangSet,
		programPrioritasDaerahSet,
		pokokPikiranSet,
		userSet,
		opdSet,
		pegawaiSet,
		lembagaSet,
		pokinOpdSet,
		pokinOpdStrategicSet,
		pokinOpdTacticalSet,
		pokinOpdOperationalSet,
		pokinOpdOperationalNSet,
		indikatorPokinOpdSet,
		indikatorPokinOpdStrategicSet,
		indikatorPokinOpdTacticalSet,
		indikatorPokinOpdOperationalSet,
		indikatorPokinOpdOperationalNSet,
		tujuanPokinOpdSet,
		targetPokinOpdSet,
		targetPokinOpdStrategicSet,
		targetPokinOpdTacticalSet,
		targetPokinOpdOperationalSet,
		targetPokinOpdOperationalNSet,
		bidangUrusanSet,
		urusanSet,
		kegiatanSet,
		programSet,
		pohonKinerjaSet,
		app.NewRouter,
	)
	return nil
}
