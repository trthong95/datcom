package main

import (
	"git.d.foundation/datcom/backend/app"

	_ "github.com/lib/pq"
)

func main() {
	app := &app.App{}
	app.Init()
	app.RunServer(":3000")
}
