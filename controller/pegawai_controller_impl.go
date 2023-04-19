package controller

import (
	"context"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/nuraziz04/echo-restful-api-v2/helper"
	"github.com/nuraziz04/echo-restful-api-v2/model/web"
	"github.com/nuraziz04/echo-restful-api-v2/service"
)

type PegawaiControllerImpl struct {
	PegawaiService service.PegawaiService
}

func NewPegawaiController(pegawaiService service.PegawaiService) PegawaiController {
	return &PegawaiControllerImpl{
		PegawaiService: pegawaiService,
	}
}

func (controller *PegawaiControllerImpl) Create(c echo.Context) error {
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	pg := new(web.PegawaiCreateRequest)
	helper.ReadFromRequestBody(c, pg)

	pegawaiResponse, err := controller.PegawaiService.Create(ctx, *pg)
	helper.PanicIfError(err)

	response := helper.WriteToResponseBody(pegawaiResponse)

	return c.JSON(http.StatusOK, response)
}

func (controller *PegawaiControllerImpl) Update(c echo.Context) error {
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	pg := new(web.PegawaiUpdateRequest)
	helper.ReadFromRequestBody(c, pg)

	pg.Id, _ = strconv.Atoi(c.Param("id"))

	pegawaiResponse, err := controller.PegawaiService.Update(ctx, *pg)
	if err != nil {
		return err
	}

	response := helper.WriteToResponseBody(pegawaiResponse)

	return c.JSON(http.StatusOK, response)
}

func (controller *PegawaiControllerImpl) Delete(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	err := controller.PegawaiService.Delete(ctx, id)
	helper.PanicIfError(err)

	response := helper.WriteToResponseBody(nil)

	return c.JSON(http.StatusOK, response)
}

func (controller *PegawaiControllerImpl) FindById(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	pegawaiResponse, err := controller.PegawaiService.FindById(ctx, id)
	helper.PanicIfError(err)

	response := helper.WriteToResponseBody(pegawaiResponse)

	return c.JSON(http.StatusOK, response)
}

func (controller *PegawaiControllerImpl) FindAll(c echo.Context) error {
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	pegawaiResponses, err := controller.PegawaiService.FindAll(ctx)
	helper.PanicIfError(err)

	response := helper.WriteToResponseBody(pegawaiResponses)

	return c.JSON(http.StatusOK, response)
}

func (controller *PegawaiControllerImpl) CreateLoop(c echo.Context) error {
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	pgs := new(web.PegawaiCreateRequestLoop)
	helper.ReadFromRequestBody(c, pgs)

	pegawaiResponse, err := controller.PegawaiService.CreateLoop(ctx, *pgs)
	helper.PanicIfError(err)

	return c.JSON(http.StatusOK, pegawaiResponse)
}
