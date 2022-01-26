package database

import "database/sql"

type Database struct {
	db      sql.DB
	address string
}
