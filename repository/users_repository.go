package repository

import (
	"context"
	"database/sql"

	"github.com/nuraziz04/echo-restful-api-v2/model/domain"
)

type UsersRepository interface {
	Create(ctx context.Context, tx *sql.Tx, user domain.Users) domain.Users
}
