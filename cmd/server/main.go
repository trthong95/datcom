package main

import (
	"log"

	"git.d.foundation/datcom/backend/src/app"

	_ "github.com/lib/pq"
)

func main() {
	app := &app.App{}
	app, err := app.NewApp()
	if err != nil {
		log.Fatal(err)
	}
	app.RunServer(":3000")
}
