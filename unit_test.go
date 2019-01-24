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
	//Build test data
	db := createDBWithInfo()
	db.connectToDatabase()
	contact := Contact{
		FirstName: "Alexander",
		LastName:  "Reno",
		Email:     "ajreno952@gmail.com",
	}

	//Operate on test data
	db.insertNewContact(contact)
	selectedContact := db.grabContact(contact)

	//Check test data
	if selectedContact.FirstName != "Alexander" || selectedContact.LastName != "Reno" || selectedContact.Email != "ajreno952@gmail.com" {
		t.Errorf("Contact has not been successfully created")
	}
}

func TestDeletingContact(t *testing.T) {
	//Build test data
	db := createDBWithInfo()
	db.connectToDatabase()
	contact := Contact{
		FirstName: "john",
		LastName:  "doe",
		Email:     "test@gmail.com",
	}

	//Operate on test data
	db.insertNewContact(contact)

	//Check test data
	if db.deleteContact(contact) != nil {
		t.Errorf("Contact has not been successfully deleted")
	}
}

func TestUpdatingContact(t *testing.T) {
	//Build test data
	db := createDBWithInfo()
	db.connectToDatabase()
	contact := Contact{
		FirstName: "john",
		LastName:  "doe",
		Email:     "1234@gmail.com",
	}

	//Operate on test data
	db.insertNewContact(contact)
	contact.FirstName = "jane"
	db.updateContact(contact)

	grabbedContact := db.grabContactByEmail(contact)

	//Check test data
	if grabbedContact.FirstName != "jane" || grabbedContact.LastName != "doe" || grabbedContact.Email != "1234@gmail.com" {
		t.Errorf("Contact has not been successfully updated")
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
