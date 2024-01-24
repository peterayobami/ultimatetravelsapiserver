package main

import (
	// "fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/peterayobami/ultimatetravelsapiserver/controllers"
	"github.com/peterayobami/ultimatetravelsapiserver/database"
	"github.com/peterayobami/ultimatetravelsapiserver/initializers"
	"github.com/peterayobami/ultimatetravelsapiserver/routes"
)

func init() {
	initializers.LoadConfiguration()
	database.Migrate()
}

func middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
	})
}

func main() {
	// Create a router instance
	var r = mux.NewRouter()

	// Map controllers
	r.Handle(routes.PersistFlightRequest, middleware(http.HandlerFunc(controllers.PersistFlightRequest))).Methods(http.MethodPost)
	r.Handle(routes.InitializeFlightBooking, middleware(http.HandlerFunc(controllers.InitializeBooking))).Methods(http.MethodPost)

	// Start the server with the router
	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}
