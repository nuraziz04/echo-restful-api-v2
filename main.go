package main

import (
	"github.com/nuraziz04/echo-restful-api-v2/database"
	"github.com/nuraziz04/echo-restful-api-v2/routes"
)

func main() {
	database.Init()

	e := routes.Init()

	e.Logger.Fatal(e.Start(":1323"))
}
