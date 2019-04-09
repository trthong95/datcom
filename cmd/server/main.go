package main

import (
	"fmt"

	_ "github.com/lib/pq"
	"gitlab.com/thongnt/lunch-order/src/store"
)

func main() {
	connectionString := "user=postgres dbname=datcom sslmode=disable password=datcom host=localhost port=5432"
	store, err := store.NewPostgresStore(connectionString)
	if err != nil {
		fmt.Println(err)
	}

	users, err := store.GetAllUsers()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(users)
}
