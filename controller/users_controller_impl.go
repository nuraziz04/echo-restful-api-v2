package controller

import (
	"context"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/nuraziz04/echo-restful-api-v2/helper"
	"github.com/nuraziz04/echo-restful-api-v2/model/web"
	"github.com/nuraziz04/echo-restful-api-v2/service"
)

type UsersControllerImpl struct {
	UsersService service.UsersService
}

func NewUsersController(userService service.UsersService) UsersController {
	return &UsersControllerImpl{
		UsersService: userService,
	}
}

func (controller *UsersControllerImpl) Save(c echo.Context) error {
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	user := new(web.UsersCreateRequest)
	helper.ReadFromRequestBody(c, user)

	userResponse, err := controller.UsersService.Create(ctx, *user)
	helper.PanicIfError(err)

	response := helper.WriteToResponseBody(userResponse)

	return c.JSON(http.StatusOK, response)
}

func (controller *UsersControllerImpl) FindUserById(c echo.Context) error {
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	id, err := strconv.Atoi(c.Param("id"))
	helper.PanicIfError(err)

	userResponse, err := controller.UsersService.FindUserById(ctx, id)
	helper.PanicIfError(err)

	response := helper.WriteToResponseBody(userResponse)

	return c.JSON(http.StatusOK, response)
}

func (controller *UsersControllerImpl) UpdatePassword(c echo.Context) error {
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	id, err := strconv.Atoi(c.Param("id"))
	helper.PanicIfError(err)

	user := new(web.UsersUpdatePasswordRequest)
	helper.ReadFromRequestBody(c, user)

	user.Id = id

	userResponse, err := controller.UsersService.UpdatePassword(ctx, *user)
	helper.PanicIfError(err)

	response := helper.WriteToResponseBody(userResponse)

	return c.JSON(http.StatusOK, response)
}
