package modul

import (
	"context"
	"errors"
	"fmt"
	"os"

	model "github.com/Febriand1/be_ku/model"
	"github.com/aiteung/atdb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var MongoString string = os.Getenv("MONGOSTRING")

var MongoInfo = atdb.DBInfo{
	DBString: MongoString,
	DBName:   "db_ku",
}

var MongoConn = atdb.MongoConnect(MongoInfo)

func InsertOneDoc(db *mongo.Database, collection string, doc interface{}) (insertedID interface{}) {
	insertResult, err := db.Collection(collection).InsertOne(context.TODO(), doc)
	if err != nil {
		fmt.Printf("InsertOneDoc: %v\n", err)
	}
	return insertResult.InsertedID
}

// insert
func InsertMembers(db *mongo.Database, col string, m_nama string, m_study string) (insertedID primitive.ObjectID, err error) {
	members := bson.M{
		"m_nama":  m_nama,
		"m_study": m_study,
	}
	result, err := db.Collection(col).InsertOne(context.Background(), members)
	if err != nil {
		fmt.Printf("InsertMembers: %v\n", err)
		return
	}
	insertedID = result.InsertedID.(primitive.ObjectID)
	return insertedID, nil
}

func InsertCustomers(db *mongo.Database, col string, c_nama string, c_study string) (insertedID primitive.ObjectID, err error) {
	customers := bson.M{
		"c_nama":  c_nama,
		"c_study": c_study,
	}
	result, err := db.Collection(col).InsertOne(context.Background(), customers)
	if err != nil {
		fmt.Printf("InsertCustomers: %v\n", err)
		return
	}
	insertedID = result.InsertedID.(primitive.ObjectID)
	return insertedID, nil
}

func InsertIncomes(db *mongo.Database, col string, qty int, halaman int, uang int) (insertedID primitive.ObjectID, err error) {
	incomes := bson.M{
		"qty":     qty,
		"halaman": halaman,
		"uang":    uang,
	}
	result, err := db.Collection(col).InsertOne(context.Background(), incomes)
	if err != nil {
		fmt.Printf("InsertIncomes: %v\n", err)
		return
	}
	insertedID = result.InsertedID.(primitive.ObjectID)
	return insertedID, nil
}

// get
func GetMembers(db *mongo.Database, col string) (data []model.Members) {
	members := db.Collection(col)
	filter := bson.M{}
	cursor, err := members.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("GetMembers :", err)
	}
	err = cursor.All(context.TODO(), &data)
	if err != nil {
		fmt.Println(err)
	}
	return
}

func GetCustomers(db *mongo.Database, col string) (data []model.Customers) {
	customers := db.Collection(col)
	filter := bson.M{}
	cursor, err := customers.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("GetCustomers :", err)
	}
	err = cursor.All(context.TODO(), &data)
	if err != nil {
		fmt.Println(err)
	}
	return
}

func GetIncomes(db *mongo.Database, col string) (data []model.Incomes) {
	incomes := db.Collection(col)
	filter := bson.M{}
	cursor, err := incomes.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("GetIncomes :", err)
	}
	err = cursor.All(context.TODO(), &data)
	if err != nil {
		fmt.Println(err)
	}
	return
}

func GetMembersByID(_id primitive.ObjectID, db *mongo.Database, col string) (data model.Members, errs error) {
	members := db.Collection(col)
	filter := bson.M{"_id": _id}
	err := members.FindOne(context.TODO(), filter).Decode(&data)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return data, fmt.Errorf("no data found for ID %s", _id)
		}
		return data, fmt.Errorf("error retrieving data for ID %s: %s", _id, err.Error())
	}
	return data, nil
}

func GetCustomersByID(_id primitive.ObjectID, db *mongo.Database, col string) (data model.Customers, errs error) {
	customers := db.Collection(col)
	filter := bson.M{"_id": _id}
	err := customers.FindOne(context.TODO(), filter).Decode(&data)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return data, fmt.Errorf("no data found for ID %s", _id)
		}
		return data, fmt.Errorf("error retrieving data for ID %s: %s", _id, err.Error())
	}
	return data, nil
}

func GetIncomesByID(_id primitive.ObjectID, db *mongo.Database, col string) (data model.Incomes, errs error) {
	incomes := db.Collection(col)
	filter := bson.M{"_id": _id}
	err := incomes.FindOne(context.TODO(), filter).Decode(&data)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return data, fmt.Errorf("no data found for ID %s", _id)
		}
		return data, fmt.Errorf("error retrieving data for ID %s: %s", _id, err.Error())
	}
	return data, nil
}

// update
func UpdateMembers(db *mongo.Database, col string, id primitive.ObjectID, m_nama string, m_study string) (err error) {
	filter := bson.M{"_id": id}
	update := bson.M{
		"$set": bson.M{
			"m_nama":  m_nama,
			"m_study": m_study,
		},
	}
	result, err := db.Collection(col).UpdateOne(context.Background(), filter, update)
	if err != nil {
		fmt.Printf("UpdateMembers: %v\n", err)
		return
	}
	if result.ModifiedCount == 0 {
		err = errors.New("No data has been changed with the specified ID")
		return
	}
	return nil
}

func UpdateCustomers(db *mongo.Database, col string, id primitive.ObjectID, c_nama string, c_study string) (err error) {
	filter := bson.M{"_id": id}
	update := bson.M{
		"$set": bson.M{
			"c_nama":  c_nama,
			"c_study": c_study,
		},
	}
	result, err := db.Collection(col).UpdateOne(context.Background(), filter, update)
	if err != nil {
		fmt.Printf("UpdateCustomers: %v\n", err)
		return
	}
	if result.ModifiedCount == 0 {
		err = errors.New("No data has been changed with the specified ID")
		return
	}
	return nil
}

func UpdateIncomes(db *mongo.Database, col string, id primitive.ObjectID, qty int, halaman int, uang int) (err error) {
	filter := bson.M{"_id": id}
	update := bson.M{
		"$set": bson.M{
			"qty":     qty,
			"halaman": halaman,
			"uang":    uang,
		},
	}
	result, err := db.Collection(col).UpdateOne(context.Background(), filter, update)
	if err != nil {
		fmt.Printf("UpdateIncomes: %v\n", err)
		return
	}
	if result.ModifiedCount == 0 {
		err = errors.New("No data has been changed with the specified ID")
		return
	}
	return nil
}

// delete
func DeleteMembersByID(_id primitive.ObjectID, db *mongo.Database, col string) error {
	members := db.Collection(col)
	filter := bson.M{"_id": _id}

	result, err := members.DeleteOne(context.TODO(), filter)
	if err != nil {
		return fmt.Errorf("error deleting data for ID %s: %s", _id, err.Error())
	}

	if result.DeletedCount == 0 {
		return fmt.Errorf("data with ID %s not found", _id)
	}

	return nil
}

func DeleteCustomersByID(_id primitive.ObjectID, db *mongo.Database, col string) error {
	customers := db.Collection(col)
	filter := bson.M{"_id": _id}

	result, err := customers.DeleteOne(context.TODO(), filter)
	if err != nil {
		return fmt.Errorf("error deleting data for ID %s: %s", _id, err.Error())
	}

	if result.DeletedCount == 0 {
		return fmt.Errorf("data with ID %s not found", _id)
	}

	return nil
}

func DeleteIncomesByID(_id primitive.ObjectID, db *mongo.Database, col string) error {
	incomes := db.Collection(col)
	filter := bson.M{"_id": _id}

	result, err := incomes.DeleteOne(context.TODO(), filter)
	if err != nil {
		return fmt.Errorf("error deleting data for ID %s: %s", _id, err.Error())
	}

	if result.DeletedCount == 0 {
		return fmt.Errorf("data with ID %s not found", _id)
	}

	return nil
}

// login
func InsertAdmin(db *mongo.Database, col string, username string, password string) (insertedID primitive.ObjectID, err error) {
	admin := bson.M{
		"username": username,
		"password": password,
	}
	result, err := db.Collection(col).InsertOne(context.Background(), admin)
	if err != nil {
		fmt.Printf("InsertAdmin: %v\n", err)
		return
	}
	insertedID = result.InsertedID.(primitive.ObjectID)
	return insertedID, nil
}

func LoginAdmin(db *mongo.Database, col string, username string, password string) (authenticated bool, err error) {
	filter := bson.M{
		"username": username,
		"password": password,
	}

	result, err := db.Collection(col).CountDocuments(context.Background(), filter)
	if err != nil {
		fmt.Printf("LoginAdmin: %v\n", err)
		return false, err
	}

	if result == 1 {
		return true, nil
	}

	return false, nil
}
