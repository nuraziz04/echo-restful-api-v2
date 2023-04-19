package repository

import (
	"context"
	"database/sql"

	"github.com/nuraziz04/echo-restful-api-v2/model/domain"
)

type PegawaiRepository interface {
	Save(ctx context.Context, tx *sql.Tx, pg domain.Pegawai) domain.Pegawai
	Update(ctx context.Context, tx *sql.Tx, pg domain.Pegawai) domain.Pegawai
	Delete(ctx context.Context, tx *sql.Tx, pg domain.Pegawai)
	FindById(ctx context.Context, tx *sql.Tx, pegawaiId int) (domain.Pegawai, error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Pegawai
}
