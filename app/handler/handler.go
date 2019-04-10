package handler

import (
	"datcom/backend/app/model"
	"datcom/backend/src/store"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func connect() store.Store {
	const connectionString = "user=postgres dbname=datcom sslmode=disable password=datcom host=localhost port=5432"
	store, err := store.NewPostgresStore(connectionString)
	if err != nil {
		fmt.Println(err)
	}
	return store
}

// InsertItems ..
func InsertItems(w http.ResponseWriter, r *http.Request) {
	store := connect()
	params := mux.Vars(r)
	var i model.Item
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&i)
	var response model.Response
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		// Bad JSON or unrecognized JSON field
		w.WriteHeader(http.StatusBadRequest)
		response.Status = http.StatusBadRequest
		response.Message = "Failed to decode JSON!"
	} else {
		err = store.InsertItemsToDB(i.ItemNames, params["menu_name"])
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			response.Status = http.StatusInternalServerError
			response.Message = "Insert Items Failed!"
		} else {
			w.WriteHeader(http.StatusCreated)
			response.Status = http.StatusCreated
			response.Message = "Items Inserted Successfully!"
		}
	}
	json.NewEncoder(w).Encode(response)
}
