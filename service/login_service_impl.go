package service

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/nuraziz04/echo-restful-api-v2/helper"
	"github.com/nuraziz04/echo-restful-api-v2/model/web"
	"github.com/nuraziz04/echo-restful-api-v2/repository"
)

type LoginServiceImpl struct {
	LoginRepository repository.LoginRepository
	DB              *sql.DB
}

func NewLoginService(loginRepository repository.LoginRepository, db *sql.DB) LoginService {
	return &LoginServiceImpl{
		LoginRepository: loginRepository,
		DB:              db,
	}
}

func (service *LoginServiceImpl) CheckLogin(ctx context.Context, request web.UserLoginRequest) (web.UserLoginResponse, error) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	usr, err := service.LoginRepository.CheckLogin(ctx, tx, request.Username)
	helper.PanicIfError(err)

	// cek password doesn't match
	match, err := helper.CheckPasswordHash(usr.Password, request.Password)
	if !match {
		helper.PanicIfError(errors.New("Hash and password doesn't match."))
	}

	// logic generate token
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = usr.Username
	claims["level"] = "application"
	claims["exp"] = time.Now().Add(time.Minute * 1).Unix()

	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		helper.PanicIfError(errors.New("error generate token"))
	}

	return web.UserLoginResponse{Token: t}, nil
}
