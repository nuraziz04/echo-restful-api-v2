package middlewares

import (
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
)

var log = logrus.New()

// log.SetFormatter(&logrus.JSONFormatter{})

func SetupLoggerFile() func(next echo.HandlerFunc) echo.HandlerFunc {
	// log.SetFormatter(&logrus.JSONFormatter{})
	return middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `{"time":"${time_rfc3339_nano}","id":"${id}","remote_ip":"${remote_ip}",` +
			`"host":"${host}","method":"${method}","uri":"${uri}","user_agent":"${user_agent}",` +
			`"status":${status},"error":"${error}"}` + "\n",
		CustomTimeFormat: "2006-01-02 15:04:05.00000",
		Output: func() *os.File {
			file, err := os.OpenFile("logs/app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
			if err != nil {
				log.Fatal(err)
			}
			return file
		}(),
	})
}

func SetupLoggerConsole() func(next echo.HandlerFunc) echo.HandlerFunc {
	// log.SetFormatter(&logrus.JSONFormatter{})
	return middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "time=${time_rfc3339_nano}, status=${status}, method=${method}, uri=${uri}\n",
		// Format: `{"time":"${time_rfc3339_nano}","id":"${id}","remote_ip":"${remote_ip}",` +
		// 	`"host":"${host}","method":"${method}","uri":"${uri}","user_agent":"${user_agent}",` +
		// 	`"status":${status},"error":"${error}"}` + "\n",
		// CustomTimeFormat: "2006-01-02 15:04:05.00000",
	})
}

func SetupLogger() func(next echo.HandlerFunc) echo.HandlerFunc {
	return middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:    true,
		LogStatus: true,
		LogValuesFunc: func(c echo.Context, values middleware.RequestLoggerValues) error {
			log.WithFields(logrus.Fields{
				"URI":    values.URI,
				"status": values.Status,
			}).Error("request")

			return nil
		},
	})
}
