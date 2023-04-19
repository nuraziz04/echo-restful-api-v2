package exception

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/nuraziz04/echo-restful-api-v2/model/web"
)

func ErrorPanicHandler(err error, c echo.Context) {
	if ValidationErrors(c, err) {
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

// func validationErrors(c echo.Context, err interface{}) bool {
// 	exception, ok := err.(validator.ValidationErrors)
// 	if ok {
// 		webResponse := web.WebResponse{
// 			Code:   http.StatusBadRequest,
// 			Status: "BAD REQUEST",
// 			Data: map[string]string{
// 				"message": exception.Error(),
// 			},
// 		}

// 		c.JSON(http.StatusBadRequest, webResponse)
// 		return true
// 	} else {
// 		return false
// 	}
// }

func ValidationErrors(c echo.Context, err error) bool {
	report, ok := err.(*echo.HTTPError)
	if !ok {
		report = echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	if castedObject, ok := err.(validator.ValidationErrors); ok {
		for _, err := range castedObject {
			switch err.Tag() {
			case "required":
				report.Message = fmt.Sprintf("%s is required",
					err.Field())
			case "email":
				report.Message = fmt.Sprintf("%s is not valid email",
					err.Field())
			case "gte":
				report.Message = fmt.Sprintf("%s value must be greater than %s",
					err.Field(), err.Param())
			case "lte":
				report.Message = fmt.Sprintf("%s value must be lower than %s",
					err.Field(), err.Param())
			}

			break
		}
		webResponse := web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   report,
		}
		c.JSON(http.StatusBadRequest, webResponse)
		return true
	} else {
		return false
	}
}
