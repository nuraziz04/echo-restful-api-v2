package controller

import (
	"context"
	"net/http"

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
