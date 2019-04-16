package store_test

import (
	"datcom/backend/src/store"
	"testing"

	_ "github.com/lib/pq"
)

func ConnectTestDB() (store.Store, error) {
	connectionString := "user=postgres dbname=datcom_test sslmode=disable password=datcom host=localhost port=5432"
	store, err := store.NewPostgresStore(connectionString)
	return store, err
}
func TestAddOrdersExist(t *testing.T) {
	// Connect to test DB
	storeif, err := ConnectTestDB()
	if err != nil {
		t.Error("Failed to connect DB")
	}
	// Order of Item 4 already exist
	err = storeif.AddOrders(1, 4)
	if err == nil {
		t.Error("Test failed")
	}
}
func TestAddOrdersNotExist(t *testing.T) {
	// Connect to test DB
	storeif, err := ConnectTestDB()
	if err != nil {
		t.Error("Failed to connect DB")
	}
	// Order of Item 5 does not exist
	err = storeif.AddOrders(1, 5)
	if err != nil {
		t.Error(err)
	}
}

func TestDeleteOrdersExist(t *testing.T) {
	// Connect to test DB
	storeif, err := ConnectTestDB()
	if err != nil {
		t.Error("Failed to connect DB")
	}
	// Order of Item 4 exists
	err = storeif.DeleteOrders(1, 4)
	if err != nil {
		t.Error(err)
	}
}

func TestDeleteOrdersNotExist(t *testing.T) {
	// Connect to test DB
	storeif, err := ConnectTestDB()
	if err != nil {
		t.Error("Failed to connect DB")
	}
	// Order of Item 6 does not exist
	err = storeif.DeleteOrders(1, 6)
	if err == nil {
		t.Error("Test failed")

	}
}

func TestGetOrder(t *testing.T) {
	// Connect to test DB
	storeif, err := ConnectTestDB()
	if err != nil {
		t.Error("Failed to connect DB")
	}
	// Get Orders of User 1
	_, err = storeif.GetOrder(1)
	if err != nil {
		t.Error("Test failed")
	}

}
