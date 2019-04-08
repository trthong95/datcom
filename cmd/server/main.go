package main

import (
	"fmt"

	"github.com/hieunmce/datcom/src/store"
	_ "github.com/lib/pq"
)

func main() {
	connectionString := "user=postgres dbname=datcom sslmode=disable password=datcom host=localhost port=5432"
	store, err := store.NewPostgresStore(connectionString)
	if err != nil {
		fmt.Println(err)
	}
	err = store.SetUserAndEmail("Thong Nguyen", "nanolove95@gmail.com")
	if err != nil {
		fmt.Println(err)
	}

	users, err := store.GetAllUsers()
	if err != nil {
		fmt.Println(err)
	}
	for _, user := range users {
		fmt.Printf("Username: %v\n", user.Name)
		fmt.Printf("Email: %v\n", user.Email)
	}

}
