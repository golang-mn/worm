package main

import (
	"database/sql"
	"fmt"
)

type DB struct {
	ConnectionString string // DSN msql ex: username:password@protocol(address)/dbname?param=value
	DB               *sql.DB
}

func (db *DB) Connect() error {
	sqlDB, err := sql.Open("mysql", db.ConnectionString)
	if err != nil {
		fmt.Println("Error opening database connection: ", err.Error())
		return nil
	}

	db.DB = sqlDB

	return nil
}

func main() {

}
