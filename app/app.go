package app

import (
	"datcom/backend/src/store"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// App struct for router and db
type App struct {
	Router *mux.Router
	DB     store.Store
}

func (a *App) Init() {
	connectionString := "user=postgres dbname=datcom sslmode=disable password=datcom host=localhost port=5432"
	storeif, err := store.NewPostgresStore(connectionString)
	if err != nil {
		log.Fatal("Could not connect to database")
	}
	a.DB = storeif
	a.Router = mux.NewRouter()
	a.SetRouters()
}

func (a *App) SetRouters() {
	a.Router.HandleFunc("/users", a.GetAllUsers).Methods("GET")
	a.Router.HandleFunc("/users", a.CreateUser).Methods("POST")
	a.Router.HandleFunc("/menus", a.GetAllMenus).Methods("GET")
	a.Router.HandleFunc("/menus", a.CreateMenuFromItems).Methods("POST")
	a.Router.HandleFunc("/menus/{name}", a.GetItemsFromMenu).Methods("GET")
	a.Router.HandleFunc("/menus/{name}", a.AddItemsToMenu).Methods("POST")
	a.Router.HandleFunc("/menus/{name}", a.DeleteMenu).Methods("DELETE")
	a.Router.HandleFunc("/menus/{name}/{item}", a.DeleteItemFromMenu).Methods("DELETE")
	a.Router.HandleFunc("/menus/{name}/people_in_charge", a.GetPeopleInCharge).Methods("GET")
	a.Router.HandleFunc("/menus/{name}/people_in_charge", a.AddPeopleInCharge).Methods("POST")
	a.Router.HandleFunc("/menus/{name}/summary", a.GetMenuSummary).Methods("GET")
	a.Router.HandleFunc("/orders", a.GetOrdersOfAllUsers).Methods("GET")
	a.Router.HandleFunc("/orders/{user}", a.GetOrderOfUser).Methods("GET")
	a.Router.HandleFunc("/orders/{user}", a.EditUserOrder).Methods("PUT")
	a.Router.HandleFunc("/orders/{user}", a.CancelUserOrders).Methods("DELETE")
}

func (a *App) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	// handler.GetAllUsers(a.DB, w, r)
	w.WriteHeader(http.StatusOK)
}

func (a *App) CreateUser(w http.ResponseWriter, r *http.Request) {
	// handler.CreateUser(a.DB, w, r)
	w.WriteHeader(http.StatusOK)
}

func (a *App) GetAllMenus(w http.ResponseWriter, r *http.Request) {
	// handler.GetAllMenus(a.DB, w, r)
	w.WriteHeader(http.StatusOK)
}

func (a *App) CreateMenuFromItems(w http.ResponseWriter, r *http.Request) {
	// handler.CreateMenuFromItems(a.DB, w, r)
	w.WriteHeader(http.StatusOK)
}

func (a *App) GetItemsFromMenu(w http.ResponseWriter, r *http.Request) {
	// handler.GetItemsFromMenu(a.DB, w, r)
	w.WriteHeader(http.StatusOK)
}

func (a *App) AddItemsToMenu(w http.ResponseWriter, r *http.Request) {
	// handler.AddItemsToMenu(a.DB, w, r)
	w.WriteHeader(http.StatusOK)
}

func (a *App) DeleteMenu(w http.ResponseWriter, r *http.Request) {
	// handler.DeleteMenu(a.DB, w, r)
	w.WriteHeader(http.StatusOK)
}

func (a *App) DeleteItemFromMenu(w http.ResponseWriter, r *http.Request) {
	// handler.DeleteItemFromMenu(a.DB, w, r)
	w.WriteHeader(http.StatusOK)
}

func (a *App) GetPeopleInCharge(w http.ResponseWriter, r *http.Request) {
	// handler.GetPeopleInCharge(a.DB, w, r)
	w.WriteHeader(http.StatusOK)
}

func (a *App) AddPeopleInCharge(w http.ResponseWriter, r *http.Request) {
	// handler.AddPeopleInCharge(a.DB, w, r)
	w.WriteHeader(http.StatusOK)
}

func (a *App) GetMenuSummary(w http.ResponseWriter, r *http.Request) {
	// handler.GetMenuSummary(a.DB, w, r)
	w.WriteHeader(http.StatusOK)
}

func (a *App) GetOrdersOfAllUsers(w http.ResponseWriter, r *http.Request) {
	// handler.GetOrdersOfAllUsers(a.DB, w, r)
	w.WriteHeader(http.StatusOK)
}

func (a *App) GetOrderOfUser(w http.ResponseWriter, r *http.Request) {
	// handler.GetOrderOfUser(a.DB, w, r)
	w.WriteHeader(http.StatusOK)
}

func (a *App) EditUserOrder(w http.ResponseWriter, r *http.Request) {
	// handler.EditUserOrder(a.DB, w, r)
	w.WriteHeader(http.StatusOK)
}

func (a *App) CancelUserOrders(w http.ResponseWriter, r *http.Request) {
	// handler.CancelUserOrders(a.DB, w, r)
	w.WriteHeader(http.StatusOK)
}

func (a *App) RunServer(host string) {
	log.Fatal(http.ListenAndServe(host, a.Router))
}
