package main

import "testing"

func TestConnectingToDatabase(t *testing.T) {
	db := createDBWithInfo()
	db.connectToDatabase()

	if db.DB == nil {
		t.Errorf("The database has not been established")
	}
}

func TestInsertingNewContacts(t *testing.T) {
	db := createDBWithInfo()
	db.connectToDatabase()

	contact := Contact{
		first_name: "Alexander",
		last_name:  "Reno",
		email:      "ajreno952@gmail.com",
	}

	db.insertNewContact(contact)

	selectedContact := db.grabContact(contact)

	if selectedContact.first_name != "Alexander" && selectedContact.last_name != "Reno" && selectedContact.email != "ajreno952@gmail.com" {
		t.Errorf("Contact has not been successfully created")
	}

}

func createDBWithInfo() DB {
	db := DB{
		name:     "contact_manager",
		user:     "postgres",
		password: "reno",
		host:     "localhost",
		port:     5432,
	}

	return db
}
