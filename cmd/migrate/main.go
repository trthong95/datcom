package main

import (
	"fmt"

	"github.com/hieunmce/datcom/src/store"
	_ "github.com/lib/pq"
)

func main() {
	connectionString := "user=postgres dbname=datcom sslmode=disable password=datcom host=localhost port=5432"
	store, err := store.NewPostgresMigrator(connectionString)
	if err != nil {
		fmt.Println(err)
	}

	err = store.MigrateDB()
	if err != nil {
		fmt.Println(err)
	}
}
