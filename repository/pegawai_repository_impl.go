package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/nuraziz04/echo-restful-api-v2/helper"
	"github.com/nuraziz04/echo-restful-api-v2/model/domain"
)

type PegawaiRepositoryImpl struct {
}

func NewPegawaiRepository() PegawaiRepository {
	return &PegawaiRepositoryImpl{}
}

func (repository *PegawaiRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, pg domain.Pegawai) domain.Pegawai {
	SQL := "insert into pegawai (name, alamat, telepon) values (?, ?, ?)"
	result, err := tx.ExecContext(ctx, SQL, pg.Name, pg.Alamat, pg.Telepon)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	pg.Id = int(id)
	return pg
}

func (repository *PegawaiRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, pg domain.Pegawai) domain.Pegawai {
	SQL := "UPDATE pegawai SET name = ?, alamat = ?, telepon = ? WHERE id = ?"
	_, err := tx.ExecContext(ctx, SQL, pg.Name, pg.Alamat, pg.Telepon, pg.Id)
	helper.PanicIfError(err)

	return pg
}

func (repository *PegawaiRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, pg domain.Pegawai) {
	SQL := "DELETE FROM pegawai WHERE id = ?"
	_, err := tx.ExecContext(ctx, SQL, pg.Id)
	helper.PanicIfError(err)
}

func (repository *PegawaiRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, pegawaiId int) (domain.Pegawai, error) {
	SQL := "SELECT id, name, alamat, telepon FROM pegawai WHERE id = ?"
	rows, err := tx.QueryContext(ctx, SQL, pegawaiId)
	helper.PanicIfError(err)

	defer rows.Close()

	pg := domain.Pegawai{}
	if rows.Next() {
		err := rows.Scan(&pg.Id, &pg.Name, &pg.Alamat, &pg.Telepon)
		helper.PanicIfError(err)
		return pg, nil
	} else {
		return pg, errors.New("Pegawai is not found bos")
	}
}

func (repository *PegawaiRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Pegawai {
	SQL := "SELECT id, name, alamat, telepon FROM pegawai"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)

	defer rows.Close()

	pgs := []domain.Pegawai{}
	for rows.Next() {
		pg := domain.Pegawai{}
		err := rows.Scan(&pg.Id, &pg.Name, &pg.Alamat, &pg.Telepon)
		helper.PanicIfError(err)

		pgs = append(pgs, pg)
	}
	return pgs
}
