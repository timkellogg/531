package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/timkellogg/531/server/config"
	"github.com/timkellogg/531/server/controllers"
)

// Route is a REST endpoint that maps common methods
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes is a collection of routes
type Routes []Route

var routes = Routes{
	Route{"Registration", "POST", "/api/users", controllers.UsersCreate},
	Route{"Login", "POST", "/api/auth/login", controllers.AuthLogin},
	Route{"Logout", "POST", "/api/auth/logout", controllers.AuthLogout},
}

// Cors - enable logging
type Cors struct {
	Log *log.Logger
}

// NewRouter establishes the root application router
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	for _, route := range routes {
		var handler http.Handler

		handler = route.HandlerFunc
		handler = config.Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}
