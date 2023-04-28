package routes

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/nuraziz04/echo-restful-api-v2/controller"
	"github.com/nuraziz04/echo-restful-api-v2/database"
	"github.com/nuraziz04/echo-restful-api-v2/middlewares"
	"github.com/nuraziz04/echo-restful-api-v2/repository"
	"github.com/nuraziz04/echo-restful-api-v2/service"
)

func Init() *echo.Echo {
	e := echo.New()
	v := validator.New()

	db := database.CreateConn()

	userRepository := repository.NewUsersRepository()
	userService := service.NewUsersService(userRepository, db, v)
	usersController := controller.NewUsersController(userService)

	pegawaiRepository := repository.NewPegawaiRepository()
	pegawaiService := service.NewPegawaiService(pegawaiRepository, db, v)
	pegawaiController := controller.NewPegawaiController(pegawaiService)

	loginRepository := repository.NewLoginRepository()
	loginService := service.NewLoginService(loginRepository, db)
	LoginController := controller.NewLoginController(loginService)

	e.Use(middlewares.SetupLoggerConsole())
	e.Use(middlewares.SetupLoggerFile())

	e.Use(middlewares.ErrorPanicHandler)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, this is echo!")
	})

	e.POST("/golang/api/pegawai", pegawaiController.Create, middlewares.IsAuthenticated)
	e.PUT("/golang/api/pegawai/:id", pegawaiController.Update, middlewares.IsAuthenticated)
	e.DELETE("/golang/api/pegawai/:id", pegawaiController.Delete, middlewares.IsAuthenticated)
	e.GET("/golang/api/pegawai/:id", pegawaiController.FindById, middlewares.IsAuthenticated)
	e.GET("/golang/api/pegawai", pegawaiController.FindAll, middlewares.IsAuthenticated)
	e.POST("/golang/api/pegawai-loop", pegawaiController.CreateLoop, middlewares.IsAuthenticated)

	e.POST("/golang/api/users", usersController.Save)
	e.PUT("/golang/api/users/:id", usersController.UpdatePassword, middlewares.IsAuthenticated)
	e.GET("/golang/api/users/:id", usersController.FindUserById, middlewares.IsAuthenticated)

	e.POST("/golang/api/login", LoginController.CheckLogin)

	return e
}
