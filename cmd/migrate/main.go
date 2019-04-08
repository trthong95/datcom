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
	err = store.OneMigrateDB()
	_ = err
	// switch flag {
	// case "up":
	// 	err = store.MigrateDB()
	// 	if err != nil {
	// 		fmt.Println(err)
	// 		return
	// 	}
	// case "down":
	// 	err = store.ReverseDB()
	// 	if err != nil {
	// 		fmt.Println(err)
	// 		return
	// 	}
	// }
}
