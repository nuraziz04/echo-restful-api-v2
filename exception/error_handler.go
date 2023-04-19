package exception

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/nuraziz04/echo-restful-api-v2/model/web"
)

func ErrorPanicHandler(err error, c echo.Context) {
	notfound(err, c)
}

func notfound(err error, c echo.Context) {
	fmt.Println("PRINT ERROR :", err)
	webResponse := web.WebResponse{
		Code:   http.StatusNotFound,
		Status: "NOT FOUND",
		Data: map[string]string{
			"message": err.Error(),
		},
	}

	c.JSON(http.StatusNotFound, webResponse)
}
