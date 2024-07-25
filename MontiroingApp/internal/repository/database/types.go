package database

import "database/sql"

type DBRepository struct {
	db *sql.DB
}
