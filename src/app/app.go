package app

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"git.d.foundation/datcom/backend/src/service"
)

// App struct for router and db
type App struct {
	Router  *mux.Router
	Service *service.Service
}

func (a *App) NewApp() (*App, error) {
	connectionString := "user=postgres dbname=datcom sslmode=disable password=datcom host=localhost port=5432"
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, errors.New("could not connect to database")
	}

	a.Service = service.NewService(db)

	a.Router = mux.NewRouter()
	a.SetRouters()

	return a, nil
}

func (a *App) SetRouters() {
	a.Router.HandleFunc("/auth/google/login-url", a.GetGoogleLoginURL).Methods("GET")
	a.Router.HandleFunc("/auth/login/google", a.VerifyGoogleUserLogin).Methods("POST")
	a.Router.HandleFunc("/menus", a.GetAllMenus).Methods("GET")
	a.Router.HandleFunc("/menus", a.CreateMenuFromItems).Methods("POST")
	a.Router.HandleFunc("/menus/{MenuID}", a.GetItemsFromMenu).Methods("GET")
	a.Router.HandleFunc("/menus/{MenuID}", a.AddItemsToMenu).Methods("POST")
	a.Router.HandleFunc("/menus/{MenuID}", a.DeleteMenu).Methods("DELETE")
	a.Router.HandleFunc("/menus/{MenuID}/{ItemID}", a.DeleteItemFromMenu).Methods("DELETE")
	a.Router.HandleFunc("/menus/{MenuID}/people-in-charge", a.GetPeopleInCharge).Methods("GET")
	a.Router.HandleFunc("/menus/{MenuID}/people-in-charge", a.AddPeopleInCharge).Methods("POST")
	a.Router.HandleFunc("/menus/{MenuID}/summary", a.GetMenuSummary).Methods("GET")
	a.Router.HandleFunc("/orders", a.GetOrdersOfAllUsers).Methods("GET")
	a.Router.HandleFunc("/orders/{UserID}", a.GetOrderOfUser).Methods("GET")
	a.Router.HandleFunc("/orders/{UserID}", a.EditUserOrder).Methods("PUT")
	a.Router.HandleFunc("/orders/{UserID}", a.DeleteOrders).Methods("DELETE")
}

func (a *App) GetGoogleLoginURL(w http.ResponseWriter, r *http.Request) {
	// handler.GetGoogleLoginURL(a.DB, w, r)
	w.WriteHeader(http.StatusOK)
}

func (a *App) VerifyGoogleUserLogin(w http.ResponseWriter, r *http.Request) {
	// handler.VerifyGoogleUserLogin(a.DB, w, r)
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

func (a *App) DeleteOrders(w http.ResponseWriter, r *http.Request) {
	// handler.DeleteOrders(a.DB, w, r)
	w.WriteHeader(http.StatusOK)
}

func (a *App) RunServer(host string) {
	fmt.Printf("server in running at %s\n", host)
	log.Fatal(http.ListenAndServe(host, a.Router))
}
