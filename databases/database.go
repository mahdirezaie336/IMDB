package databases

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type Database struct {
	mariadb *sql.DB
	address string
}

func New(address string) (Database, error) {
	db, err := sql.Open("mysql", address)
	if err != nil {
		return Database{}, err
	}
	return Database{
		mariadb: db,
		address: address,
	}, nil
}

func (d *Database) Close() error {
	return d.mariadb.Close()
}

func (d *Database) Query(s string) (*sql.Rows, error) {
	return d.mariadb.Query(s)
}
