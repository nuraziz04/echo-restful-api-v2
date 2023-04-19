package controller

import "github.com/labstack/echo/v4"

type LoginController interface {
	CheckLogin(c echo.Context) error
}
