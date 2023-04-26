package controller

import "github.com/labstack/echo/v4"

type UsersController interface {
	Save(c echo.Context) error
	FindUserById(c echo.Context) error
	UpdatePassword(c echo.Context) error
}
