package middlewares

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/nuraziz04/echo-restful-api-v2/exception"
	"github.com/nuraziz04/echo-restful-api-v2/model/web"
)

var IsAuthenticated = middleware.JWTWithConfig(middleware.JWTConfig{
	SigningKey: []byte("secret"),
	ErrorHandler: func(err error) error {
		response := web.WebResponse{
			Code:   http.StatusUnauthorized,
			Status: "Unauthorization",
			Data: map[string]string{
				"message": "Invalid Token",
			},
		}
		return echo.NewHTTPError(http.StatusUnauthorized, response)
	},
})

var ErrorPanicHandler = func(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		defer func() {
			if err := recover(); err != nil {
				exception.ErrorPanicHandler(err.(error), c)
			}
		}()
		return next(c)
	}
}
