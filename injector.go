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

var indikatorPokinOpdSet = wire.NewSet(
	repository.NewIndikatorPokinOpdRepositoryImpl,
	wire.Bind(new(repository.IndikatorPokinOpdRepository), new(*repository.IndikatorPokinOpdRepositoryImpl)),
	service.NewIndikatorPokinOpdServiceImpl,
	wire.Bind(new(service.IndikatorPokinOpdService), new(*service.IndikatorPokinOpdServiceImpl)),
	controller.NewIndikatorPokinOpdControllerImpl,
	wire.Bind(new(controller.IndikatorPokinOpdController), new(*controller.IndikatorPokinOpdControllerImpl)),
)

func InitializedServer() *echo.Echo {
	wire.Build(
		app.GetConnection,
		wire.Value([]validator.Option{}),
		validator.New,
		pokinOpdSet,
		indikatorPokinOpdSet,
		app.NewRouter,
	)
	return nil
}
