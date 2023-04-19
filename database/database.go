package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/nuraziz04/echo-restful-api-v2/config"
	"github.com/nuraziz04/echo-restful-api-v2/helper"
)

var db *sql.DB
var err error

func Init() {
	conf := config.GetConfig()

	connectionString := conf.DB_USERNAME + "@tcp(" + conf.DB_HOST + ":" + conf.DB_PORT + ")/" + conf.DB_NAME
	// connectionString := "root@tcp(localhost:3306)/belajar_golang_restful_api"

	db, err = sql.Open("mysql", connectionString)
	helper.PanicIfError(err)

	err = db.Ping()
	helper.PanicIfError(err)
}

func CreateConn() *sql.DB {
	return db
}
