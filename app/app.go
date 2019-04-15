package main

import (
	"datcom/backend/app/handler"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/menus/{menu_name}", handler.InsertItems).Methods("POST")
	http.Handle("/", r)
	fmt.Println("Server is running...")
	http.ListenAndServe(":3000", nil)
}
