package service

import (
	"context"

	"github.com/nuraziz04/echo-restful-api-v2/model/web"
)

type UsersService interface {
	Create(ctx context.Context, request web.UsersCreateRequest) (web.UsersCreateResponse, error)
	FindUserById(ctx context.Context, id int) (web.UsersCreateResponse, error)
	UpdatePassword(ctx context.Context, request web.UsersUpdatePasswordRequest) (web.UsersCreateResponse, error)
}
