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

type UsersServiceImpl struct {
	UsersRepository repository.UsersRepository
	DB              *sql.DB
	Validate        *validator.Validate
}

func NewUsersService(userRepository repository.UsersRepository, db *sql.DB, validate *validator.Validate) UsersService {
	return &UsersServiceImpl{
		UsersRepository: userRepository,
		DB:              db,
		Validate:        validate,
	}
}

func (service *UsersServiceImpl) Create(ctx context.Context, request web.UsersCreateRequest) (web.UsersCreateResponse, error) {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	hashPassword, err := helper.HashPassword(request.Password)
	helper.PanicIfError(err)

	user := domain.Users{
		Username: request.Username,
		Email:    request.Email,
		Password: hashPassword,
	}

	user = service.UsersRepository.Create(ctx, tx, user)

	return helper.ToUsersResponse(user), nil
}
