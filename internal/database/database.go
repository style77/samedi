package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

const file string = "samedi.db"
const schemaFile string = "internal/database/schema.sql"

type Database struct {
	db *sql.DB
}

func Init() (*Database, error) {
	db, err := sql.Open("sqlite3", file)
	if err != nil {
		fmt.Println("Error opening database:", err)
		return nil, err
	}

	database := &Database{db: db}
	err = database.CreateTables()
	if err != nil {
		return nil, err
	}

	return database, nil
}

func (d *Database) CreateTables() error {
	schemaSQL, err := os.ReadFile(schemaFile)
	if err != nil {
		fmt.Println("Error reading schema file:", err)
		return err
	}

	_, err = d.db.Exec(string(schemaSQL))
	if err != nil {
		fmt.Println("Error creating tables:", err)
		return err
	}

	return nil
}

func (d *Database) Close() error {
	return d.db.Close()
}

func (d *Database) Query(query string, args ...interface{}) (*sql.Rows, error) {
	return d.db.Query(query, args...)
}

func (d *Database) QueryRow(query string, args ...interface{}) *sql.Row {
	return d.db.QueryRow(query, args...)
}

func (d *Database) Exec(query string, args ...interface{}) (sql.Result, error) {
	return d.db.Exec(query, args...)
}
