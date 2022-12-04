package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/craguilar/demo-cars-fn/cmd"
	appHttp "github.com/craguilar/demo-cars-fn/cmd/http"
	"github.com/craguilar/demo-cars-fn/internal/app/mock"
)

func main() {
	log.Printf("Server started")

	// Create car service and provide it to handler
	carService := mock.NewCarService()
	handler := appHttp.NewCarServiceHandler(carService)
	router := appHttp.NewRouter(handler)
	// Start server
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", cmd.GetConfig("PORT")), router))
}
