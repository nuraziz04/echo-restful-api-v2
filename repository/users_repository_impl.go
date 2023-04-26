package repository

import (
	"context"
	"database/sql"
	"errors"

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

func (repository *UserRepositoryImpl) FindUserById(ctx context.Context, tx *sql.Tx, id int) (domain.Users, error) {
	SQL := "SELECT id, username, email, password FROM users WHERE id = ?"
	rows, err := tx.QueryContext(ctx, SQL, id)
	helper.PanicIfError(err)

	defer rows.Close()

	usr := domain.Users{}

	if rows.Next() {
		err := rows.Scan(&usr.Id, &usr.Username, &usr.Email, &usr.Password)
		helper.PanicIfError(err)
		return usr, nil
	} else {
		return usr, errors.New("User not found")
	}
}

func (repository *UserRepositoryImpl) UpdatePassword(ctx context.Context, tx *sql.Tx, user domain.Users) domain.Users {
	SQL := "UPDATE users SET password = ? WHERE username = ?"
	_, err := tx.ExecContext(ctx, SQL, user.Password, user.Username)
	helper.PanicIfError(err)

	return user
}
