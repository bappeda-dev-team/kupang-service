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
	repository.NewPokinOpdStrategicRepositoryImpl,
	wire.Bind(new(repository.PokinOpdStrategicRepository), new(*repository.PokinOpdStrategicRepositoryImpl)),
	service.NewIndikatorPokinOpdStrategicServiceImpl,
	wire.Bind(new(service.IndikatorPokinOpdStrategicService), new(*service.IndikatorPokinOpdStrategicServiceImpl)),
	controller.NewIndikatorPokinOpdStrategicControllerImpl,
	wire.Bind(new(controller.IndikatorPokinOpdStrategicController), new(*controller.IndikatorPokinOpdStrategicControllerImpl)),
)

var indikatorPokinOpdTacticalSet = wire.NewSet(
	repository.NewIndikatorPokinOpdTacticalRepositoryImpl,
	wire.Bind(new(repository.IndikatorPokinOpdTacticalRepository), new(*repository.IndikatorPokinOpdTacticalRepositoryImpl)),
	repository.NewPokinOpdTacticalRepositoryImpl,
	wire.Bind(new(repository.PokinOpdTacticalRepository), new(*repository.PokinOpdTacticalRepositoryImpl)),
	service.NewIndikatorPokinOpdTacticalServiceImpl,
	wire.Bind(new(service.IndikatorPokinOpdTacticalService), new(*service.IndikatorPokinOpdTacticalServiceImpl)),
	controller.NewIndikatorPokinOpdTacticalControllerImpl,
	wire.Bind(new(controller.IndikatorPokinOpdTacticalController), new(*controller.IndikatorPokinOpdTacticalControllerImpl)),
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
		pokinOpdSet,
		pokinOpdStrategicSet,
		pokinOpdTacticalSet,
		indikatorPokinOpdSet,
		indikatorPokinOpdStrategicSet,
		indikatorPokinOpdTacticalSet,
		tujuanPokinOpdSet,
		targetPokinOpdSet,
		targetPokinOpdStrategicSet,
		targetPokinOpdTacticalSet,
		pohonKinerjaSet,
		app.NewRouter,
	)
	return nil
}
