package main

import (
	"testing"

	"github.com/aws/aws-lambda-go/events"
	appHttp "github.com/craguilar/demo-cars-fn/cmd/http"
	"github.com/craguilar/demo-cars-fn/internal/app/mock"
	"github.com/stretchr/testify/assert"
)

func TestHandler(t *testing.T) {

	request := events.APIGatewayProxyRequest{}
	request.Path = "/"
	expectedResponse := events.APIGatewayProxyResponse{
		StatusCode: 200,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		Body: "404 page not found",
	}

	handler := createMockHandler()
	response, err := handler.Handler(request)

	// assert.Equal(t, response.Headers, expectedResponse.Headers)
	assert.Contains(t, response.Body, expectedResponse.Body)
	assert.Equal(t, err, nil)

}

func createMockHandler() *LambaHandler {
	// Create car service and provide it to handler
	carService := mock.NewCarService()
	carHandler := appHttp.NewCarServiceHandler(carService)
	router := appHttp.NewRouter(carHandler)
	return NewLambaHandler(router)
}
