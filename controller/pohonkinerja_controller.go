package controller

import "github.com/labstack/echo/v4"

type PohonKinerjaController interface {
	FindByKodeOpdAndTahun(c echo.Context) error
}
