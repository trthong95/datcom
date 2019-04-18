package main

import (
	"database/sql"

	"github.com/k0kubun/pp"

	"git.d.foundation/datcom/backend/src/domain"
	"git.d.foundation/datcom/backend/src/service"

	_ "github.com/lib/pq"
)

func main() {
	connectionString := "user=postgres dbname=datcom sslmode=disable password=datcom host=localhost port=5432"
	db, _ := sql.Open("postgres", connectionString)
	ser := service.NewService(db)
	p := domain.PersonInfo{"Thong Nguyen", "nguyentrthong95@gmail.com", "xyz"}
	ser.CreateUser(&p)
	users, _ := ser.GetAllUser()
	for _, user := range users {
		pp.Println(user)
	}
}
