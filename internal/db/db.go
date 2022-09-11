package db

import (
	"database/sql"
	"os"
)

type DB struct {
	Conn *sql.DB
}

func (d *DB) Init(dbPath string) error {
	if _, err := os.Stat(dbPath); os.IsNotExist(err) {
		f, err := os.Create(dbPath)
		if err != nil {
			return err
		}
		defer f.Close()
	}

	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return err
	}

	err = db.Ping()
	if err != nil {
		return err
	}

	d.Conn = db

	err = d.CreateTables()
	if err != nil {
		return err
	}

	return nil
}

func (d *DB) CreateTables() error {
	// users table
	stmt := `CREATE TABLE if not exists users (
		"uid" TEXT primary key,
		"login" TEXT,
		"pass" TEXT
	);`
	_, err := d.Conn.Exec(stmt)
	if err != nil {
		return err
	}

	return nil
}
