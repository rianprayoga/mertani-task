package main

import (
	"device-service/cmd/httpserver"
	"device-service/internal/repository"
	"device-service/internal/repository/pg"
	"flag"
)

type application struct {
	DSN      string
	HttpPort string
	Db       repository.Repo
}

func main() {

	var app application

	flag.StringVar(&app.DSN, "dsn", "host=localhost port=5432 user=mertani password=mertani dbname=mertani_devices timezone=UTC", "Postgres connection string")
	flag.StringVar(&app.HttpPort, "http-port", "8081", "Port for http inventories service")
	flag.Parse()

	conn, err := app.connectDb()
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	app.Db = &pg.PgRepo{
		DB: conn,
	}

	httpServer := httpserver.NewHttpServer(app.HttpPort, app.Db)
	err = httpServer.Run()
	if err != nil {
		panic(err)
	}

}
