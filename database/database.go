package database

import "database/sql"

type Database struct {
	db      *sql.DB
	address string
}

func New(address string) (Database, error) {
	db, err := sql.Open("mysql", "root@tcp(172.17.0.2:3306)/imdb")
	if err != nil {
		return Database{}, err
	}
	return Database{
		db:      db,
		address: address,
	}, nil
}

func (d *Database) Close() error {
	return d.db.Close()
}
