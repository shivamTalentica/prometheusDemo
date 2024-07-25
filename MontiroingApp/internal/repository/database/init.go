package database

import "database/sql"

func InitDB() (*DBRepository, error) {
	// Todo: Add host and creds of DB in config file
	connStr := "postgres://postgres:forDb@231@localhost:5432/postgres?sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}

	return &DBRepository{db: db}, nil
}
