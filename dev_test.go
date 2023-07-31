package be_ku

import (
	"context"
	"fmt"
	"reflect"
	"testing"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	model "github.com/Febriand1/be_ku/model"
	modul "github.com/Febriand1/be_ku/modul"
)

// insert
func TestInsertMembers(t *testing.T) {
	m_nama := "Aurelia Syadyra W."
	m_study := "UPN Yogyakarta"

	insertedID, err := modul.InsertMembers(modul.MongoConn, "members", m_nama, m_study)
	if err != nil {
		t.Errorf("Error inserting data: %v", err)
	}
	fmt.Printf("Data berhasil disimpan dengan id %s", insertedID.Hex())
}

func TestInsertCustomers(t *testing.T) {
	c_nama := "Aurelia Syadyra W."
	c_study := "UPN Yogyakarta"

	insertedID, err := modul.InsertCustomers(modul.MongoConn, "customers", c_nama, c_study)
	if err != nil {
		t.Errorf("Error inserting data: %v", err)
	}
	fmt.Printf("Data berhasil disimpan dengan id %s", insertedID.Hex())
}

func TestInsertIncomes(t *testing.T) {
	qty := 2
	halaman := 2
	uang := 4000

	insertedID, err := modul.InsertIncomes(modul.MongoConn, "incomes", qty, halaman, uang)
	if err != nil {
		t.Errorf("Error inserting data: %v", err)
	}
	fmt.Printf("Data berhasil disimpan dengan id %s", insertedID.Hex())
}

// get
func TestGetMembers(t *testing.T) {
	data := modul.GetMembers(modul.MongoConn, "members")
	fmt.Println(data)
}

func TestGetCustomers(t *testing.T) {
	data := modul.GetCustomers(modul.MongoConn, "customers")
	fmt.Println(data)
}

func TestGetIncomes(t *testing.T) {
	data := modul.GetIncomes(modul.MongoConn, "incomes")
	fmt.Println(data)
}

func TestGetMembersByID(t *testing.T) {
	id := "64c7bdb1123a2ff919286ef7"
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		t.Fatalf("error converting id to ObjectID: %v", err)
	}
	data, err := modul.GetMembersByID(objectID, modul.MongoConn, "members")
	if err != nil {
		t.Fatalf("error calling GetMembersByID: %v", err)
	}
	fmt.Println(data)
}

func TestGetCustomersByID(t *testing.T) {
	id := "64c7bdd8fc3e2a88629e3724"
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		t.Fatalf("error converting id to ObjectID: %v", err)
	}
	data, err := modul.GetCustomersByID(objectID, modul.MongoConn, "customers")
	if err != nil {
		t.Fatalf("error calling GetCustomersByID: %v", err)
	}
	fmt.Println(data)
}

func TestGetIncomesByID(t *testing.T) {
	id := "64c7be14fc301c32dfc224b6"
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		t.Fatalf("error converting id to ObjectID: %v", err)
	}
	data, err := modul.GetIncomesByID(objectID, modul.MongoConn, "incomes")
	if err != nil {
		t.Fatalf("error calling GetIncomesByID: %v", err)
	}
	fmt.Println(data)
}

// update
func TestUpdateMembers(t *testing.T) {
	col := "members"

	// Define a test document
	doc := model.Members{
		ID:      primitive.NewObjectID(),
		M_Nama:  "Aurelia Syadyra W.",
		M_Study: "UPN Yogyakarta",
	}

	// Insert the test document into the collection
	if _, err := modul.MongoConn.Collection(col).InsertOne(context.Background(), doc); err != nil {
		t.Fatalf("Failed to insert test document: %v", err)
	}

	// Define the fields to update
	m_nama := "Aurelia Syadyra W."
	m_study := "UPN Yogyakarta"

	// Call UpdateNilai with the test document ID and updated fields
	if err := modul.UpdateMembers(modul.MongoConn, col, doc.ID, m_nama, m_study); err != nil {
		t.Fatalf("UpdateMembers failed: %v", err)
	}

	// Retrieve the updated document from the collection
	var updatedDoc model.Members
	if err := modul.MongoConn.Collection(col).FindOne(context.Background(), bson.M{"_id": doc.ID}).Decode(&updatedDoc); err != nil {
		t.Fatalf("Failed to retrieve updated document: %v", err)
	}

	if !reflect.DeepEqual(updatedDoc.M_Nama, m_nama) ||
		!reflect.DeepEqual(updatedDoc.M_Study, m_study) {
		t.Fatalf("Document was not updated as expected")
	}
}

func TestUpdateCustomers(t *testing.T) {
	col := "customers"

	// Define a test document
	doc := model.Customers{
		ID:      primitive.NewObjectID(),
		C_Nama:  "Aurelia Syadyra W.",
		C_Study: "UPN Yogyakarta",
	}

	// Insert the test document into the collection
	if _, err := modul.MongoConn.Collection(col).InsertOne(context.Background(), doc); err != nil {
		t.Fatalf("Failed to insert test document: %v", err)
	}

	// Define the fields to update
	c_nama := "Aurelia Syadyra W."
	c_study := "UPN Yogyakarta"

	// Call UpdateNilai with the test document ID and updated fields
	if err := modul.UpdateCustomers(modul.MongoConn, col, doc.ID, c_nama, c_study); err != nil {
		t.Fatalf("UpdateCustomers failed: %v", err)
	}

	// Retrieve the updated document from the collection
	var updatedDoc model.Customers
	if err := modul.MongoConn.Collection(col).FindOne(context.Background(), bson.M{"_id": doc.ID}).Decode(&updatedDoc); err != nil {
		t.Fatalf("Failed to retrieve updated document: %v", err)
	}

	if !reflect.DeepEqual(updatedDoc.C_Nama, c_nama) ||
		!reflect.DeepEqual(updatedDoc.C_Study, c_study) {
		t.Fatalf("Document was not updated as expected")
	}
}

func TestUpdateIncomes(t *testing.T) {
	col := "incomes"

	// Define a test document
	doc := model.Incomes{
		ID:      primitive.NewObjectID(),
		Qty:     2,
		Halaman: 2,
		Uang:    4000,
	}

	// Insert the test document into the collection
	if _, err := modul.MongoConn.Collection(col).InsertOne(context.Background(), doc); err != nil {
		t.Fatalf("Failed to insert test document: %v", err)
	}

	// Define the fields to update
	qty := 2
	halaman := 2
	uang := 4000

	// Call UpdateNilai with the test document ID and updated fields
	if err := modul.UpdateIncomes(modul.MongoConn, col, doc.ID, qty, halaman, uang); err != nil {
		t.Fatalf("UpdateIncomes failed: %v", err)
	}

	// Retrieve the updated document from the collection
	var updatedDoc model.Incomes
	if err := modul.MongoConn.Collection(col).FindOne(context.Background(), bson.M{"_id": doc.ID}).Decode(&updatedDoc); err != nil {
		t.Fatalf("Failed to retrieve updated document: %v", err)
	}

	if !reflect.DeepEqual(updatedDoc.Qty, qty) ||
		!reflect.DeepEqual(updatedDoc.Halaman, halaman) ||
		!reflect.DeepEqual(updatedDoc.Uang, uang) {
		t.Fatalf("Document was not updated as expected")
	}
}

// delete
func TestDeleteMembersByID(t *testing.T) {
	id := "64c7bdb1123a2ff919286ef7" // ID data yang ingin dihapus
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		t.Fatalf("error converting id to ObjectID: %v", err)
	}

	err = modul.DeleteMembersByID(objectID, modul.MongoConn, "members")
	if err != nil {
		t.Fatalf("error calling DeleteMembersByID: %v", err)
	}

	// Verifikasi bahwa data telah dihapus dengan melakukan pengecekan menggunakan GetPresensiFromID
	_, err = modul.GetMembersByID(objectID, modul.MongoConn, "members")
	if err == nil {
		t.Fatalf("expected data to be deleted, but it still exists")
	}
}

func TestDeleteCustomersByID(t *testing.T) {
	id := "64c7bdd8fc3e2a88629e3724" // ID data yang ingin dihapus
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		t.Fatalf("error converting id to ObjectID: %v", err)
	}

	err = modul.DeleteCustomersByID(objectID, modul.MongoConn, "customers")
	if err != nil {
		t.Fatalf("error calling DeleteCustomersByID: %v", err)
	}

	// Verifikasi bahwa data telah dihapus dengan melakukan pengecekan menggunakan GetPresensiFromID
	_, err = modul.GetCustomersByID(objectID, modul.MongoConn, "customers")
	if err == nil {
		t.Fatalf("expected data to be deleted, but it still exists")
	}
}

func TestDeleteIncomesByID(t *testing.T) {
	id := "64c7be14fc301c32dfc224b6" // ID data yang ingin dihapus
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		t.Fatalf("error converting id to ObjectID: %v", err)
	}

	err = modul.DeleteIncomesByID(objectID, modul.MongoConn, "incomes")
	if err != nil {
		t.Fatalf("error calling DeleteIncomesByID: %v", err)
	}

	// Verifikasi bahwa data telah dihapus dengan melakukan pengecekan menggunakan GetPresensiFromID
	_, err = modul.GetIncomesByID(objectID, modul.MongoConn, "incomes")
	if err == nil {
		t.Fatalf("expected data to be deleted, but it still exists")
	}
}

// login
func TestInsertAdmin(t *testing.T) {
	username := "admin"
	password := "admin"

	insertedID, err := modul.InsertAdmin(modul.MongoConn, "admin", username, password)
	if err != nil {
		t.Errorf("Error inserting data: %v", err)
	}
	fmt.Printf("Data berhasil disimpan dengan id %s", insertedID.Hex())
}

func TestLoginAdmin(t *testing.T) {
	username := "admin"
	password := "admin"

	authenticated, err := modul.LoginAdmin(modul.MongoConn, "admin", username, password)
	if err != nil {
		t.Errorf("Error authenticating admin: %v", err)
	}

	if authenticated {
		fmt.Println("Admin authenticated successfully")
	} else {
		t.Errorf("Admin authentication failed")
	}
}
