package main

import (
	"log"

	"github.com/aws/aws-lambda-go/lambda"

	appHttp "github.com/craguilar/demo-cars-fn/cmd/http"

	"github.com/craguilar/demo-cars-fn/internal/app/mock"
)

func main() {
	log.Printf("Lambda started")
	// Create car service and provide it to handler
	carService := mock.NewCarService()
	carHandler := appHttp.NewCarServiceHandler(carService)
	router := appHttp.NewRouter(carHandler)
	handler := NewLambaHandler(router)
	// Start lambda
	lambda.Start(handler.Handler)
}
