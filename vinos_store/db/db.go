package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func GetDBConnection() *sql.DB {

	const (
		host     = "localhost"
		port     = 5432
		user     = "postgres"
		password = "dock@2022"
		dbname   = "vinos_store"
	)

	connectionUrl := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", connectionUrl)

	if err != nil {
		panic(err)
	}

	return db

}
