package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/nuraziz04/echo-restful-api-v2/helper"
	"github.com/nuraziz04/echo-restful-api-v2/model/domain"
)

type LoginRepositoryImpl struct{}

func NewLoginRepository() LoginRepository {
	return &LoginRepositoryImpl{}
}

func (repository *LoginRepositoryImpl) CheckLogin(ctx context.Context, tx *sql.Tx, username string) (domain.Users, error) {
	SQL := "SELECT id, username, password FROM users WHERE username = ?"
	rows, err := tx.QueryContext(ctx, SQL, username)
	helper.PanicIfError(err)

	defer rows.Close()

	usr := domain.Users{}

	if rows.Next() {
		err := rows.Scan(&usr.Id, &usr.Username, &usr.Password)
		helper.PanicIfError(err)
		return usr, nil
	} else {
		return usr, errors.New("Users not found")
	}
}
