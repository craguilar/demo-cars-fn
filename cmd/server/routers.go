package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/20200201/",
		Index,
	},

	Route{
		"AddCar",
		strings.ToUpper("Post"),
		"/20200201/cars",
		AddCar,
	},

	Route{
		"GetCar",
		strings.ToUpper("Get"),
		"/20200201/cars/{carId}",
		GetCar,
	},

	Route{
		"ListCars",
		strings.ToUpper("Get"),
		"/20200201/cars",
		ListCars,
	},

	Route{
		"UpdateCar",
		strings.ToUpper("Put"),
		"/20200201/cars",
		UpdateCar,
	},
}

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Demo Cars API!")
}

func Logger(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		inner.ServeHTTP(w, r)

		log.Printf(
			"%s %s %s %s",
			r.Method,
			r.RequestURI,
			name,
			time.Since(start),
		)
	})
}
