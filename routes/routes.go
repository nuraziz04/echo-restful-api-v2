package routes

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/nuraziz04/echo-restful-api-v2/controller"
	"github.com/nuraziz04/echo-restful-api-v2/database"
	"github.com/nuraziz04/echo-restful-api-v2/exception"
	"github.com/nuraziz04/echo-restful-api-v2/middleware"
	"github.com/nuraziz04/echo-restful-api-v2/repository"
	"github.com/nuraziz04/echo-restful-api-v2/service"
)

func Init() *echo.Echo {
	e := echo.New()
	v := validator.New()

	db := database.CreateConn()

	pegawaiRepository := repository.NewPegawaiRepository()
	pegawaiService := service.NewPegawaiService(pegawaiRepository, db, v)
	pegawaiController := controller.NewPegawaiController(pegawaiService)

	loginRepository := repository.NewLoginRepository()
	loginService := service.NewLoginService(loginRepository, db)
	LoginController := controller.NewLoginController(loginService)

	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			defer func() {
				if err := recover(); err != nil {
					exception.ErrorPanicHandler(err.(error), c)
				}
			}()
			return next(c)
		}
	})

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, this is echo!")
	})

	e.POST("/pegawai", pegawaiController.Create, middleware.IsAuthenticated)
	e.PUT("/pegawai/:id", pegawaiController.Update, middleware.IsAuthenticated)
	e.DELETE("/pegawai/:id", pegawaiController.Delete, middleware.IsAuthenticated)
	e.GET("/pegawai/:id", pegawaiController.FindById, middleware.IsAuthenticated)
	e.GET("/pegawai", pegawaiController.FindAll, middleware.IsAuthenticated)

	e.POST("/login", LoginController.CheckLogin)

	return e
}
