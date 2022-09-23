package http

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

func NewRouter(handler *CarsServiceHandler) *mux.Router {
	var routes = []Route{
		{
			"Index",
			"GET",
			"/20200201/",
			Index,
		}, {
			"AddCar",
			strings.ToUpper("Post"),
			"/20200201/cars",
			handler.AddCar,
		}, {
			"GetCar",
			strings.ToUpper("Get"),
			"/20200201/cars/{carId}",
			handler.GetCar,
		}, {
			"ListCars",
			strings.ToUpper("Get"),
			"/20200201/cars",
			handler.ListCars,
		}, {
			"UpdateCar",
			strings.ToUpper("Put"),
			"/20200201/cars",
			handler.UpdateCar,
		},
	}
	//
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = SetupGlobalMiddleware(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}
	// OPTIONS Method no op handler
	router.
		Methods("OPTIONS").
		Name("OptionsNoOp").
		PathPrefix("/20200201").
		Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		}))

	return router
}

// Go Server API Index API!
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "DemoCars API is UP!")
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func SetupGlobalMiddleware(handler http.Handler, name string) http.Handler {
	return LoggerMiddleWare(JsonContentTypeMiddleWare(Cors(handler)), name)
}

// Set application/json for all Responses in this Server
func JsonContentTypeMiddleWare(inner http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		inner.ServeHTTP(w, r)
	})
}

// Log all Requests in this server
func LoggerMiddleWare(inner http.Handler, name string) http.Handler {
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

func Cors(inner http.Handler) http.Handler {
	// Where ORIGIN_ALLOWED is like `scheme://dns[:port]`, or `*` (insecure)
	headersOk := handlers.AllowedHeaders([]string{"*"})
	originsOk := handlers.AllowedOrigins([]string{"localhost:3000"})
	methodsOk := handlers.AllowedMethods([]string{"*"})

	return handlers.CORS(originsOk, headersOk, methodsOk)(inner)
}
