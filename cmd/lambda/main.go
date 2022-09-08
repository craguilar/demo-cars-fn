package main

import (
	"encoding/json"
	"io/ioutil"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	cars "github.com/craguilar/demo-cars-fn/internal"
)

// Handler is executed by AWS Lambda in the main function. Once the request
// is processed, it returns an Amazon API Gateway response object to AWS Lambda
func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	if request.Path == "/api" {
		index, err := ioutil.ReadFile("api/swagger-carsdemo.json")
		if err != nil {
			return events.APIGatewayProxyResponse{}, err
		}
		return events.APIGatewayProxyResponse{
			StatusCode: 200,
			Body:       string(index),
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
		}, nil
	}
	// Function Handlers
	car, error := cars.GetCar("GDL-123")
	if error != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       "InternalServerError",
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
		}, nil
	}
	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       string(serializeData(car)),
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	}, nil
}

func serializeData(data interface{}) []byte {
	result, error := json.Marshal(data)
	if error != nil {
		return nil
	}
	return result
}

func main() {
	lambda.Start(Handler)
}
