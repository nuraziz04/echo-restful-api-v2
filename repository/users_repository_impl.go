package repository

import (
	"context"
	"database/sql"

	"github.com/nuraziz04/echo-restful-api-v2/helper"
	"github.com/nuraziz04/echo-restful-api-v2/model/domain"
)

type UserRepositoryImpl struct {
}

func NewUsersRepository() UsersRepository {
	return &UserRepositoryImpl{}
}

func (repository *UserRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, user domain.Users) domain.Users {
	SQL := "INSERT INTO users (username, email, password) values (?, ?, ?)"
	result, err := tx.ExecContext(ctx, SQL, user.Username, user.Email, user.Password)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	user.Id = int(id)

	return user
}
