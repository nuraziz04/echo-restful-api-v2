package service

import (
	"context"
	"database/sql"
	"errors"

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

func (service *UsersServiceImpl) FindUserById(ctx context.Context, id int) (web.UsersCreateResponse, error) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	usr, err := service.UsersRepository.FindUserById(ctx, tx, id)
	helper.PanicIfError(err)

	return helper.ToUsersResponse(usr), nil
}

func (service *UsersServiceImpl) UpdatePassword(ctx context.Context, request web.UsersUpdatePasswordRequest) (web.UsersCreateResponse, error) {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	// cek user apakah ada atau tidak
	usr, err := service.UsersRepository.FindUserById(ctx, tx, request.Id)
	helper.PanicIfError(err)

	// cek password doesn't match
	match, err := helper.CheckPasswordHash(usr.Password, request.OldPassword)
	if !match {
		helper.PanicIfError(errors.New("Hash and password doesn't match."))
	}

	// cek confirm password
	if request.NewPassword != request.ConfirmPassword {
		helper.PanicIfError(errors.New("New Password and Confirm Password doesn't match."))
	}

	hashPassword, err := helper.HashPassword(request.NewPassword)

	usr.Password = hashPassword

	usr = service.UsersRepository.UpdatePassword(ctx, tx, usr)

	return helper.ToUsersResponse(usr), nil
}
