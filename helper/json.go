package helper

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/nuraziz04/echo-restful-api-v2/model/web"
)

func ReadFromRequestBody(c echo.Context, pg interface{}) {
	if err := c.Bind(pg); err != nil {
		panic(err)
	}
}

func WriteToResponseBody(response interface{}) web.WebResponse {
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "Success",
		Data:   response,
	}

	return webResponse
}
