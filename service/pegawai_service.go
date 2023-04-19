package service

import (
	"context"

	"github.com/nuraziz04/echo-restful-api-v2/model/web"
)

type PegawaiService interface {
	Create(ctx context.Context, request web.PegawaiCreateRequest) (web.PegawaiResponse, error)
	Update(ctx context.Context, request web.PegawaiUpdateRequest) (web.PegawaiResponse, error)
	Delete(ctx context.Context, pegawaiId int) error
	FindById(ctx context.Context, pegawaiId int) (web.PegawaiResponse, error)
	FindAll(ctx context.Context) ([]web.PegawaiResponse, error)
}
