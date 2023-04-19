package repository

import (
	"context"
	"database/sql"

	"github.com/nuraziz04/echo-restful-api-v2/model/domain"
)

type LoginRepository interface {
	CheckLogin(ctx context.Context, tx *sql.Tx, username string) (domain.Users, error)
}
