package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type DB struct {
	DB       *sql.DB
	name     string
	user     string
	password string
	host     string
	port     int
}

func (d *DB) setDatabaseName(dbName string) {
	d.name = dbName
}

func (d *DB) setDatabasePort(portNum int) {
	d.port = portNum
}

func (d *DB) setDatabaseHost(host string) {
	d.host = host
}

func (d *DB) setUsernameAndPassword(user, password string) {
	d.user = user
	d.password = password
}

func (d *DB) connectToDatabase() {
	connectionString := d.generateConnectionString()

	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	} else {
		d.DB = db
	}
}

func (d *DB) generateConnectionString() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		d.host, d.port, d.user, d.password, d.name)
}
