package main

import (
	"fmt"

	"datcom/backend/src/store"

	_ "github.com/lib/pq"
)

// type person Struct{
// 	Name string
// 	Email string
// 	token string
// }

func main() {
	connectionString := "user=postgres dbname=datcom sslmode=disable password=datcom host=localhost port=5432"
	storeif, err := store.NewPostgresStore(connectionString)
	if err != nil {
		fmt.Println(err)
	}

	err = store.AddOrders(1, 2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("OK")
	}
}
