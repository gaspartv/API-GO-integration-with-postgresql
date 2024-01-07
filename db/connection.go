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

	conn, err := sql.Open("postgres", sc)
	if err != nil {
			return nil, fmt.Errorf("error opening database connection: %w", err)
	}

	err = conn.Ping()
	if err != nil {
			return nil, fmt.Errorf("error pinging database: %w", err)
	}

	return conn, nil
}
