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

func InitializedServer() *echo.Echo {
	wire.Build(
		app.GetConnection,
		wire.Value([]validator.Option{}),
		validator.New,
		pokinOpdSet,
		app.NewRouter,
	)
	return nil
}
