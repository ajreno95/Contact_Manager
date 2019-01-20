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

	err = db.Ping()
	if err != nil {
		panic(err)
	} else {
		d.DB = db
	}
}

func (d *DB) insertNewContact(c Contact) {
	var query string
	if d.DB != nil {
		query = d.generateInsertContactString()
	} else {
		fmt.Println("There is no connection to the database, cannot insert contact")
		return
	}
	_, err := d.DB.Exec(query, c.first_name, c.last_name, c.email)
	if err != nil {
		panic(err)
	}
}

func (d *DB) grabContact(c Contact) Contact {
	query := d.generateSelectContactString()
	var grabbedContact Contact

	err := d.DB.QueryRow(query, c.first_name, c.last_name, c.email).Scan(&grabbedContact.id, &grabbedContact.first_name, &grabbedContact.last_name, &grabbedContact.email)

	if err != nil {
		panic(err)
	}

	return grabbedContact
}

func (d *DB) generateConnectionString() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		d.host, d.port, d.user, d.password, d.name)
}

func (d *DB) generateInsertContactString() string {
	return `INSERT INTO contacts (first_name, last_name, email) VALUES ($1, $2, $3)`
}

func (d *DB) generateSelectContactString() string {
	return `SELECT * from contacts WHERE first_name=$1 AND last_name=$2 AND email=$3`
}
