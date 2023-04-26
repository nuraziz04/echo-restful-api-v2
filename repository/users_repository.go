package repository

import (
	"context"
	"database/sql"

	"github.com/nuraziz04/echo-restful-api-v2/model/domain"
)

type UsersRepository interface {
	Create(ctx context.Context, tx *sql.Tx, user domain.Users) domain.Users
	FindUserById(ctx context.Context, tx *sql.Tx, id int) (domain.Users, error)
	UpdatePassword(ctx context.Context, tx *sql.Tx, user domain.Users) domain.Users
}
