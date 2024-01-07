package db

import (
	"database/sql"
	"fmt"

	"github.com/gaspartv/API-GO-integration-with-postgresql/configs"
	_ "github.com/lib/pq"
)

func OpenConnection() (*sql.DB, error) {
	conf := configs.GetDb()

	sc := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		conf.Host, conf.Port, conf.User, conf.Pass, conf.Database,
	)

	conn, err := sql.Open("postgresql", sc)
	if err != nil {
		panic(err)
	}

	err = conn.Ping()

	return conn, err
}
