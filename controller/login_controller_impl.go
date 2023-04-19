package controller

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/nuraziz04/echo-restful-api-v2/helper"
	"github.com/nuraziz04/echo-restful-api-v2/model/web"
	"github.com/nuraziz04/echo-restful-api-v2/service"
)

type LoginControllerImpl struct {
	LoginService service.LoginService
}

func NewLoginController(loginService service.LoginService) LoginController {
	return &LoginControllerImpl{
		LoginService: loginService,
	}
}

func (controller *LoginControllerImpl) CheckLogin(c echo.Context) error {
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	usr := new(web.UserLoginRequest)
	helper.ReadFromRequestBody(c, usr)

	userLoginResponse, err := controller.LoginService.CheckLogin(ctx, *usr)
	helper.PanicIfError(err)

	response := helper.WriteToResponseBody(userLoginResponse)

	return c.JSON(http.StatusOK, response)
}
