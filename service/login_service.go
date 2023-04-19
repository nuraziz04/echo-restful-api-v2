package service

import (
	"context"

	"github.com/nuraziz04/echo-restful-api-v2/model/web"
)

type LoginService interface {
	CheckLogin(ctx context.Context, request web.UserLoginRequest) (web.UserLoginResponse, error)
}
