package db

import "github.com/jmoiron/sqlx"
import _ "github.com/lib/pq"

func ConnectToDB() (*sqlx.DB, error) {
	// TODO: use env
	db, err := sqlx.Connect("postgres",
		"user=hen host=postgres dbname=messenger password=verysecure sslmode=disable")

	if err != nil {
		return nil, err
	}

	return db, err
}
