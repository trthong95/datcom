package main

import (
	"fmt"

	"datcom/backend/src/store"

	"github.com/k0kubun/pp"
	_ "github.com/lib/pq"
)

func main() {
	connectionString := "user=postgres dbname=datcom sslmode=disable password=datcom host=localhost port=5432"
	storeif, err := store.NewPostgresStore(connectionString)
	if err != nil {
		fmt.Println(err)
	}

	p := store.PersonInfo{"Thong Nguyen", "nguyentrthong95@gmail.com", "xyz"}
	err = storeif.CreateUser(&p)
	if err != nil {
		fmt.Println(err)
	}

	users, err := storeif.GetAllUsers()
	if err != nil {
		fmt.Println(err)
	}
	for _, user := range users {
		pp.Println(user)
	}
}
