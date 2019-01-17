package main

import "testing"

func TestConnectingToDatabase(t *testing.T) {
	db := DB{}
	db.setDatabaseName("contact_manager")
	db.setDatabasePort(5432)
	db.setUsernameAndPassword("postgres", "reno")
	db.setDatabaseHost("localhost")
	db.connectToDatabase()

	if db.DB == nil {
		t.Errorf("The database has not been established")
	}

}

//func TestConnectingToDatabase(t *Testing.T) {}
