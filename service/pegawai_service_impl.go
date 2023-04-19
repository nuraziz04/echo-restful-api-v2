package service

import (
	"context"
	"database/sql"

	"github.com/go-playground/validator/v10"
	"github.com/nuraziz04/echo-restful-api-v2/helper"
	"github.com/nuraziz04/echo-restful-api-v2/model/domain"
	"github.com/nuraziz04/echo-restful-api-v2/model/web"
	"github.com/nuraziz04/echo-restful-api-v2/repository"
)

type PegawaiServiceImpl struct {
	PegawaiRepository repository.PegawaiRepository
	DB                *sql.DB
	Validate          *validator.Validate
}

func NewPegawaiService(pegawaiRespository repository.PegawaiRepository, db *sql.DB, validate *validator.Validate) PegawaiService {
	return &PegawaiServiceImpl{
		PegawaiRepository: pegawaiRespository,
		DB:                db,
		Validate:          validate,
	}
}

func (service *PegawaiServiceImpl) Create(ctx context.Context, request web.PegawaiCreateRequest) (web.PegawaiResponse, error) {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	pg := domain.Pegawai{
		Name:    request.Name,
		Alamat:  request.Alamat,
		Telepon: request.Telepon,
	}

	pg = service.PegawaiRepository.Save(ctx, tx, pg)

	return helper.ToPegawaiResponse(pg), nil
}

func (service *PegawaiServiceImpl) Update(ctx context.Context, request web.PegawaiUpdateRequest) (web.PegawaiResponse, error) {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	pg, err := service.PegawaiRepository.FindById(ctx, tx, request.Id)
	helper.PanicIfError(err)

	pg.Name = request.Name
	pg.Alamat = request.Alamat
	pg.Telepon = request.Telepon

	pg = service.PegawaiRepository.Update(ctx, tx, pg)

	return helper.ToPegawaiResponse(pg), nil
}

func (service *PegawaiServiceImpl) Delete(ctx context.Context, pegawaiId int) error {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	pg, err := service.PegawaiRepository.FindById(ctx, tx, pegawaiId)
	helper.PanicIfError(err)

	service.PegawaiRepository.Delete(ctx, tx, pg)

	return nil
}

func (service *PegawaiServiceImpl) FindById(ctx context.Context, pegawaiId int) (web.PegawaiResponse, error) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	pg, err := service.PegawaiRepository.FindById(ctx, tx, pegawaiId)
	helper.PanicIfError(err)

	return helper.ToPegawaiResponse(pg), nil
}

func (service *PegawaiServiceImpl) FindAll(ctx context.Context) ([]web.PegawaiResponse, error) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	pgs := service.PegawaiRepository.FindAll(ctx, tx)

	return helper.ToPegawaiResponses(pgs), nil
}
