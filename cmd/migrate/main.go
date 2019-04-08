package main

import (
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"

	"github.com/hieunmce/datcom/src/store"
)

func main() {
	connectionString := "user=postgres dbname=datcom sslmode=disable password=datcom host=localhost port=5432"
	store, err := store.NewPostgresMigrator(connectionString)
	if err != nil {
		fmt.Println(err)
	}
	_ = store

	flag := os.Args[1]
	if flag != "up" && flag != "down" {
		log.Fatalln("up down")
	}

	switch flag {
	case "up":
		// err = store.MigrateDB()
		err = store.MigrateOneDB()
		if err != nil {
			fmt.Println(err)
			return
		}
	case "down":
		// err = store.ReverseDB()
		err = store.ReverseOneDB()
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}
