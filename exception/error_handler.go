package exception

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/nuraziz04/echo-restful-api-v2/model/web"
)

func ErrorPanicHandler(err error, c echo.Context) {
	if validationErrors(c, err) {
		return
	}
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

func validationErrors(c echo.Context, err interface{}) bool {
	exception, ok := err.(validator.ValidationErrors)
	if ok {
		webResponse := web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data: map[string]string{
				"message": exception.Error(),
			},
		}

		c.JSON(http.StatusBadRequest, webResponse)
		return true
	} else {
		return false
	}
}
