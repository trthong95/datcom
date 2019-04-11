package handler

import (
	"datcom/backend/app/model"
	"datcom/backend/src/store"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

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

// DeleteItem ..
func DeleteItem(w http.ResponseWriter, r *http.Request) {
	store := connect()
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	err := store.DeleteItemFromDB(id)
	var response model.Response
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		if err.Error() == "non-exist" {
			w.WriteHeader(http.StatusNotFound)
			response.Status = http.StatusNotFound
			response.Message = "Item Doesn't Exist!"
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			response.Status = http.StatusInternalServerError
			response.Message = "Delete Item Failed!"
		}
	} else {
		w.WriteHeader(http.StatusOK)
		response.Status = http.StatusOK
		response.Message = "Item Deleted Successfully!"
	}
	json.NewEncoder(w).Encode(response)
}
