package databases

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type Database struct {
	Mariadb *sql.DB
	address string
}

func New(address string) (Database, error) {
	db, err := sql.Open("mysql", address)
	if err != nil {
		return Database{}, err
	}
	return Database{
		Mariadb: db,
		address: address,
	}, nil
}

func (d *Database) Close() error {
	return d.Mariadb.Close()
}

func (d *Database) Query(s string, args ...interface{}) (*sql.Rows, error) {
	return d.Mariadb.Query(s, args)
}
